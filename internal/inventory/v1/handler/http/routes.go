package http

import (
	"github.com/gofiber/fiber/v2"
	"main/internal/inventory/v1/domain/ports"
)

func MapRoutes(h ports.IHandlers, router fiber.Router) {
	router.Get("/inventory", func(c *fiber.Ctx) error {
		return c.JSON("inventory")
	})
}
