/*
 * @Description:
 * @Author: Caoxiang
 * @Date: 2021-01-18 17:47:21
 * @LastEditors: Caoxiang
 */
package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/newinternetboy/poor_union/pkg/e"
)

func Response(code int, data interface{}) gin.H {
	return gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	}
}

//接口参数输出
func Render(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, Response(code, data))
}
