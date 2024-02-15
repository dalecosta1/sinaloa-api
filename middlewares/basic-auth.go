package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/dalecosta1/sinaloa-api/helpers"
)

// BasicAuth takes as argument a map[string]string where
// the key is the user name and the value is the password.
func BasicAuth() gin.HandlerFunc {
	helpers.LoadConfig()
	return gin.BasicAuth(gin.Accounts{
		helpers.AppConfig.BASIC_AUTH_USER: helpers.AppConfig.BASIC_AUTH_PASSWORD,
	})
}
