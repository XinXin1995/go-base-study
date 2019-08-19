package blogs

import (
	"blog/models"
	"blog/pkg/e"
	"blog/pkg/util"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
)

func Articles(c *gin.Context) {
	name := c.Query("name")
	pageNo, _ := com.StrTo(c.Query("pageNo")).Int()
	pageSize, _ := com.StrTo(c.Query("pageSize")).Int()
	valid := validation.Validation{}
	valid.Min(pageNo, 1, "pageNo").Message("页码最小为1")
	valid.Min(pageSize, 1, "pageSize").Message("每页数据最小为1")
	code := e.INVALID_PARAMS
	data := make(map[string]interface{})
	if !valid.HasErrors() {
		code = e.SUCCESS
		data["list"], data["total"] = models.GetArticles(pageSize, pageNo, name)
	} else {
		util.LoopLog(valid.Errors)
	}
	res := &util.Res{
		Code: code,
		Msg:  e.MsgUser[code],
		Data: data,
	}
	c.JSON(http.StatusOK, res)
}

func AddArticle(c *gin.Context) {
	articleParam := &models.ArticleParam{}
	err := c.BindJSON(articleParam)
	code := e.INVALID_PARAMS
	if err != nil {
		log.Println(err)
		res := &util.Res{
			Code: code,
			Msg:  e.MsgUser[code],
		}
		c.JSON(http.StatusOK, res)
	} else {
		claims, _ := util.ParseToken(c.GetHeader("Authorization"))
		tags := models.GetTagsById(articleParam.Tags)
		article := &models.Article{
			Name:         articleParam.Name,
			CategoryUuid: uuid.FromStringOrNil(articleParam.CategoryUuid),
			Content:      articleParam.Content,
			Creator:      claims.Uuid,
			Tags:         tags,
		}
		b := models.AddArticle(article)
		if b {
			code = e.SUCCESS
		} else {
			code = e.ERROR
		}
		res := &util.Res{
			Code: code,
			Msg:  e.MsgUser[code],
			Data: nil,
		}
		c.JSON(http.StatusOK, res)
	}
}

func EditArticle(c *gin.Context) {
	article := &models.Article{}
	err := c.BindJSON(article)
	code := e.INVALID_PARAMS
	if err != nil {
		log.Println(err)
		res := &util.Res{
			Code: code,
			Msg:  e.MsgUser[code],
		}
		c.JSON(http.StatusOK, res)
	} else {
		b := models.EditArticle(article)
		if b {
			code = e.SUCCESS
		} else {
			code = e.ERROR
		}
		res := &util.Res{
			Code: code,
			Msg:  e.MsgUser[code],
			Data: nil,
		}
		c.JSON(http.StatusOK, res)
	}

}

func DelArticle(c *gin.Context) {
	id := c.Param("id")
	valid := validation.Validation{}
	valid.Required(id, "id").Message("文章ID不能为空")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		b := models.DelArticle(id)
		if b {
			code = e.SUCCESS
		} else {
			code = e.ERROR
		}
	} else {
		util.LoopLog(valid.Errors)
	}
	res := &util.Res{
		Code: code,
		Msg:  e.MsgUser[code],
		Data: nil,
	}
	c.JSON(http.StatusOK, res)
}
