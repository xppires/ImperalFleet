package repository

import (
	"app/internal/common"
	pb "app/internal/genproto/users"
	"app/internal/models"
	"context"
	"fmt"
)

var UsersLocal = []models.User{
	// Sample users
	// In a real application, you would fetch these from a database
	// For simplicity, we are using hardcoded values here
	{UID: "TK1", Username: "neo", Password: "$2a$12$1kB7EE06HKOEsC23Sadm8.gaQZFCH9FtlQVV9ob4JSG9Ei8oLmWy2", Email: "", Role: "Technician"},
	{UID: "MG1", Username: "morpheus", Password: "$2a$12$cHVLmITk0X1sNi7nkTfwd.IxlmEeMnliBSAMM0eUIVBE.B4hh2a52", Email: "", Role: "Manager"},
}

type UmsRepositoryLocal struct {
	users []models.User
}

func NewUmsRepositoryLocal() *UmsRepositoryLocal {
	return &UmsRepositoryLocal{users: UsersLocal}
}

func (a *UmsRepositoryLocal) Authenticate(ctx context.Context, authRequest *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error) {
	for _, user := range a.users {
		if user.Username == authRequest.Username && common.CheckPasswordHash(authRequest.Password, user.Password) {
			fmt.Println(user.UID, user.Username, user.Password, user.Role)
			return &pb.AuthenticateResponse{
				UID:  user.UID,
				Role: user.Role,
			}, nil
		}
	}
	return nil, fmt.Errorf("incorrect credantials") // or return an error if preferred
}

func (a *UmsRepositoryLocal) GetUsers(ctx context.Context, userRequest *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {

	var pbUsers []*pb.User
	for _, user := range a.users {
		pbUsers = append(pbUsers, &pb.User{
			UID:      user.UID,
			Username: user.Username,
			Email:    user.Email,
			Role:     user.Role,
		})
	}

	return &pb.GetUsersResponse{
		Users: pbUsers,
	}, nil
}

func (a *UmsRepositoryLocal) GetUserById(ctx context.Context, userRequest *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {
	for _, user := range a.users {
		if user.UID == userRequest.UID {
			return &pb.GetUserByIdResponse{
				User: &pb.User{
					UID:      user.UID,
					Username: user.Username,
					Email:    user.Email,
					Role:     user.Role,
				},
			}, nil
		}
	}
	return nil, nil
}
