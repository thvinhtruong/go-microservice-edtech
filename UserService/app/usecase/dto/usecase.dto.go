package dto

type UserRequest struct {
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Gender   string `json:"gender"`
	Password string `json:"password"`
}

type UserResponse struct {
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Gender   string `json:"gender"`
}

type UserAuthRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
