package middleware

import (
	"strings"

	"go-fiber-boilerplate/internal/models"
	"go-fiber-boilerplate/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RequireAuth(db *gorm.DB, jwtSecret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Check for JWT token in Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization header required",
			})
		}

		// Extract token from "Bearer <token>"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid authorization format",
			})
		}

		// Validate token
		userID, err := services.ValidateJWT(tokenString, jwtSecret)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		// Get user from database
		var user models.User
		if err := db.First(&user, "id = ?", userID).Error; err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "User not found",
			})
		}

		// Store user in context
		c.Locals("user", &user)
		return c.Next()
	}
}

func RequireAdmin(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	if !user.Admin {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Admin access required",
		})
	}
	return c.Next()
}

func OptionalAuth(db *gorm.DB, jwtSecret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader != "" {
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString != authHeader {
				if userID, err := services.ValidateJWT(tokenString, jwtSecret); err == nil {
					var user models.User
					if err := db.First(&user, "id = ?", userID).Error; err == nil {
						c.Locals("user", &user)
					}
				}
			}
		}
		return c.Next()
	}
}
