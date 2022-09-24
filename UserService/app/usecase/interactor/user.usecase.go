package interactor

import (
	entity "SepFirst/UserService/app/domain/entities"
	"SepFirst/UserService/app/domain/repository"
	"SepFirst/UserService/app/usecase/dto"
	"context"

	"github.com/jinzhu/copier"
)

type UserUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.Repository) *UserUsecase {
	return &UserUsecase{userRepo: userRepo}
}

func (u *UserUsecase) RegisterUser(ctx context.Context, user dto.UserRequest) (int, error) {
	var record entity.User
	err := copier.Copy(&record, &user)
	if err != nil {
		return 0, err
	}

	return u.userRepo.CreateUser(ctx, record, user.Password)
}

func (u *UserUsecase) GetUserById(ctx context.Context, id int) (dto.UserResponse, error) {
	var record dto.UserResponse
	user, err := u.userRepo.GetUserById(ctx, id)
	if err != nil {
		return record, err
	}

	err = copier.Copy(&record, &user)
	if err != nil {
		return record, err
	}

	return record, nil
}

func (u *UserUsecase) UpdateUserPassword(ctx context.Context, userId int, newPassword string) error {
	return u.userRepo.UpdateUserPassword(ctx, userId, newPassword)
}
