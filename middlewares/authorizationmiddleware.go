package middlewares

import (
	"github.com/draco121/common/clients"
	"github.com/draco121/common/constants"
	"github.com/draco121/common/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func AuthMiddleware(actions ...constants.Action) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationServiceApiClient := clients.NewAuthorizationServiceApiClient(os.Getenv("AUTHORIZATION_SERVICE_BASEURL"))
		requestData := models.AuthorizationInput{
			Token:   c.GetHeader("Authorization"),
			Actions: actions,
		}
		authResponse, err := authorizationServiceApiClient.Authorize(requestData)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			c.Abort()
			return
		} else {
			if authResponse.Grant == constants.Rejected {
				c.Status(http.StatusUnauthorized)
				c.Abort()
			} else {
				c.Set("UserId", authResponse.UserId)
				c.Next()
			}
		}
	}
}
