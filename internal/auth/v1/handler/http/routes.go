package http

import (
	"github.com/gofiber/fiber/v2"
	"main/internal/auth/v1/domain/ports"
)

func MapRoutes(h ports.IHandlers, router fiber.Router) {
	router.Get("/auth", func(c *fiber.Ctx) error {
		return c.JSON("auth")
	})
}
