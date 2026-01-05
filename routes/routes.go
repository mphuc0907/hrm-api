package routes

import (
	"hrm-api/controllers"
	"hrm-api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	api := r.Group("/api")
	{
		api.POST("/login", controllers.Login)
		api.POST("/register", controllers.Register)

		protected := api.Group("/")
		protected.Use(middlewares.AuthMiddleware())
		{
			protected.GET("/employees", controllers.GetEmployees)
			protected.POST("/employees", controllers.CreateEmployee)
			protected.GET("/employees/:id", controllers.GetEmployee)
			protected.PUT("/employees/:id", controllers.UpdateEmployee)
			protected.DELETE("/employees/:id", controllers.DeleteEmployee)
		}
	}
}
