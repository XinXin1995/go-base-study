package v1

import (
	"blog/models/v1"
	"blog/pkg/e"
	"blog/pkg/util"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
)

func AddModule(c *gin.Context) {
	name := c.PostForm("name")
	parentId := c.DefaultPostForm("parentId", "")
	icon := c.DefaultPostForm("icon", "")
	router := c.DefaultPostForm("router", "")
	valid := validation.Validation{}
	valid.Required(name, "name").Message("模块名称未填写")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		module := &v1.Module{
			Name:     name,
			ParentId: parentId,
			Router:   router,
			Icon:     icon,
		}
		b := v1.AddModule(module)
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

func GetModules(c *gin.Context) {
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
		data["list"], data["total"] = v1.GetModules(pageSize, pageNo, name)
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

func GetAllModules(c *gin.Context) {
	modules := v1.GetAllModules()
	code := e.SUCCESS
	res := &util.Res{
		Code: code,
		Msg:  e.MsgUser[code],
		Data: modules,
	}
	c.JSON(http.StatusOK, res)
}

func DeleteModule(c *gin.Context) {
	id := c.Param("id")
	valid := validation.Validation{}
	valid.Required(id, "id").Message("分组ID不能为空")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		b := v1.DeleteModule(id)
		if b == 0 {
			code = e.SUCCESS
		} else if b == 1 {
			code = e.ERROR_EXISR
		} else if b == 2 {
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

func EditModule(c *gin.Context) {
	id := c.Param("id")
	name := c.PostForm("name")
	icon := c.DefaultPostForm("icon", "")
	parentId := c.DefaultPostForm("parentId", "")
	router := c.DefaultPostForm("router", "")
	valid := validation.Validation{}
	valid.Required(id, "id").Message("角色ID不能为空")
	valid.Required(name, "name").Message("角色名称为空")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		module := &v1.Module{
			Name:     name,
			ParentId: parentId,
			Router:   router,
			Icon:     icon,
		}
		b := v1.EditModule(module, id)
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

func AddModuleApis(c *gin.Context) {
	MA := v1.ModuleApis{}
	err := c.BindJSON(&MA)
	code := e.INVALID_PARAMS
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &util.Res{
			Code: code,
			Msg:  e.MsgUser[code],
		})
		return
	} else {
		id := uuid.FromStringOrNil(MA.Id)
		apis := MA.Apis
		b := v1.AddModuleApis(apis, id)
		if b {
			code = e.SUCCESS
		} else {
			code = e.ERROR
		}
	}
	res := &util.Res{
		Code: code,
		Msg:  e.MsgUser[code],
	}
	c.JSON(http.StatusOK, res)
}

func ModuleApis(c *gin.Context) {
	id := c.Query("id")
	valid := validation.Validation{}
	valid.Required(id, "id").Message("模块ID不能为空")
	code := e.INVALID_PARAMS
	res := &util.Res{}
	if !valid.HasErrors() {
		code = e.SUCCESS

		res.Data = v1.GetModuleApis(uuid.FromStringOrNil(id))
	} else {
		util.LoopLog(valid.Errors)
	}
	res.Code = code
	res.Msg = e.MsgUser[code]
	c.JSON(http.StatusOK, res)
}
