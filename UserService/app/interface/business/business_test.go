package business

import (
	"SepFirst/UserService/app/usecase/dto"
	"SepFirst/UserService/config"
	"context"
	"fmt"
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
	fmt.Println(Instance)
	result, err := Instance.RegisterUser(context.Background(), dto.UserRequest{
		FullName: "cjanjcacssa",
		Password: "cjabnjscbsac",
		Email:    "cjabcjabca",
		Phone:    "197312371931",
		Gender:   "male",
	})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(result)

}
