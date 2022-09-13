package handlers

import (
	GrpcUserService "SepFirst/MainService/GrpcClients/UserService"
	"SepFirst/MainService/business"
	"encoding/json"
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
