package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/ehsaniara/go-crash/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code string

		code = "SUCCESS"
		token := c.Query("token")

		fmt.Printf(">>>>>>>>>>>> token: %s \n", token)

		if token == "" {
			code = "INVALID_PARAMS"
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = "ERROR_AUTH_CHECK_TOKEN_TIMEOUT"
				default:
					code = "ERROR_AUTH_CHECK_TOKEN_FAIL"
				}
			}
			fmt.Printf("claims: %v\n", claims)
		}

		if code != "SUCCESS" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "Unauthorized",
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
