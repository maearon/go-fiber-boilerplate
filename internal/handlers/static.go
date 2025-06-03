package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type StaticHandler struct{}

func NewStaticHandler() *StaticHandler {
	return &StaticHandler{}
}

func (h *StaticHandler) Home(c *fiber.Ctx) error {
	return c.Render("home", fiber.Map{
		"Title": "Home",
	})
}

func (h *StaticHandler) About(c *fiber.Ctx) error {
	return c.Render("about", fiber.Map{
		"Title": "About",
	})
}

func (h *StaticHandler) Help(c *fiber.Ctx) error {
	return c.Render("help", fiber.Map{
		"Title": "Help",
	})
}

func (h *StaticHandler) Contact(c *fiber.Ctx) error {
	return c.Render("contact", fiber.Map{
		"Title": "Contact",
	})
}
