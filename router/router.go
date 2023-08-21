package router

import (
	"assignment-project-new/handler"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	studentRouter := router.Group("/students")
	{
		studentRouter.POST("/", handler.CreateStudent)
		studentRouter.GET("/", handler.GetAllStudents)
		studentRouter.PUT("/:id", handler.UpdateStudent)
		studentRouter.DELETE("/:id", handler.DeleteStudent)
	}

	scoreRouter := router.Group("/scores")
	{
		scoreRouter.GET("/", handler.GetAllScoreByStudent)

		scoreRouter.POST("/", handler.CreateScore)
		scoreRouter.PUT("/:id", handler.UpdateScore)
		scoreRouter.DELETE("/:id", handler.DeleteScore)
	}

	return router
}
