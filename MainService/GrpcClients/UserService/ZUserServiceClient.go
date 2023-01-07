package GrpcUserService

import (
	"context"
	"fmt"
	"log"
	"server/MainService/config"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ZUserServiceClient struct {
	innerClient UserServiceClient
	Config      config.Config
}

var Instance ZUserServiceClient

func init() {
	Config := config.GetInstance()
	target := fmt.Sprintf(
		"%v:%v",
		Config.GetConfig(config.USER_SERVICE_HOST),
		Config.GetConfig(config.USER_SERVICE_PORT),
	)
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Connect to UserService failed")
		log.Fatal(err)
	}

	fmt.Println("Connect to UserService success")
	innerClient := NewUserServiceClient(conn)
	Instance = ZUserServiceClient{
		innerClient: innerClient,
		Config:      Config,
	}
}

func (c *ZUserServiceClient) LoginTutor(request *LoginTutorRequest) *LoginTutorResponse {
	response, _ := c.innerClient.LoginTutor(context.Background(), request)
	return response
}

func (c *ZUserServiceClient) LoginUser(request *LoginUserRequest) *LoginUserResponse {
	response, _ := c.innerClient.LoginUser(context.Background(), request)
	return response
}

func (c *ZUserServiceClient) RegisterUser(request *RegisterUserRequest) *RegisterUserResponse {
	response, _ := c.innerClient.RegisterUser(context.Background(), request)
	return response
}

func (c *ZUserServiceClient) RegisterTutor(request *RegisterTutorRequest) *RegisterTutorResponse {
	response, _ := c.innerClient.RegisterTutor(context.Background(), request)
	return response
}
