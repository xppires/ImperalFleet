package repository

import (
	"log"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"app/internal/genproto/users"
)

type AuthRepositoryGRPCClient struct {
	conn *grpc.ClientConn
}

func  NewAuthRepositoryGRPCClient(addr string) *AuthRepositoryGRPCClient {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return  &AuthRepositoryGRPCClient{
		conn: conn,
	}
}

func (c  AuthRepositoryGRPCClient) Authenticate(username, password string) (bool, string, string, error) {
	client := users.NewUserServiceClient(c.conn)
	req := &users.AuthenticateRequest{
		Username: username,
		Password: password,
	}

	resp, err := client.Authenticate(context.Background(), req)
	if err != nil {
		log.Printf("error during authentication: %v", err)
		return false, "", "", err
	}

	return true, resp.UID, resp.Role, nil
}