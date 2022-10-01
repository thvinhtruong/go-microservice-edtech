package handlers

import (
	GrpcUserService "SepFirst/MainService/GrpcClients/UserService"
	"SepFirst/MainService/business"
	"SepFirst/MainService/errors"
	_struct "SepFirst/MainService/struct"
	"SepFirst/MainService/utils"
	"encoding/json"
	"fmt"
	"net/http"
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
	message := _struct.ApiMessage{
		ErrorCode: errors.UNSUPPORTED_YET,
		Message:   "",
		Data:      "",
	}
	returnString := utils.Convert(message)
	
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	fmt.Fprintf(w, returnString)
}
