package http

import (
	"github.com/gofiber/fiber/v2"
	"main/internal/common/v1/domain/ports"
)

func MapRoutes(h ports.IHandlers, router fiber.Router) {
	router.Get("/common", func(c *fiber.Ctx) error {
		return c.JSON("common")
	})
}
