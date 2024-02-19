package users

import (
	"fmt"

	"github.com/gin-gonic/gin"
	response "github.com/impleopus/go-rest-api/pkg/custom-response"
)

func CreateUser(ctx *gin.Context) {
	response.Created(ctx, "User created successfully.")
}

func FetchUsers(ctx *gin.Context) {
	response.Ok(ctx, "Users fetched successfully.")
}

func FetchUser(ctx *gin.Context) {
	id := ctx.Param("id")
	response.Ok(ctx, fmt.Sprintf("User fetched by id %s successfully.", id))
}

func UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	response.Ok(ctx, fmt.Sprintf("User updated by id %s successfully.", id))
}

func PartialUpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	response.Ok(ctx, fmt.Sprintf("User partially updated by id %s successfully.", id))
}

func DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	response.Ok(ctx, fmt.Sprintf("User deleted by id %s successfully.", id))
}
