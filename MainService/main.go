package main

import (
	"fmt"
	"log"
	"net/http"
	"server/MainService/config"
	"server/MainService/handlers"
)

func main() {

	configuration := config.GetInstance()
	http.HandleFunc("/api/LoginUser", handlers.LoginUser)
	http.HandleFunc("/api/RegisterUser", handlers.RegisterUser)

	port := configuration.GetConfig(config.MAIN_SERVICE_PORT)

	fmt.Printf(
		"Server is listening at %v:%v",
		configuration.GetConfig(config.MAIN_SERVICE_HOST),
		configuration.GetConfig(config.MAIN_SERVICE_PORT),
	)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
