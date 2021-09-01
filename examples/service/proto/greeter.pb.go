// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/chinahtl/go-micro/examples/v3/service/proto/greeter.proto

package greeter

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Request struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_f752d72c38fc9ff3, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Greeting             string   `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_f752d72c38fc9ff3, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterType((*Response)(nil), "Response")
}

func init() {
	proto.RegisterFile("github.com/chinahtl/go-micro/examples/v3/service/proto/greeter.proto", fileDescriptor_f752d72c38fc9ff3)
}

var fileDescriptor_f752d72c38fc9ff3 = []byte{
	// 159 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8d, 0x41, 0x0a, 0xc2, 0x30,
	0x14, 0x44, 0x2d, 0xa8, 0x6d, 0xff, 0x32, 0x2b, 0x29, 0x28, 0xd2, 0x85, 0xe8, 0x26, 0x01, 0xc5,
	0x33, 0xe8, 0xba, 0x37, 0x68, 0xcb, 0x50, 0x03, 0x4d, 0x7e, 0x4c, 0x52, 0xf1, 0xf8, 0x42, 0xac,
	0xee, 0xe6, 0xc1, 0xcc, 0x1b, 0xba, 0x0e, 0x3a, 0x3e, 0xa6, 0x4e, 0xf6, 0x6c, 0x94, 0xd1, 0xbd,
	0x67, 0x85, 0x77, 0x6b, 0xdc, 0x88, 0xa0, 0x02, 0xfc, 0x4b, 0xf7, 0x50, 0xce, 0x73, 0x64, 0x35,
	0x78, 0x20, 0xc2, 0xcb, 0x44, 0xf5, 0x96, 0xf2, 0x06, 0xcf, 0x09, 0x21, 0x0a, 0x41, 0x4b, 0xdb,
	0x1a, 0x6c, 0xb2, 0x7d, 0x76, 0x2c, 0x9b, 0x94, 0xeb, 0x03, 0x15, 0x0d, 0x82, 0x63, 0x1b, 0x20,
	0x2a, 0x2a, 0xd2, 0x56, 0xdb, 0x61, 0xee, 0xfc, 0xf9, 0x7c, 0xa2, 0xfc, 0xf6, 0xf5, 0x8a, 0x1d,
	0xad, 0xee, 0x18, 0x47, 0x16, 0x85, 0x9c, 0xcd, 0x55, 0x29, 0x7f, 0x92, 0x7a, 0xd1, 0xad, 0xd3,
	0xf1, 0xe5, 0x13, 0x00, 0x00, 0xff, 0xff, 0x92, 0xab, 0x44, 0xf7, 0xb1, 0x00, 0x00, 0x00,
}
