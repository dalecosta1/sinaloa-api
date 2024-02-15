package api

import (
    "net/http"
	"encoding/json"
    "github.com/dalecosta1/sinaloa-api/controller"
	"github.com/dalecosta1/sinaloa-api/helpers"
    "github.com/gin-gonic/gin"
)

type ArgocdManagerApi struct {
	argocdManagerController controller.ArgocdManagerController
}

func NewArgocdManagerAPI(argocdManagerController controller.ArgocdManagerController) *ArgocdManagerApi {
	return &ArgocdManagerApi{
		argocdManagerController: argocdManagerController,
	}
}

// Create godoc
// @Security bearerAuth
// @Summary Create new items
// @Description Create a new items
// @Tags argocdManager,create
// @Accept json
// @Produce json
// @Param obj body json.RawMessage true "Create"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /argocd-manager [post]
func (api *ArgocdManagerApi) Create(ctx *gin.Context) {
    // Check input json
	var rawJSON json.RawMessage
    if err := ctx.ShouldBindJSON(&rawJSON); err != nil {
        // Error occurred while binding JSON, so we call HandleResponse with an error
        response := helpers.HandleResponse(false, "400", "Bad request", err, struct{}{})
        ctx.JSON(http.StatusBadRequest, response)
        return
    }    
	// Handle the JSON binding and controller interaction correctly...
    jsonData, err := api.argocdManagerController.Create(ctx)
    if err != nil {
        response := helpers.HandleResponse(false, "500", "Internal Server Error", err, struct{}{})
        ctx.JSON(http.StatusInternalServerError, response)
        return
    }
	// Return the response
    response := helpers.HandleResponse(true, "200", "Operation executed correctly", nil, jsonData)
    ctx.JSON(http.StatusOK, response)
}
