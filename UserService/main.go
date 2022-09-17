package main

import (
	GrpcUserService "SepFirst/UserService/cmd/GrpcUserService"
	config "SepFirst/UserService/config"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	configuration := config.GetInstance()
	host := fmt.Sprintf(
		"%v:%v",
		configuration.GetConfig(config.USER_SERVICE_HOST),
		configuration.GetConfig(config.USER_SERVICE_PORT))

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
