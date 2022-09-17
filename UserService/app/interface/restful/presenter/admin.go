package presenter

import "github.com/gofiber/fiber"

type AdminAuthRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func BindAdminAuthRequest(c *fiber.Ctx) (AdminAuthRequest, error) {
	var admin AdminAuthRequest
	err := c.BodyParser(&admin)
	if err != nil {
		c.Status(400).SendString(err.Error())
		return admin, err
	}
	if err := ValidateUserAuth(admin.Username, admin.Email, admin.Password); err != nil {
		c.Status(400).SendString(err.Error())
		return admin, err
	}

	return admin, nil
}
