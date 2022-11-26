package business

import (
	repository "SepFirst/UserService/app/interface/db/mysql/sqlc"
	"SepFirst/UserService/app/interface/sqlconnection"
	"SepFirst/UserService/app/registry"
	"SepFirst/UserService/app/usecase/dto"
	"context"

	"github.com/jinzhu/copier"
)

type Business struct {
}

var Instance *Business

func init() {
	Instance = &Business{}
}

func (b *Business) RegisterUser(ctx context.Context, user dto.UserRequest) (dto.UserResponse, error) {
	access := registry.BuildUserAccessPoint(sqlconnection.DBConn)
	var record repository.RegisterUserParams
	err := copier.Copy(&record, &user)
	if err != nil {
		return dto.UserResponse{}, err
	}
	result, err := access.Service.RegisterUser(ctx, record)
	if err != nil {
		return dto.UserResponse{}, err
	}

	return result, nil
}
