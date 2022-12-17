package usecase

import (
	"context"
	repository "server/UserService/app/db/mysql/sqlc"
	"server/UserService/app/dto"
)

type UserService interface {
	RegisterUser(ctx context.Context, user repository.RegisterUserParams) (dto.UserResponse, error)
	//UpdateUserPassword(ctx context.Context, id int, password string) error
}

type AdminService interface {
	DeleteUser(ctx context.Context, id int) error
	BlockUser(ctx context.Context, id int) error
}
