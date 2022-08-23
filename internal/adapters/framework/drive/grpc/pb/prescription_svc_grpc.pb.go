// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// PrescriptionServiceClient is the client API for PrescriptionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PrescriptionServiceClient interface {
	GetLatestPrescriptions(ctx context.Context, in *GetLatestPrescriptionsParameters, opts ...grpc.CallOption) (*LatestPrescriptionResult, error)
}

type prescriptionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPrescriptionServiceClient(cc grpc.ClientConnInterface) PrescriptionServiceClient {
	return &prescriptionServiceClient{cc}
}

func (c *prescriptionServiceClient) GetLatestPrescriptions(ctx context.Context, in *GetLatestPrescriptionsParameters, opts ...grpc.CallOption) (*LatestPrescriptionResult, error) {
	out := new(LatestPrescriptionResult)
	err := c.cc.Invoke(ctx, "/pb.PrescriptionService/GetLatestPrescriptions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PrescriptionServiceServer is the server API for PrescriptionService service.
// All implementations should embed UnimplementedPrescriptionServiceServer
// for forward compatibility
type PrescriptionServiceServer interface {
	GetLatestPrescriptions(context.Context, *GetLatestPrescriptionsParameters) (*LatestPrescriptionResult, error)
}

// UnimplementedPrescriptionServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPrescriptionServiceServer struct {
}

func (UnimplementedPrescriptionServiceServer) GetLatestPrescriptions(context.Context, *GetLatestPrescriptionsParameters) (*LatestPrescriptionResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestPrescriptions not implemented")
}

// UnsafePrescriptionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PrescriptionServiceServer will
// result in compilation errors.
type UnsafePrescriptionServiceServer interface {
	mustEmbedUnimplementedPrescriptionServiceServer()
}

func RegisterPrescriptionServiceServer(s grpc.ServiceRegistrar, srv PrescriptionServiceServer) {
	s.RegisterService(&PrescriptionService_ServiceDesc, srv)
}

func _PrescriptionService_GetLatestPrescriptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLatestPrescriptionsParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrescriptionServiceServer).GetLatestPrescriptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PrescriptionService/GetLatestPrescriptions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrescriptionServiceServer).GetLatestPrescriptions(ctx, req.(*GetLatestPrescriptionsParameters))
	}
	return interceptor(ctx, in, info, handler)
}

// PrescriptionService_ServiceDesc is the grpc.ServiceDesc for PrescriptionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PrescriptionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.PrescriptionService",
	HandlerType: (*PrescriptionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLatestPrescriptions",
			Handler:    _PrescriptionService_GetLatestPrescriptions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "prescription_svc.proto",
}