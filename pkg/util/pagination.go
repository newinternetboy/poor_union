/*
 * @Description:
 * @Author: Caoxiang
 * @Date: 2021-01-18 14:49:38
 * @LastEditors: Caoxiang
 */
package util

import (
	"github.com/gin-gonic/gin"
	"github.com/newinternetboy/poor_union/pkg/setting"
	"github.com/unknwon/com"
)

func GetPageOffset(c *gin.Context) int {
	offset := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		offset = (page - 1) * setting.AppSetting.PageSize
	}
	return offset
}
