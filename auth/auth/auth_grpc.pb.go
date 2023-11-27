package auth

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

const _ = grpc.SupportPackageIsVersion7

const (
	Greeter_Auth_FullMethodName = "/auth.Greeter/Auth"
)

type GreeterClient interface {
	Auth(ctx context.Context, in *Auth_Check_Request, opts ...grpc.CallOption) (*Auth_Check_Reply, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) Auth(ctx context.Context, in *Auth_Check_Request, opts ...grpc.CallOption) (*Auth_Check_Reply, error) {
	out := new(Auth_Check_Reply)
	err := c.cc.Invoke(ctx, Greeter_Auth_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type GreeterServer interface {
	Auth(context.Context, *Auth_Check_Request) (*Auth_Check_Reply, error)
	mustEmbedUnimplementedGreeterServer()
}

type UnimplementedGreeterServer struct {
}

func (UnimplementedGreeterServer) Auth(context.Context, *Auth_Check_Request) (*Auth_Check_Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}

type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	s.RegisterService(&Greeter_ServiceDesc, srv)
}

func _Greeter_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth_Check_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greeter_Auth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Auth(ctx, req.(*Auth_Check_Request))
	}
	return interceptor(ctx, in, info, handler)
}

var Greeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Auth",
			Handler:    _Greeter_Auth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "/auth/auth.proto",
}
