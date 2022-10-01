package presenter

import (
	"SepFirst/UserService/app/apperror"

	"github.com/gofiber/fiber/v2"
)

type AppCtx struct {
	FiberCtx *fiber.Ctx
}

//haha

type ErrorResponse struct {
	ErrorCode  int32  `json:"error_code"`
	ErrorMsg   string `json:"error_msg"`
	ErrorField string `json:"error_field"`
}

func (c *AppCtx) Response(httpCode int, ReceivedError error, data interface{}) error {
	var msg string
	if ReceivedError != nil {
		msg = ReceivedError.Error()
	} else {
		msg = "success"
	}

	field := getErrorField(ReceivedError)
	if field != "" {
		msg = "success"
	}
	code := apperror.GetCode(ReceivedError)

	c.FiberCtx.Set("Content-Type", "application/json")
	return c.FiberCtx.Status(httpCode).JSON(fiber.Map{
		"data": data,
		"error": ErrorResponse{
			ErrorCode:  code,
			ErrorMsg:   msg,
			ErrorField: field,
		},
	})
}
