package handlers
import ( 
	"context"
	"app/internal/services" 
	pb "app/internal/genproto/users"  
	"google.golang.org/grpc"
)
 
type umsgRPCHandler struct {
	service *services.UmsService
	// unimplementedOrderServiceServer pb.UnimplementedOrderServiceServer
}	
func NewGrpcUmsHandler(grpcServer *grpc.Server, service *services.UmsService)  *umsgRPCHandler {
	h := &umsgRPCHandler{service: service}
	// Register the handler with the gRPC server

	pb.RegisterUserServiceServer(grpcServer, h)	
	return h
}	

func (h *umsgRPCHandler) Authenticate( ctx context.Context, userRequest *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error) {

	return h.service.Authenticate(ctx, userRequest)
}

func (h *umsgRPCHandler) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {
 
	return h.service.GetUserById(ctx, req)
}	

func (h *umsgRPCHandler) GetUsers(ctx context.Context, userRequest *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {

	return  h.service.GetUsers(ctx, userRequest)
}

func (h *umsgRPCHandler) mustEmbedUnimplementedUserServiceServer() {
	// This method is required to implement the gRPC service interface.
	// It can be left empty if you are not using the generated code.
}
