package main

import (
	v1 "GRPCDemo/pkg/api/v1"
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	address := flag.String("server", "", "gRPC server in format host:port")
	email := flag.String("email", "", "Email of the user")
	firstName := flag.String("FirstName", "", "First name of the user")
	lastName := flag.String("LastName", "", "Last name of the user")
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := v1.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Call Create
	req1 := v1.UserRequest{
		User: &v1.User{
			Email:     *email,
			Firstname: *firstName,
			Lastname:  *lastName,
		},
	}
	res1, err := c.Registeruser(ctx, &req1)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Response: <%+v>\n", res1)
}
