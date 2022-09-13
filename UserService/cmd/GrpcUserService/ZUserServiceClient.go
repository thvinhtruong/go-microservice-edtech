package GrpcUserService

import (
	"SepFirst/UserService/utils"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

type ZUserServiceClient struct {
	innerClient UserServiceClient
}

func (c *ZUserServiceClient) Init() {
	target := fmt.Sprintf(
		"%v:%d",
		utils.Configuration.UserServiceServerHost,
		utils.Configuration.UserServiceServerPort)
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Connect to UserService failed")
		fmt.Printf("%v\n", err)
	}
	c.innerClient = NewUserServiceClient(conn)
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

func (c *ZUserServiceClient) RegisterTutor(request *RegisterUserRequest) *RegisterUserResponse {
	response, _ := c.innerClient.RegisterUser(context.Background(), request)
	return response
}
