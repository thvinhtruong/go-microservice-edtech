package GrpcUserService

import (
	"context"
	"server/UserService/app/apperror"
	"server/UserService/app/business"
	"server/UserService/app/dto"

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
	var result *RegisterUserResponse
	user := dto.UserRequest{
		FullName: request.FullName,
		Password: request.Password,
		Phone:    request.Phone,
		Gender:   request.Gender,
	}

	record, err := business.Instance.RegisterUser(ctx, user)
	if err != nil {
		return &RegisterUserResponse{
			UserId:    -1,
			ErrorCode: apperror.GetCode(err),
		}, nil
	}

	err = copier.Copy(&result, &record)
	if err != nil {
		return &RegisterUserResponse{
			UserId:    -1,
			ErrorCode: apperror.GetCode(err),
		}, nil
	}

	return result, nil
}

func (s *ZUserServiceServer) mustEmbedUnimplementedUserServiceServer() {
}
