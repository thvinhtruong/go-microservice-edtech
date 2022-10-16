package interactor

import (
	repository "SepFirst/UserService/app/interface/db/mysql/sqlc"
	"SepFirst/UserService/app/usecase/dto"
	"SepFirst/UserService/pkg/hasher"
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
func (u *UserUsecase) RegisterUser(ctx context.Context, user dto.UserRequest) (dto.UserResponse, error) {
	var user_record repository.CreateUserParams
	var password_record repository.CreateUserPasswordParams
	var user_response dto.UserResponse
	err := copier.Copy(&user_record, &user)
	if err != nil {
		return dto.UserResponse{}, err
	}

	err = u.userRepo.EnableTx(func() error {
		user_result, err := u.userRepo.CreateUser(ctx, user_record)
		if err != nil {
			return err
		}

		user_id, err := user_result.LastInsertId()
		if err != nil {
			return err
		}

		hashed, err := hasher.HashPassword(user.Password)
		if err != nil {
			return err
		}

		password_record.UserID = int32(user_id)
		password_record.Password = hashed

		err = u.userRepo.CreateUserPassword(ctx, password_record)
		if err != nil {
			return err
		}

		err = copier.Copy(&user_response, &user_result)
		if err != nil {
			return err
		}

		return nil
	})

	return user_response, nil
}
