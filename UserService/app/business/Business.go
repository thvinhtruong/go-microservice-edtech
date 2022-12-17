package business

import (
	"context"
	db "server/UserService/app/db/mysql/sqlc"
	"server/UserService/app/dto"
	"server/UserService/app/presenter"
	"server/UserService/app/registry"
	"server/UserService/app/sqlconnection"

	"github.com/jinzhu/copier"
)

type Business struct {
}

var Instance *Business

func init() {
	Instance = &Business{}
}

func (b *Business) RegisterUser(ctx context.Context, user dto.UserRequest) (presenter.UserResponse, error) {
	access := registry.BuildUserAccessPoint(sqlconnection.DBConn)
	var record db.RegisterUserParams
	err := copier.Copy(&record, &user)
	if err != nil {
		return presenter.UserResponse{}, err
	}
	result, err := access.Service.RegisterUser(ctx, record)
	if err != nil {
		return presenter.UserResponse{}, err
	}

	res := presenter.UserResponse{
		ID:    result.ID,
		Phone: result.Phone,
	}

	return res, nil

}
