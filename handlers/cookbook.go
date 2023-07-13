package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func HandleCookBookPage(c *fiber.Ctx) error {
	return c.Render("cookbook/index", fiber.Map{})
}




