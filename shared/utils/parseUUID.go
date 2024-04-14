package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func ParseUUID(c *fiber.Ctx, id string) uuid.UUID {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return uuid.UUID{}
	}

	return parsedUUID
}
