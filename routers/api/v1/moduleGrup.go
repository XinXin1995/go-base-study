package v1

import (
	"blog/models/v1"
	"blog/pkg/e"
	"blog/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddModuleGroup(c *gin.Context) {
	name := c.PostForm("name")
	router := c.DefaultPostForm("router", "")
	valid := validation.Validation{}
	valid.Required(name, "name").Message("分组名称未填写")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		mg := &v1.ModuleGroup{
			Name:   name,
			Router: router,
		}
		b := v1.AddModuleGroup(mg)
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

func GetAllGroups(c *gin.Context) {
	mgs := v1.GetAllModules()
	res := util.Res{
		Code: e.SUCCESS,
		Msg:  e.MsgUser[e.SUCCESS],
		Data: mgs,
	}
	c.JSON(http.StatusOK, res)
}
