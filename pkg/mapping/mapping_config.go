package mapping

import "github.com/gofiber/fiber/v2"

type handler struct {
	Logger Logger
}

func NewHandler(logger Logger) Handler {
	return &handler{
		Logger: logger,
	}
}

func (h handler) AddRouteMapping(fiberCtx *fiber.Ctx) error {
	return nil
}

func (h handler) DeleteRouteMapping(fiberCtx *fiber.Ctx) error {
	return nil
}

func (h handler) UpdateRouteMapping(fiberCtx *fiber.Ctx) error {
	return nil
}

func (h handler) GetRouteMapping(fiberCtx *fiber.Ctx) error {
	return nil
}
