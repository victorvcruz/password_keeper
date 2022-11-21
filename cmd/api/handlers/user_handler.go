package handlers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"password_warehouse.com/internal/user/service"
)

type UserHandlerClient interface {
	CreateUser(*fiber.Ctx) error
}

type UserHandler struct {
}

func NewUser() UserHandlerClient {
	return &UserHandler{}
}

func (u *UserHandler) CreateUser(c *fiber.Ctx) error {
	var request service.UserRequest

	err := json.Unmarshal(c.Body(), &request)
	if err != nil {
		return err
	}

	return nil
}
