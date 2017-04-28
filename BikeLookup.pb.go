// Code generated by protoc-gen-go.
// source: BikeLookup.proto
// DO NOT EDIT!

/*

It is generated from these files:
	BikeLookup.proto

It has these top-level messages:
	BikeLookupRequest
	BikeLookupReply
	BikeRack
	Point
*/
package main

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The request message containing the user's name.
type BikeLookupRequest struct {
	Suburb      string `protobuf:"bytes,1,opt,name=suburb" json:"suburb,omitempty"`
	MinCapacity int32  `protobuf:"varint,2,opt,name=min_capacity,json=minCapacity" json:"min_capacity,omitempty"`
}

func (m *BikeLookupRequest) Reset()                    { *m = BikeLookupRequest{} }
func (m *BikeLookupRequest) String() string            { return proto.CompactTextString(m) }
func (*BikeLookupRequest) ProtoMessage()               {}
func (*BikeLookupRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// The response message containing the greetings
type BikeLookupReply struct {
	BikeRack  []*BikeRack `protobuf:"bytes,1,rep,name=bike_rack,json=bikeRack" json:"bike_rack,omitempty"`
	Timestamp int64       `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *BikeLookupReply) Reset()                    { *m = BikeLookupReply{} }
func (m *BikeLookupReply) String() string            { return proto.CompactTextString(m) }
func (*BikeLookupReply) ProtoMessage()               {}
func (*BikeLookupReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BikeLookupReply) GetBikeRack() []*BikeRack {
	if m != nil {
		return m.BikeRack
	}
	return nil
}

type BikeRack struct {
	Suburb   string `protobuf:"bytes,1,opt,name=suburb" json:"suburb,omitempty"`
	Address  string `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	Capacity int32  `protobuf:"varint,3,opt,name=capacity" json:"capacity,omitempty"`
	Position *Point `protobuf:"bytes,4,opt,name=position" json:"position,omitempty"`
}

func (m *BikeRack) Reset()                    { *m = BikeRack{} }
func (m *BikeRack) String() string            { return proto.CompactTextString(m) }
func (*BikeRack) ProtoMessage()               {}
func (*BikeRack) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BikeRack) GetPosition() *Point {
	if m != nil {
		return m.Position
	}
	return nil
}

type Point struct {
	Latitude  float64 `protobuf:"fixed64,1,opt,name=latitude" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude" json:"longitude,omitempty"`
}

func (m *Point) Reset()                    { *m = Point{} }
func (m *Point) String() string            { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()               {}
func (*Point) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*BikeLookupRequest)(nil), "helloworld.BikeLookupRequest")
	proto.RegisterType((*BikeLookupReply)(nil), "helloworld.BikeLookupReply")
	proto.RegisterType((*BikeRack)(nil), "helloworld.BikeRack")
	proto.RegisterType((*Point)(nil), "helloworld.Point")
}

func init() { proto.RegisterFile("BikeLookup.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x52, 0xdf, 0x4a, 0xfb, 0x30,
	0x14, 0xfe, 0x75, 0xfb, 0x6d, 0xb6, 0x67, 0xc2, 0x5c, 0x10, 0x29, 0xd3, 0x81, 0xf6, 0x6a, 0x37,
	0x16, 0xac, 0x4f, 0xe0, 0x04, 0xf1, 0x42, 0x64, 0x04, 0xc1, 0xcb, 0x91, 0xb6, 0x61, 0x86, 0xa5,
	0x3d, 0x35, 0x4d, 0x95, 0x3d, 0x81, 0xaf, 0x6d, 0x1a, 0xbb, 0x76, 0x0a, 0xbb, 0xfb, 0xfe, 0x24,
	0x5f, 0xbf, 0x73, 0x1a, 0x38, 0x59, 0x88, 0x0d, 0x7f, 0x42, 0xdc, 0x54, 0x45, 0x58, 0x28, 0xd4,
	0x48, 0xe0, 0x8d, 0x4b, 0x89, 0x9f, 0xa8, 0x64, 0x1a, 0x3c, 0xc3, 0xa4, 0xf3, 0x29, 0x7f, 0xaf,
	0x78, 0xa9, 0xc9, 0x19, 0x0c, 0xcb, 0x2a, 0xae, 0x54, 0xec, 0x3b, 0x97, 0xce, 0xdc, 0xa3, 0x0d,
	0x23, 0x57, 0x70, 0x9c, 0x89, 0x7c, 0x95, 0xb0, 0x82, 0x25, 0x42, 0x6f, 0xfd, 0x9e, 0x71, 0x07,
	0x74, 0x64, 0xb4, 0xfb, 0x46, 0x0a, 0x62, 0x18, 0xef, 0xe7, 0x15, 0x72, 0x4b, 0x6e, 0xc0, 0x8b,
	0x8d, 0xb4, 0x52, 0x2c, 0xd9, 0x98, 0xc0, 0xfe, 0x7c, 0x14, 0x9d, 0x86, 0x5d, 0x85, 0xb0, 0x3e,
	0x4f, 0x8d, 0x47, 0xdd, 0xb8, 0x41, 0xe4, 0x02, 0x3c, 0x2d, 0x32, 0x53, 0x85, 0x65, 0x85, 0xfd,
	0x4a, 0x9f, 0x76, 0x42, 0xf0, 0xe5, 0x80, 0xbb, 0xbb, 0x74, 0xb0, 0xab, 0x0f, 0x47, 0x2c, 0x4d,
	0x15, 0x2f, 0x4b, 0x1b, 0xe0, 0xd1, 0x1d, 0x25, 0x53, 0x70, 0xdb, 0x09, 0xfa, 0x76, 0x82, 0x96,
	0x93, 0x6b, 0x70, 0x0b, 0x2c, 0x85, 0x16, 0x98, 0xfb, 0xff, 0x8d, 0x37, 0x8a, 0x26, 0xfb, 0x55,
	0x97, 0x28, 0x72, 0x4d, 0xdb, 0x23, 0xc1, 0x1d, 0x0c, 0xac, 0x54, 0x67, 0x4a, 0xa6, 0x85, 0xae,
	0x52, 0x6e, 0x7b, 0x38, 0xb4, 0xe5, 0xf5, 0x30, 0x12, 0xf3, 0xf5, 0x8f, 0xd9, 0xb3, 0x66, 0x27,
	0x44, 0x2f, 0x00, 0xdd, 0xc2, 0xc8, 0x03, 0x0c, 0x1b, 0x34, 0xfb, 0xbb, 0xa2, 0x5f, 0xbf, 0x68,
	0x7a, 0x7e, 0xc8, 0x36, 0x1b, 0x0f, 0xfe, 0x2d, 0x22, 0x98, 0x25, 0x98, 0x85, 0xa5, 0x90, 0x1f,
	0x5c, 0x29, 0x26, 0x64, 0x88, 0xd9, 0x7a, 0xef, 0xca, 0x62, 0xfc, 0x58, 0xe3, 0xd7, 0x1a, 0x2f,
	0xeb, 0x47, 0xb1, 0x74, 0xe2, 0xa1, 0x7d, 0x1d, 0xb7, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x45,
	0x54, 0xbe, 0x4e, 0x31, 0x02, 0x00, 0x00,
}
