package main

import (
	"embed"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed all:ui/dist/*
var dist embed.FS

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	app := fiber.New()
	app.Use(filesystem.New(filesystem.Config{
		Root:         http.FS(dist),
		Index:        "index.html",
		NotFoundFile: "index.html",
		PathPrefix:   "ui/dist",
	}))
	//handle fallback routes
	app.Use("*", func(ctx *fiber.Ctx) error {
		return filesystem.SendFile(ctx, http.FS(dist), "ui/dist/index.html")
	})

	log.Fatal(app.Listen(":" + port))
}
