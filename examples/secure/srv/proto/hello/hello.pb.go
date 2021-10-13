// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go-micro.dev/examples/v4/secure/srv/proto/hello/hello.proto

/*
Package go_micro_srv_greeter is a generated protocol buffer package.

It is generated from these files:
	go-micro.dev/examples/v4/secure/srv/proto/hello/hello.proto

It has these top-level messages:
	Request
	Response
*/
package go_micro_srv_greeter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.greeter.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.greeter.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Say service

type SayClient interface {
	Hello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type sayClient struct {
	cc *grpc.ClientConn
}

func NewSayClient(cc *grpc.ClientConn) SayClient {
	return &sayClient{cc}
}

func (c *sayClient) Hello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/go.micro.srv.greeter.Say/Hello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Say service

type SayServer interface {
	Hello(context.Context, *Request) (*Response, error)
}

func RegisterSayServer(s *grpc.Server, srv SayServer) {
	s.RegisterService(&_Say_serviceDesc, srv)
}

func _Say_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SayServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.srv.greeter.Say/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SayServer).Hello(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Say_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go.micro.srv.greeter.Say",
	HandlerType: (*SayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Say_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go-micro.dev/examples/v4/secure/srv/proto/hello/hello.proto",
}

func init() {
	proto.RegisterFile("go-micro.dev/examples/v4/secure/srv/proto/hello/hello.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xce, 0x4f, 0xcb, 0x82, 0x40,
	0x10, 0x06, 0xf0, 0x57, 0x7c, 0xff, 0x35, 0xa7, 0x58, 0x3a, 0x44, 0x64, 0x84, 0xa7, 0x4e, 0xb3,
	0x50, 0xd7, 0x3e, 0x80, 0xb7, 0xc0, 0x3e, 0x81, 0xca, 0xb0, 0x0a, 0xae, 0x6b, 0x33, 0xab, 0xd4,
	0xb7, 0x0f, 0x37, 0x8f, 0x75, 0x19, 0x1e, 0xe6, 0x61, 0xf8, 0x0d, 0x9c, 0x4d, 0xe3, 0xeb, 0xa1,
	0xc4, 0xca, 0x59, 0x6d, 0x9b, 0x8a, 0x9d, 0xa6, 0x7b, 0x61, 0xfb, 0x96, 0x44, 0x0b, 0x55, 0x03,
	0x93, 0x16, 0x1e, 0x75, 0xcf, 0xce, 0x3b, 0x5d, 0x53, 0xdb, 0xce, 0x13, 0xc3, 0x46, 0xad, 0x8c,
	0xc3, 0x70, 0x85, 0xc2, 0x23, 0x1a, 0x26, 0xf2, 0xc4, 0x69, 0x02, 0x7f, 0x39, 0xdd, 0x06, 0x12,
	0xaf, 0x14, 0x7c, 0x77, 0x85, 0xa5, 0x75, 0xb4, 0x8f, 0x0e, 0x8b, 0x3c, 0xe4, 0x74, 0x0b, 0xff,
	0x39, 0x49, 0xef, 0x3a, 0x21, 0xb5, 0x84, 0xd8, 0x8a, 0x99, 0xeb, 0x29, 0x1e, 0x2f, 0x10, 0x5f,
	0x8b, 0x87, 0xca, 0xe0, 0x27, 0x9b, 0x20, 0x95, 0xe0, 0x3b, 0x03, 0x67, 0x60, 0xb3, 0xfb, 0x54,
	0xbf, 0x80, 0xf4, 0xab, 0xfc, 0x0d, 0xaf, 0x9e, 0x9e, 0x01, 0x00, 0x00, 0xff, 0xff, 0x61, 0xd8,
	0xd1, 0xf2, 0xea, 0x00, 0x00, 0x00,
}
