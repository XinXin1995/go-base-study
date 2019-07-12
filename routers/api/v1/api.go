package v1

import (
	"blog/models/v1"
	"blog/pkg/e"
	"blog/pkg/util"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddApi(c *gin.Context) {
	name := c.PostForm("name")
	path := c.PostForm("path")
	valid := validation.Validation{}
	valid.Required(name, "name").Message("api名称未填写")
	valid.Required(path, "path").Message("api路劲未填写")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		api := &v1.Api{
			Name: name,
			Path: path,
		}
		b := v1.AddApi(api)
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
	}
	c.JSON(http.StatusOK, res)
}

func GetApis(c *gin.Context) {
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
		data["list"], data["total"] = v1.GetApis(pageSize, pageNo, name)
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

func DeleteApi(c *gin.Context) {
	id := c.Param("id")
	valid := validation.Validation{}
	valid.Required(id, "id").Message("分组ID不能为空")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		b := v1.DeleteApi(id)
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

func EditApi(c *gin.Context) {
	id := c.Param("id")
	name := c.PostForm("name")
	path := c.PostForm("path")
	valid := validation.Validation{}
	valid.Required(id, "id").Message("apiID不能为空")
	valid.Required(name, "name").Message("api名称为空")
	valid.Required(path, "path").Message("api路径不能为空")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		api := &v1.Api{
			Name: name,
			Path: path,
		}
		b := v1.EditApi(api, id)
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
