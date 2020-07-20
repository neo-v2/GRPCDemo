package v1

import (
	v1 "GRPCDemo/pkg/api/v1"
	"context"

	uuid "github.com/satori/go.uuid"
)

type userServiceServer struct{}

//NewUserServiceServer implementation
func NewUserServiceServer() v1.UserServiceServer {
	return &userServiceServer{}
}
func genUUIDv4() string {
	id, _ := uuid.NewV4()
	return id.String()
}
func (s *userServiceServer) Registeruser(ctx context.Context, req *v1.UserRequest) (*v1.UserResponse, error) {
	return &v1.UserResponse{
		Msg: "Welcome " + req.GetUser().Firstname,

		Token: genUUIDv4(),
	}, nil
}
