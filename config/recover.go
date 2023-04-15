package config

import (
	fRecover "github.com/gofiber/fiber/v2/middleware/recover"
)

func recoverConfig() fRecover.Config {
	return fRecover.Config{}
}
