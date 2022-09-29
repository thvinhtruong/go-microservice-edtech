package business

import (
	"SepFirst/UserService/app/usecase/dto"
	"context"
)

type Business struct {
}

var Instance *Business

func init() {
	Instance = &Business{}
}

func (b *Business) RegisterUser(ctx context.Context, request dto.UserRequest) dto.UserResponse {
	return dto.UserResponse{
		UserId:    -1,
		ErrorCode: 0,
	}
}
