package model

import (
	"github.com/gofiber/fiber/v3"
)

type (
	Service interface {
		Start() error
		Stop() error
	}

	Controller interface {
		Setup(*fiber.App) error
	}
)
