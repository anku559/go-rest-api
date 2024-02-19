package auth

import (
	"github.com/gin-gonic/gin"
	response "github.com/impleopus/go-rest-api/pkg/custom-response"
)

func SignUp(ctx *gin.Context) {
	response.Created(ctx, "User signed up successfully.")
}

func Login(ctx *gin.Context) {
	response.Ok(ctx, "User logged in successfully.")
}
