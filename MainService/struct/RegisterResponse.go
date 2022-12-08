package _struct

import (
	GrpcUserService "SepFirst/MainService/GrpcClients/UserService"
)

type RegisterUserResponse struct {
	UserId int32 `json:"UserId"`
}

func GetRegisterUserResponse(resp *GrpcUserService.RegisterUserResponse) RegisterUserResponse {
	return RegisterUserResponse{
		UserId: resp.UserId,
	}
}
