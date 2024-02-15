package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"github.com/dalecosta1/sinaloa-api/service"
)

type ArgocdManagerController interface {
	Create(ctx *gin.Context) (interface{}, error)
}

type controller struct {
	service service.ArgocdManagerService
}

var validate *validator.Validate

func NewArgocdManagerController(service service.ArgocdManagerService) ArgocdManagerController {
	validate = validator.New()
	return &controller{
		service: service,
	}
}

func (c *controller) Create(ctx *gin.Context) (interface{}, error) {
    var rawJSON json.RawMessage
    if err := ctx.ShouldBindJSON(&rawJSON); err != nil {
        return nil, err
    }
    return c.service.MultiActions(rawJSON) // Adjusted to match the updated service method
}
