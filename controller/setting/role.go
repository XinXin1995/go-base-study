package setting

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
		data["list"], data["total"] = models.GetRoles(pageSize, pageNo, name)
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
		role := &models.Role{
			Name: name,
		}
		b := models.AddRole(role)
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
		role := &models.Role{
			Name: name,
		}
		b := models.EditRole(role, id)
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
		b := models.DeleteRole(id)
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
	roles := models.GetAllRoles()
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
		res.Data = models.GetRoleModules(uuid.FromStringOrNil(id))
	} else {
		util.LoopLog(valid.Errors)
	}
	res.Code = code
	c.JSON(http.StatusOK, res)
}

func AddRoleModules(c *gin.Context) {
	RM := models.RoleModules{}
	err := c.BindJSON(&RM)
	code := e.INVALID_PARAMS
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &util.Res{
			Code: code,
			Msg:  e.MsgUser[code],
		})
		return
	} else {
		id := uuid.FromStringOrNil(RM.Id)
		modules := RM.Modules
		b := models.AddRoleModules(id, modules)
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
