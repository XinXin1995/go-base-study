package middleware

import (
	"blog/pkg/e"
	"blog/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int

		code = e.SUCCESS
		token := c.GetHeader("Authorization")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {

			res := &util.Res{
				Code: code,
				Msg:  e.MsgUser[code],
			}
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}
		c.Next()
	}
}
