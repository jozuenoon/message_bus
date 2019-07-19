// Code generated by protoc-gen-go. DO NOT EDIT.
// source: query.proto

package queryserver

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
	proto.RegisterType((*DetectionEvent)(nil), "queryserver.DetectionEvent")
	proto.RegisterType((*Range)(nil), "queryserver.Range")
	proto.RegisterType((*GetEventsRequest)(nil), "queryserver.GetEventsRequest")
	proto.RegisterType((*GetEventsResponse)(nil), "queryserver.GetEventsResponse")
	proto.RegisterType((*GetBatteryVoltageRequest)(nil), "queryserver.GetBatteryVoltageRequest")
	proto.RegisterType((*GetBatteryVoltageResponse)(nil), "queryserver.GetBatteryVoltageResponse")
	proto.RegisterType((*VoltageReport)(nil), "queryserver.VoltageReport")
}

func init() { proto.RegisterFile("query.proto", fileDescriptor_5c6ac9b241082464) }

var fileDescriptor_5c6ac9b241082464 = []byte{
	// 458 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xb5, 0x49, 0x93, 0xe2, 0x71, 0xca, 0x9f, 0x11, 0x07, 0xe3, 0x0a, 0x6a, 0x8c, 0x40,
	0x3e, 0xb9, 0x55, 0xca, 0x13, 0x20, 0x50, 0x69, 0x85, 0x84, 0xba, 0x45, 0xdc, 0x50, 0xe4, 0xd4,
	0x93, 0xc8, 0x52, 0x9c, 0x4d, 0x77, 0x27, 0x96, 0x7a, 0xe0, 0x19, 0x78, 0x0c, 0x5e, 0x13, 0x79,
	0xd7, 0x0e, 0x71, 0x55, 0x9a, 0x13, 0xc7, 0xfd, 0xf6, 0xe7, 0x6f, 0xe6, 0x9b, 0x1d, 0x83, 0x7f,
	0xb3, 0x26, 0x7d, 0x9b, 0xae, 0xb4, 0x62, 0x85, 0xee, 0x60, 0x48, 0x57, 0xa4, 0xc3, 0xa3, 0xb9,
	0x52, 0xf3, 0x05, 0x1d, 0xdb, 0xab, 0xe9, 0x7a, 0x76, 0xcc, 0x45, 0x49, 0x86, 0xb3, 0x72, 0xe5,
	0xe8, 0xf8, 0x07, 0x3c, 0xfe, 0x48, 0x4c, 0xd7, 0x5c, 0xa8, 0xe5, 0xa7, 0x8a, 0x96, 0x8c, 0x29,
	0xec, 0xd5, 0x50, 0x20, 0xa2, 0x7e, 0xe2, 0x8f, 0xc3, 0xd4, 0x39, 0xa4, 0xad, 0x43, 0xfa, 0xad,
	0x75, 0x90, 0x96, 0xc3, 0x43, 0xf0, 0x72, 0xaa, 0x8a, 0x6b, 0x9a, 0x14, 0x79, 0xd0, 0x8b, 0x44,
	0xe2, 0xc9, 0x47, 0x4e, 0x38, 0xcf, 0xe3, 0x12, 0x06, 0x32, 0x5b, 0xce, 0x09, 0x4f, 0x60, 0x90,
	0xcd, 0x98, 0x74, 0x20, 0x22, 0xb1, 0xc3, 0xd6, 0x81, 0x38, 0x86, 0xe1, 0x94, 0x66, 0x4a, 0x93,
	0x35, 0x7d, 0xf8, 0x93, 0x86, 0x8c, 0x7f, 0x0b, 0x78, 0x7a, 0x46, 0x6c, 0x83, 0x18, 0x49, 0x37,
	0x6b, 0x32, 0x8c, 0xaf, 0x61, 0x94, 0xdb, 0x88, 0x4a, 0x4f, 0x8a, 0xdc, 0xd8, 0x60, 0x9e, 0xf4,
	0x5b, 0xed, 0x3c, 0x37, 0x98, 0xc0, 0x40, 0xd7, 0x6d, 0x36, 0xa5, 0x30, 0xdd, 0x9a, 0x61, 0x6a,
	0x03, 0x48, 0x07, 0xe0, 0x1b, 0x38, 0x60, 0x9d, 0x55, 0xb4, 0x30, 0x93, 0x45, 0x51, 0x16, 0x1c,
	0xf4, 0x23, 0x91, 0xf4, 0xe5, 0xa8, 0x11, 0xbf, 0xd4, 0x9a, 0xab, 0x68, 0x47, 0xe2, 0x98, 0x3d,
	0xcb, 0xf8, 0x4e, 0xb3, 0x48, 0xfc, 0x19, 0x9e, 0x6d, 0x35, 0x6a, 0x56, 0x6a, 0x69, 0x08, 0x4f,
	0x61, 0x48, 0x56, 0x69, 0x86, 0x7f, 0xd8, 0xe9, 0xa3, 0xfb, 0x4e, 0xb2, 0x41, 0xe3, 0x9f, 0x10,
	0x9c, 0x11, 0x7f, 0xc8, 0x98, 0x49, 0xdf, 0x7e, 0x57, 0x0b, 0xce, 0xe6, 0xf4, 0x5f, 0xa2, 0x3f,
	0x87, 0xc1, 0x76, 0x64, 0x77, 0x88, 0x2f, 0xe1, 0xc5, 0x3d, 0xe5, 0x9b, 0x40, 0xef, 0x61, 0x5f,
	0xd3, 0x4a, 0xe9, 0x4d, 0xa2, 0xb0, 0x63, 0xbf, 0xc1, 0x6b, 0x44, 0xb6, 0x68, 0x7c, 0x01, 0x07,
	0x9d, 0x1b, 0x3c, 0x02, 0x7f, 0x2b, 0x86, 0x5d, 0x21, 0x4f, 0xc2, 0xdf, 0x14, 0x18, 0xc0, 0x7e,
	0xe5, 0xbe, 0x08, 0x7a, 0x51, 0x3f, 0x11, 0xb2, 0x3d, 0x8e, 0x7f, 0xf5, 0x60, 0x74, 0x59, 0x97,
	0xbc, 0x22, 0x5d, 0x4f, 0x1f, 0x2f, 0xc0, 0xdb, 0x0c, 0x1e, 0x5f, 0x76, 0xda, 0xb9, 0xbb, 0x39,
	0xe1, 0xab, 0x7f, 0x5d, 0x37, 0xf1, 0xbe, 0xc2, 0x93, 0x8d, 0x78, 0xc5, 0x9a, 0xb2, 0x72, 0x97,
	0xe3, 0x43, 0x2f, 0x7a, 0x22, 0x70, 0x6a, 0xb7, 0xa2, 0x3b, 0x4c, 0x7c, 0x7b, 0xd7, 0xf2, 0xde,
	0xb7, 0x0e, 0xdf, 0xed, 0xc2, 0x5c, 0xd3, 0xd3, 0xa1, 0xfd, 0x7f, 0x4e, 0xff, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x45, 0x1f, 0x51, 0x1f, 0x35, 0x04, 0x00, 0x00,
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
	GetEventsStream(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (QueryService_GetEventsStreamClient, error)
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
	err := c.cc.Invoke(ctx, "/queryserver.QueryService/GetEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) GetEventsStream(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (QueryService_GetEventsStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_QueryService_serviceDesc.Streams[0], "/queryserver.QueryService/GetEventsStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &queryServiceGetEventsStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type QueryService_GetEventsStreamClient interface {
	Recv() (*DetectionEvent, error)
	grpc.ClientStream
}

type queryServiceGetEventsStreamClient struct {
	grpc.ClientStream
}

func (x *queryServiceGetEventsStreamClient) Recv() (*DetectionEvent, error) {
	m := new(DetectionEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *queryServiceClient) GetBatteryVoltage(ctx context.Context, in *GetBatteryVoltageRequest, opts ...grpc.CallOption) (*GetBatteryVoltageResponse, error) {
	out := new(GetBatteryVoltageResponse)
	err := c.cc.Invoke(ctx, "/queryserver.QueryService/GetBatteryVoltage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServiceServer is the server API for QueryService service.
type QueryServiceServer interface {
	GetEvents(context.Context, *GetEventsRequest) (*GetEventsResponse, error)
	GetEventsStream(*GetEventsRequest, QueryService_GetEventsStreamServer) error
	GetBatteryVoltage(context.Context, *GetBatteryVoltageRequest) (*GetBatteryVoltageResponse, error)
}

// UnimplementedQueryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServiceServer struct {
}

func (*UnimplementedQueryServiceServer) GetEvents(ctx context.Context, req *GetEventsRequest) (*GetEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvents not implemented")
}
func (*UnimplementedQueryServiceServer) GetEventsStream(req *GetEventsRequest, srv QueryService_GetEventsStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetEventsStream not implemented")
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
		FullMethod: "/queryserver.QueryService/GetEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).GetEvents(ctx, req.(*GetEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_GetEventsStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetEventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QueryServiceServer).GetEventsStream(m, &queryServiceGetEventsStreamServer{stream})
}

type QueryService_GetEventsStreamServer interface {
	Send(*DetectionEvent) error
	grpc.ServerStream
}

type queryServiceGetEventsStreamServer struct {
	grpc.ServerStream
}

func (x *queryServiceGetEventsStreamServer) Send(m *DetectionEvent) error {
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
		FullMethod: "/queryserver.QueryService/GetBatteryVoltage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).GetBatteryVoltage(ctx, req.(*GetBatteryVoltageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QueryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "queryserver.QueryService",
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
			StreamName:    "GetEventsStream",
			Handler:       _QueryService_GetEventsStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "query.proto",
}
