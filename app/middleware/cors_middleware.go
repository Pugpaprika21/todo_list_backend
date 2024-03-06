package middleware

import (
	"github.com/gofiber/fiber/v2"
)

type corsMiddleware struct{}

func NewCorsMiddleware() *corsMiddleware {
	return &corsMiddleware{}
}

func Cors() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ctx.Set("Access-Control-Allow-Origin", "*")
		ctx.Set("Access-Control-Max-Age", "86400")
		ctx.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		ctx.Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		ctx.Set("Access-Control-Allow-Credentials", "true")
		ctx.Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0, post-check=0, pre-check=0")
		ctx.Set("Pragma", "no-cache")

		if ctx.Method() == "OPTIONS" {
			return ctx.SendStatus(fiber.StatusNoContent)
		}
		return ctx.Next()
	}
}
