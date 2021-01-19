/*
 * @Description:
 * @Author: Caoxiang
 * @Date: 2021-01-18 16:07:59
 * @LastEditors: Caoxiang
 */
package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/newinternetboy/poor_union/routers/api/v1"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//注册路由
	apiv1 := router.Group("/api/v1")
	{
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//article
		apiv1.GET("/articles", v1.GetArticles)
		apiv1.GET("/articles/:id", v1.GetArticle)
		apiv1.POST("/articles", v1.AddArticle)
		apiv1.PUT("/articles/:id", v1.EditArticle)
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}
	return router
}