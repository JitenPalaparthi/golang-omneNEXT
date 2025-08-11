package handlers

import (
	"demo/models"
	"demo/utils"
	"log"
	"math/rand/v2"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	FileName string
}

func NewUserHandler(fileName string) *UserHandler {
	return &UserHandler{fileName}
}

func (uh *UserHandler) Create(c *fiber.Ctx) error {
	user := new(models.User)
	//c.Request().Body()
	err := c.BodyParser(user)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	if err := user.Validate(); err != nil {
		return err
	}

	user.Id = uint(rand.IntN(10000))
	user.Status = "active"
	user.LastModified = time.Now().Unix()
	bytes, err := user.ToBytes()
	if err != nil {
		return err
	}
	utils.ChUser <- bytes
	return nil
}
