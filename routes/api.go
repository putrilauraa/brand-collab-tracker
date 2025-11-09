package routes

import (
	"github.com/gin-gonic/gin"
	"brand-collab-tracker/controllers"
	"brand-collab-tracker/middlewares"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "status": "up",
            "message": "Brand Collaboration Tracker API is running!",
        })
    })

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

		project := authGroup.Group("/projects")
		{
			project.POST("/", controllers.CreateProjectHandler)
			project.GET("/", controllers.GetProjectsHandler)
			project.GET("/:id", controllers.GetProjectByIDHandler)
			project.PUT("/:id", controllers.UpdateProjectHandler)
			project.DELETE("/:id", controllers.DeleteProjectHandler)

			project.GET("/:id/tasks", controllers.GetTaskByProjectHandler)

			project.GET("/:id/attachments", controllers.GetAttachmentsByProjectHandler)
		}

		task := authGroup.Group("/tasks")
		{
			task.POST("/", controllers.CreateTaskHandler)
			task.GET("/:id", controllers.GetTaskByIDHandler)
			task.PUT("/:id", controllers.UpdateTaskHandler)
			task.DELETE(":id", controllers.DeleteTaskHandler)
		}

		attachment := authGroup.Group("/attachments")
		{
			attachment.POST("/", controllers.CreateAttachmentHandler)
			attachment.GET("/:id", controllers.GetAttachmentByIDHandler)
			attachment.PUT("/:id", controllers.UpdateAttachmentHandler)
			attachment.DELETE("/:id", controllers.DeleteAttachmentHandler)
		}
	}
	return r
}