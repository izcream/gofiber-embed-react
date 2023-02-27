package main

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed all:frontend/dist/*
var dist embed.FS

func main() {
	app := fiber.New()
	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(dist),
		PathPrefix: "/frontend/dist",
	}))

	app.Listen(":3000")
}
