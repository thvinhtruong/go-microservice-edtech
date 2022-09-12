package main

import (
	GrpcUserService "SepFirst/UserService/cmd/GrpcUserService"
	"SepFirst/UserService/utils"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	utils.LoadConfig()

	host := fmt.Sprintf(
		"%v:%d",
		utils.Configuration.UserServiceServerHost,
		utils.Configuration.UserServiceServerHost)

	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("Failed to listen: %v", host)
	}

	s := grpc.NewServer()
	GrpcUserService.RegisterUserServiceServer(s, &GrpcUserService.ZUserServiceServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server %v", err)
	}
	fmt.Printf("ZUserServiceServer is listening at %v", host)
}
