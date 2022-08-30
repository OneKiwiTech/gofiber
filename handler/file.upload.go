package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// https://github.com/gofiber/recipes/issues/47
func Upload(ctx *fiber.Ctx) error {
	// Get first file from form field "document":
	file, err := ctx.FormFile("document")
	if err != nil {
		return err
	}
	// Save file to root directory:
	//err = ctx.SaveFile(file, fmt.Sprintf("./%s", file.Filename))

	// Check for errors:
	if err == nil {
		// 👷 Save file to root directory:
		ctx.SaveFile(file, fmt.Sprintf("./%s", file.Filename))
		// 👷 Save file inside uploads folder under current working directory:
		ctx.SaveFile(file, fmt.Sprintf("./uploads/%s", file.Filename))
		// 👷 Save file using a relative path:
		ctx.SaveFile(file, fmt.Sprintf("/tmp/uploads_relative/%s", file.Filename))
	}
	return ctx.SendStatus(fiber.StatusOK)
}
