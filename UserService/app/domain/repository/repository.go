package repository

import (
	entity "SepFirst/UserService/app/domain/entities"
	"context"
)

type Repository interface {
	// User
	UserRepository
	// Admin
	AdminRepository
}

type UserRepository interface {
	CreateUser(ctx context.Context, user entity.User, password string) (int, error)
	GetUserById(ctx context.Context, userId int) (entity.User, error)
	UpdateUserPassword(ctx context.Context, userId int, password string) error
}

type AdminRepository interface {
	GetAllUsers(ctx context.Context) ([]entity.User, error)
	DeleteUser(ctx context.Context, userId int) error
	BlockUser(ctx context.Context, userId int) error
}
