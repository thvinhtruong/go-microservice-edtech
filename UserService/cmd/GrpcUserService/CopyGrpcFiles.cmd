set container=../../../%1/GrpcClients/UserService/
if not exist "%container%" mkdir "%container%"
copy "UserService.pb.go" "%container%UserService.pb.go"
copy "UserService_grpc.pb.go" "%container%UserService_grpc.pb.go"
copy "ZUserServiceClient.go" "%container%ZUserServiceClient.go"
