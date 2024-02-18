package mapping

import "github.com/gofiber/fiber/v2"

type Handler interface {
	GetRouteMapping(ctx *fiber.Ctx) error
	UpdateRouteMapping(ctx *fiber.Ctx) error
	DeleteRouteMapping(ctx *fiber.Ctx) error
	AddRouteMapping(ctx *fiber.Ctx) error
}

type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
	Debug(args ...interface{})
}

func RegisterRoutes(router *fiber.App, logger Logger) {

}
