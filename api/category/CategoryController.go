package category

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	response "github.com/impleopus/go-rest-api/pkg/custom-response"
)

func CreateCategory(ctx *gin.Context) {

	var reqData CreateCategoryValidation

	errors := []string{}

	if err := ctx.Bind(&reqData); err != nil {
		log.Println("Validation failed:", err)
		response.BadRequest(ctx, err.Error())
		return
	}

	if err := reqData.Validate(); err != nil {
		log.Println("Validation failed:", err)

		for _, e := range err.(validator.ValidationErrors) {

			fmt.Println(e)

			log.Printf("Field: %s, Error: %s\n", e.Field(), e.Tag())

			errors = append(
				errors,
				fmt.Sprintf("Field [%s]: %s", e.Field(), e),
			)
		}

		response.ValidationBadRequest(ctx, errors)
		return
	}

	fmt.Printf("%+v\n", reqData)

	response.Created(ctx, "Category created successfully.")
}

func FetchCategories(ctx *gin.Context) {
	response.Ok(ctx, "Categorys fetched successfully.")
}

func FetchCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	response.Ok(ctx, fmt.Sprintf("Category fetched by id %s successfully.", id))
}

func UpdateCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	response.Ok(ctx, fmt.Sprintf("Category updated by id %s successfully.", id))
}

func PartialUpdateCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	response.Ok(ctx, fmt.Sprintf("Category partially updated by id %s successfully.", id))
}

func DeleteCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	response.Ok(ctx, fmt.Sprintf("Category deleted by id %s successfully.", id))
}
