package page_model

import (
	"link-shortener/src/database"
	"link-shortener/src/models"

	"github.com/google/uuid"
)

func Create(page *models.Page) error {
	uuid := uuid.New()
	page.Alias = uuid.String()[0:8]
	result := database.DB.Create(page)
	return result.Error
}

func GetByAlias(page *models.Page, alias string) error {
	result := database.DB.First(page, "alias = ?", alias)
	return result.Error
}
