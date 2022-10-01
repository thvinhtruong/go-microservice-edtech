package grpc

import (
	"SepFirst/UserService/app/apperror"
	"SepFirst/UserService/app/interface/business"
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
		FullName: request.FullName,
		Password: request.Password,
		Email:    request.Email,
		Gender:   request.Gender,
		Username: request.Username,
	}
	return business.Instance.RegisterUser(ctx, user), nil
}

func (s *ZUserServiceServer) mustEmbedUnimplementedUserServiceServer() {

}
