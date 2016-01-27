// Code generated by protoc-gen-go.
// source: flow/flow.proto
// DO NOT EDIT!

/*
Package flow is a generated protocol buffer package.

It is generated from these files:
	flow/flow.proto

It has these top-level messages:
	FlowEndpointStatistics
	FlowEndpointsStatistics
	FlowStatistics
	Flow
*/
package flow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type FlowEndpointType int32

const (
	FlowEndpointType_ETHERNET FlowEndpointType = 0
	FlowEndpointType_IPV4     FlowEndpointType = 1
	FlowEndpointType_TCPPORT  FlowEndpointType = 2
	FlowEndpointType_UDPPORT  FlowEndpointType = 3
	FlowEndpointType_SCTPPORT FlowEndpointType = 4
)

var FlowEndpointType_name = map[int32]string{
	0: "ETHERNET",
	1: "IPV4",
	2: "TCPPORT",
	3: "UDPPORT",
	4: "SCTPPORT",
}
var FlowEndpointType_value = map[string]int32{
	"ETHERNET": 0,
	"IPV4":     1,
	"TCPPORT":  2,
	"UDPPORT":  3,
	"SCTPPORT": 4,
}

func (x FlowEndpointType) String() string {
	return proto.EnumName(FlowEndpointType_name, int32(x))
}
func (FlowEndpointType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type FlowEndpointStatistics struct {
	Type    FlowEndpointType `protobuf:"varint,1,opt,name=Type,enum=flow.FlowEndpointType" json:"Type,omitempty"`
	Value   string           `protobuf:"bytes,2,opt,name=Value" json:"Value,omitempty"`
	Packets uint64           `protobuf:"varint,5,opt,name=Packets" json:"Packets,omitempty"`
	Bytes   uint64           `protobuf:"varint,6,opt,name=Bytes" json:"Bytes,omitempty"`
}

func (m *FlowEndpointStatistics) Reset()                    { *m = FlowEndpointStatistics{} }
func (m *FlowEndpointStatistics) String() string            { return proto.CompactTextString(m) }
func (*FlowEndpointStatistics) ProtoMessage()               {}
func (*FlowEndpointStatistics) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type FlowEndpointsStatistics struct {
	Type FlowEndpointType        `protobuf:"varint,1,opt,name=Type,enum=flow.FlowEndpointType" json:"Type,omitempty"`
	AB   *FlowEndpointStatistics `protobuf:"bytes,3,opt,name=AB" json:"AB,omitempty"`
	BA   *FlowEndpointStatistics `protobuf:"bytes,4,opt,name=BA" json:"BA,omitempty"`
}

func (m *FlowEndpointsStatistics) Reset()                    { *m = FlowEndpointsStatistics{} }
func (m *FlowEndpointsStatistics) String() string            { return proto.CompactTextString(m) }
func (*FlowEndpointsStatistics) ProtoMessage()               {}
func (*FlowEndpointsStatistics) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FlowEndpointsStatistics) GetAB() *FlowEndpointStatistics {
	if m != nil {
		return m.AB
	}
	return nil
}

func (m *FlowEndpointsStatistics) GetBA() *FlowEndpointStatistics {
	if m != nil {
		return m.BA
	}
	return nil
}

type FlowStatistics struct {
	Start     int64                              `protobuf:"varint,1,opt,name=Start" json:"Start,omitempty"`
	Last      int64                              `protobuf:"varint,2,opt,name=Last" json:"Last,omitempty"`
	Endpoints map[int32]*FlowEndpointsStatistics `protobuf:"bytes,3,rep,name=Endpoints" json:"Endpoints,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *FlowStatistics) Reset()                    { *m = FlowStatistics{} }
func (m *FlowStatistics) String() string            { return proto.CompactTextString(m) }
func (*FlowStatistics) ProtoMessage()               {}
func (*FlowStatistics) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FlowStatistics) GetEndpoints() map[int32]*FlowEndpointsStatistics {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

type Flow struct {
	UUID           string          `protobuf:"bytes,1,opt,name=UUID" json:"UUID,omitempty"`
	LayersPath     string          `protobuf:"bytes,2,opt,name=LayersPath" json:"LayersPath,omitempty"`
	ProbeGraphPath string          `protobuf:"bytes,11,opt,name=ProbeGraphPath" json:"ProbeGraphPath,omitempty"`
	Statistics     *FlowStatistics `protobuf:"bytes,3,opt,name=Statistics" json:"Statistics,omitempty"`
	// useless, should be removed
	IfSrcGraphPath    string `protobuf:"bytes,14,opt,name=IfSrcGraphPath" json:"IfSrcGraphPath,omitempty"`
	IfDstGraphPath    string `protobuf:"bytes,19,opt,name=IfDstGraphPath" json:"IfDstGraphPath,omitempty"`
	DebugTimestamp    int64  `protobuf:"varint,90,opt,name=DebugTimestamp" json:"DebugTimestamp,omitempty"`
	DebugKeyNet       uint64 `protobuf:"varint,91,opt,name=DebugKeyNet" json:"DebugKeyNet,omitempty"`
	DebugKeyTransport uint64 `protobuf:"varint,92,opt,name=DebugKeyTransport" json:"DebugKeyTransport,omitempty"`
}

func (m *Flow) Reset()                    { *m = Flow{} }
func (m *Flow) String() string            { return proto.CompactTextString(m) }
func (*Flow) ProtoMessage()               {}
func (*Flow) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Flow) GetStatistics() *FlowStatistics {
	if m != nil {
		return m.Statistics
	}
	return nil
}

func init() {
	proto.RegisterType((*FlowEndpointStatistics)(nil), "flow.FlowEndpointStatistics")
	proto.RegisterType((*FlowEndpointsStatistics)(nil), "flow.FlowEndpointsStatistics")
	proto.RegisterType((*FlowStatistics)(nil), "flow.FlowStatistics")
	proto.RegisterType((*Flow)(nil), "flow.Flow")
	proto.RegisterEnum("flow.FlowEndpointType", FlowEndpointType_name, FlowEndpointType_value)
}

var fileDescriptor0 = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x25, 0x8d, 0xbb, 0xb5, 0x37, 0x23, 0x0b, 0x1e, 0x2a, 0x06, 0x81, 0x84, 0x02, 0x0f, 0x15,
	0x42, 0x45, 0x2a, 0x08, 0x21, 0xde, 0xd6, 0x25, 0x40, 0xc4, 0xb4, 0x45, 0xad, 0xbb, 0x07, 0xe0,
	0xc5, 0x2d, 0x2e, 0x8b, 0xb6, 0x26, 0x91, 0xed, 0x81, 0xf2, 0x1d, 0xfc, 0x0b, 0x1f, 0xc4, 0x97,
	0x60, 0x7b, 0x6d, 0x13, 0x54, 0x24, 0xb4, 0x97, 0x24, 0xf7, 0xdc, 0xe3, 0x73, 0x8f, 0x8f, 0x1d,
	0xd8, 0x5f, 0x5c, 0x16, 0x3f, 0x5e, 0x98, 0xc7, 0xa0, 0x14, 0x85, 0x2a, 0x30, 0x32, 0xdf, 0xe1,
	0x12, 0x7a, 0xef, 0xf4, 0x3b, 0xce, 0xbf, 0x96, 0x45, 0x96, 0xab, 0x89, 0x62, 0x2a, 0x93, 0x2a,
	0x9b, 0x4b, 0xfc, 0x14, 0x10, 0xad, 0x4a, 0x4e, 0x9c, 0xc7, 0x4e, 0xdf, 0x1f, 0xf6, 0x06, 0x76,
	0x69, 0x93, 0x6b, 0xba, 0xf8, 0x36, 0xb4, 0xcf, 0xd8, 0xe5, 0x15, 0x27, 0x2d, 0x4d, 0xeb, 0xe2,
	0x7d, 0xd8, 0x4d, 0xd9, 0xfc, 0x82, 0x2b, 0x49, 0xda, 0x1a, 0x40, 0xa6, 0x3f, 0xaa, 0x14, 0x97,
	0x64, 0xc7, 0x94, 0xe1, 0x4f, 0x07, 0xee, 0x35, 0x35, 0xe4, 0x8d, 0x07, 0xf6, 0xa1, 0x75, 0x38,
	0x22, 0xae, 0xe6, 0x78, 0xc3, 0x87, 0xdb, 0x9c, 0x86, 0x9e, 0x66, 0x8e, 0x0e, 0x09, 0xfa, 0x3f,
	0x33, 0xfc, 0xe5, 0x80, 0x6f, 0x5a, 0x8d, 0xc5, 0xda, 0xb7, 0xae, 0x84, 0xb2, 0x6e, 0x5c, 0xbc,
	0x07, 0xe8, 0x98, 0x49, 0x65, 0x77, 0xe9, 0xe2, 0xd7, 0xd0, 0xdd, 0x6c, 0x40, 0x5b, 0x71, 0xf5,
	0x80, 0x27, 0xf5, 0x80, 0x5a, 0x65, 0xb0, 0x61, 0xc5, 0xb9, 0x12, 0xd5, 0x83, 0x53, 0xf0, 0xff,
	0x46, 0xb0, 0x07, 0xee, 0x05, 0xaf, 0xec, 0x90, 0x36, 0x7e, 0x0e, 0xed, 0xef, 0x9b, 0x2c, 0xbd,
	0xe1, 0xa3, 0x6d, 0xcf, 0x8d, 0xb8, 0xde, 0xb6, 0xde, 0x38, 0xe1, 0x6f, 0x07, 0x90, 0xe9, 0x1b,
	0x7f, 0xd3, 0x69, 0x12, 0x59, 0xa1, 0x2e, 0xc6, 0x00, 0xc7, 0xac, 0xe2, 0x42, 0xa6, 0x4c, 0x9d,
	0xaf, 0x4e, 0xa6, 0x07, 0x7e, 0x2a, 0x8a, 0x19, 0x7f, 0x2f, 0x58, 0x79, 0x6e, 0x71, 0xcf, 0xe2,
	0x7d, 0x80, 0x5a, 0x74, 0x95, 0xeb, 0xdd, 0x7f, 0x6d, 0xc6, 0x28, 0x24, 0x8b, 0x89, 0x98, 0xd7,
	0x0a, 0xfe, 0x5a, 0x39, 0x59, 0x44, 0x52, 0xd5, 0xf8, 0xc1, 0x1a, 0x8f, 0xf8, 0xec, 0xea, 0x1b,
	0xcd, 0x96, 0x5c, 0x2a, 0xb6, 0x2c, 0xc9, 0x27, 0x9b, 0xde, 0x01, 0x78, 0x16, 0xff, 0xc8, 0xab,
	0x13, 0xae, 0xc8, 0x67, 0x7b, 0x4f, 0xee, 0xc3, 0x9d, 0x35, 0x48, 0x05, 0xcb, 0x65, 0x59, 0xe8,
	0xec, 0xbf, 0x98, 0xd6, 0xb3, 0x31, 0x04, 0x5b, 0xb7, 0x60, 0x0f, 0x3a, 0x31, 0xfd, 0x10, 0x8f,
	0x4f, 0x62, 0x1a, 0xdc, 0xc2, 0x1d, 0x40, 0x49, 0x7a, 0xf6, 0x2a, 0x70, 0x74, 0x9e, 0xbb, 0xf4,
	0x28, 0x4d, 0x4f, 0xc7, 0x34, 0x68, 0x99, 0x62, 0x1a, 0x5d, 0x17, 0xe6, 0x04, 0x3b, 0x93, 0x23,
	0x7a, 0x5d, 0xa1, 0xd9, 0x8e, 0xfd, 0x07, 0x5e, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x4c,
	0xab, 0x35, 0x16, 0x03, 0x00, 0x00,
}
