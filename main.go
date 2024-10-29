package main

import (
	"link-shortener/src/database"
	"link-shortener/src/models"
	page_routes "link-shortener/src/routes/page"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()
}

func main() {
	database.Conn()
	database.DB.AutoMigrate(models.Page{})

	engine := django.New("./src/views", ".html")

	app := fiber.New(fiber.Config{Views: engine})

	//Routes
	page_routes.Router(app.Group("/"))

	app.Listen(":3000")
}
