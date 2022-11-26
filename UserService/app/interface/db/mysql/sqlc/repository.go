package db

import (
	"context"
	"database/sql"
)

type Repository interface {
	// User
	UserRepository
	// Admin
	// AdminRepository
	// Tutor
	// TutorRepository
}

type UserRepository interface {
	RegisterUser(ctx context.Context, arg RegisterUserParams) (RegisterUserResult, error)
	// GetUser(ctx context.Context, userId int32) (User, error)
	// ListUsers(ctx context.Context) ([]User, error)
	// UpdateUserInfo(ctx context.Context, arg UpdateUserInfoParams) (sql.Result, error)
	// UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) (sql.Result, error)
}

type AdminRepository interface {
	CreateAdmin(ctx context.Context) (sql.Result, error)
}

type TutorRepository interface {
	CreateTutor(ctx context.Context, arg CreateTutorParams) (sql.Result, error)
	CreateTutorPassword(ctx context.Context, arg CreateTutorPasswordParams) (sql.Result, error)
	GetTutor(ctx context.Context, tutorId int32) (Tutor, error)
	ListTutors(ctx context.Context) ([]Tutor, error)
	UpdateTutorInfo(ctx context.Context, arg UpdateTutorInfoParams) (sql.Result, error)
	UpdateTutorPassword(ctx context.Context, arg UpdateTutorPasswordParams) (sql.Result, error)
}
