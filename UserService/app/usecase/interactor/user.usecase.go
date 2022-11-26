package interactor

import (
	repository "SepFirst/UserService/app/interface/db/mysql/sqlc"
	"SepFirst/UserService/app/usecase/dto"
	"context"

	"github.com/jinzhu/copier"
)

type UserUsecase struct {
	userRepo repository.Repository
}

func NewUserUsecase(userRepo repository.Repository) *UserUsecase {
	return &UserUsecase{userRepo: userRepo}
}

// Transaction
// create User and get user id generated
// create password based on the generated user id
func (u *UserUsecase) RegisterUser(ctx context.Context, user repository.RegisterUserParams) (dto.UserResponse, error) {
	var result dto.UserResponse

	// create user
	record, err := u.userRepo.RegisterUser(ctx, user)
	if err != nil {
		return dto.UserResponse{}, err
	}

	err = copier.Copy(&result, &record)
	if err != nil {
		return dto.UserResponse{}, err
	}

	return result, nil
}
