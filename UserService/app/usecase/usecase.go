package usecase

import (
	repository "SepFirst/UserService/app/interface/db/mysql/sqlc"
	"SepFirst/UserService/app/usecase/dto"
	"context"
)

type UserService interface {
	RegisterUser(ctx context.Context, user repository.RegisterUserParams) (dto.UserResponse, error)
	//UpdateUserPassword(ctx context.Context, id int, password string) error
}

type AdminService interface {
	DeleteUser(ctx context.Context, id int) error
	BlockUser(ctx context.Context, id int) error
}
