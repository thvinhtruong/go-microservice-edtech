package presenter

import "github.com/google/go-cmp/cmp"

type UserResponse struct {
	ID    int32  `json:"id"`
	Phone string `json:"phone"`
}

type Response struct {
	Data   interface{}   `json:"data"`
	Errors ErrorResponse `json:"errors"`
}

// IsEmpty check if the struct is empty or not
func (r Response) IsEmpty() bool {
	return cmp.Equal(r, Response{})
}
