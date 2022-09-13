package main

import (
	"SepFirst/MainService/handlers"
	"SepFirst/MainService/utils"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/LoginUser", handlers.LoginUser)

	port := utils.Configuration.MainServiceServerPort

	fmt.Printf(
		"Server is listening at %v:%v",
		utils.Configuration.MainServiceServerHost,
		utils.Configuration.MainServiceServerPort)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
