package main

import (
	"fmt"
	"log"
	"net"
	grpc2 "server/UserService/app/grpc"
	config "server/UserService/config"

	"google.golang.org/grpc"
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
	grpc2.RegisterUserServiceServer(s, &grpc2.ZUserServiceServer{})

	fmt.Printf("ZUserServiceServer is listening at %v", host)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server %v", err)
	}

}
