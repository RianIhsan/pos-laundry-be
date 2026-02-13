package middleware

import (
	"github.com/RianIhsan/pos-laundry-be/pkg/httpErrors/response"
	"github.com/RianIhsan/pos-laundry-be/pkg/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Auth struct {
	Id       uint64
	Username string
	Role     string
	Name     string
}

func (mw *MiddlewareManager) AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		authorizationHeader := context.GetHeader("Authorization")
		mw.logger.WithFields(logrus.Fields{
			"Authorization": authorizationHeader,
		}).Debug("auth middleware authorization header")

		var tokenString string

		if authorizationHeader != "" {
			headerParts := strings.Split(authorizationHeader, " ")
			if len(headerParts) < 2 {
				errResponse := errors.New("Authorization header is invalid")
				utils.LogErrorResponse(context, mw.logger, errResponse)
				response.SendErrorResponse(context, http.StatusUnauthorized, "Unauthorized header")
				context.Abort()
				return
			}
			tokenString = headerParts[1]
		} else {
			cookie, err := context.Cookie("jwt-token")
			if err != nil {
				errResponse := errors.New("Authorization header is invalid (cookie)")
				utils.LogErrorResponse(context, mw.logger, errResponse)
				response.SendErrorResponse(context, http.StatusUnauthorized, "Unauthorized cookie")
				context.Abort()
				return
			}
			tokenString = cookie
		}

		if tokenString == "" || len(strings.Split(tokenString, ".")) != 3 {
			utils.LogErrorResponse(context, mw.logger, errors.New("Invalid token"))
			response.SendErrorResponse(context, http.StatusUnauthorized, "invalid token format")
			context.Abort()
			return
		}

		claims, err := utils.ValidateJwtToken(tokenString, mw.cfg)

		if err != nil {
			utils.LogErrorResponse(context, mw.logger, err)

			if strings.Contains(err.Error(), "token is expired") {
				response.SendErrorResponse(context, http.StatusUnauthorized, "token is expired")
			} else if strings.Contains(err.Error(), "signature is invalid") {
				response.SendErrorResponse(context, http.StatusUnauthorized, "invalid token, please check your token")
			} else {
				response.SendErrorResponse(context, http.StatusUnauthorized, "invalid token, please check your token")
			}
			context.Abort()
			return
		}

		auth := &Auth{
			Id:       claims.ID,
			Role:     claims.Role,
			Name:     claims.Name,
			Username: claims.Username,
		}
		context.Set("auth", auth)

		context.Next()
	}
}
