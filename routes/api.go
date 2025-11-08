package routes

import (
	"github.com/gin-gonic/gin"
	"brand-collab-tracker/controllers"
	"brand-collab-tracker/middlewares"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		users := api.Group("/users")
		{
			users.POST("/register", controllers.RegisterHandler)
			users.POST("/login", controllers.LoginHandler)
		}

		authGroup := api.Group("/")
		authGroup.Use(middlewares.AuthJWTMiddleware())
        
		category := authGroup.Group("/categories")
		{
			category.POST("/", controllers.CreateCategoryHandler)
			category.GET("/", controllers.GetCategoriesHandler)
			category.PUT("/:id", controllers.UpdateCategoryHandler)
			category.DELETE("/:id", controllers.DeleteCategoryHandler)
		}

		brand := authGroup.Group("/brands")
		{
			brand.POST("/", controllers.CreateBrandHandler)
			brand.GET("/", controllers.GetBrandsHandler)
			brand.GET("/:id", controllers.GetBrandByIDHandler)
			brand.PUT("/:id", controllers.UpdateBrandHandler)
			brand.DELETE(":id", controllers.DeleteBrandHandler)
		}
	}
	return r
}