package main

import (
	"SepFirst/MainService/config"
	"SepFirst/MainService/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {

	configuration := config.GetInstance()
	http.HandleFunc("/api/LoginUser", handlers.LoginUser)

	port := configuration.GetConfig(config.MAIN_SERVICE_PORT)

	fmt.Printf(
		"Server is listening at %v:%v",
		configuration.GetConfig(config.MAIN_SERVICE_HOST),
		configuration.GetConfig(config.MAIN_SERVICE_PORT),
	)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
