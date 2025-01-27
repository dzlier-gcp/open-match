// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/api/synchronizer.proto

package ipb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
	pb "open-match.dev/open-match/pkg/pb"
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

type RegisterRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_35ff6b85fea1c4b7, []int{0}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

type RegisterResponse struct {
	// Identifier for this request valid for the current synchronization cycle.
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_35ff6b85fea1c4b7, []int{1}
}

func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (m *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(m, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

func (m *RegisterResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type EvaluateProposalsRequest struct {
	// List of proposals to evaluate in the current synchronization cycle.
	Match *pb.Match `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	// Identifier for this request issued during request registration.
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EvaluateProposalsRequest) Reset()         { *m = EvaluateProposalsRequest{} }
func (m *EvaluateProposalsRequest) String() string { return proto.CompactTextString(m) }
func (*EvaluateProposalsRequest) ProtoMessage()    {}
func (*EvaluateProposalsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_35ff6b85fea1c4b7, []int{2}
}

func (m *EvaluateProposalsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvaluateProposalsRequest.Unmarshal(m, b)
}
func (m *EvaluateProposalsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvaluateProposalsRequest.Marshal(b, m, deterministic)
}
func (m *EvaluateProposalsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvaluateProposalsRequest.Merge(m, src)
}
func (m *EvaluateProposalsRequest) XXX_Size() int {
	return xxx_messageInfo_EvaluateProposalsRequest.Size(m)
}
func (m *EvaluateProposalsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EvaluateProposalsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EvaluateProposalsRequest proto.InternalMessageInfo

func (m *EvaluateProposalsRequest) GetMatch() *pb.Match {
	if m != nil {
		return m.Match
	}
	return nil
}

func (m *EvaluateProposalsRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type EvaluateProposalsResponse struct {
	// Results from evaluating proposals for this request.
	Match                *pb.Match `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *EvaluateProposalsResponse) Reset()         { *m = EvaluateProposalsResponse{} }
func (m *EvaluateProposalsResponse) String() string { return proto.CompactTextString(m) }
func (*EvaluateProposalsResponse) ProtoMessage()    {}
func (*EvaluateProposalsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_35ff6b85fea1c4b7, []int{3}
}

func (m *EvaluateProposalsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvaluateProposalsResponse.Unmarshal(m, b)
}
func (m *EvaluateProposalsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvaluateProposalsResponse.Marshal(b, m, deterministic)
}
func (m *EvaluateProposalsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvaluateProposalsResponse.Merge(m, src)
}
func (m *EvaluateProposalsResponse) XXX_Size() int {
	return xxx_messageInfo_EvaluateProposalsResponse.Size(m)
}
func (m *EvaluateProposalsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EvaluateProposalsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EvaluateProposalsResponse proto.InternalMessageInfo

func (m *EvaluateProposalsResponse) GetMatch() *pb.Match {
	if m != nil {
		return m.Match
	}
	return nil
}

func init() {
	proto.RegisterType((*RegisterRequest)(nil), "openmatch.internal.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "openmatch.internal.RegisterResponse")
	proto.RegisterType((*EvaluateProposalsRequest)(nil), "openmatch.internal.EvaluateProposalsRequest")
	proto.RegisterType((*EvaluateProposalsResponse)(nil), "openmatch.internal.EvaluateProposalsResponse")
}

func init() { proto.RegisterFile("internal/api/synchronizer.proto", fileDescriptor_35ff6b85fea1c4b7) }

var fileDescriptor_35ff6b85fea1c4b7 = []byte{
	// 551 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0x96, 0x5d, 0x28, 0x70, 0x20, 0x68, 0x6f, 0x40, 0x6d, 0x40, 0xc2, 0x75, 0x51, 0x55, 0x45,
	0x8d, 0x2f, 0x0d, 0x99, 0x82, 0x90, 0x5a, 0x4a, 0x86, 0x4a, 0xe5, 0x87, 0x5c, 0x89, 0x81, 0xed,
	0x62, 0x3f, 0xec, 0x43, 0xf6, 0xdd, 0x71, 0xef, 0x9c, 0x16, 0x46, 0x36, 0x56, 0xd8, 0x98, 0xd9,
	0xd9, 0xf8, 0x47, 0x98, 0xd9, 0xf8, 0x43, 0x90, 0xed, 0xb8, 0x49, 0x9b, 0x80, 0xba, 0xd8, 0xf2,
	0xbd, 0xef, 0x7d, 0xf7, 0x7d, 0xdf, 0xdd, 0x33, 0x79, 0x20, 0xa4, 0x05, 0x23, 0x79, 0xc6, 0xb8,
	0x16, 0x0c, 0x3f, 0xc8, 0x28, 0x35, 0x4a, 0x8a, 0x8f, 0x60, 0x02, 0x6d, 0x94, 0x55, 0x94, 0x2a,
	0x0d, 0x32, 0xe7, 0x36, 0x4a, 0x83, 0x06, 0xda, 0xa2, 0x25, 0x36, 0x07, 0x44, 0x9e, 0x00, 0xd6,
	0xb8, 0xd6, 0xfd, 0x44, 0xa9, 0x24, 0x83, 0x8a, 0x86, 0x4b, 0xa9, 0x2c, 0xb7, 0x42, 0xc9, 0xa6,
	0xba, 0x53, 0xbd, 0xa2, 0x4e, 0x02, 0xb2, 0x83, 0x27, 0x3c, 0x49, 0xc0, 0x30, 0xa5, 0x2b, 0xc4,
	0x3c, 0xda, 0x5f, 0x25, 0x77, 0x42, 0x48, 0x04, 0x5a, 0x30, 0x21, 0xbc, 0x2f, 0x00, 0xad, 0xef,
	0x93, 0x95, 0xe9, 0x12, 0x6a, 0x25, 0x11, 0xe8, 0x6d, 0xe2, 0x8a, 0x78, 0xcd, 0xf1, 0x9c, 0xed,
	0x1b, 0xa1, 0x2b, 0x62, 0x3f, 0x24, 0x6b, 0xc3, 0x31, 0xcf, 0x0a, 0x6e, 0xe1, 0x95, 0x51, 0x5a,
	0x21, 0xcf, 0x70, 0xd2, 0x4f, 0xb7, 0xc8, 0xd5, 0xca, 0x44, 0x05, 0xbf, 0xd9, 0x5b, 0x09, 0xa6,
	0xb6, 0x9e, 0x97, 0xcf, 0xb0, 0x2e, 0x4f, 0x38, 0xdd, 0x33, 0xce, 0x03, 0xb2, 0xbe, 0x80, 0x73,
	0x22, 0xe0, 0x92, 0xa4, 0xbd, 0x1f, 0x2e, 0xb9, 0x75, 0x3c, 0x13, 0x2d, 0x3d, 0x25, 0xd7, 0x1b,
	0x37, 0x74, 0x33, 0x98, 0x4f, 0x38, 0xb8, 0x60, 0xbf, 0xf5, 0xf0, 0xff, 0xa0, 0x5a, 0x8f, 0xbf,
	0xf1, 0xe9, 0xd7, 0x9f, 0xaf, 0xee, 0x3d, 0xba, 0xce, 0xc6, 0xbb, 0xe7, 0xce, 0x92, 0x99, 0x66,
	0xb7, 0xef, 0x0e, 0x59, 0x9d, 0x33, 0x44, 0x77, 0x16, 0xd1, 0xff, 0x2b, 0xcb, 0x56, 0xe7, 0x92,
	0xe8, 0x89, 0xaa, 0xa0, 0x52, 0xb5, 0xed, 0x6f, 0xce, 0xa9, 0xd2, 0x0d, 0x76, 0x00, 0x93, 0xee,
	0x81, 0xd3, 0xde, 0x76, 0xba, 0xce, 0xd3, 0xcf, 0x4b, 0x5f, 0xf6, 0x7f, 0xbb, 0xf4, 0xa7, 0x73,
	0x3e, 0x38, 0xff, 0x90, 0x90, 0x97, 0x1a, 0xa4, 0x57, 0xc5, 0x4b, 0xef, 0xa6, 0xd6, 0x6a, 0x1c,
	0x30, 0x56, 0xea, 0xe9, 0xd4, 0x82, 0x62, 0x18, 0xb7, 0x36, 0xa7, 0xdf, 0x9d, 0x58, 0x60, 0x54,
	0x20, 0xee, 0xd5, 0xd7, 0x33, 0x31, 0xaa, 0xd0, 0x18, 0x44, 0x2a, 0x6f, 0xbf, 0x26, 0x74, 0x5f,
	0xf3, 0x28, 0x05, 0xaf, 0x17, 0x74, 0xbd, 0x23, 0x11, 0x41, 0x79, 0xa4, 0x7b, 0x0d, 0x65, 0x22,
	0x6c, 0x5a, 0x8c, 0x4a, 0x24, 0xab, 0x5b, 0xdf, 0x2a, 0x93, 0xf0, 0x1c, 0x70, 0x66, 0x33, 0x36,
	0xca, 0xd4, 0x88, 0xe5, 0xbc, 0x4c, 0x96, 0x1d, 0x1d, 0x1e, 0x0c, 0x5f, 0x1c, 0x0f, 0x7b, 0x4b,
	0xbb, 0x41, 0xb7, 0xed, 0x3a, 0x6e, 0x6f, 0x85, 0x6b, 0x9d, 0x89, 0xa8, 0xba, 0xd9, 0xec, 0x1d,
	0x2a, 0x39, 0x98, 0x5b, 0x09, 0x1f, 0x93, 0xa5, 0x7e, 0xb7, 0x4f, 0xfb, 0xa4, 0x1d, 0x82, 0x2d,
	0x8c, 0x84, 0xd8, 0x3b, 0x49, 0x41, 0x7a, 0x36, 0x05, 0xcf, 0x00, 0xaa, 0xc2, 0x44, 0xe0, 0xc5,
	0x0a, 0xd0, 0x93, 0xca, 0x7a, 0x70, 0x2a, 0xd0, 0x06, 0x74, 0x99, 0x5c, 0xf9, 0xe6, 0x3a, 0xd7,
	0xcc, 0x13, 0xb2, 0x36, 0x0d, 0xc3, 0x7b, 0xa6, 0xa2, 0x22, 0x07, 0x59, 0x4f, 0x12, 0xdd, 0x58,
	0x1c, 0x0d, 0x43, 0x61, 0x81, 0xc5, 0x2a, 0x42, 0xf6, 0x66, 0xeb, 0x42, 0x69, 0xc6, 0xd7, 0xd9,
	0x4f, 0x41, 0xe8, 0xd1, 0x68, 0xb9, 0x1a, 0xca, 0x47, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xbb,
	0x4d, 0x17, 0xa3, 0x2b, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SynchronizerClient is the client API for Synchronizer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SynchronizerClient interface {
	// Register associates this request with the current synchronization cycle and
	// returns an identifier for this registration. The caller returns this
	// identifier back in the evaluation request. This enables synchronizer to
	// identify stale evaluation requests belonging to a prior cycle.
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	// EvaluateProposals accepts a list of proposals and a registration identifier
	// for this request. If the synchronization cycle to which the request was
	// registered is completed, this request fails otherwise the proposals are
	// added to the list of proposals to be evaluated in the current cycle. At the
	//  end of the cycle, the user defined evaluation method is triggered and the
	// matches accepted by it are returned as results.
	EvaluateProposals(ctx context.Context, opts ...grpc.CallOption) (Synchronizer_EvaluateProposalsClient, error)
}

type synchronizerClient struct {
	cc *grpc.ClientConn
}

func NewSynchronizerClient(cc *grpc.ClientConn) SynchronizerClient {
	return &synchronizerClient{cc}
}

func (c *synchronizerClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/openmatch.internal.Synchronizer/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synchronizerClient) EvaluateProposals(ctx context.Context, opts ...grpc.CallOption) (Synchronizer_EvaluateProposalsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Synchronizer_serviceDesc.Streams[0], "/openmatch.internal.Synchronizer/EvaluateProposals", opts...)
	if err != nil {
		return nil, err
	}
	x := &synchronizerEvaluateProposalsClient{stream}
	return x, nil
}

type Synchronizer_EvaluateProposalsClient interface {
	Send(*EvaluateProposalsRequest) error
	Recv() (*EvaluateProposalsResponse, error)
	grpc.ClientStream
}

type synchronizerEvaluateProposalsClient struct {
	grpc.ClientStream
}

func (x *synchronizerEvaluateProposalsClient) Send(m *EvaluateProposalsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *synchronizerEvaluateProposalsClient) Recv() (*EvaluateProposalsResponse, error) {
	m := new(EvaluateProposalsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SynchronizerServer is the server API for Synchronizer service.
type SynchronizerServer interface {
	// Register associates this request with the current synchronization cycle and
	// returns an identifier for this registration. The caller returns this
	// identifier back in the evaluation request. This enables synchronizer to
	// identify stale evaluation requests belonging to a prior cycle.
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	// EvaluateProposals accepts a list of proposals and a registration identifier
	// for this request. If the synchronization cycle to which the request was
	// registered is completed, this request fails otherwise the proposals are
	// added to the list of proposals to be evaluated in the current cycle. At the
	//  end of the cycle, the user defined evaluation method is triggered and the
	// matches accepted by it are returned as results.
	EvaluateProposals(Synchronizer_EvaluateProposalsServer) error
}

// UnimplementedSynchronizerServer can be embedded to have forward compatible implementations.
type UnimplementedSynchronizerServer struct {
}

func (*UnimplementedSynchronizerServer) Register(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedSynchronizerServer) EvaluateProposals(srv Synchronizer_EvaluateProposalsServer) error {
	return status.Errorf(codes.Unimplemented, "method EvaluateProposals not implemented")
}

func RegisterSynchronizerServer(s *grpc.Server, srv SynchronizerServer) {
	s.RegisterService(&_Synchronizer_serviceDesc, srv)
}

func _Synchronizer_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynchronizerServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openmatch.internal.Synchronizer/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynchronizerServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Synchronizer_EvaluateProposals_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SynchronizerServer).EvaluateProposals(&synchronizerEvaluateProposalsServer{stream})
}

type Synchronizer_EvaluateProposalsServer interface {
	Send(*EvaluateProposalsResponse) error
	Recv() (*EvaluateProposalsRequest, error)
	grpc.ServerStream
}

type synchronizerEvaluateProposalsServer struct {
	grpc.ServerStream
}

func (x *synchronizerEvaluateProposalsServer) Send(m *EvaluateProposalsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *synchronizerEvaluateProposalsServer) Recv() (*EvaluateProposalsRequest, error) {
	m := new(EvaluateProposalsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Synchronizer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openmatch.internal.Synchronizer",
	HandlerType: (*SynchronizerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Synchronizer_Register_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EvaluateProposals",
			Handler:       _Synchronizer_EvaluateProposals_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "internal/api/synchronizer.proto",
}
