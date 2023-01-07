package main

import (
	"context"
	db "server/UserService/app/db/mysql/sqlc"
	GrpcUserService "server/UserService/app/grpc"
	"server/UserService/app/sqlconnection"
	"server/UserService/pkg/random"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRegisterTutor(t *testing.T) {
	dbconn := sqlconnection.GetDB()
	repository := db.NewRepository(dbconn)

	user := db.RegisterUserParams{
		Phone:    random.RandomPhone(),
		Password: "123456789",
		Fullname: "Bao",
		Gender:   "male",
	}

	userServiceServer := GrpcUserService.NewZUserServiceServer(repository)
	request := GrpcUserService.RegisterTutorRequest{
		Phone:    user.Phone,
		Password: user.Password,
		Fullname: user.Fullname,
		Gender:   user.Gender,
	}

	result, err := userServiceServer.RegisterTutor(context.Background(), &request)
	require.NoError(t, err)
	require.NotEmpty(t, result)

}
