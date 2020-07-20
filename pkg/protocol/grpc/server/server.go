package grpc

import (
	v1 "GRPCDemo/pkg/api/v1"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

//RunServer implementation
func RunServer(ctx context.Context, userService v1.UserServiceServer) error {

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}

	//register Service
	server := grpc.NewServer()
	v1.RegisterUserServiceServer(server, userService)

	log.Println("Starting grpc server on 8080....")
	return server.Serve(listen)
}
