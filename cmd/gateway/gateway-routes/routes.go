package gatewayroutes

import (
	"github.com/gin-gonic/gin"
	"github.com/impleopus/go-rest-api/api/auth"
	"github.com/impleopus/go-rest-api/api/category"
	"github.com/impleopus/go-rest-api/api/users"
	"github.com/impleopus/go-rest-api/internal/config/env"
	"github.com/impleopus/go-rest-api/pkg/constants"
)

func SetupRouter() *gin.Engine {

	engine := gin.Default()

	// API Prefix
	router := engine.Group(env.API_PREFIX)

	// Auth Routes
	authRoutes := router.Group(constants.AuthRoute)
	authRoutes.POST("/sign-up", auth.SignUp)
	authRoutes.POST("/login", auth.Login)

	// User Routes
	userRoutes := router.Group(constants.UserRoutes)
	userRoutes.POST("/", users.CreateUser)
	userRoutes.GET("/", users.FetchUsers)
	userRoutes.GET("/:id", users.FetchUser)
	userRoutes.PUT("/:id", users.UpdateUser)
	userRoutes.PATCH("/:id", users.PartialUpdateUser)
	userRoutes.DELETE("/:id", users.DeleteUser)

	// Category Routes
	categoryRoutes := router.Group(constants.CategoryRoute)
	categoryRoutes.POST("/", category.CreateCategory)
	categoryRoutes.GET("/", category.FetchCategories)
	categoryRoutes.GET("/:id", category.FetchCategory)
	categoryRoutes.PUT("/:id", category.UpdateCategory)
	categoryRoutes.PATCH("/:id", category.PartialUpdateCategory)
	categoryRoutes.DELETE("/:id", category.DeleteCategory)

	return engine
}
