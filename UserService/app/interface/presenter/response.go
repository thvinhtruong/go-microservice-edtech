package presenter

import "sync"

var (
	once      sync.Once
	singleton *Response
)

type Response struct {
	Data   interface{}   `json:"data"`
	Errors ErrorResponse `json:"errors"`
}

func SetResponse(res *Response) *Response {
	once.Do(func() {
		singleton = res
	})
	return singleton
}
