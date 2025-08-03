package gRPCServer

import (
	"log"
	"net"
	"google.golang.org/grpc"
)

type umsgRPCServer struct {
	addr string 
	GrpcServer *grpc.Server
}

func NewUmsGRPCServer(addr string) *umsgRPCServer {
	
	server :=  grpc.NewServer()
	return &umsgRPCServer{
		addr: addr, 
		GrpcServer: server,
		}
}

func (s *umsgRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

	// grpcServer := grpc.NewServer()
	// register our grpc services
	// umsService := service.NewUmsService()
	// handler.NewGrpcUmsHandler(grpcServer, s.umsService)

	log.Println("Starting UMS gRPC server on", s.addr)

	return s.GrpcServer.Serve(lis)
}