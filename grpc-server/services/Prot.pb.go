// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Prot.proto

package services

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ProdAres int32

const (
	ProdAres_A ProdAres = 0
	ProdAres_B ProdAres = 1
	ProdAres_C ProdAres = 2
)

var ProdAres_name = map[int32]string{
	0: "A",
	1: "B",
	2: "C",
}

var ProdAres_value = map[string]int32{
	"A": 0,
	"B": 1,
	"C": 2,
}

func (x ProdAres) String() string {
	return proto.EnumName(ProdAres_name, int32(x))
}

func (ProdAres) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a611e38168e61a91, []int{0}
}

type ProdRequest struct {
	ProdId               int32    `protobuf:"varint,1,opt,name=prod_id,json=prodId,proto3" json:"prod_id,omitempty"`
	ProdArea             ProdAres `protobuf:"varint,2,opt,name=prod_area,json=prodArea,proto3,enum=services.ProdAres" json:"prod_area,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProdRequest) Reset()         { *m = ProdRequest{} }
func (m *ProdRequest) String() string { return proto.CompactTextString(m) }
func (*ProdRequest) ProtoMessage()    {}
func (*ProdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a611e38168e61a91, []int{0}
}

func (m *ProdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProdRequest.Unmarshal(m, b)
}
func (m *ProdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProdRequest.Marshal(b, m, deterministic)
}
func (m *ProdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProdRequest.Merge(m, src)
}
func (m *ProdRequest) XXX_Size() int {
	return xxx_messageInfo_ProdRequest.Size(m)
}
func (m *ProdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProdRequest proto.InternalMessageInfo

func (m *ProdRequest) GetProdId() int32 {
	if m != nil {
		return m.ProdId
	}
	return 0
}

func (m *ProdRequest) GetProdArea() ProdAres {
	if m != nil {
		return m.ProdArea
	}
	return ProdAres_A
}

type ProdResponse struct {
	ProdStock            int32    `protobuf:"varint,1,opt,name=prod_stock,json=prodStock,proto3" json:"prod_stock,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProdResponse) Reset()         { *m = ProdResponse{} }
func (m *ProdResponse) String() string { return proto.CompactTextString(m) }
func (*ProdResponse) ProtoMessage()    {}
func (*ProdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a611e38168e61a91, []int{1}
}

func (m *ProdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProdResponse.Unmarshal(m, b)
}
func (m *ProdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProdResponse.Marshal(b, m, deterministic)
}
func (m *ProdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProdResponse.Merge(m, src)
}
func (m *ProdResponse) XXX_Size() int {
	return xxx_messageInfo_ProdResponse.Size(m)
}
func (m *ProdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProdResponse proto.InternalMessageInfo

func (m *ProdResponse) GetProdStock() int32 {
	if m != nil {
		return m.ProdStock
	}
	return 0
}

type QuerySize struct {
	Size                 int32    `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuerySize) Reset()         { *m = QuerySize{} }
func (m *QuerySize) String() string { return proto.CompactTextString(m) }
func (*QuerySize) ProtoMessage()    {}
func (*QuerySize) Descriptor() ([]byte, []int) {
	return fileDescriptor_a611e38168e61a91, []int{2}
}

func (m *QuerySize) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuerySize.Unmarshal(m, b)
}
func (m *QuerySize) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuerySize.Marshal(b, m, deterministic)
}
func (m *QuerySize) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySize.Merge(m, src)
}
func (m *QuerySize) XXX_Size() int {
	return xxx_messageInfo_QuerySize.Size(m)
}
func (m *QuerySize) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySize.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySize proto.InternalMessageInfo

func (m *QuerySize) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type ProdResponseList struct {
	Prodres              []*ProdResponse `protobuf:"bytes,1,rep,name=prodres,proto3" json:"prodres,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ProdResponseList) Reset()         { *m = ProdResponseList{} }
func (m *ProdResponseList) String() string { return proto.CompactTextString(m) }
func (*ProdResponseList) ProtoMessage()    {}
func (*ProdResponseList) Descriptor() ([]byte, []int) {
	return fileDescriptor_a611e38168e61a91, []int{3}
}

func (m *ProdResponseList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProdResponseList.Unmarshal(m, b)
}
func (m *ProdResponseList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProdResponseList.Marshal(b, m, deterministic)
}
func (m *ProdResponseList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProdResponseList.Merge(m, src)
}
func (m *ProdResponseList) XXX_Size() int {
	return xxx_messageInfo_ProdResponseList.Size(m)
}
func (m *ProdResponseList) XXX_DiscardUnknown() {
	xxx_messageInfo_ProdResponseList.DiscardUnknown(m)
}

var xxx_messageInfo_ProdResponseList proto.InternalMessageInfo

func (m *ProdResponseList) GetProdres() []*ProdResponse {
	if m != nil {
		return m.Prodres
	}
	return nil
}

func init() {
	proto.RegisterEnum("services.ProdAres", ProdAres_name, ProdAres_value)
	proto.RegisterType((*ProdRequest)(nil), "services.ProdRequest")
	proto.RegisterType((*ProdResponse)(nil), "services.ProdResponse")
	proto.RegisterType((*QuerySize)(nil), "services.QuerySize")
	proto.RegisterType((*ProdResponseList)(nil), "services.ProdResponseList")
}

func init() { proto.RegisterFile("Prot.proto", fileDescriptor_a611e38168e61a91) }

var fileDescriptor_a611e38168e61a91 = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x9b, 0x6a, 0x6b, 0x3b, 0xa9, 0x12, 0xa6, 0x54, 0x4b, 0x54, 0x5a, 0x72, 0x2a, 0x82,
	0x8d, 0xd6, 0xa3, 0xa7, 0x56, 0x41, 0x0a, 0x0a, 0x35, 0x3d, 0x08, 0x5e, 0x24, 0x36, 0x63, 0x09,
	0x96, 0x6c, 0xdc, 0xdd, 0x16, 0xac, 0x78, 0xf1, 0x15, 0x7c, 0x34, 0x5f, 0xc1, 0x93, 0x4f, 0x21,
	0xd9, 0x4d, 0xb4, 0x15, 0x7b, 0xda, 0x99, 0xd9, 0x9f, 0x6f, 0x7e, 0x7e, 0x06, 0x60, 0xc0, 0x99,
	0x6c, 0xc7, 0x9c, 0x49, 0x86, 0x25, 0x41, 0x7c, 0x16, 0x8e, 0x48, 0xd8, 0x7b, 0x63, 0xc6, 0xc6,
	0x13, 0x72, 0xfd, 0x38, 0x74, 0xfd, 0x28, 0x62, 0xd2, 0x97, 0x21, 0x8b, 0x84, 0xd6, 0xd9, 0xe6,
	0x15, 0x0b, 0x68, 0xa2, 0x1b, 0xe7, 0x06, 0xcc, 0x01, 0x67, 0x81, 0x47, 0x4f, 0x53, 0x12, 0x12,
	0x77, 0x60, 0x23, 0xe6, 0x2c, 0xb8, 0x0b, 0x83, 0xba, 0xd1, 0x34, 0x5a, 0x05, 0xaf, 0x98, 0xb4,
	0xfd, 0x00, 0x5d, 0x28, 0xab, 0x0f, 0x9f, 0x93, 0x5f, 0xcf, 0x37, 0x8d, 0xd6, 0x56, 0x07, 0xdb,
	0xd9, 0xc2, 0x76, 0x82, 0xe8, 0x72, 0x12, 0x5e, 0x29, 0xd6, 0x95, 0xef, 0x1c, 0x42, 0x45, 0x83,
	0x45, 0xcc, 0x22, 0x41, 0xb8, 0x0f, 0xa0, 0x00, 0x42, 0xb2, 0xd1, 0x63, 0x0a, 0x57, 0xc8, 0x61,
	0x32, 0x70, 0x1a, 0x50, 0xbe, 0x9e, 0x12, 0x7f, 0x1e, 0x86, 0x73, 0x42, 0x84, 0x75, 0x11, 0xce,
	0x29, 0x55, 0xa9, 0xda, 0x39, 0x07, 0x6b, 0x91, 0x77, 0x19, 0x0a, 0x89, 0x47, 0xda, 0x2d, 0x27,
	0x51, 0x37, 0x9a, 0x6b, 0x2d, 0xb3, 0xb3, 0xbd, 0x6c, 0x29, 0x13, 0x7b, 0x99, 0xec, 0xa0, 0x01,
	0xa5, 0xcc, 0x2b, 0x16, 0xc0, 0xe8, 0x5a, 0xb9, 0xe4, 0xe9, 0x59, 0x46, 0xf2, 0x9c, 0x59, 0xf9,
	0xce, 0x97, 0xa1, 0x03, 0x19, 0x6a, 0x0e, 0xde, 0x42, 0xe5, 0x82, 0xe4, 0x20, 0xf3, 0x89, 0xb5,
	0xbf, 0x1b, 0x54, 0x6e, 0xf6, 0x8a, 0xc5, 0xce, 0xee, 0xdb, 0xc7, 0xe7, 0x7b, 0xbe, 0x86, 0x55,
	0x77, 0x76, 0xec, 0xc6, 0x89, 0x0b, 0xf7, 0x25, 0x0d, 0xf8, 0x15, 0x7b, 0xb0, 0xb9, 0xc8, 0x16,
	0x58, 0xfd, 0xa5, 0xfc, 0x84, 0x61, 0xdb, 0xff, 0xa3, 0x93, 0x00, 0x9c, 0x1c, 0x9e, 0x82, 0x99,
	0x32, 0xfa, 0xd1, 0x03, 0x5b, 0x65, 0xaf, 0xba, 0x3c, 0x56, 0x07, 0xe0, 0xe4, 0xee, 0x8b, 0xea,
	0x06, 0x4e, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x2f, 0x7e, 0x80, 0x46, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProdServiceClient is the client API for ProdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProdServiceClient interface {
	GetProdStock(ctx context.Context, in *ProdRequest, opts ...grpc.CallOption) (*ProdResponse, error)
	GetProdStocks(ctx context.Context, in *QuerySize, opts ...grpc.CallOption) (*ProdResponseList, error)
	GetProdInfo(ctx context.Context, in *ProdRequest, opts ...grpc.CallOption) (*ProdModel, error)
}

type prodServiceClient struct {
	cc *grpc.ClientConn
}

func NewProdServiceClient(cc *grpc.ClientConn) ProdServiceClient {
	return &prodServiceClient{cc}
}

func (c *prodServiceClient) GetProdStock(ctx context.Context, in *ProdRequest, opts ...grpc.CallOption) (*ProdResponse, error) {
	out := new(ProdResponse)
	err := c.cc.Invoke(ctx, "/services.ProdService/GetProdStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prodServiceClient) GetProdStocks(ctx context.Context, in *QuerySize, opts ...grpc.CallOption) (*ProdResponseList, error) {
	out := new(ProdResponseList)
	err := c.cc.Invoke(ctx, "/services.ProdService/GetProdStocks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prodServiceClient) GetProdInfo(ctx context.Context, in *ProdRequest, opts ...grpc.CallOption) (*ProdModel, error) {
	out := new(ProdModel)
	err := c.cc.Invoke(ctx, "/services.ProdService/GetProdInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProdServiceServer is the server API for ProdService service.
type ProdServiceServer interface {
	GetProdStock(context.Context, *ProdRequest) (*ProdResponse, error)
	GetProdStocks(context.Context, *QuerySize) (*ProdResponseList, error)
	GetProdInfo(context.Context, *ProdRequest) (*ProdModel, error)
}

// UnimplementedProdServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProdServiceServer struct {
}

func (*UnimplementedProdServiceServer) GetProdStock(ctx context.Context, req *ProdRequest) (*ProdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProdStock not implemented")
}
func (*UnimplementedProdServiceServer) GetProdStocks(ctx context.Context, req *QuerySize) (*ProdResponseList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProdStocks not implemented")
}
func (*UnimplementedProdServiceServer) GetProdInfo(ctx context.Context, req *ProdRequest) (*ProdModel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProdInfo not implemented")
}

func RegisterProdServiceServer(s *grpc.Server, srv ProdServiceServer) {
	s.RegisterService(&_ProdService_serviceDesc, srv)
}

func _ProdService_GetProdStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProdServiceServer).GetProdStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.ProdService/GetProdStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProdServiceServer).GetProdStock(ctx, req.(*ProdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProdService_GetProdStocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySize)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProdServiceServer).GetProdStocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.ProdService/GetProdStocks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProdServiceServer).GetProdStocks(ctx, req.(*QuerySize))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProdService_GetProdInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProdServiceServer).GetProdInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.ProdService/GetProdInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProdServiceServer).GetProdInfo(ctx, req.(*ProdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProdService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.ProdService",
	HandlerType: (*ProdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProdStock",
			Handler:    _ProdService_GetProdStock_Handler,
		},
		{
			MethodName: "GetProdStocks",
			Handler:    _ProdService_GetProdStocks_Handler,
		},
		{
			MethodName: "GetProdInfo",
			Handler:    _ProdService_GetProdInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Prot.proto",
}
