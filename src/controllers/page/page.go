package page_controller

import (
	"fmt"
	"link-shortener/src/models"
	page_model "link-shortener/src/models/page"
	"os"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{}, "layouts/layout")
}

func Create(c *fiber.Ctx) error {
	url := c.FormValue("url")

	if url == "" {
		return c.Redirect("/")
	}

	page := new(models.Page)
	page.Url = url
	err := page_model.Create(page)
	if err != nil {
		return c.Redirect("/")
	}

	domain := os.Getenv("domain")
	shortUrl := domain + page.Alias

	return c.Render("create", fiber.Map{"url": shortUrl}, "layouts/layout")
}

func GoToUrl(c *fiber.Ctx) error {
	alias := c.Params("alias")
	fmt.Printf(alias)
	page := new(models.Page)

	err := page_model.GetByAlias(page, alias)

	if err != nil {
		fmt.Printf(err.Error())
		return c.Redirect("/")
	}

	if page.ID == 0 {
		return c.Redirect("/")
	}

	return c.Redirect(page.Url)
}
