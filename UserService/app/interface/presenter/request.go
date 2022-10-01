package presenter

type UserRequest struct {
	Fullname string `json:"fullname"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
}

type UserAuthRequest struct {
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TutorRequest struct {
	Fullname string `json:"fullname"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
}

type TutorAuthRequest struct {
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AdminAuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
