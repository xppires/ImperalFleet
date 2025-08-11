package handlers

import (
	pb "app/internal/genproto/users"
	"app/internal/services"
	"context"
	"google.golang.org/grpc"
)

type UmsgRPCHandler struct {
	service *services.UmsService
}

func NewGrpcUmsHandler(grpcServer *grpc.Server, service *services.UmsService) *UmsgRPCHandler {
	h := &UmsgRPCHandler{service: service}
	// Register the handler with the gRPC server

	pb.RegisterUserServiceServer(grpcServer, h)
	return h
}

func (h *UmsgRPCHandler) Authenticate(ctx context.Context, userRequest *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error) {

	return h.service.Authenticate(ctx, userRequest)
}

func (h *UmsgRPCHandler) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {

	return h.service.GetUserById(ctx, req)
}

func (h *UmsgRPCHandler) GetUsers(ctx context.Context, userRequest *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {

	return h.service.GetUsers(ctx, userRequest)
}

func (h *UmsgRPCHandler) mustEmbedUnimplementedUserServiceServer() {
	// This method is required to implement the gRPC service interface.
	// It can be left empty if you are not using the generated code.
}
