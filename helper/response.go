package helper

import "github.com/gofiber/fiber/v2"

func SuccessResponse(ctx *fiber.Ctx, code int, data interface{}) error {
	return ctx.Status(code).JSON(fiber.Map{
		"status": "Success",
		"code":   code,
		"data":   data,
	})
}

func ErrorResponse(ctx *fiber.Ctx, code int, message string) error {
	return ctx.Status(code).JSON(fiber.Map{
		"status":  "Error",
		"code":    code,
		"message": message,
	})
}
