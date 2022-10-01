package handlers

import (
	"SepFirst/UserService/app/interface/persistence/rdbms/sqlconnection"
	"SepFirst/UserService/app/interface/restful/presenter"
	"SepFirst/UserService/app/registry"
	"SepFirst/UserService/app/usecase/dto"
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

func UserHandler(app fiber.Router) {
	app.Post("/register", registerUser)
	app.Get("/:userId", getUserByID)
	app.Patch("/password/:userId", changePassword)
}

func registerUser(c *fiber.Ctx) error {
	app := presenter.AppCtx{FiberCtx: c}
	ctx := context.Background()
	data, err := presenter.BindUserRequest(c, true)
	if err != nil {
		return err
	}

	log.Println("data: ", data)

	var record dto.UserRequest
	err = copier.Copy(&record, &data)
	if err != nil {
		return app.Response(http.StatusInternalServerError, err, nil)
	}

	access := registry.BuildUserAccessPoint(false, sqlconnection.DBConn)

	id, err := access.Service.RegisterUser(ctx, record)
	if err != nil {
		return app.Response(http.StatusInternalServerError, err, nil)
	}

	return app.Response(200, err, id)
}

func getUserByID(c *fiber.Ctx) error {
	app := presenter.AppCtx{FiberCtx: c}
	ctx := context.Background()
	id, err := strconv.Atoi(c.Params("userId"))
	if err != nil {
		return app.Response(http.StatusBadRequest, err, nil)
	}

	access := registry.BuildUserAccessPoint(false, sqlconnection.DBConn)

	user, err := access.Service.GetUserById(ctx, id)
	if err != nil {
		return app.Response(http.StatusInternalServerError, err, nil)
	}

	return app.Response(200, err, user)
}

func changePassword(c *fiber.Ctx) error {
	app := presenter.AppCtx{FiberCtx: c}
	ctx := context.Background()
	id, err := strconv.Atoi(c.Params("userId"))
	if err != nil {
		return app.Response(http.StatusBadRequest, err, nil)
	}

	data, err := presenter.BindUserRequest(c, false)
	if err != nil {
		return app.Response(http.StatusBadRequest, err, nil)
	}
	if err != nil {
		return app.Response(http.StatusInternalServerError, err, nil)
	}

	access := registry.BuildUserAccessPoint(false, sqlconnection.DBConn)

	err = access.Service.UpdateUserPassword(ctx, id, data.Password)
	if err != nil {
		return app.Response(http.StatusInternalServerError, err, nil)
	}

	return app.Response(200, err, nil)
}
