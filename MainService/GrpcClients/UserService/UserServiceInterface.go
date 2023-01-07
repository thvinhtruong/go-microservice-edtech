package GrpcUserService

type UserServiceRepository interface {
	LoginUser(request *LoginUserRequest) *LoginUserResponse
	RegisterUser(request *RegisterUserRequest) *RegisterUserResponse
	LoginTutor(request *LoginTutorRequest) *LoginTutorResponse
	RegisterTutor(request *RegisterTutorRequest) *RegisterTutorResponse
}
