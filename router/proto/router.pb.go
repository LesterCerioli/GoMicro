// Code generated by protoc-gen-go. DO NOT EDIT.
// source: router.proto

package go_micro_router

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

// AdvertType defines the type of advert
type AdvertType int32

const (
	AdvertType_AdvertAnnounce AdvertType = 0
	AdvertType_AdvertUpdate   AdvertType = 1
)

var AdvertType_name = map[int32]string{
	0: "AdvertAnnounce",
	1: "AdvertUpdate",
}

var AdvertType_value = map[string]int32{
	"AdvertAnnounce": 0,
	"AdvertUpdate":   1,
}

func (x AdvertType) String() string {
	return proto.EnumName(AdvertType_name, int32(x))
}

func (AdvertType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{0}
}

// EventType defines the type of event
type EventType int32

const (
	EventType_Create EventType = 0
	EventType_Delete EventType = 1
	EventType_Update EventType = 2
)

var EventType_name = map[int32]string{
	0: "Create",
	1: "Delete",
	2: "Update",
}

var EventType_value = map[string]int32{
	"Create": 0,
	"Delete": 1,
	"Update": 2,
}

func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}

func (EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{1}
}

// Empty request
type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{0}
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

// Empty response
type Response struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{1}
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

// ListResponse is returned by List
type ListResponse struct {
	Routes               []*Route `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{2}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

// LookupRequest is made to Lookup
type LookupRequest struct {
	Query                *Query   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LookupRequest) Reset()         { *m = LookupRequest{} }
func (m *LookupRequest) String() string { return proto.CompactTextString(m) }
func (*LookupRequest) ProtoMessage()    {}
func (*LookupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{3}
}

func (m *LookupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LookupRequest.Unmarshal(m, b)
}
func (m *LookupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LookupRequest.Marshal(b, m, deterministic)
}
func (m *LookupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LookupRequest.Merge(m, src)
}
func (m *LookupRequest) XXX_Size() int {
	return xxx_messageInfo_LookupRequest.Size(m)
}
func (m *LookupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LookupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LookupRequest proto.InternalMessageInfo

func (m *LookupRequest) GetQuery() *Query {
	if m != nil {
		return m.Query
	}
	return nil
}

// LookupResponse is returned by Lookup
type LookupResponse struct {
	Routes               []*Route `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LookupResponse) Reset()         { *m = LookupResponse{} }
func (m *LookupResponse) String() string { return proto.CompactTextString(m) }
func (*LookupResponse) ProtoMessage()    {}
func (*LookupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{4}
}

func (m *LookupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LookupResponse.Unmarshal(m, b)
}
func (m *LookupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LookupResponse.Marshal(b, m, deterministic)
}
func (m *LookupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LookupResponse.Merge(m, src)
}
func (m *LookupResponse) XXX_Size() int {
	return xxx_messageInfo_LookupResponse.Size(m)
}
func (m *LookupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LookupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LookupResponse proto.InternalMessageInfo

func (m *LookupResponse) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

// QueryRequest queries Table for Routes
type QueryRequest struct {
	Query                *Query   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{5}
}

func (m *QueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryRequest.Unmarshal(m, b)
}
func (m *QueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryRequest.Marshal(b, m, deterministic)
}
func (m *QueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequest.Merge(m, src)
}
func (m *QueryRequest) XXX_Size() int {
	return xxx_messageInfo_QueryRequest.Size(m)
}
func (m *QueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequest proto.InternalMessageInfo

func (m *QueryRequest) GetQuery() *Query {
	if m != nil {
		return m.Query
	}
	return nil
}

// QueryResponse is returned by Query
type QueryResponse struct {
	Routes               []*Route `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryResponse) Reset()         { *m = QueryResponse{} }
func (m *QueryResponse) String() string { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()    {}
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{6}
}

func (m *QueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryResponse.Unmarshal(m, b)
}
func (m *QueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryResponse.Marshal(b, m, deterministic)
}
func (m *QueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResponse.Merge(m, src)
}
func (m *QueryResponse) XXX_Size() int {
	return xxx_messageInfo_QueryResponse.Size(m)
}
func (m *QueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResponse proto.InternalMessageInfo

func (m *QueryResponse) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

// WatchRequest is made to Watch Router
type WatchRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WatchRequest) Reset()         { *m = WatchRequest{} }
func (m *WatchRequest) String() string { return proto.CompactTextString(m) }
func (*WatchRequest) ProtoMessage()    {}
func (*WatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{7}
}

func (m *WatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WatchRequest.Unmarshal(m, b)
}
func (m *WatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WatchRequest.Marshal(b, m, deterministic)
}
func (m *WatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchRequest.Merge(m, src)
}
func (m *WatchRequest) XXX_Size() int {
	return xxx_messageInfo_WatchRequest.Size(m)
}
func (m *WatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WatchRequest proto.InternalMessageInfo

// Advert is router advertsement streamed by Watch
type Advert struct {
	// id of the advertising router
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// type of advertisement
	Type AdvertType `protobuf:"varint,2,opt,name=type,proto3,enum=go.micro.router.AdvertType" json:"type,omitempty"`
	// unix timestamp of the advertisement
	Timestamp int64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// TTL of the Advert
	Ttl int64 `protobuf:"varint,4,opt,name=ttl,proto3" json:"ttl,omitempty"`
	// events is a list of advertised events
	Events               []*Event `protobuf:"bytes,5,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Advert) Reset()         { *m = Advert{} }
func (m *Advert) String() string { return proto.CompactTextString(m) }
func (*Advert) ProtoMessage()    {}
func (*Advert) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{8}
}

func (m *Advert) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Advert.Unmarshal(m, b)
}
func (m *Advert) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Advert.Marshal(b, m, deterministic)
}
func (m *Advert) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Advert.Merge(m, src)
}
func (m *Advert) XXX_Size() int {
	return xxx_messageInfo_Advert.Size(m)
}
func (m *Advert) XXX_DiscardUnknown() {
	xxx_messageInfo_Advert.DiscardUnknown(m)
}

var xxx_messageInfo_Advert proto.InternalMessageInfo

func (m *Advert) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Advert) GetType() AdvertType {
	if m != nil {
		return m.Type
	}
	return AdvertType_AdvertAnnounce
}

func (m *Advert) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Advert) GetTtl() int64 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

func (m *Advert) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

// Solicit solicits routes
type Solicit struct {
	// id of the soliciting router
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Solicit) Reset()         { *m = Solicit{} }
func (m *Solicit) String() string { return proto.CompactTextString(m) }
func (*Solicit) ProtoMessage()    {}
func (*Solicit) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{9}
}

func (m *Solicit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Solicit.Unmarshal(m, b)
}
func (m *Solicit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Solicit.Marshal(b, m, deterministic)
}
func (m *Solicit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Solicit.Merge(m, src)
}
func (m *Solicit) XXX_Size() int {
	return xxx_messageInfo_Solicit.Size(m)
}
func (m *Solicit) XXX_DiscardUnknown() {
	xxx_messageInfo_Solicit.DiscardUnknown(m)
}

var xxx_messageInfo_Solicit proto.InternalMessageInfo

func (m *Solicit) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// ProcessResponse is returned by Process
type ProcessResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcessResponse) Reset()         { *m = ProcessResponse{} }
func (m *ProcessResponse) String() string { return proto.CompactTextString(m) }
func (*ProcessResponse) ProtoMessage()    {}
func (*ProcessResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{10}
}

func (m *ProcessResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessResponse.Unmarshal(m, b)
}
func (m *ProcessResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessResponse.Marshal(b, m, deterministic)
}
func (m *ProcessResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessResponse.Merge(m, src)
}
func (m *ProcessResponse) XXX_Size() int {
	return xxx_messageInfo_ProcessResponse.Size(m)
}
func (m *ProcessResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessResponse proto.InternalMessageInfo

// CreateResponse is returned by Create
type CreateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{11}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

// DeleteResponse is returned by Delete
type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{12}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

// UpdateResponse is returned by Update
type UpdateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{13}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

// Event is routing table event
type Event struct {
	// type of event
	Type EventType `protobuf:"varint,1,opt,name=type,proto3,enum=go.micro.router.EventType" json:"type,omitempty"`
	// unix timestamp of event
	Timestamp int64 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// service route
	Route                *Route   `protobuf:"bytes,3,opt,name=route,proto3" json:"route,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{14}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetType() EventType {
	if m != nil {
		return m.Type
	}
	return EventType_Create
}

func (m *Event) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Event) GetRoute() *Route {
	if m != nil {
		return m.Route
	}
	return nil
}

// Query is passed in a LookupRequest
type Query struct {
	// service to lookup
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	// gateway to lookup
	Gateway string `protobuf:"bytes,2,opt,name=gateway,proto3" json:"gateway,omitempty"`
	// network to lookup
	Network              string   `protobuf:"bytes,3,opt,name=network,proto3" json:"network,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{15}
}

func (m *Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query.Unmarshal(m, b)
}
func (m *Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query.Marshal(b, m, deterministic)
}
func (m *Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query.Merge(m, src)
}
func (m *Query) XXX_Size() int {
	return xxx_messageInfo_Query.Size(m)
}
func (m *Query) XXX_DiscardUnknown() {
	xxx_messageInfo_Query.DiscardUnknown(m)
}

var xxx_messageInfo_Query proto.InternalMessageInfo

func (m *Query) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Query) GetGateway() string {
	if m != nil {
		return m.Gateway
	}
	return ""
}

func (m *Query) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

// Route is a service route
type Route struct {
	// service for the route
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	// the address that advertise this route
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// gateway as the next hop
	Gateway string `protobuf:"bytes,3,opt,name=gateway,proto3" json:"gateway,omitempty"`
	// the network for this destination
	Network string `protobuf:"bytes,4,opt,name=network,proto3" json:"network,omitempty"`
	// router if the router id
	Router string `protobuf:"bytes,5,opt,name=router,proto3" json:"router,omitempty"`
	// the network link
	Link string `protobuf:"bytes,6,opt,name=link,proto3" json:"link,omitempty"`
	// the metric / score of this route
	Metric               int64    `protobuf:"varint,7,opt,name=metric,proto3" json:"metric,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{16}
}

func (m *Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route.Unmarshal(m, b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route.Marshal(b, m, deterministic)
}
func (m *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(m, src)
}
func (m *Route) XXX_Size() int {
	return xxx_messageInfo_Route.Size(m)
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Route) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Route) GetGateway() string {
	if m != nil {
		return m.Gateway
	}
	return ""
}

func (m *Route) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *Route) GetRouter() string {
	if m != nil {
		return m.Router
	}
	return ""
}

func (m *Route) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

func (m *Route) GetMetric() int64 {
	if m != nil {
		return m.Metric
	}
	return 0
}

type Status struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{17}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Status) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type StatusResponse struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusResponse) Reset()         { *m = StatusResponse{} }
func (m *StatusResponse) String() string { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()    {}
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{18}
}

func (m *StatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusResponse.Unmarshal(m, b)
}
func (m *StatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusResponse.Marshal(b, m, deterministic)
}
func (m *StatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusResponse.Merge(m, src)
}
func (m *StatusResponse) XXX_Size() int {
	return xxx_messageInfo_StatusResponse.Size(m)
}
func (m *StatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StatusResponse proto.InternalMessageInfo

func (m *StatusResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterEnum("go.micro.router.AdvertType", AdvertType_name, AdvertType_value)
	proto.RegisterEnum("go.micro.router.EventType", EventType_name, EventType_value)
	proto.RegisterType((*Request)(nil), "go.micro.router.Request")
	proto.RegisterType((*Response)(nil), "go.micro.router.Response")
	proto.RegisterType((*ListResponse)(nil), "go.micro.router.ListResponse")
	proto.RegisterType((*LookupRequest)(nil), "go.micro.router.LookupRequest")
	proto.RegisterType((*LookupResponse)(nil), "go.micro.router.LookupResponse")
	proto.RegisterType((*QueryRequest)(nil), "go.micro.router.QueryRequest")
	proto.RegisterType((*QueryResponse)(nil), "go.micro.router.QueryResponse")
	proto.RegisterType((*WatchRequest)(nil), "go.micro.router.WatchRequest")
	proto.RegisterType((*Advert)(nil), "go.micro.router.Advert")
	proto.RegisterType((*Solicit)(nil), "go.micro.router.Solicit")
	proto.RegisterType((*ProcessResponse)(nil), "go.micro.router.ProcessResponse")
	proto.RegisterType((*CreateResponse)(nil), "go.micro.router.CreateResponse")
	proto.RegisterType((*DeleteResponse)(nil), "go.micro.router.DeleteResponse")
	proto.RegisterType((*UpdateResponse)(nil), "go.micro.router.UpdateResponse")
	proto.RegisterType((*Event)(nil), "go.micro.router.Event")
	proto.RegisterType((*Query)(nil), "go.micro.router.Query")
	proto.RegisterType((*Route)(nil), "go.micro.router.Route")
	proto.RegisterType((*Status)(nil), "go.micro.router.Status")
	proto.RegisterType((*StatusResponse)(nil), "go.micro.router.StatusResponse")
}

func init() { proto.RegisterFile("router.proto", fileDescriptor_367072455c71aedc) }

var fileDescriptor_367072455c71aedc = []byte{
	// 716 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x4d, 0x4f, 0xdb, 0x4c,
	0x10, 0xb6, 0x93, 0xd8, 0x79, 0x33, 0x6f, 0x08, 0xe9, 0xa8, 0x02, 0x93, 0x16, 0x88, 0x7c, 0x42,
	0x08, 0x99, 0x2a, 0xbd, 0xf6, 0x83, 0x40, 0xa9, 0x2a, 0x95, 0x43, 0x6b, 0x40, 0x3d, 0x1b, 0x7b,
	0x45, 0x2d, 0x12, 0xaf, 0xd9, 0xdd, 0x80, 0x72, 0xee, 0x8f, 0xa9, 0x7a, 0xe9, 0xa5, 0x7f, 0xb0,
	0xf2, 0xee, 0x3a, 0x84, 0x38, 0x8b, 0x04, 0xea, 0x29, 0x3b, 0x5f, 0xcf, 0xcc, 0xec, 0x3e, 0x33,
	0x0e, 0xb4, 0x19, 0x9d, 0x08, 0xc2, 0x82, 0x9c, 0x51, 0x41, 0x71, 0xf5, 0x92, 0x06, 0xe3, 0x34,
	0x66, 0x34, 0x50, 0x6a, 0xbf, 0x05, 0xcd, 0x90, 0x5c, 0x4f, 0x08, 0x17, 0x3e, 0xc0, 0x7f, 0x21,
	0xe1, 0x39, 0xcd, 0x38, 0xf1, 0xdf, 0x41, 0xfb, 0x24, 0xe5, 0xa2, 0x94, 0x31, 0x00, 0x57, 0x06,
	0x70, 0xcf, 0xee, 0xd7, 0x77, 0xfe, 0x1f, 0xac, 0x05, 0x0b, 0x40, 0x41, 0x58, 0xfc, 0x84, 0xda,
	0xcb, 0x7f, 0x0b, 0x2b, 0x27, 0x94, 0x5e, 0x4d, 0x72, 0x0d, 0x8e, 0x7b, 0xe0, 0x5c, 0x4f, 0x08,
	0x9b, 0x7a, 0x76, 0xdf, 0x5e, 0x1a, 0xff, 0xb5, 0xb0, 0x86, 0xca, 0xc9, 0x3f, 0x80, 0x4e, 0x19,
	0xfe, 0xc4, 0x02, 0xde, 0x40, 0x5b, 0x21, 0x3e, 0x29, 0xff, 0x7b, 0x58, 0xd1, 0xd1, 0x4f, 0x4c,
	0xdf, 0x81, 0xf6, 0xb7, 0x48, 0xc4, 0xdf, 0xcb, 0xbb, 0xfd, 0x65, 0x83, 0x3b, 0x4c, 0x6e, 0x08,
	0x13, 0xd8, 0x81, 0x5a, 0x9a, 0xc8, 0x32, 0x5a, 0x61, 0x2d, 0x4d, 0x70, 0x1f, 0x1a, 0x62, 0x9a,
	0x13, 0xaf, 0xd6, 0xb7, 0x77, 0x3a, 0x83, 0x17, 0x15, 0x60, 0x15, 0x76, 0x36, 0xcd, 0x49, 0x28,
	0x1d, 0xf1, 0x25, 0xb4, 0x44, 0x3a, 0x26, 0x5c, 0x44, 0xe3, 0xdc, 0xab, 0xf7, 0xed, 0x9d, 0x7a,
	0x78, 0xa7, 0xc0, 0x2e, 0xd4, 0x85, 0x18, 0x79, 0x0d, 0xa9, 0x2f, 0x8e, 0x45, 0xed, 0xe4, 0x86,
	0x64, 0x82, 0x7b, 0x8e, 0xa1, 0xf6, 0xe3, 0xc2, 0x1c, 0x6a, 0x2f, 0x7f, 0x03, 0x9a, 0xa7, 0x74,
	0x94, 0xc6, 0x69, 0xa5, 0x56, 0xff, 0x19, 0xac, 0x7e, 0x61, 0x34, 0x26, 0x9c, 0xcf, 0x98, 0xd2,
	0x85, 0xce, 0x11, 0x23, 0x91, 0x20, 0xf3, 0x9a, 0x0f, 0x64, 0x44, 0xee, 0x6b, 0xce, 0xf3, 0x64,
	0xde, 0xe7, 0x87, 0x0d, 0x8e, 0xcc, 0x8a, 0x81, 0x6e, 0xdf, 0x96, 0xed, 0xf7, 0x96, 0xd7, 0x66,
	0xea, 0xbe, 0xb6, 0xd8, 0xfd, 0x1e, 0x38, 0x32, 0x4e, 0xde, 0x8b, 0xf9, 0x99, 0x94, 0x93, 0x7f,
	0x0e, 0x8e, 0x7c, 0x66, 0xf4, 0xa0, 0xc9, 0x09, 0xbb, 0x49, 0x63, 0xa2, 0x9b, 0x2d, 0xc5, 0xc2,
	0x72, 0x19, 0x09, 0x72, 0x1b, 0x4d, 0x65, 0xb2, 0x56, 0x58, 0x8a, 0x85, 0x25, 0x23, 0xe2, 0x96,
	0xb2, 0x2b, 0x99, 0xac, 0x15, 0x96, 0xa2, 0xff, 0xc7, 0x06, 0x47, 0xe6, 0x79, 0x18, 0x37, 0x4a,
	0x12, 0x46, 0x38, 0x2f, 0x71, 0xb5, 0x38, 0x9f, 0xb1, 0x6e, 0xcc, 0xd8, 0xb8, 0x97, 0x11, 0xd7,
	0x34, 0x3d, 0x99, 0xe7, 0x48, 0x83, 0x96, 0x10, 0xa1, 0x31, 0x4a, 0xb3, 0x2b, 0xcf, 0x95, 0x5a,
	0x79, 0x2e, 0x7c, 0xc7, 0x44, 0xb0, 0x34, 0xf6, 0x9a, 0xf2, 0xf6, 0xb4, 0xe4, 0x0f, 0xc0, 0x3d,
	0x15, 0x91, 0x98, 0xf0, 0x22, 0x2a, 0xa6, 0x49, 0x59, 0xb2, 0x3c, 0xe3, 0x73, 0x70, 0x08, 0x63,
	0x94, 0xe9, 0x6a, 0x95, 0xe0, 0x0f, 0xa1, 0xa3, 0x62, 0x66, 0x83, 0xb2, 0x0f, 0x2e, 0x97, 0x1a,
	0x3d, 0x68, 0xeb, 0x95, 0x17, 0xd0, 0x01, 0xda, 0x6d, 0x77, 0x00, 0x70, 0xc7, 0x70, 0x44, 0xe8,
	0x28, 0x69, 0x98, 0x65, 0x74, 0x92, 0xc5, 0xa4, 0x6b, 0x61, 0x17, 0xda, 0x4a, 0xa7, 0x38, 0xd4,
	0xb5, 0x77, 0xf7, 0xa1, 0x35, 0xa3, 0x05, 0x02, 0xb8, 0x8a, 0x80, 0x5d, 0xab, 0x38, 0x2b, 0xea,
	0x75, 0xed, 0xe2, 0xac, 0x03, 0x6a, 0x83, 0xdf, 0x75, 0x70, 0x43, 0x75, 0x25, 0x9f, 0xc1, 0x55,
	0xab, 0x05, 0xb7, 0x2a, 0xa5, 0xdd, 0x5b, 0x59, 0xbd, 0x6d, 0xa3, 0x5d, 0x93, 0xd8, 0xc2, 0x43,
	0x70, 0xe4, 0x98, 0xe3, 0x66, 0xc5, 0x77, 0x7e, 0xfc, 0x7b, 0x86, 0x91, 0xf3, 0xad, 0x57, 0x36,
	0x1e, 0x42, 0x4b, 0xb5, 0x97, 0x72, 0x82, 0x5e, 0x95, 0xb0, 0x1a, 0x62, 0xdd, 0xb0, 0x18, 0x24,
	0xc6, 0xc1, 0xdd, 0xc8, 0x9a, 0x11, 0x36, 0x96, 0x58, 0x66, 0x9d, 0x7c, 0x84, 0xa6, 0x9e, 0x6c,
	0x34, 0x65, 0xea, 0xf5, 0x2b, 0x86, 0xc5, 0x65, 0x60, 0xe1, 0xf1, 0x8c, 0x45, 0xe6, 0x42, 0xb6,
	0x4d, 0x9c, 0x98, 0xc1, 0x0c, 0x7e, 0xd6, 0xc1, 0x39, 0x8b, 0x2e, 0x46, 0x04, 0x8f, 0xca, 0xe7,
	0x45, 0xc3, 0x30, 0x2f, 0x81, 0x5b, 0x58, 0x48, 0x56, 0x01, 0xa2, 0x78, 0xf1, 0x08, 0x90, 0x85,
	0x1d, 0x26, 0x41, 0x14, 0xa1, 0x1e, 0x01, 0xb2, 0xb0, 0xf6, 0x2c, 0x1c, 0x42, 0xa3, 0xf8, 0xb0,
	0x3e, 0x70, 0x3b, 0x55, 0x2a, 0xcd, 0x7f, 0x89, 0x7d, 0x0b, 0x3f, 0x95, 0x5b, 0x6b, 0xd3, 0xf0,
	0x11, 0xd3, 0x40, 0x5b, 0x26, 0xf3, 0xbf, 0xa4, 0xef, 0x85, 0x2b, 0xff, 0x58, 0xbc, 0xfe, 0x1b,
	0x00, 0x00, 0xff, 0xff, 0xb6, 0xaf, 0xdf, 0xff, 0x68, 0x08, 0x00, 0x00,
}
