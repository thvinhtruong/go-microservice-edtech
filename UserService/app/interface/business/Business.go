package business

import (
	"SepFirst/UserService/app/interface/presenter"
	"SepFirst/UserService/app/interface/sqlconnection"
	"SepFirst/UserService/app/registry"
	"SepFirst/UserService/app/usecase/dto"
	"context"
)

type Business struct {
}

var Instance *Business

func init() {
	Instance = &Business{}
}

func (b *Business) RegisterUser(ctx context.Context, user dto.UserRequest) (*presenter.Response, error) {
	access := registry.BuildUserAccessPoint(sqlconnection.DBConn)
	userId, err := access.Service.RegisterUser(ctx, user)
	if err != nil {
		return &presenter.Response{
			Data:   nil,
			Errors: presenter.GetErrorResponse(err),
		}, err
	}

	data := presenter.SetResponse(&presenter.Response{
		Data:   userId,
		Errors: presenter.GetErrorResponse(err),
	})

	return data, nil

}
