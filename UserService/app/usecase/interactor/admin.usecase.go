package interactor

import (
	"SepFirst/UserService/app/domain/repository"
	"SepFirst/UserService/app/usecase/dto"
	"context"
)

type AdminUsecase struct {
	adminRepo repository.AdminRepository
}

func NewAdminUsecase(adminRepo repository.Repository) *AdminUsecase {
	return &AdminUsecase{adminRepo: adminRepo}
}

func (a *AdminUsecase) GetAllUsers(ctx context.Context) ([]dto.UserResponse, error) {
	var results []dto.UserResponse
	users, err := a.adminRepo.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		results = append(results, dto.UserResponse{
			FullName: user.FullName,
			Email:    user.Email,
			Username: user.Username,
			Gender:   user.Gender,
		})
	}
	return results, nil
}

func (a *AdminUsecase) DeleteUser(ctx context.Context, userId int) error {
	return a.adminRepo.DeleteUser(ctx, userId)
}

func (a *AdminUsecase) BlockUser(ctx context.Context, userId int) error {
	return a.adminRepo.BlockUser(ctx, userId)
}
