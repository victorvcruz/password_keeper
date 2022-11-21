package handlers

import "github.com/gofiber/fiber/v2"

type AuthHandlerClient interface {
	Auth(*fiber.Ctx) error
}

type AuthHandler struct {
}

func NewAuth() AuthHandlerClient {
	return &AuthHandler{}
}

func (u *AuthHandler) Auth(*fiber.Ctx) error {

	return nil
}
