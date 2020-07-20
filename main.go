package main

import (
	server "GRPCDemo/pkg/protocol/grpc/server"
	service "GRPCDemo/pkg/service/v1"
	"context"
	"fmt"
	"os"
)

func main() {
	ctx := context.Background()
	if err := server.RunServer(ctx, service.NewUserServiceServer()); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}
