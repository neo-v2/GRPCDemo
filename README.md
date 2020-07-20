# GRPCDemo
This is a simple GRPC Demonstration

### To Run GRPC Server
From terminal - go run GRPCDemo

### To Run GRPC Client
cd GRPCDemo/pkg/protocol/grpc/client
go buid .  
./client --server=localhost:8080 --FirstName=TestUser1 --LastName=TestUser2 --email=a@me.com

### Response
Response: <msg:"Welcome TestUser1"  token:"a5c6b45d-6a3a-4c61-9481-f44017f3fa19">
