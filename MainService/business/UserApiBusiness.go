package business

import GrpcUserService "SepFirst/MainService/GrpcClients/UserService"

func LoginUser(request *GrpcUserService.LoginUserRequest) *GrpcUserService.LoginUserResponse {
	return GrpcUserService.INSTANCE.LoginUser(request)
}

func LoginTutor(request *GrpcUserService.LoginTutorRequest) *GrpcUserService.LoginTutorResponse {
	return GrpcUserService.INSTANCE.LoginTutor(request)
}

func RegisterUser(request *GrpcUserService.RegisterUserRequest) *GrpcUserService.RegisterUserResponse {
	return GrpcUserService.INSTANCE.RegisterUser(request)
}

func RegisterTutor(request *GrpcUserService.RegisterTutorRequest) *GrpcUserService.RegisterTutorResponse {
	return GrpcUserService.INSTANCE.RegisterTutor(request)
}
