/*
 * @Description:
 * @Author: Caoxiang
 * @Date: 2021-01-19 16:46:25
 * @LastEditors: Caoxiang
 */
package jwt

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/newinternetboy/poor_union/pkg/e"
	"github.com/newinternetboy/poor_union/pkg/util"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.Query("token")
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
			c.JSON(http.StatusUnauthorized, util.Response(code, data))
			c.Abort()
			return
		}

		c.Next()
	}
}
