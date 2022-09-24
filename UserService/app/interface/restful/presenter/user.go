package presenter

import "github.com/gofiber/fiber/v2"

type UserRequest struct {
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
}

type UserAuthRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func BindUserRequest(c *fiber.Ctx, withRegister bool) (UserRequest, error) {
	var user UserRequest
	err := c.BodyParser(&user)
	if err != nil {
		c.Status(400).SendString(err.Error())
		return user, err
	}
	if withRegister {
		if err := ValidateUserEmptyField(user); err != nil {
			c.Status(400).SendString(err.Error())
			return user, err
		}
	}

	return user, nil
}

func BindUserAuthRequest(c *fiber.Ctx) (UserAuthRequest, error) {
	var user UserAuthRequest
	err := c.BodyParser(&user)
	if err != nil {
		c.Status(400).SendString(err.Error())
		return user, err
	}
	if err := ValidateUserAuth(user.Username, user.Email, user.Password); err != nil {
		c.Status(400).SendString(err.Error())
		return user, err
	}

	return user, nil
}
