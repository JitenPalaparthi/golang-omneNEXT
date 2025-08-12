package handlers

import (
	"time"
	"users-service/database"
	"users-service/models"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	database.IUserDB // prmoted field
}

type IUserHandler interface {
	CreateUser(c *fiber.Ctx) error
}

func NewUserHandler(iuserdb database.IUserDB) IUserHandler {
	return &UserHandler{iuserdb}
}

func (uh *UserHandler) CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	err := c.BodyParser(user)
	if err != nil {
		return err
	}

	err = user.Validate()
	if err != nil {
		return err
	}

	user.Status = "active"
	user.LastModified = time.Now().Unix()

	user, err = uh.Create(user)
	if err != nil {
		return err
	}
	return c.JSON(user)
}
