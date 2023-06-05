package controller_example

import "github.com/gin-gonic/gin"

type ExampleController interface {
	Create(ctx *gin.Context)
}

type exampleController struct {
}

func NewExampleController() ExampleController {
	return &exampleController{}
}

func (e exampleController) Create(ctx *gin.Context) {
}
