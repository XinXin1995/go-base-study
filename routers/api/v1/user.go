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

func GetUsers(c *gin.Context) {
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
		data["list"], data["total"] = v1.GetUsers(pageSize, pageNo, name)

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

func AddUser(c *gin.Context) {
	name := c.PostForm("name")
	password := c.DefaultPostForm("password", "123456")
	avatar := c.PostForm("avatar")
	valid := validation.Validation{}
	valid.Required(name, "name").Message("登录名为空")
	valid.Required(password, "password").Message("密码不能为空")
	valid.Required(avatar, "avatar").Message("头像不能为空")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		password, salt := util.GeneratePwd(password)
		user := &v1.User{
			Name:     name,
			Password: password,
			Avatar:   avatar,
			Salt:     salt,
		}
		b := v1.AddUser(user)
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

func EditUser(c *gin.Context) {
	id := c.PostForm("id")
	name := c.PostForm("name")
	avatar := c.PostForm("avatar")
	valid := validation.Validation{}
	valid.Required(id, "id").Message("用户ID不能为空")
	valid.Required(name, "name").Message("登录名为空")
	valid.Required(avatar, "avatar").Message("头像不能为空")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		user := &v1.User{
			Name:   name,
			Avatar: avatar,
		}
		b := v1.EditUser(user, id)
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

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	valid := validation.Validation{}
	valid.Required(id, "id").Message("用户ID不能为空")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		b := v1.DeleteUser(id)
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
