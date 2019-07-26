package blogs

import (
	"blog/models"
	"blog/pkg/e"
	"blog/pkg/util"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Tags(c *gin.Context) {
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
		data["list"], data["total"] = models.GetTags(pageSize, pageNo, name)
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

func AddTag(c *gin.Context) {
	name := c.PostForm("name")
	valid := validation.Validation{}
	valid.Required(name, "name").Message("登录名为空")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		tag := &models.Tag{
			Name: name,
		}
		b := models.AddTag(tag)
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

func EditTag(c *gin.Context) {
	id := c.PostForm("id")
	name := c.PostForm("name")
	valid := validation.Validation{}
	valid.Required(id, "id").Message("标签ID不能为空")
	valid.Required(name, "name").Message("标签名为空")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		tag := &models.Tag{
			Name: name,
		}
		b := models.EditTag(tag, id)
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

func DelTag(c *gin.Context) {
	id := c.Param("id")
	valid := validation.Validation{}
	valid.Required(id, "id").Message("标签ID不能为空")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		b := models.DelTag(id)
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
