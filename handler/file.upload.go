package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// https://github.com/gofiber/recipes/issues/47
// https://github.com/gofiber/fiber/issues/221

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
		//ctx.SaveFile(file, fmt.Sprintf("./%s", file.Filename))
		// 👷 Save file inside uploads folder under current working directory:
		ctx.SaveFile(file, fmt.Sprintf("./uploads/%s", file.Filename))
		// 👷 Save file using a relative path:
		//ctx.SaveFile(file, fmt.Sprintf("/tmp/uploads_relative/%s", file.Filename))
	}
	return ctx.SendStatus(fiber.StatusOK)
}

func Uploads(ctx *fiber.Ctx) error {
	// Parse the multipart form:
	form, err := ctx.MultipartForm()
	if err != nil {
		return err
	}
	// => *multipart.Form

	// Get all files from "documents" key:
	files := form.File["documents"]
	// => []*multipart.FileHeader

	// Loop through files:
	for _, file := range files {
		fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
		// => "tutorial.pdf" 360641 "application/pdf"

		// Save the files to disk:
		err := ctx.SaveFile(file, fmt.Sprintf("./%s", file.Filename))

		// Check for errors
		if err != nil {
			return err
		}
	}
	return ctx.SendStatus(fiber.StatusOK)
}
