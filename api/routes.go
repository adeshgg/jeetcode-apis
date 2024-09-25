package api

import (
	"jeetcode-apis/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter(problemService *service.ProblemService) *gin.Engine {
	problemHandler := NewProblemHandler(problemService)

	router := gin.Default()

	// Define routes
	router.GET("/problems", problemHandler.GetProblems)
	router.POST("/problems", problemHandler.CreateProblems)

	return router
}
