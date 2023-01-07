package main

import (
	"fmt"
	"log"
	"net/http"
	GrpcUserService "server/MainService/GrpcClients/UserService"
	"server/MainService/config"
	"server/MainService/handlers"
)

func main() {
	configuration := config.GetInstance()
	userApiHandler := handlers.NewUserApiHanlder(configuration, &GrpcUserService.Instance)
	http.HandleFunc("/api/LoginUser", userApiHandler.LoginUser)
	http.HandleFunc("/api/RegisterUser", userApiHandler.RegisterUser)

	port := configuration.GetConfig(config.MAIN_SERVICE_PORT)

	fmt.Printf(
		"Server is listening at %v:%v",
		configuration.GetConfig(config.MAIN_SERVICE_HOST),
		configuration.GetConfig(config.MAIN_SERVICE_PORT),
	)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
