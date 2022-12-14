// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: doctor_svc.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DoctorServicesClient is the client API for DoctorServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DoctorServicesClient interface {
	RegisterDoctor(ctx context.Context, in *RegisterDoctorRequest, opts ...grpc.CallOption) (*RegisterDoctorResponse, error)
}

type doctorServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewDoctorServicesClient(cc grpc.ClientConnInterface) DoctorServicesClient {
	return &doctorServicesClient{cc}
}

func (c *doctorServicesClient) RegisterDoctor(ctx context.Context, in *RegisterDoctorRequest, opts ...grpc.CallOption) (*RegisterDoctorResponse, error) {
	out := new(RegisterDoctorResponse)
	err := c.cc.Invoke(ctx, "/pb.DoctorServices/RegisterDoctor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DoctorServicesServer is the server API for DoctorServices service.
// All implementations should embed UnimplementedDoctorServicesServer
// for forward compatibility
type DoctorServicesServer interface {
	RegisterDoctor(context.Context, *RegisterDoctorRequest) (*RegisterDoctorResponse, error)
}

// UnimplementedDoctorServicesServer should be embedded to have forward compatible implementations.
type UnimplementedDoctorServicesServer struct {
}

func (UnimplementedDoctorServicesServer) RegisterDoctor(context.Context, *RegisterDoctorRequest) (*RegisterDoctorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterDoctor not implemented")
}

// UnsafeDoctorServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DoctorServicesServer will
// result in compilation errors.
type UnsafeDoctorServicesServer interface {
	mustEmbedUnimplementedDoctorServicesServer()
}

func RegisterDoctorServicesServer(s grpc.ServiceRegistrar, srv DoctorServicesServer) {
	s.RegisterService(&DoctorServices_ServiceDesc, srv)
}

func _DoctorServices_RegisterDoctor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterDoctorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServicesServer).RegisterDoctor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DoctorServices/RegisterDoctor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServicesServer).RegisterDoctor(ctx, req.(*RegisterDoctorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DoctorServices_ServiceDesc is the grpc.ServiceDesc for DoctorServices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DoctorServices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.DoctorServices",
	HandlerType: (*DoctorServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterDoctor",
			Handler:    _DoctorServices_RegisterDoctor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "doctor_svc.proto",
}
