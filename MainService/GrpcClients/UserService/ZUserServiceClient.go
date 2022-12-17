package GrpcUserService

import (
	"context"
	"fmt"
	"log"
	"server/MainService/config"

	"google.golang.org/grpc"
)

type ZUserServiceClient struct {
	innerClient UserServiceClient
}

var INSTANCE *ZUserServiceClient

func init() {
	INSTANCE = &ZUserServiceClient{}
	INSTANCE.Init()
}

func (c *ZUserServiceClient) Init() {
	configuration := config.GetInstance()
	target := fmt.Sprintf(
		"%v:%v",
		configuration.GetConfig(config.USER_SERVICE_HOST),
		configuration.GetConfig(config.USER_SERVICE_PORT),
	)
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
	log.Println("RegisterUser response: ", response)
	return response
}

func (c *ZUserServiceClient) RegisterTutor(request *RegisterTutorRequest) *RegisterTutorResponse {
	response, _ := c.innerClient.RegisterTutor(context.Background(), request)
	return response
}
