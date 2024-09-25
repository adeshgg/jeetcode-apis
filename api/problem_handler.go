package api

import (
	"jeetcode-apis/internal/service"
	"jeetcode-apis/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProblemHandler struct {
	service *service.ProblemService
}

func NewProblemHandler(service *service.ProblemService) *ProblemHandler {
	return &ProblemHandler{service: service}
}

func (p *ProblemHandler) GetProblems(c *gin.Context) {
	problems, err := p.service.GetProblems()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, problems)
}

func (p *ProblemHandler) CreateProblems(c *gin.Context) {
	var problemInput model.ProblemCreateInput

	if err := c.BindJSON(&problemInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	problem := model.Problem{
		Name: problemInput.Name,
		Link: problemInput.Link,
	}

	if err := p.service.CreateProblems(problem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, problem)
}
