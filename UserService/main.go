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
		"%v:%v",
		utils.Configuration.UserServiceServerHost,
		utils.Configuration.UserServiceServerPort)

	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("Failed to listen: %v", host)
	}

	s := grpc.NewServer()
	GrpcUserService.RegisterUserServiceServer(s, &GrpcUserService.ZUserServiceServer{})

	fmt.Printf("ZUserServiceServer is listening at %v", host)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server %v", err)
	}

}
