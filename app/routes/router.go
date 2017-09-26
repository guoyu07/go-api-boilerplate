package routes

import (
	"github.com/w3tecch/go-api-boilerplate/app/controllers"
	"github.com/w3tecch/go-api-boilerplate/app/middlewares"
	"github.com/w3tecch/go-api-boilerplate/app/repositories"
	"github.com/w3tecch/go-api-boilerplate/app/services"
	"github.com/w3tecch/go-api-boilerplate/config/database"

	"github.com/gin-gonic/gin"
)

// Router ...
func Router() *gin.Engine {
	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// ------------------------------
	// Logger middleware will write the logs to gin.DefaultWriter even you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// Custom middlewares
	r.Use(middlewares.Secure())
	r.Use(middlewares.CORS())
	r.Use(middlewares.NoCache())
	r.Use(middlewares.Version())
	r.Use(middlewares.RequestID())
	// r.Use(middlewares.Authentication())

	// Define repos and services
	// ------------------------------
	db := database.Connection()

	// Repositories
	userRepository := repositories.NewUserRepository(db)

	// Services
	userService := services.NewUserService(userRepository)

	// Controller Routes
	// ------------------------------
	// Return the api information to the user
	apiController := new(controllers.APIController)
	r.GET("/api/info", apiController.GetInfo)
	r.GET("/api/seeding", apiController.Seeding)

	// Users endpoints
	users := r.Group("/api/users")
	{
		userController := controllers.NewUserController(userService)
		users.GET("/", userController.GetAll)
		users.POST("/", userController.Create)
		users.GET("/:id", userController.GetByID)
		users.PUT("/:id", userController.Update)
		users.DELETE("/:id", userController.Delete)
	}

	return r
}
