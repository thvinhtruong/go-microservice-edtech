package grpc

import (
	"SepFirst/UserService/app/apperror"
	"SepFirst/UserService/app/interface/persistence/rdbms/sqlconnection"
	"SepFirst/UserService/app/registry"
	"SepFirst/UserService/app/usecase/dto"
	"context"
)

type ZUserServiceServer struct {
}

func (s *ZUserServiceServer) LoginUser(ctx context.Context, request *LoginUserRequest) (*LoginUserResponse, error) {

	return &LoginUserResponse{
		UserId:    -1,
		ErrorCode: apperror.GetCode(apperror.ErrorNotSupportedYet),
	}, nil
}

func (s *ZUserServiceServer) LoginTutor(ctx context.Context, request *LoginTutorRequest) (*LoginTutorResponse, error) {
	return &LoginTutorResponse{
		TutorId:   -1,
		ErrorCode: apperror.GetCode(apperror.ErrorNotSupportedYet),
	}, nil
}

func (s *ZUserServiceServer) RegisterTutor(ctx context.Context, request *RegisterTutorRequest) (*RegisterTutorResponse, error) {
	return &RegisterTutorResponse{
		TutorId:   -1,
		ErrorCode: apperror.GetCode(apperror.ErrorNotSupportedYet),
	}, nil
}
func (s *ZUserServiceServer) RegisterUser(ctx context.Context, request *RegisterUserRequest) (*RegisterUserResponse, error) {
	user := dto.UserRequest{
		FullName: "Tam",
		Email:    "Tam",
		Username: "Tam",
		Password: "Tam",
		Gender:   "Female",
	}
	access := registry.BuildUserAccessPoint(false, sqlconnection.DBConn)
	userId, err := access.Service.RegisterUser(ctx, user)
	if err != nil {

	}
	return &RegisterUserResponse{
		UserId:    int32(userId),
		ErrorCode: apperror.GetCode(apperror.ErrorNotSupportedYet),
	}, nil
}

func (s *ZUserServiceServer) mustEmbedUnimplementedUserServiceServer() {

}
