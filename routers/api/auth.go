/*
 * @Description:
 * @Author: Caoxiang
 * @Date: 2021-01-19 17:31:45
 * @LastEditors: Caoxiang
 */
package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/newinternetboy/poor_union/models"
	"github.com/newinternetboy/poor_union/pkg/e"
	"github.com/newinternetboy/poor_union/pkg/util"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		util.ValidErrorPrintf(valid.Errors)
	}
	util.Render(c, code, data)
}
