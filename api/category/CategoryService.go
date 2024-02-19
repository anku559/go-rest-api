package category

import (
	"github.com/impleopus/go-rest-api/internal/config/db"
	"github.com/impleopus/go-rest-api/pkg/models"
)

func DBCreateCategory(reqData CreateCategoryValidation) {

	category := &models.Category{
		Name: reqData.Name,
		Slug: reqData.Slug,
	}
	db.DB.Create(category)
}
