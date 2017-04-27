package main

import (
  "encoding/csv"
  "flag"
  "net/http"
  "log"
)

var sourcePath = flag.String("url",
  "https://www.data.brisbane.qld.gov.au/data/dataset/404d9718-33c2-4ec2-8a8f-ed0195a151d0/resource/4a67a16d-ffc7-4831-a77b-64d8ac42459e/download/CBD-bike-racks.csv",
  "URL to download CSV data from")

func main() {
  flag.Parse()

  log.Println("Downloading CSV data from", *sourcePath)
  resp, err := http.Get(*sourcePath)
  if err != nil {
    log.Panic("Error downloading CSV data from ", *sourcePath)
  }
  defer resp.Body.Close()

  reader := csv.NewReader(resp.Body)

  header, err := reader.Read()
  if err != nil {
    log.Panic("Error reading header values: ", err)
  }

  suburbCol := findColumn("Suburb", header)
  addressCol := findColumn("Address", header)
  sizeCol := findColumn("Capacity", header)
  latCol := findColumn("Latitude", header)
  longCol := findColumn("Longitude", header)

  dataRecords, err := reader.ReadAll()
  if err != nil {
    log.Panic("Error reading data records ", err)
  }

  log.Println("Loaded", len(dataRecords), "data records")

  for _, record := range dataRecords {
    suburb := record[suburbCol]
    address := record[addressCol]
    size := record[sizeCol]
    lat := record[latCol]
    long := record[longCol]

    log.Println(suburb, address, size, lat, long)
  }
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
