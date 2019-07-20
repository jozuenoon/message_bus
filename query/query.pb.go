// Code generated by protoc-gen-go. DO NOT EDIT.
// source: query.proto

package query

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type DetectionEvent struct {
	Time                 []*timestamp.Timestamp `protobuf:"bytes,1,rep,name=time,proto3" json:"time,omitempty"`
	DeviceId             string                 `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *DetectionEvent) Reset()         { *m = DetectionEvent{} }
func (m *DetectionEvent) String() string { return proto.CompactTextString(m) }
func (*DetectionEvent) ProtoMessage()    {}
func (*DetectionEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{0}
}

func (m *DetectionEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetectionEvent.Unmarshal(m, b)
}
func (m *DetectionEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetectionEvent.Marshal(b, m, deterministic)
}
func (m *DetectionEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetectionEvent.Merge(m, src)
}
func (m *DetectionEvent) XXX_Size() int {
	return xxx_messageInfo_DetectionEvent.Size(m)
}
func (m *DetectionEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_DetectionEvent.DiscardUnknown(m)
}

var xxx_messageInfo_DetectionEvent proto.InternalMessageInfo

func (m *DetectionEvent) GetTime() []*timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *DetectionEvent) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

type Range struct {
	After                *timestamp.Timestamp `protobuf:"bytes,1,opt,name=after,proto3" json:"after,omitempty"`
	Before               *timestamp.Timestamp `protobuf:"bytes,2,opt,name=before,proto3" json:"before,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Range) Reset()         { *m = Range{} }
func (m *Range) String() string { return proto.CompactTextString(m) }
func (*Range) ProtoMessage()    {}
func (*Range) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{1}
}

func (m *Range) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Range.Unmarshal(m, b)
}
func (m *Range) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Range.Marshal(b, m, deterministic)
}
func (m *Range) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Range.Merge(m, src)
}
func (m *Range) XXX_Size() int {
	return xxx_messageInfo_Range.Size(m)
}
func (m *Range) XXX_DiscardUnknown() {
	xxx_messageInfo_Range.DiscardUnknown(m)
}

var xxx_messageInfo_Range proto.InternalMessageInfo

func (m *Range) GetAfter() *timestamp.Timestamp {
	if m != nil {
		return m.After
	}
	return nil
}

func (m *Range) GetBefore() *timestamp.Timestamp {
	if m != nil {
		return m.Before
	}
	return nil
}

type GetEventsRequest struct {
	DetectorIds          []string `protobuf:"bytes,1,rep,name=detector_ids,json=detectorIds,proto3" json:"detector_ids,omitempty"`
	Range                *Range   `protobuf:"bytes,2,opt,name=range,proto3" json:"range,omitempty"`
	TravelsLimit         int64    `protobuf:"varint,3,opt,name=travels_limit,json=travelsLimit,proto3" json:"travels_limit,omitempty"`
	DeviceLimit          int64    `protobuf:"varint,4,opt,name=device_limit,json=deviceLimit,proto3" json:"device_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEventsRequest) Reset()         { *m = GetEventsRequest{} }
func (m *GetEventsRequest) String() string { return proto.CompactTextString(m) }
func (*GetEventsRequest) ProtoMessage()    {}
func (*GetEventsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{2}
}

func (m *GetEventsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventsRequest.Unmarshal(m, b)
}
func (m *GetEventsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventsRequest.Marshal(b, m, deterministic)
}
func (m *GetEventsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventsRequest.Merge(m, src)
}
func (m *GetEventsRequest) XXX_Size() int {
	return xxx_messageInfo_GetEventsRequest.Size(m)
}
func (m *GetEventsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventsRequest proto.InternalMessageInfo

func (m *GetEventsRequest) GetDetectorIds() []string {
	if m != nil {
		return m.DetectorIds
	}
	return nil
}

func (m *GetEventsRequest) GetRange() *Range {
	if m != nil {
		return m.Range
	}
	return nil
}

func (m *GetEventsRequest) GetTravelsLimit() int64 {
	if m != nil {
		return m.TravelsLimit
	}
	return 0
}

func (m *GetEventsRequest) GetDeviceLimit() int64 {
	if m != nil {
		return m.DeviceLimit
	}
	return 0
}

type GetEventsResponse struct {
	Events               []*DetectionEvent `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetEventsResponse) Reset()         { *m = GetEventsResponse{} }
func (m *GetEventsResponse) String() string { return proto.CompactTextString(m) }
func (*GetEventsResponse) ProtoMessage()    {}
func (*GetEventsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{3}
}

func (m *GetEventsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventsResponse.Unmarshal(m, b)
}
func (m *GetEventsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventsResponse.Marshal(b, m, deterministic)
}
func (m *GetEventsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventsResponse.Merge(m, src)
}
func (m *GetEventsResponse) XXX_Size() int {
	return xxx_messageInfo_GetEventsResponse.Size(m)
}
func (m *GetEventsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventsResponse proto.InternalMessageInfo

func (m *GetEventsResponse) GetEvents() []*DetectionEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

type GetBatteryVoltageRequest struct {
	DetectorIds          []string `protobuf:"bytes,1,rep,name=detector_ids,json=detectorIds,proto3" json:"detector_ids,omitempty"`
	Range                *Range   `protobuf:"bytes,2,opt,name=range,proto3" json:"range,omitempty"`
	Limit                int64    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBatteryVoltageRequest) Reset()         { *m = GetBatteryVoltageRequest{} }
func (m *GetBatteryVoltageRequest) String() string { return proto.CompactTextString(m) }
func (*GetBatteryVoltageRequest) ProtoMessage()    {}
func (*GetBatteryVoltageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{4}
}

func (m *GetBatteryVoltageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBatteryVoltageRequest.Unmarshal(m, b)
}
func (m *GetBatteryVoltageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBatteryVoltageRequest.Marshal(b, m, deterministic)
}
func (m *GetBatteryVoltageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBatteryVoltageRequest.Merge(m, src)
}
func (m *GetBatteryVoltageRequest) XXX_Size() int {
	return xxx_messageInfo_GetBatteryVoltageRequest.Size(m)
}
func (m *GetBatteryVoltageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBatteryVoltageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBatteryVoltageRequest proto.InternalMessageInfo

func (m *GetBatteryVoltageRequest) GetDetectorIds() []string {
	if m != nil {
		return m.DetectorIds
	}
	return nil
}

func (m *GetBatteryVoltageRequest) GetRange() *Range {
	if m != nil {
		return m.Range
	}
	return nil
}

func (m *GetBatteryVoltageRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GetBatteryVoltageResponse struct {
	Reports              []*VoltageReport `protobuf:"bytes,1,rep,name=reports,proto3" json:"reports,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetBatteryVoltageResponse) Reset()         { *m = GetBatteryVoltageResponse{} }
func (m *GetBatteryVoltageResponse) String() string { return proto.CompactTextString(m) }
func (*GetBatteryVoltageResponse) ProtoMessage()    {}
func (*GetBatteryVoltageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{5}
}

func (m *GetBatteryVoltageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBatteryVoltageResponse.Unmarshal(m, b)
}
func (m *GetBatteryVoltageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBatteryVoltageResponse.Marshal(b, m, deterministic)
}
func (m *GetBatteryVoltageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBatteryVoltageResponse.Merge(m, src)
}
func (m *GetBatteryVoltageResponse) XXX_Size() int {
	return xxx_messageInfo_GetBatteryVoltageResponse.Size(m)
}
func (m *GetBatteryVoltageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBatteryVoltageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBatteryVoltageResponse proto.InternalMessageInfo

func (m *GetBatteryVoltageResponse) GetReports() []*VoltageReport {
	if m != nil {
		return m.Reports
	}
	return nil
}

type VoltageReport struct {
	DetectorId           string    `protobuf:"bytes,1,opt,name=detector_id,json=detectorId,proto3" json:"detector_id,omitempty"`
	Voltage              []float64 `protobuf:"fixed64,2,rep,packed,name=voltage,proto3" json:"voltage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *VoltageReport) Reset()         { *m = VoltageReport{} }
func (m *VoltageReport) String() string { return proto.CompactTextString(m) }
func (*VoltageReport) ProtoMessage()    {}
func (*VoltageReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{6}
}

func (m *VoltageReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoltageReport.Unmarshal(m, b)
}
func (m *VoltageReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoltageReport.Marshal(b, m, deterministic)
}
func (m *VoltageReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoltageReport.Merge(m, src)
}
func (m *VoltageReport) XXX_Size() int {
	return xxx_messageInfo_VoltageReport.Size(m)
}
func (m *VoltageReport) XXX_DiscardUnknown() {
	xxx_messageInfo_VoltageReport.DiscardUnknown(m)
}

var xxx_messageInfo_VoltageReport proto.InternalMessageInfo

func (m *VoltageReport) GetDetectorId() string {
	if m != nil {
		return m.DetectorId
	}
	return ""
}

func (m *VoltageReport) GetVoltage() []float64 {
	if m != nil {
		return m.Voltage
	}
	return nil
}

func init() {
	proto.RegisterType((*DetectionEvent)(nil), "query.DetectionEvent")
	proto.RegisterType((*Range)(nil), "query.Range")
	proto.RegisterType((*GetEventsRequest)(nil), "query.GetEventsRequest")
	proto.RegisterType((*GetEventsResponse)(nil), "query.GetEventsResponse")
	proto.RegisterType((*GetBatteryVoltageRequest)(nil), "query.GetBatteryVoltageRequest")
	proto.RegisterType((*GetBatteryVoltageResponse)(nil), "query.GetBatteryVoltageResponse")
	proto.RegisterType((*VoltageReport)(nil), "query.VoltageReport")
}

func init() { proto.RegisterFile("query.proto", fileDescriptor_5c6ac9b241082464) }

var fileDescriptor_5c6ac9b241082464 = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xd5, 0x36, 0x75, 0x8a, 0xc7, 0x2e, 0x82, 0x55, 0x11, 0x4b, 0x38, 0xc4, 0x2c, 0x97, 0x5c,
	0x70, 0xab, 0x70, 0x47, 0xa8, 0x02, 0x55, 0x05, 0x2e, 0x6c, 0x51, 0x6f, 0x28, 0x72, 0xea, 0x49,
	0x64, 0x29, 0xce, 0xa6, 0xbb, 0x93, 0xa0, 0x7e, 0x0e, 0x7f, 0xc7, 0x67, 0x20, 0xef, 0xae, 0xd3,
	0x04, 0x02, 0xb9, 0xf4, 0x38, 0x6f, 0x9e, 0xdf, 0xbc, 0x37, 0xb3, 0x86, 0xe4, 0x76, 0x89, 0xe6,
	0x2e, 0x5f, 0x18, 0x4d, 0x9a, 0x47, 0xae, 0xe8, 0xf5, 0xa7, 0x5a, 0x4f, 0x67, 0x78, 0xea, 0xc0,
	0xf1, 0x72, 0x72, 0x4a, 0x55, 0x8d, 0x96, 0x8a, 0x7a, 0xe1, 0x79, 0xf2, 0x3b, 0x3c, 0xfe, 0x80,
	0x84, 0x37, 0x54, 0xe9, 0xf9, 0xc7, 0x15, 0xce, 0x89, 0xe7, 0x70, 0xd8, 0x90, 0x04, 0xcb, 0x3a,
	0x83, 0x64, 0xd8, 0xcb, 0xbd, 0x42, 0xde, 0x2a, 0xe4, 0xdf, 0x5a, 0x05, 0xe5, 0x78, 0xfc, 0x25,
	0xc4, 0x25, 0xae, 0xaa, 0x1b, 0x1c, 0x55, 0xa5, 0x38, 0xc8, 0xd8, 0x20, 0x56, 0x8f, 0x3c, 0x70,
	0x59, 0xca, 0x1a, 0x22, 0x55, 0xcc, 0xa7, 0xc8, 0xcf, 0x20, 0x2a, 0x26, 0x84, 0x46, 0xb0, 0x8c,
	0xed, 0x91, 0xf5, 0x44, 0x3e, 0x84, 0xee, 0x18, 0x27, 0xda, 0xa0, 0x13, 0xfd, 0xff, 0x27, 0x81,
	0x29, 0x7f, 0x32, 0x78, 0x72, 0x81, 0xe4, 0x82, 0x58, 0x85, 0xb7, 0x4b, 0xb4, 0xc4, 0x5f, 0x41,
	0x5a, 0xba, 0x88, 0xda, 0x8c, 0xaa, 0xd2, 0xba, 0x60, 0xb1, 0x4a, 0x5a, 0xec, 0xb2, 0xb4, 0x5c,
	0x42, 0x64, 0x1a, 0x9b, 0x61, 0x54, 0x9a, 0xfb, 0x55, 0x3a, 0xeb, 0xca, 0xb7, 0xf8, 0x6b, 0x38,
	0x26, 0x53, 0xac, 0x70, 0x66, 0x47, 0xb3, 0xaa, 0xae, 0x48, 0x74, 0x32, 0x36, 0xe8, 0xa8, 0x34,
	0x80, 0x5f, 0x1a, 0xcc, 0xcf, 0x72, 0xcb, 0xf0, 0x9c, 0x43, 0xc7, 0x49, 0x3c, 0xe6, 0x28, 0xf2,
	0x1c, 0x9e, 0x6e, 0x58, 0xb4, 0x0b, 0x3d, 0xb7, 0xc8, 0xdf, 0x40, 0x17, 0x1d, 0x12, 0xd6, 0xfe,
	0x2c, 0x38, 0xd8, 0xbe, 0x8d, 0x0a, 0x24, 0xf9, 0x03, 0xc4, 0x05, 0xd2, 0x79, 0x41, 0x84, 0xe6,
	0xee, 0x5a, 0xcf, 0xa8, 0x98, 0xe2, 0x03, 0xc7, 0x3d, 0x81, 0x68, 0x33, 0xa6, 0x2f, 0xe4, 0x67,
	0x78, 0xb1, 0x63, 0x70, 0x08, 0x91, 0xc3, 0x91, 0xc1, 0x85, 0x36, 0xeb, 0x14, 0x27, 0x41, 0x78,
	0x4d, 0x6c, 0x9a, 0xaa, 0x25, 0xc9, 0x4f, 0x70, 0xbc, 0xd5, 0xe1, 0x7d, 0x48, 0x36, 0xac, 0xbb,
	0xa7, 0x12, 0x2b, 0xb8, 0x77, 0xce, 0x05, 0x1c, 0xad, 0xfc, 0x17, 0xe2, 0x20, 0xeb, 0x0c, 0x98,
	0x6a, 0xcb, 0xe1, 0x2f, 0x06, 0xe9, 0xd7, 0x66, 0xd8, 0x15, 0x9a, 0x66, 0xd7, 0xfc, 0x1d, 0xc4,
	0xeb, 0x35, 0xf3, 0xe7, 0xc1, 0xc8, 0x9f, 0x6f, 0xa3, 0x27, 0xfe, 0x6e, 0x84, 0x30, 0xef, 0x21,
	0xbd, 0x22, 0x83, 0x45, 0xbd, 0x4f, 0x62, 0xf7, 0xa9, 0xce, 0x18, 0xbf, 0x76, 0x87, 0xde, 0xde,
	0x15, 0xef, 0xdf, 0xcb, 0xec, 0x3c, 0x5f, 0x2f, 0xfb, 0x37, 0xc1, 0x3b, 0x1b, 0x77, 0xdd, 0x0f,
	0xf0, 0xf6, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x6a, 0x82, 0xda, 0xf0, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryServiceClient is the client API for QueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryServiceClient interface {
	GetEvents(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (*GetEventsResponse, error)
	StreamEvents(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (QueryService_StreamEventsClient, error)
	GetBatteryVoltage(ctx context.Context, in *GetBatteryVoltageRequest, opts ...grpc.CallOption) (*GetBatteryVoltageResponse, error)
}

type queryServiceClient struct {
	cc *grpc.ClientConn
}

func NewQueryServiceClient(cc *grpc.ClientConn) QueryServiceClient {
	return &queryServiceClient{cc}
}

func (c *queryServiceClient) GetEvents(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (*GetEventsResponse, error) {
	out := new(GetEventsResponse)
	err := c.cc.Invoke(ctx, "/query.QueryService/GetEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) StreamEvents(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (QueryService_StreamEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_QueryService_serviceDesc.Streams[0], "/query.QueryService/StreamEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &queryServiceStreamEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type QueryService_StreamEventsClient interface {
	Recv() (*DetectionEvent, error)
	grpc.ClientStream
}

type queryServiceStreamEventsClient struct {
	grpc.ClientStream
}

func (x *queryServiceStreamEventsClient) Recv() (*DetectionEvent, error) {
	m := new(DetectionEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *queryServiceClient) GetBatteryVoltage(ctx context.Context, in *GetBatteryVoltageRequest, opts ...grpc.CallOption) (*GetBatteryVoltageResponse, error) {
	out := new(GetBatteryVoltageResponse)
	err := c.cc.Invoke(ctx, "/query.QueryService/GetBatteryVoltage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServiceServer is the server API for QueryService service.
type QueryServiceServer interface {
	GetEvents(context.Context, *GetEventsRequest) (*GetEventsResponse, error)
	StreamEvents(*GetEventsRequest, QueryService_StreamEventsServer) error
	GetBatteryVoltage(context.Context, *GetBatteryVoltageRequest) (*GetBatteryVoltageResponse, error)
}

// UnimplementedQueryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServiceServer struct {
}

func (*UnimplementedQueryServiceServer) GetEvents(ctx context.Context, req *GetEventsRequest) (*GetEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvents not implemented")
}
func (*UnimplementedQueryServiceServer) StreamEvents(req *GetEventsRequest, srv QueryService_StreamEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamEvents not implemented")
}
func (*UnimplementedQueryServiceServer) GetBatteryVoltage(ctx context.Context, req *GetBatteryVoltageRequest) (*GetBatteryVoltageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBatteryVoltage not implemented")
}

func RegisterQueryServiceServer(s *grpc.Server, srv QueryServiceServer) {
	s.RegisterService(&_QueryService_serviceDesc, srv)
}

func _QueryService_GetEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).GetEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/query.QueryService/GetEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).GetEvents(ctx, req.(*GetEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_StreamEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetEventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QueryServiceServer).StreamEvents(m, &queryServiceStreamEventsServer{stream})
}

type QueryService_StreamEventsServer interface {
	Send(*DetectionEvent) error
	grpc.ServerStream
}

type queryServiceStreamEventsServer struct {
	grpc.ServerStream
}

func (x *queryServiceStreamEventsServer) Send(m *DetectionEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _QueryService_GetBatteryVoltage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBatteryVoltageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).GetBatteryVoltage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/query.QueryService/GetBatteryVoltage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).GetBatteryVoltage(ctx, req.(*GetBatteryVoltageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QueryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "query.QueryService",
	HandlerType: (*QueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEvents",
			Handler:    _QueryService_GetEvents_Handler,
		},
		{
			MethodName: "GetBatteryVoltage",
			Handler:    _QueryService_GetBatteryVoltage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamEvents",
			Handler:       _QueryService_StreamEvents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "query.proto",
}
