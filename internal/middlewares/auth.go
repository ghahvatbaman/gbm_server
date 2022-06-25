package middlewares

import (
	"strings"

	"github.com/ghahvatbaman/gbm/internal/services/jwt"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	h := c.Get("Authorization")

	if h == "" {
		return fiber.ErrUnauthorized
	}

	chunks := strings.Split(h, " ")
	if len(chunks) < 2 {
		return fiber.ErrUnauthorized
	}

	userId, err := jwt.Verify(chunks[1])
	if err != nil {
		return fiber.ErrUnauthorized
	}

	c.Locals("userId", userId)
	return c.Next()
}
