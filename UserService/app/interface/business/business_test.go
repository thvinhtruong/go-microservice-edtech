package business

import (
	"SepFirst/UserService/app/interface/grpc"
	"SepFirst/UserService/app/usecase/dto"
	"SepFirst/UserService/config"
	"context"
	"testing"
)

type TConfig struct {
}

func (t *TConfig) GetConfig(key config.ConfigKey) interface{} {
	if key == config.MYSQL_USERNAME {
		return "SepFi"
	} else if key == config.MYSQL_PASSWORD {
		return "SepFi"
	} else if key == config.MYSQL_HOST {
		return "localhost"
	} else if key == config.MYSQL_PORT {
		return "3306"
	}
	return nil
}

func TestBusiness_RegisterUser(t *testing.T) {
	got := Instance.RegisterUser(context.Background(), dto.UserRequest{
		FullName: "Tam",
		Password: "Tam",
		Email:    "Tam",
		Username: "Tam",
		Gender:   "Female",
	})
	want := grpc.RegisterUserResponse{
		ErrorCode: 0,
		UserId:    1,
	}
	if !equalsRegisterUserResponse(*got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func equalsRegisterUserResponse(got grpc.RegisterUserResponse, want grpc.RegisterUserResponse) bool {
	if got.ErrorCode != want.ErrorCode {
		return false
	}
	return true
}
