// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculate.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type AddResponse_ServerType int32

const (
	AddResponse_INVALID AddResponse_ServerType = 0
	AddResponse_GOLANG  AddResponse_ServerType = 1
	AddResponse_PYTHON  AddResponse_ServerType = 2
	AddResponse_JAVA    AddResponse_ServerType = 3
	AddResponse_CPP     AddResponse_ServerType = 4
)

var AddResponse_ServerType_name = map[int32]string{
	0: "INVALID",
	1: "GOLANG",
	2: "PYTHON",
	3: "JAVA",
	4: "CPP",
}

var AddResponse_ServerType_value = map[string]int32{
	"INVALID": 0,
	"GOLANG":  1,
	"PYTHON":  2,
	"JAVA":    3,
	"CPP":     4,
}

func (x AddResponse_ServerType) String() string {
	return proto.EnumName(AddResponse_ServerType_name, int32(x))
}

func (AddResponse_ServerType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1d8b31321fcad8b1, []int{1, 0}
}

type AddRequest struct {
	A                    int32    `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    int32    `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d8b31321fcad8b1, []int{0}
}

func (m *AddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRequest.Unmarshal(m, b)
}
func (m *AddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRequest.Marshal(b, m, deterministic)
}
func (m *AddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRequest.Merge(m, src)
}
func (m *AddRequest) XXX_Size() int {
	return xxx_messageInfo_AddRequest.Size(m)
}
func (m *AddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRequest proto.InternalMessageInfo

func (m *AddRequest) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *AddRequest) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

type AddResponse struct {
	Result               int32                  `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	ServerType           AddResponse_ServerType `protobuf:"varint,2,opt,name=server_type,json=serverType,proto3,enum=proto.AddResponse_ServerType" json:"server_type,omitempty"`
	ServerId             *AddResponse_ServerId  `protobuf:"bytes,3,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *AddResponse) Reset()         { *m = AddResponse{} }
func (m *AddResponse) String() string { return proto.CompactTextString(m) }
func (*AddResponse) ProtoMessage()    {}
func (*AddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d8b31321fcad8b1, []int{1}
}

func (m *AddResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddResponse.Unmarshal(m, b)
}
func (m *AddResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddResponse.Marshal(b, m, deterministic)
}
func (m *AddResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddResponse.Merge(m, src)
}
func (m *AddResponse) XXX_Size() int {
	return xxx_messageInfo_AddResponse.Size(m)
}
func (m *AddResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddResponse proto.InternalMessageInfo

func (m *AddResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func (m *AddResponse) GetServerType() AddResponse_ServerType {
	if m != nil {
		return m.ServerType
	}
	return AddResponse_INVALID
}

func (m *AddResponse) GetServerId() *AddResponse_ServerId {
	if m != nil {
		return m.ServerId
	}
	return nil
}

type AddResponse_ServerId struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddResponse_ServerId) Reset()         { *m = AddResponse_ServerId{} }
func (m *AddResponse_ServerId) String() string { return proto.CompactTextString(m) }
func (*AddResponse_ServerId) ProtoMessage()    {}
func (*AddResponse_ServerId) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d8b31321fcad8b1, []int{1, 0}
}

func (m *AddResponse_ServerId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddResponse_ServerId.Unmarshal(m, b)
}
func (m *AddResponse_ServerId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddResponse_ServerId.Marshal(b, m, deterministic)
}
func (m *AddResponse_ServerId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddResponse_ServerId.Merge(m, src)
}
func (m *AddResponse_ServerId) XXX_Size() int {
	return xxx_messageInfo_AddResponse_ServerId.Size(m)
}
func (m *AddResponse_ServerId) XXX_DiscardUnknown() {
	xxx_messageInfo_AddResponse_ServerId.DiscardUnknown(m)
}

var xxx_messageInfo_AddResponse_ServerId proto.InternalMessageInfo

func (m *AddResponse_ServerId) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterEnum("proto.AddResponse_ServerType", AddResponse_ServerType_name, AddResponse_ServerType_value)
	proto.RegisterType((*AddRequest)(nil), "proto.AddRequest")
	proto.RegisterType((*AddResponse)(nil), "proto.AddResponse")
	proto.RegisterType((*AddResponse_ServerId)(nil), "proto.AddResponse.ServerId")
}

func init() { proto.RegisterFile("calculate.proto", fileDescriptor_1d8b31321fcad8b1) }

var fileDescriptor_1d8b31321fcad8b1 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xc1, 0x4b, 0xf3, 0x40,
	0x10, 0xc5, 0xbf, 0x4d, 0xda, 0x34, 0x9d, 0x7c, 0xd6, 0x38, 0x07, 0x09, 0x11, 0xa1, 0xe4, 0x94,
	0x83, 0xe4, 0x10, 0x2f, 0x7a, 0x11, 0x96, 0x16, 0x6a, 0xa4, 0xa4, 0x61, 0x2d, 0x05, 0x4f, 0x92,
	0x74, 0xf7, 0x10, 0x08, 0x26, 0x66, 0x37, 0x42, 0x8f, 0xfe, 0xe7, 0xd2, 0x6d, 0xaa, 0x05, 0xf1,
	0xb4, 0xef, 0x0d, 0xf3, 0xdb, 0x79, 0x3c, 0x38, 0xdf, 0xe6, 0xd5, 0xb6, 0xab, 0x72, 0x25, 0xa2,
	0xa6, 0xad, 0x55, 0x8d, 0x43, 0xfd, 0x04, 0x21, 0x00, 0xe5, 0x9c, 0x89, 0xf7, 0x4e, 0x48, 0x85,
	0xff, 0x81, 0xe4, 0x1e, 0x99, 0x92, 0x70, 0xc8, 0x48, 0xbe, 0x77, 0x85, 0x67, 0x1c, 0x5c, 0x11,
	0x7c, 0x1a, 0xe0, 0xe8, 0x55, 0xd9, 0xd4, 0x6f, 0x52, 0xe0, 0x25, 0x58, 0xad, 0x90, 0x5d, 0xa5,
	0x7a, 0xa0, 0x77, 0xf8, 0x00, 0x8e, 0x14, 0xed, 0x87, 0x68, 0x5f, 0xd5, 0xae, 0x11, 0x9a, 0x9f,
	0xc4, 0xd7, 0x87, 0xab, 0xd1, 0xc9, 0x07, 0xd1, 0xb3, 0xde, 0x5a, 0xef, 0x1a, 0xc1, 0x40, 0x7e,
	0x6b, 0xbc, 0x83, 0x71, 0xcf, 0x97, 0xdc, 0x33, 0xa7, 0x24, 0x74, 0xe2, 0xab, 0x3f, 0xe9, 0x84,
	0x33, 0x5b, 0xf6, 0xca, 0xf7, 0xc1, 0x3e, 0x4e, 0x71, 0x02, 0x46, 0xc9, 0x75, 0xb2, 0x33, 0x66,
	0x94, 0x3c, 0x98, 0x03, 0xfc, 0xdc, 0x43, 0x07, 0x46, 0x49, 0xba, 0xa1, 0xcb, 0x64, 0xee, 0xfe,
	0x43, 0x00, 0x6b, 0xb1, 0x5a, 0xd2, 0x74, 0xe1, 0x92, 0xbd, 0xce, 0x5e, 0xd6, 0x8f, 0xab, 0xd4,
	0x35, 0xd0, 0x86, 0xc1, 0x13, 0xdd, 0x50, 0xd7, 0xc4, 0x11, 0x98, 0xb3, 0x2c, 0x73, 0x07, 0xf1,
	0x3d, 0x8c, 0x67, 0xc7, 0x1e, 0xf1, 0x06, 0x4c, 0xca, 0x39, 0x5e, 0x9c, 0x86, 0xd3, 0x35, 0xfa,
	0xf8, 0x3b, 0x6f, 0x61, 0xe9, 0xd1, 0xed, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x34, 0x58,
	0x23, 0x89, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculateClient is the client API for Calculate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculateClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
}

type calculateClient struct {
	cc *grpc.ClientConn
}

func NewCalculateClient(cc *grpc.ClientConn) CalculateClient {
	return &calculateClient{cc}
}

func (c *calculateClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/proto.Calculate/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculateServer is the server API for Calculate service.
type CalculateServer interface {
	Add(context.Context, *AddRequest) (*AddResponse, error)
}

// UnimplementedCalculateServer can be embedded to have forward compatible implementations.
type UnimplementedCalculateServer struct {
}

func (*UnimplementedCalculateServer) Add(ctx context.Context, req *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}

func RegisterCalculateServer(s *grpc.Server, srv CalculateServer) {
	s.RegisterService(&_Calculate_serviceDesc, srv)
}

func _Calculate_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculateServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Calculate/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculateServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calculate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Calculate",
	HandlerType: (*CalculateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Calculate_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calculate.proto",
}