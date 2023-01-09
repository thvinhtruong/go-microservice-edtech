package main

import (
	"context"
	"log"
	"os"
	db "server/UserService/app/db/mysql/sqlc"
	GrpcUserService "server/UserService/app/grpc"
	"server/UserService/app/sqlconnection"
	"server/UserService/pkg/random"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	dbconn := sqlconnection.GetDB()
	if dbconn == nil {
		panic("database is nil")
	}
	if err := dbconn.Ping(); err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func CreateRandomUser() db.RegisterTutorParams {
	return db.RegisterTutorParams{
		Phone:    random.RandomPhone(),
		Password: "123456789",
		Fullname: random.RandomName(),
		Gender:   random.RandomGender(),
		Topic:    random.RandomTopic(),
		City:     random.RandomCity(),
		Country:  random.RandomCountry(),
		Age:      int32(random.RandomAge()),
	}
}

func TestRegisterTutor(t *testing.T) {
	dbconn := sqlconnection.GetDB()
	repository := db.NewRepository(dbconn)

	for i := 0; i < 30; i++ {
		user := CreateRandomUser()
		userServiceServer := GrpcUserService.NewZUserServiceServer(repository)
		request := GrpcUserService.RegisterTutorRequest{
			Phone:    user.Phone,
			Password: user.Password,
			Fullname: user.Fullname,
			Gender:   user.Gender,
			Topic:    user.Topic,
			City:     user.City,
			Country:  user.Country,
			Age:      user.Age,
		}
		result, err := userServiceServer.RegisterTutor(context.Background(), &request)
		require.NoError(t, err)
		require.NotEmpty(t, result)
	}
}

func TestRegisterUser(t *testing.T) {
	dbconn := sqlconnection.GetDB()
	repository := db.NewRepository(dbconn)

	user := db.RegisterUserParams{
		Phone:    random.RandomPhone(),
		Password: "123456789",
		Fullname: "Thang",
		Gender:   random.RandomGender(),
	}

	userServiceServer := GrpcUserService.NewZUserServiceServer(repository)
	request := GrpcUserService.RegisterUserRequest{
		Phone:    user.Phone,
		Password: user.Password,
		Fullname: user.Fullname,
		Gender:   user.Gender,
	}

	result, err := userServiceServer.RegisterUser(context.Background(), &request)
	log.Println(result)
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestFilterTutors(t *testing.T) {
	dbconn := sqlconnection.GetDB()
	repository := db.NewRepository(dbconn)

	userServiceServer := GrpcUserService.NewZUserServiceServer(repository)
	request := GrpcUserService.FilterTutorRequest{
		Gender: "female",
		Topic:  "Math",
	}

	result, err := userServiceServer.FilterTutors(context.Background(), &request)
	log.Println(result)
	require.NoError(t, err)
	require.NotEmpty(t, result)
}
