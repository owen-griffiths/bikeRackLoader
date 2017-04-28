package main

import (
  "encoding/csv"
  "flag"
  "net/http"
  "log"
  "strconv"
  "github.com/golang/protobuf/proto"
  "time"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/s3"
  "github.com/aws/aws-sdk-go/aws"
  "bytes"
)

var sourcePath = flag.String("url",
  "https://www.data.brisbane.qld.gov.au/data/dataset/404d9718-33c2-4ec2-8a8f-ed0195a151d0/resource/4a67a16d-ffc7-4831-a77b-64d8ac42459e/download/CBD-bike-racks.csv",
  "URL to download CSV data from")

var s3region = flag.String("s3region", "ap-southeast-2", "S3 region that s3 bucket is in")
var s3Bucket = flag.String("s3bucket", "experimental-syd", "Name of S3 bucket to upload protobuf data set to")
var outputName = flag.String("outputName", "bikeRacks.pb", "Name of file to save data to")

func main() {
  flag.Parse()

  var header []string
  var dataRecords [][]string
  downloadBikeCsv(*sourcePath, &header, &dataRecords)

  recordSet := csvToProto(header, dataRecords)

  data, err := proto.Marshal(recordSet)
  if err != nil {
    log.Panic("Error saving record set: ", err)
  }
  log.Println("Data size", len(data), "[b]")

  uploadDataToS3(data, *s3region, *s3Bucket, *outputName)
}

func downloadBikeCsv(url string, header *[]string, dataRecords *[][]string) {
  log.Println("Downloading CSV data from", *sourcePath)
  resp, err := http.Get(*sourcePath)
  if err != nil {
    log.Panic("Error downloading CSV data from ", *sourcePath)
  }
  defer resp.Body.Close()

  log.Println("Response Length:", resp.ContentLength, "[b]")
  reader := csv.NewReader(resp.Body)

  *header, err = reader.Read()
  if err != nil {
    log.Panic("Error reading header values: ", err)
  }

  *dataRecords, err = reader.ReadAll()
  if err != nil {
    log.Panic("Error reading data records ", err)
  }

  log.Println("Loaded", len(*dataRecords), "data records")
}

func csvToProto(header[] string, dataRecords [][]string) *BikeLookupReply {
  suburbCol := findColumn("Suburb", header)
  addressCol := findColumn("Address", header)
  sizeCol := findColumn("Capacity", header)
  latCol := findColumn("Latitude", header)
  longCol := findColumn("Longitude", header)

  recordSet := &BikeLookupReply{Timestamp: time.Now().Unix() * 1000}

  for _, record := range dataRecords {
    cap, err := strconv.ParseInt(record[sizeCol], 10, 32)
    if err != nil {
      cap = 0
    }

    pos := parsePoint(record[latCol], record[longCol])

    pbRecord := &BikeRack {
      Suburb: record[suburbCol],
      Address: record[addressCol],
      Capacity: int32(cap),
      Position: pos,
    }

    recordSet.BikeRack = append(recordSet.BikeRack, pbRecord)
  }

  log.Println("Record Set has", len(recordSet.BikeRack), "records")
  return recordSet
}

func findColumn(name string, header []string) int {
  for i, h := range header {
    if h == name {
      log.Println("Column", name, "is at", i)
      return i;
    }
  }
  log.Panic("Unable to find column named ", name)
  return -1;
}

// Result will be nil if lat or long don't convert to float64
func parsePoint(lat string, long string) *Point {
  latVal, latErr := strconv.ParseFloat(lat, 64)
  longVal, longErr := strconv.ParseFloat(long, 64)

  if (latErr == nil) && (longErr == nil) {
    return &Point {
      Latitude: latVal,
      Longitude: longVal,
    }
  }
  return nil
}

func uploadDataToS3(data []byte, region string, bucket string, outputName string) {
  log.Println("Attempting to save", len(data), "[b] to S3 Region:", region, "Bucket:", bucket, "Path:", outputName)

  sess := session.Must(session.NewSessionWithOptions(session.Options{
            SharedConfigState: session.SharedConfigEnable,
            Config: aws.Config{Region: aws.String(region)},
         }))
  svc := s3.New(sess)

  putReq := &s3.PutObjectInput {
    Bucket: aws.String(bucket),
    Key: aws.String(outputName),
    Body: bytes.NewReader(data),
  }
  _, err := svc.PutObject(putReq)
  if err != nil {
    log.Panic("Unable to save ", len(data), "[b] to S3 Bucket: ", s3Bucket, " Path: ", outputName, " - ", err)
  }

  log.Println("Successfully saved", len(data), "[b] to S3 Bucket:", bucket, "Path:", outputName)
}
