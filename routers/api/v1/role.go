package v1

import (
	"blog/models/v1"
	"blog/pkg/e"
	"blog/pkg/util"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetRoles(c *gin.Context) {
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
		data["list"], data["total"] = v1.GetRoles(pageSize, pageNo, name)
	} else {
		for _, err := range valid.Errors {
			log.Fatalln(err.Key, err.Message)
		}
	}
	res := &util.Res{
		Code: code,
		Msg:  e.MsgUser[code],
		Data: data,
	}

	c.JSON(http.StatusOK, res)
}

func AddRole(c *gin.Context) {
	name := c.PostForm("name")
	valid := validation.Validation{}
	valid.Required(name, "name").Message("角色名称为空")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		role := &v1.Role{
			Name: name,
		}
		b := v1.AddRole(role)
		if b {
			code = e.SUCCESS
		} else {
			code = e.ERROR
		}
	} else {
		for _, err := range valid.Errors {
			log.Fatalln(err.Key, err.Message)
		}
	}
	res := &util.Res{
		Code: code,
		Msg:  e.MsgUser[code],
		Data: nil,
	}
	c.JSON(http.StatusOK, res)
}

func EditRole(c *gin.Context) {
	id := c.Param("id")
	name := c.PostForm("name")
	valid := validation.Validation{}
	valid.Required(id, "id").Message("角色ID不能为空")
	valid.Required(name, "name").Message("角色名称为空")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		role := &v1.Role{
			Name: name,
		}
		b := v1.EditRole(role, id)
		if b {
			code = e.SUCCESS
		} else {
			code = e.ERROR
		}
	} else {
		for _, err := range valid.Errors {
			log.Fatalln(err.Key, err.Message)
		}
	}
	res := &util.Res{
		Code: code,
		Msg:  e.MsgUser[code],
		Data: nil,
	}
	c.JSON(http.StatusOK, res)
}

func DeleteRole(c *gin.Context) {
	id := c.Param("id")
	valid := validation.Validation{}
	valid.Required(id, "id").Message("角色ID不能为空")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		b := v1.DeleteRole(id)
		if b {
			code = e.SUCCESS
		} else {
			code = e.ERROR
		}
	} else {
		for _, err := range valid.Errors {
			log.Fatalln(err.Key, err.Message)
		}
	}
	res := &util.Res{
		Code: code,
		Msg:  e.MsgUser[code],
		Data: nil,
	}
	c.JSON(http.StatusOK, res)

}

func GetAllRoles(c *gin.Context) {
	roles := v1.GetAllRoles()
	res := &util.Res{
		Code: e.SUCCESS,
		Msg:  e.MsgUser[e.SUCCESS],
		Data: roles,
	}
	c.JSON(http.StatusOK, res)
}

func GetRoleModules(c *gin.Context) {
	id := c.Query("id")
	valid := validation.Validation{}
	valid.Required(id, "id").Message("模块id不能为空")
	code := e.INVALID_PARAMS
	res := util.Res{}
	if !valid.HasErrors() {
		code = e.SUCCESS
		res.Msg = e.MsgUser[code]
		res.Data = v1.GetRoleModules(id)
	} else {
		util.LoopLog(valid.Errors)
	}
	res.Code = code
	c.JSON(http.StatusOK, res)
}
