package router

import (
	"assignment-project-new/handler"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	router.GET("/students", handler.GetAllStudents)
	studentRouter := router.Group("/student")
	{
		studentRouter.POST("/", handler.CreateStudent)

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
