package presenter

type UserRequest struct {
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
}

type UserAuthRequest struct {
	Password string `json:"password"`
}
