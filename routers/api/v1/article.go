/*
 * @Description:
 * @Author: Caoxiang
 * @Date: 2021-01-19 10:22:43
 * @LastEditors: Caoxiang
 */
package v1

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/newinternetboy/poor_union/models"
	"github.com/newinternetboy/poor_union/pkg/e"
	"github.com/newinternetboy/poor_union/pkg/logging"
	"github.com/newinternetboy/poor_union/pkg/setting"
	"github.com/newinternetboy/poor_union/pkg/util"
	"github.com/unknwon/com"
)

//获取单篇文章
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	//认证
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			data = models.GetArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			logging.Error(fmt.Sprintf("err.Key:%s;err.Message:%s", err.Key, err.Message))
		}
	}
	c.JSON(http.StatusOK, util.Response(code, data))
}

//获取文章列表
func GetArticles(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})

	valid := validation.Validation{}
	var state int = -1
	if argv := c.Query("state"); argv != "" {
		state = com.StrTo(argv).MustInt()
		maps["state"] = state
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	var tagId int = -1
	if argv := c.Query("tag_id"); argv != "" {
		tagId = com.StrTo(argv).MustInt()
		maps["tag_id"] = tagId
		valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	}

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		data["lists"] = models.GetArticles(util.GetPageOffset(c), setting.PageSize, maps)
		data["total"] = models.GetArticleTotal(maps)
	} else {
		for _, err := range valid.Errors {
			logging.Error(fmt.Sprintf("err.Key:%s;err.Message:%s", err.Key, err.Message))
		}
	}
	c.JSON(http.StatusOK, util.Response(code, data))
}

//添加文章
func AddArticle(c *gin.Context) {
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	createdBy := c.Query("created_by")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()

	valid := validation.Validation{}
	valid.Min(tagId, 1, "tag_id").Message("标题ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(desc, "desc").Message("简述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistTagByID(tagId) {
			data := make(map[string]interface{})
			data["tag_id"] = tagId
			data["title"] = title
			data["desc"] = desc
			data["content"] = content
			data["created_by"] = createdBy
			data["state"] = state

			models.AddArticle(data)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			logging.Error(fmt.Sprintf("err.Key:%s;err.Message:%s", err.Key, err.Message))
		}
	}
	c.JSON(http.StatusOK, util.Response(code, make(map[string]interface{})))
}

//编辑文章
func EditArticle(c *gin.Context) {
	valid := validation.Validation{}

	id := com.StrTo(c.Query("id")).MustInt()
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	modifiedBy := c.Query("modified_by")

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.MaxSize(title, 100, "title").Message("标题最长100字符")
	valid.Required(desc, "desc").Message("描述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.MaxSize(content, 65535, "content").Message("内容最长65535字符")
	valid.MaxSize(desc, 100, "desc").Message("描述最长100字符")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长字符为100字符")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			if models.ExistTagByID(tagId) {
				data := make(map[string]interface{})
				if tagId > 0 {
					data["tag_id"] = tagId
				}
				data["title"] = title
				data["desc"] = desc
				data["content"] = content
				data["modified_by"] = modifiedBy

				models.EditArticle(id, data)
				code = e.SUCCESS
			} else {
				code = e.ERROR_NOT_EXIST_TAG
			}
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		util.ValidErrorPrintf(valid.Errors)
	}
	util.Render(c, code, make(map[string]interface{}))
}

//删除文章
func DeleteArticle(c *gin.Context) {
	//api/v1/articles/:id
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("文章id最小为1")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			models.DeleteArticleByID(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		util.ValidErrorPrintf(valid.Errors)
	}
	util.Render(c, code, make(map[string]interface{}))
}
