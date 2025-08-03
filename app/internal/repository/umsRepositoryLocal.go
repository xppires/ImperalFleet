package repository

import (
	"app/internal/models"
	pb "app/internal/genproto/users"
	"context"
)

var UsersLocal = []models.User{
	// Sample users
	// In a real application, you would fetch these from a database
	// For simplicity, we are using hardcoded values here
	{UID: "TK1", Username: "neo", Password: "keanu", Email: "", Role: "Technician"},
	{UID: "MG1", Username: "morpheus", Password: "lawrence", Email: "", Role: "Manager"},
}

type umsRepositoryLocal struct {
	users []models.User
}

func NewUmsRepositoryLocal() *umsRepositoryLocal {
	return &umsRepositoryLocal{users: UsersLocal}
}
func (a *umsRepositoryLocal) Authenticate(ctx context.Context, authRequest *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error) {
	for _, user := range a.users {
		if user.Username == authRequest.Username && user.Password == authRequest.Password {
			return &pb.AuthenticateResponse{
				UID:  user.UID,
				Role: user.Role,
			}, nil
		}
	}
	return nil, nil // or return an error if preferred
}

func  (a *umsRepositoryLocal) GetUsers(ctx context.Context, userRequest *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {

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

func  (a *umsRepositoryLocal) GetUserById(ctx context.Context, userRequest *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {
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
