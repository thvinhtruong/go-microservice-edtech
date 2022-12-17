package repository

import (
	"context"
	"database/sql"
	db "server/UserService/app/db/mysql/sqlc"
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
	RegisterUser(ctx context.Context, arg db.RegisterUserParams) (db.RegisterUserResult, error)
	// GetUser(ctx context.Context, userId int32) (User, error)
	// ListUsers(ctx context.Context) ([]User, error)
	// UpdateUserInfo(ctx context.Context, arg UpdateUserInfoParams) (sql.Result, error)
	// UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) (sql.Result, error)
}

type AdminRepository interface {
	CreateAdmin(ctx context.Context) (sql.Result, error)
}

type TutorRepository interface {
	CreateTutor(ctx context.Context, arg db.CreateTutorParams) (sql.Result, error)
	CreateTutorPassword(ctx context.Context, arg db.CreateTutorPasswordParams) (sql.Result, error)
	GetTutor(ctx context.Context, tutorId int32) (db.Tutor, error)
	ListTutors(ctx context.Context) ([]db.Tutor, error)
	UpdateTutorInfo(ctx context.Context, arg db.UpdateTutorInfoParams) (sql.Result, error)
	UpdateTutorPassword(ctx context.Context, arg db.UpdateTutorPasswordParams) (sql.Result, error)
}
