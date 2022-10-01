package interactor

import (
	"SepFirst/UserService/app/domain/repository"
	db "SepFirst/UserService/app/interface/db/mysql/sqlc"
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
func (u *UserUsecase) RegisterUser(ctx context.Context, user dto.UserRequest) (int, error) {
	var user_record db.CreateUserParams
	var password_record db.CreateUserPasswordParams
	err := copier.Copy(&user_record, &user)
	if err != nil {
		return 0, err
	}

	hashed, err := hasher.HashPassword(user.Password)
	if err != nil {
		return 0, err
	}

	password_record.Password = hashed

	return u.userRepo.EnableTx(ctx, func(u.userRepo) error {
		user_id, err := u.userRepo.CreateUser(ctx, user_record)
		if err != nil {
			return err
		}

		id, err := user_id.LastInsertId()
		if err != nil {
			return err
		}

		password_record.UserID = int32(id)

		_, err = u.userRepo.CreatePassword(ctx, password_record)
		if err != nil {
			return err
		}

		return nil
	})
}

func (u *UserUsecase) UpdateUserPassword(ctx context.Context, userId int, newPassword string) error {
	return u.userRepo.UpdateUserPassword(ctx, userId, newPassword)
}
