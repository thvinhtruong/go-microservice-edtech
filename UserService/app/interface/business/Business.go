package business

import (
	"SepFirst/UserService/app/interface/grpc"
	"SepFirst/UserService/app/interface/persistence/rdbms/sqlconnection"
	"SepFirst/UserService/app/registry"
	"SepFirst/UserService/app/usecase/dto"
	"context"
)

type Business struct {
}

var Instance *Business

func init() {
	Instance = &Business{}
}

func (b *Business) RegisterUser(ctx context.Context, request *grpc.RegisterUserRequest) *grpc.RegisterUserResponse {
	access := registry.BuildUserAccessPoint(false, sqlconnection.DBConn)
	user := dto.UserRequest{
		FullName: request.FullName,
		Email:    request.Email,
		Password: request.Password,
		Gender:   request.Gender,
		Username: request.Password,
	}
	userId, err := access.Service.RegisterUser(ctx, user)
	if err != nil {
		return &grpc.RegisterUserResponse{
			ErrorCode: -1,
			UserId:    -1,
		}
	}
	return &grpc.RegisterUserResponse{
		ErrorCode: 0,
		UserId:    int32(userId),
	}
}
