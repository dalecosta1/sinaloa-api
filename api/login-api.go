package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/dalecosta1/sinaloa-api/controller"
	"github.com/dalecosta1/sinaloa-api/helpers"
	"github.com/dalecosta1/sinaloa-api/dto"
	"github.com/dalecosta1/sinaloa-api/dto/messages"
)

type LoginApi struct {
	loginController controller.LoginController
}

func NewLoginAPI(loginController controller.LoginController) *LoginApi {
	return &LoginApi {
		loginController: loginController,
	}
}

// Paths Information

// Authenticate godoc
// @Summary Provides a JSON Web Token
// @Description Authenticates a user and provides a JWT to Authorize API calls
// @ID Authentication
// @Consume application/x-www-form-urlencoded
// @Produce json
// @Param username formData string true "Username"
// @Param password formData string true "Password"
// @Success 200 {object} dto.JWT
// @Failure 401 {object} messages.ResponseToken
// @Router /auth/token [post]
func (api *LoginApi) Authenticate(ctx *gin.Context) {
	token := api.loginController.Login(ctx)
		
	if token != "" {
		var result = helpers.HandleResponse(
			true,
			"200",
			"",
			nil,
			&dto.JWT{
				Token: token,
			},
		)
		ctx.JSON(http.StatusOK, result)
	} else {
		var result = helpers.HandleResponse(
			false,
			"401",
			"Not Authorized",
			nil,
			&messages.ResponseToken{
				Message: "Not Authorized",
			},
		)
		ctx.JSON(http.StatusUnauthorized, result)
	}
}
