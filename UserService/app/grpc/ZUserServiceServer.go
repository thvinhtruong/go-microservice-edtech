package GrpcUserService

import (
	"context"
	"server/UserService/app/apperror"
	db "server/UserService/app/db/mysql/sqlc"
)

type ZUserServiceServer struct {
	repository db.Repository
}

func NewZUserServiceServer(repository db.Repository) ZUserServiceServer {
	return ZUserServiceServer{repository: repository}
}

func (s *ZUserServiceServer) LoginUser(ctx context.Context, request *LoginUserRequest) (*LoginUserResponse, error) {
	var result *LoginUserResponse
	user := db.LoginUserParams{
		Phone:    request.Phone,
		Password: request.Password,
	}

	record, err := s.repository.LoginUser(ctx, user)
	if err != nil {
		return &LoginUserResponse{
			UserId:    -1,
			ErrorCode: apperror.GetCode(err),
		}, err
	}

	result.UserId = record.ID

	return result, nil
}

func (s *ZUserServiceServer) LoginTutor(ctx context.Context, request *LoginTutorRequest) (*LoginTutorResponse, error) {
	var result LoginTutorResponse

	user := db.LoginTutorParams{
		Phone:    request.Phone,
		Password: request.Password,
	}

	record, err := s.repository.LoginTutor(ctx, user)
	if err != nil {
		return &LoginTutorResponse{
			TutorId:   -1,
			ErrorCode: apperror.GetCode(err),
		}, err
	}

	result.TutorId = record.ID

	return &result, nil
}

func (s *ZUserServiceServer) RegisterTutor(ctx context.Context, request *RegisterTutorRequest) (*RegisterTutorResponse, error) {
	var result *RegisterTutorResponse
	user := db.RegisterTutorParams{
		Fullname: request.Fullname,
		Password: request.Password,
		Phone:    request.Phone,
		Gender:   request.Gender,
	}

	record, err := s.repository.RegisterTutor(ctx, user)
	if err != nil {
		return &RegisterTutorResponse{
			TutorId:   -1,
			ErrorCode: apperror.GetCode(err),
		}, err
	}

	result.TutorId = record.ID

	return result, nil

}
func (s *ZUserServiceServer) RegisterUser(ctx context.Context, request *RegisterUserRequest) (*RegisterUserResponse, error) {
	var result *RegisterUserResponse
	user := db.RegisterUserParams{
		Fullname: request.Fullname,
		Password: request.Password,
		Phone:    request.Phone,
		Gender:   request.Gender,
	}

	record, err := s.repository.RegisterUser(ctx, user)
	if err != nil {
		return &RegisterUserResponse{
			UserId:    -1,
			ErrorCode: apperror.GetCode(err),
		}, err
	}

	result.UserId = record.ID

	return result, nil
}

func (s *ZUserServiceServer) mustEmbedUnimplementedUserServiceServer() {
}
