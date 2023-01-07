package db

type RegisterUserParams struct {
	Fullname string `json:"fullname"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
}

type RegisterUserResult struct {
	ID int32
}

type LoginUserParams struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type LoginUserResult struct {
	ID int32
}

type UpdateUserParams struct {
	ID       int32  `json:"id"`
	Fullname string `json:"fullname"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
}

type RegisterTutorParams struct {
	Fullname string `json:"fullname"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
}

type RegisterTutorResult struct {
	ID int32
}

type LoginTutorParams struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type LoginTutorResult struct {
	ID int32
}
