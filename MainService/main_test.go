package main

import (
	"fmt"
	GrpcUserService "server/MainService/GrpcClients/UserService"
	"server/MainService/config"
	"server/MainService/handlers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupSuite(tb testing.TB) func(tb testing.TB) {
	return func(tb testing.TB) {
	}
}

func TestRegisterUser(t *testing.T) {
	t.Logf("TestRegisterUser\n")
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	configuration := config.GetInstance()
	userHandler := handlers.NewUserApiHanlder(configuration, &GrpcUserService.Instance)

	fmt.Printf(
		"Server is listening at %v:%v",
		configuration.GetConfig(config.MAIN_SERVICE_HOST),
		configuration.GetConfig(config.MAIN_SERVICE_PORT),
	)

	got := userHandler.Repo.RegisterUser(&GrpcUserService.RegisterUserRequest{
		FullName: "Bao",
		Phone:    "18397193712",
		Password: "Tam",
		Gender:   "Male",
	})

	fmt.Println(got)

	assert.NotNil(t, got)
}

func equalsLoginUserResponse(got *GrpcUserService.LoginUserResponse, want GrpcUserService.LoginUserResponse) bool {
	if got == nil || got.UserId != want.UserId || got.ErrorCode != want.ErrorCode {
		return false
	}
	return true
}

func equalsLoginTutorResponse(got *GrpcUserService.LoginTutorResponse, want GrpcUserService.LoginTutorResponse) bool {
	if got == nil || got.TutorId != want.TutorId || got.ErrorCode != want.ErrorCode {
		return false
	}
	return true
}

func equalsRegisterUserResponse(got *GrpcUserService.RegisterUserResponse, want GrpcUserService.RegisterUserResponse) bool {
	if got == nil || got.UserId != want.UserId || got.ErrorCode != want.ErrorCode {
		return false
	}
	return true
}

func equalsRegisterTutorResponse(got *GrpcUserService.RegisterTutorResponse, want GrpcUserService.RegisterTutorResponse) bool {
	if got == nil || got.TutorId != want.TutorId || got.ErrorCode != want.ErrorCode {
		return false
	}
	return true
}
