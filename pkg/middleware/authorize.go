package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"github.com/hngprojects/hng_boilerplate_golang_web/utility"
)

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenStr string
		bearerToken := c.GetHeader("Authorization")
		strArr := strings.Split(bearerToken, " ")
		if len(strArr) == 2 {
			tokenStr = strArr[1]
		}

		if tokenStr == "" {
			r := utility.BuildErrorResponse(http.StatusUnauthorized, "error", "Token could not be found!", "Unauthorized", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, r)
			return
		}

		token, err := TokenValid(tokenStr)
		if err != nil {
			r := utility.BuildErrorResponse(http.StatusUnauthorized, "error", "Token is invalid!", "Unauthorized", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, r)
			return
		}

		// access user claims

		claims := token.Claims.(jwt.MapClaims)

		// check if user id exists
		_, ok := claims["user_id"].(string) //convert the interface to string
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utility.BuildErrorResponse(http.StatusUnauthorized, "error", "Token is invalid!", "Unauthorized", nil))
			return
		}

		authoriseStatus, ok := claims["authorised"].(bool) //check if token is authorised for middleware
		if !ok && !authoriseStatus {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utility.BuildErrorResponse(http.StatusUnauthorized, "error", "Token is invalid!", "Unauthorized", nil))
			return
		}

		// store user claims in Context
		// for accesiblity in controller

		c.Set("userClaims", claims)

		// call the next handler
		c.Next()

	}
}
