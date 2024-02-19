package food

import (
	"fmt"

	"github.com/gin-gonic/gin"
	response "github.com/impleopus/go-rest-api/pkg/custom-response"
)

func CreateFood(ctx *gin.Context) {
	response.Created(ctx, "Food created successfully.")
}

func FetchFoods(ctx *gin.Context) {
	response.Ok(ctx, "Foods fetched successfully.")
}

func FetchFood(ctx *gin.Context) {
	id := ctx.Param("id")
	response.Ok(ctx, fmt.Sprintf("Food fetched by id %s successfully.", id))
}

func UpdateFood(ctx *gin.Context) {
	id := ctx.Param("id")
	response.Ok(ctx, fmt.Sprintf("Food updated by id %s successfully.", id))
}

func PartialUpdateFood(ctx *gin.Context) {
	id := ctx.Param("id")
	response.Ok(ctx, fmt.Sprintf("Food partially updated by id %s successfully.", id))
}

func DeleteFood(ctx *gin.Context) {
	id := ctx.Param("id")
	response.Ok(ctx, fmt.Sprintf("Food deleted by id %s successfully.", id))
}
