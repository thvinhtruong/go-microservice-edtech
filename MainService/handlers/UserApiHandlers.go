package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	GrpcUserService "server/MainService/GrpcClients/UserService"
	"server/MainService/business"
	_struct "server/MainService/struct"
	"server/MainService/utils"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	Username := r.FormValue("Username")
	Password := r.FormValue("Password")
	loginRequest := GrpcUserService.LoginUserRequest{
		Username: Username,
		Password: Password,
	}
	res := business.LoginUser(&loginRequest)
	out, _ := json.Marshal(res)
	w.Write(out)
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	PhoneNumber := r.FormValue("PhoneNumber")
	Password := r.FormValue("Password")
	FullName := r.FormValue("FullName")
	fmt.Printf("%v %v %v\n", PhoneNumber, Password, FullName)
	registerRequest := GrpcUserService.RegisterUserRequest{
		// Phone: PhoneNumber,
		Password: Password,
		FullName: FullName,
	}
	response := business.RegisterUser(&registerRequest)
	log.Println(response)
	message := _struct.ApiMessage{
		ErrorCode: response.ErrorCode,
		Message:   "success",
		Data:      utils.Convert(_struct.GetRegisterUserResponse(response)),
	}
	returnString := utils.Convert(message)

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	//Setting content type to json
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(returnString))

}
