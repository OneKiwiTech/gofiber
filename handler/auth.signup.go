package handler

import (
	"time"

	"github.com/OneKiwiTech/gofiber/model"
	"github.com/OneKiwiTech/gofiber/service"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func Signup(ctx *fiber.Ctx) error {
	var user model.User
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.SendStatus(fiber.StatusBadGateway)
	}

	if errors := service.Validate(user); errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	user.Password = hashedPassword
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	user, _ = service.CreateUser(user)
	return ctx.JSON(fiber.Map{
		"message": "Registration Successful",
		"data":    user,
	})
}
