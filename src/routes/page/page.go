package page_routes

import (
	page_controller "link-shortener/src/controllers/page"

	"github.com/gofiber/fiber/v2"
)

func Router(router fiber.Router) {
	router.Get("/", page_controller.Index)
	router.Post("/", page_controller.Create)
	router.Get("/analytics/:alias", page_controller.Analytics)
	router.Get("/:alias", page_controller.GoToUrl)
}
