package main

import (
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/OneKiwiTech/gofiber/config"
	"github.com/OneKiwiTech/gofiber/database"
	"github.com/OneKiwiTech/gofiber/router"
)

var (
	port    = config.LoadConfig("APP_PORT")
	prod, _ = strconv.ParseBool(config.LoadConfig("APP_PROD"))
)

func main() {
	// Connect to the database
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}
	//app := fiber.New()
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		ServerHeader:  "AuthServer",
		Prefork:       prod,
	})

	app.Use(logger.New())
	app.Use(limiter.New(limiter.Config{
		Max:        20,
		Expiration: 20 * time.Second,
	}))
	app.Use(recover.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "POST",
	}))
	router.SetupRoutes(app)

	app.Listen(":8080")
}
