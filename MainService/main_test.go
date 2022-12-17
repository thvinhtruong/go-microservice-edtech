package main

import (
	GrpcUserService "server/MainService/GrpcClients/UserService"
	"server/MainService/business"
	"testing"
)

func setupSuite(tb testing.TB) func(tb testing.TB) {
	return func(tb testing.TB) {
	}
}

func TestLoginUser(t *testing.T) {
	t.Logf("TestLoginUser\n")
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	got := business.LoginUser(&GrpcUserService.LoginUserRequest{
		Phone:    "918231893713",
		Password: "Bao Yeu Tam",
	})
	want := GrpcUserService.LoginUserResponse{
		UserId:    -1,
		ErrorCode: 30005,
	}

	if !equalsLoginUserResponse(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestLoginTutor(t *testing.T) {
	t.Logf("TestLoginTutor\n")
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	got := business.LoginTutor(&GrpcUserService.LoginTutorRequest{
		Phone:    "918231893713",
		Password: "Tam",
	})
	want := GrpcUserService.LoginTutorResponse{
		TutorId:   -1,
		ErrorCode: 30005,
	}

	if !equalsLoginTutorResponse(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRegisterUser(t *testing.T) {
	t.Logf("TestRegisterUser\n")
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	got := business.RegisterUser(&GrpcUserService.RegisterUserRequest{
		FullName: "Bao",
		Phone:    "183971937121",
		Password: "Tam",
		Blocked:  true,
		Gender:   "Male",
	})
	want := GrpcUserService.RegisterUserResponse{
		UserId:    -1,
		ErrorCode: 30005,
	}

	if !equalsRegisterUserResponse(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRegisterTutor(t *testing.T) {
	t.Logf("TestRegisterTutor\n")
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	got := business.RegisterTutor(&GrpcUserService.RegisterTutorRequest{
		Phone:    "1987392137",
		Password: "Tam",
		FullName: "Bao",
		Blocked:  true,
		Gender:   "Male",
	})
	want := GrpcUserService.RegisterTutorResponse{
		TutorId:   -1,
		ErrorCode: 30005,
	}

	if !equalsRegisterTutorResponse(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
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
