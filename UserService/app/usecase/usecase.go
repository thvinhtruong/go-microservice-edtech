package usecase

import (
	"SepFirst/UserService/app/usecase/dto"
	"context"
)

type UserService interface {
	RegisterUser(ctx context.Context, user dto.UserRequest) (int, error)
	GetUserById(ctx context.Context, id int) (dto.UserResponse, error)
	UpdateUserPassword(ctx context.Context, id int, password string) error
}

type AdminService interface {
	GetAllUsers(ctx context.Context) ([]dto.UserResponse, error)
	DeleteUser(ctx context.Context, id int) error
	BlockUser(ctx context.Context, id int) error
}
