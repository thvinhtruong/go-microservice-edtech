package business

import (
	"SepFirst/UserService/config"
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
}

// func equalsRegisterUserResponse(got grpc.RegisterUserResponse, want grpc.RegisterUserResponse) bool {
// 	if got.ErrorCode != want.ErrorCode {
// 		return false
// 	}
// 	return true
// }
