package dto

import "time"

type UserRequest struct {
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
}

type UserAuthRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	FullName    string    `json:"fullname"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Gender      string    `json:"gender"`
	DateCreated time.Time `json:"date_created"`
}

type TutorRequest struct {
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
}

type TutorAuthRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type TutorResponse struct {
	FullName    string    `json:"fullname"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Gender      string    `json:"gender"`
	DateCreated time.Time `json:"date_created"`
}
