package usecase

import (
	"SepFirst/UserService/app/usecase/dto"
	"context"
)

type UserService interface {
	RegisterUser(ctx context.Context, user dto.UserRequest) (int, error)
	UpdateUserPassword(ctx context.Context, id int, password string) error
}

type AdminService interface {
	DeleteUser(ctx context.Context, id int) error
	BlockUser(ctx context.Context, id int) error
}
