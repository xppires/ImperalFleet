package gRPCServer

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

type UmsgRPCServer struct {
	addr       string
	GrpcServer *grpc.Server
}

func NewUmsGRPCServer(addr string) *UmsgRPCServer {

	server := grpc.NewServer()
	return &UmsgRPCServer{
		addr:       addr,
		GrpcServer: server,
	}
}

func (s *UmsgRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Starting UMS gRPC server on", s.addr)

	return s.GrpcServer.Serve(lis)
}
