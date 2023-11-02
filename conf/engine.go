package conf

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	customLogger "github.com/shashank-mugiwara/meowxy/logger"
)

var router *fiber.App

func GetMeowxyEngine() *fiber.App {
	return router
}

func InitEngine() {
	app := fiber.New()
	app.Use(logger.New(customLogger.GetFiberLoggerConfig()))
	router = app
}
