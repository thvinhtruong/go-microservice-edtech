package grpc

import (
	"SepFirst/UserService/app/apperror"
	"SepFirst/UserService/app/interface/business"
	"SepFirst/UserService/app/usecase/dto"
	"context"

	"github.com/jinzhu/copier"
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
	}

	user_response, err := business.Instance.RegisterUser(ctx, user)
	if err != nil {
		return &RegisterUserResponse{
			UserId:    -1,
			ErrorCode: apperror.GetCode(err),
		}, nil
	}

	var result RegisterUserResponse

	err = copier.Copy(&user_response, &result)
	if err != nil {
		return &RegisterUserResponse{
			UserId:    -1,
			ErrorCode: apperror.GetCode(err),
		}, nil
	}

	return &result, nil
}

func (s *ZUserServiceServer) mustEmbedUnimplementedUserServiceServer() {

}
