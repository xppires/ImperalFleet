package  services

import (
	"context"
	"app/internal/repository" 
	pb "app/internal/genproto/users"  
)



type UmsService struct {
	repo repository.UmsRepository 
}
func NewUmsService(repo repository.UmsRepository) *UmsService {
	return &UmsService{repo: repo}
}
func (s *UmsService) Authenticate(ctx context.Context, authRequest *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error) {
	return s.repo.Authenticate(ctx, authRequest)
}
func (s *UmsService) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {
	return s.repo.GetUserById(ctx,req)
}	
func (s *UmsService) GetUsers(ctx context.Context, userRequest *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	return s.repo.GetUsers(ctx, userRequest	)
}

// func (s *UmsService) mustEmbedUnimplementedUserServiceServer() {
// 	// This method is required to implement the gRPC service interface.
// 	// It can be left empty if you are not using the generated code.
// }
