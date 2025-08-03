package repository
import (
	"context"
	pb "app/internal/genproto/users"  
)	

type UmsRepository interface {
	Authenticate(ctx context.Context, authRequest *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error)
	GetUserById(ctx context.Context, userRequest *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error)
	GetUsers(ctx context.Context, userRequest *pb.GetUsersRequest) (*pb.GetUsersResponse, error)
}

