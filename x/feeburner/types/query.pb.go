// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: neutron/feeburner/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f540485c4b79b2ac, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f540485c4b79b2ac, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// QueryTotalBurnedNeutronsAmountRequest is request type for the
// Query/QueryTotalBurnedNeutronsAmount method.
type QueryTotalBurnedNeutronsAmountRequest struct {
}

func (m *QueryTotalBurnedNeutronsAmountRequest) Reset()         { *m = QueryTotalBurnedNeutronsAmountRequest{} }
func (m *QueryTotalBurnedNeutronsAmountRequest) String() string { return proto.CompactTextString(m) }
func (*QueryTotalBurnedNeutronsAmountRequest) ProtoMessage()    {}
func (*QueryTotalBurnedNeutronsAmountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f540485c4b79b2ac, []int{2}
}
func (m *QueryTotalBurnedNeutronsAmountRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryTotalBurnedNeutronsAmountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTotalBurnedNeutronsAmountRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryTotalBurnedNeutronsAmountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTotalBurnedNeutronsAmountRequest.Merge(m, src)
}
func (m *QueryTotalBurnedNeutronsAmountRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryTotalBurnedNeutronsAmountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTotalBurnedNeutronsAmountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTotalBurnedNeutronsAmountRequest proto.InternalMessageInfo

// QueryTotalBurnedNeutronsAmountResponse is response type for the
// Query/QueryTotalBurnedNeutronsAmount method.
type QueryTotalBurnedNeutronsAmountResponse struct {
	TotalBurnedNeutronsAmount TotalBurnedNeutronsAmount `protobuf:"bytes,1,opt,name=total_burned_neutrons_amount,json=totalBurnedNeutronsAmount,proto3" json:"total_burned_neutrons_amount"`
}

func (m *QueryTotalBurnedNeutronsAmountResponse) Reset() {
	*m = QueryTotalBurnedNeutronsAmountResponse{}
}
func (m *QueryTotalBurnedNeutronsAmountResponse) String() string { return proto.CompactTextString(m) }
func (*QueryTotalBurnedNeutronsAmountResponse) ProtoMessage()    {}
func (*QueryTotalBurnedNeutronsAmountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f540485c4b79b2ac, []int{3}
}
func (m *QueryTotalBurnedNeutronsAmountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryTotalBurnedNeutronsAmountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTotalBurnedNeutronsAmountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryTotalBurnedNeutronsAmountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTotalBurnedNeutronsAmountResponse.Merge(m, src)
}
func (m *QueryTotalBurnedNeutronsAmountResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryTotalBurnedNeutronsAmountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTotalBurnedNeutronsAmountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTotalBurnedNeutronsAmountResponse proto.InternalMessageInfo

func (m *QueryTotalBurnedNeutronsAmountResponse) GetTotalBurnedNeutronsAmount() TotalBurnedNeutronsAmount {
	if m != nil {
		return m.TotalBurnedNeutronsAmount
	}
	return TotalBurnedNeutronsAmount{}
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "neutron.feeburner.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "neutron.feeburner.QueryParamsResponse")
	proto.RegisterType((*QueryTotalBurnedNeutronsAmountRequest)(nil), "neutron.feeburner.QueryTotalBurnedNeutronsAmountRequest")
	proto.RegisterType((*QueryTotalBurnedNeutronsAmountResponse)(nil), "neutron.feeburner.QueryTotalBurnedNeutronsAmountResponse")
}

func init() { proto.RegisterFile("neutron/feeburner/query.proto", fileDescriptor_f540485c4b79b2ac) }

var fileDescriptor_f540485c4b79b2ac = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x31, 0x6f, 0xda, 0x40,
	0x14, 0xc7, 0x6d, 0xd4, 0x32, 0x5c, 0xa7, 0x5e, 0x19, 0x6a, 0x97, 0xba, 0xad, 0x25, 0x68, 0x55,
	0xb5, 0x3e, 0x01, 0x95, 0x68, 0xc7, 0xb2, 0x17, 0x35, 0x28, 0x53, 0x16, 0x74, 0x26, 0x17, 0xc7,
	0x12, 0xbe, 0x67, 0x7c, 0x67, 0x14, 0x32, 0xe6, 0x13, 0x44, 0xca, 0x9c, 0xef, 0x43, 0xa6, 0x20,
	0x65, 0xc9, 0x14, 0x45, 0x90, 0x0f, 0x12, 0x71, 0x3e, 0x50, 0x22, 0xe3, 0xa0, 0x64, 0x3b, 0xdf,
	0xff, 0xfd, 0xdf, 0xff, 0xf7, 0x9e, 0x0f, 0x7d, 0xe4, 0x2c, 0x95, 0x09, 0x70, 0x72, 0xc0, 0x98,
	0x9f, 0x26, 0x9c, 0x25, 0x64, 0x94, 0xb2, 0x64, 0xe2, 0xc5, 0x09, 0x48, 0xc0, 0x6f, 0xb5, 0xec,
	0xad, 0x65, 0xfb, 0xfb, 0x00, 0x44, 0x04, 0x82, 0xf8, 0x54, 0xb0, 0xac, 0x96, 0x8c, 0x1b, 0x3e,
	0x93, 0xb4, 0x41, 0x62, 0x1a, 0x84, 0x9c, 0xca, 0x10, 0x78, 0x66, 0xb7, 0x2b, 0x01, 0x04, 0xa0,
	0x8e, 0x64, 0x79, 0xd2, 0xb7, 0xd5, 0x00, 0x20, 0x18, 0x32, 0x42, 0xe3, 0x90, 0x50, 0xce, 0x41,
	0x2a, 0x8b, 0xd0, 0xaa, 0x93, 0x27, 0x8a, 0x69, 0x42, 0xa3, 0x95, 0xfe, 0x2b, 0xaf, 0x4b, 0x90,
	0x74, 0xd8, 0x57, 0x1f, 0xfb, 0x7d, 0x2d, 0x8b, 0x3e, 0x8d, 0x20, 0xe5, 0x32, 0x73, 0xb9, 0x15,
	0x84, 0x77, 0x96, 0xac, 0xff, 0x55, 0xab, 0x1e, 0x1b, 0xa5, 0x4c, 0x48, 0xb7, 0x8b, 0xde, 0x3d,
	0xba, 0x15, 0x31, 0x70, 0xc1, 0x70, 0x1b, 0x95, 0xb3, 0xc8, 0xf7, 0xe6, 0x67, 0xf3, 0xdb, 0x9b,
	0xa6, 0xe5, 0xe5, 0xd6, 0xe0, 0x65, 0x96, 0xce, 0xab, 0xe9, 0xcd, 0x27, 0xa3, 0xa7, 0xcb, 0xdd,
	0xaf, 0xa8, 0xa6, 0xfa, 0xed, 0x2e, 0x81, 0x3a, 0x8a, 0xa7, 0xab, 0x71, 0xfe, 0x2a, 0x9a, 0x55,
	0xf0, 0xb9, 0x89, 0xea, 0xdb, 0x2a, 0x35, 0x8c, 0x40, 0xd5, 0xa7, 0xe6, 0xd3, 0x88, 0x3f, 0x36,
	0x20, 0x16, 0xf6, 0xd6, 0xd4, 0x96, 0x2c, 0x2a, 0x68, 0x5e, 0x96, 0xd0, 0x6b, 0xc5, 0x87, 0x8f,
	0x51, 0x39, 0x1b, 0x15, 0xd7, 0x36, 0x44, 0xe4, 0x77, 0x6a, 0xd7, 0xb7, 0x95, 0x65, 0x73, 0xb9,
	0x5f, 0x4e, 0xae, 0xee, 0xce, 0x4a, 0x1f, 0xb0, 0x45, 0x8a, 0x7e, 0x38, 0xbe, 0x30, 0x91, 0x55,
	0x38, 0x04, 0xfe, 0x5d, 0x14, 0xb4, 0x6d, 0xfb, 0xf6, 0x9f, 0x17, 0x38, 0x35, 0x75, 0x5b, 0x51,
	0x37, 0x30, 0x21, 0xcf, 0x7b, 0x86, 0x9d, 0x7f, 0xd3, 0xb9, 0x63, 0xce, 0xe6, 0x8e, 0x79, 0x3b,
	0x77, 0xcc, 0xd3, 0x85, 0x63, 0xcc, 0x16, 0x8e, 0x71, 0xbd, 0x70, 0x8c, 0xbd, 0x56, 0x10, 0xca,
	0xc3, 0xd4, 0xf7, 0x06, 0x10, 0xad, 0x9a, 0xfe, 0x84, 0x24, 0x58, 0x07, 0x8c, 0x9b, 0xe4, 0xe8,
	0x61, 0xca, 0x24, 0x66, 0xc2, 0x2f, 0xab, 0x67, 0xdd, 0xba, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xec,
	0x06, 0xd3, 0x86, 0xc0, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// TotalBurnedNeutronsAmount queries total amount of burned neutron fees.
	TotalBurnedNeutronsAmount(ctx context.Context, in *QueryTotalBurnedNeutronsAmountRequest, opts ...grpc.CallOption) (*QueryTotalBurnedNeutronsAmountResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/neutron.feeburner.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TotalBurnedNeutronsAmount(ctx context.Context, in *QueryTotalBurnedNeutronsAmountRequest, opts ...grpc.CallOption) (*QueryTotalBurnedNeutronsAmountResponse, error) {
	out := new(QueryTotalBurnedNeutronsAmountResponse)
	err := c.cc.Invoke(ctx, "/neutron.feeburner.Query/TotalBurnedNeutronsAmount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// TotalBurnedNeutronsAmount queries total amount of burned neutron fees.
	TotalBurnedNeutronsAmount(context.Context, *QueryTotalBurnedNeutronsAmountRequest) (*QueryTotalBurnedNeutronsAmountResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) TotalBurnedNeutronsAmount(ctx context.Context, req *QueryTotalBurnedNeutronsAmountRequest) (*QueryTotalBurnedNeutronsAmountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TotalBurnedNeutronsAmount not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neutron.feeburner.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TotalBurnedNeutronsAmount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTotalBurnedNeutronsAmountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TotalBurnedNeutronsAmount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neutron.feeburner.Query/TotalBurnedNeutronsAmount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TotalBurnedNeutronsAmount(ctx, req.(*QueryTotalBurnedNeutronsAmountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "neutron.feeburner.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "TotalBurnedNeutronsAmount",
			Handler:    _Query_TotalBurnedNeutronsAmount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "neutron/feeburner/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryTotalBurnedNeutronsAmountRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTotalBurnedNeutronsAmountRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTotalBurnedNeutronsAmountRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryTotalBurnedNeutronsAmountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTotalBurnedNeutronsAmountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTotalBurnedNeutronsAmountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.TotalBurnedNeutronsAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryTotalBurnedNeutronsAmountRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryTotalBurnedNeutronsAmountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.TotalBurnedNeutronsAmount.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryTotalBurnedNeutronsAmountRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryTotalBurnedNeutronsAmountRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTotalBurnedNeutronsAmountRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryTotalBurnedNeutronsAmountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryTotalBurnedNeutronsAmountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTotalBurnedNeutronsAmountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalBurnedNeutronsAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalBurnedNeutronsAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
