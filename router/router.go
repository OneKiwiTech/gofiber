package router

import (
	"github.com/OneKiwiTech/gofiber/handler"
	"github.com/OneKiwiTech/gofiber/middleware"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {

	// Auth
	auth := app.Group("/auth")
	auth.Post("/signup", handler.Signup)
	auth.Post("/login", handler.Login)
	auth.Post("/logout", middleware.ProtectedRoute(), handler.Logout)
	auth.Post("/renew-token", handler.RenewToken)
	auth.Post("/reset-password", handler.ResetPassword)

	file := app.Group("/file")
	file.Post("/upload", handler.Upload)
	file.Post("/uploads", handler.Uploads)

}
