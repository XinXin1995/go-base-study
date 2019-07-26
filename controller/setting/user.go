package setting

import (
	"blog/models"
	"blog/pkg/e"
	"blog/pkg/util"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"net/http"
	"time"
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
		data["list"], data["total"] = models.GetUsers(pageSize, pageNo, name)

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

func GetUser(c *gin.Context) {
	token := c.GetHeader("Authorization")
	claims, err := util.ParseToken(token)
	res := &util.Res{}
	data := make(map[string]interface{})
	if err != nil {
		res.Code = e.INVALID_PARAMS
		res.Msg = e.MsgUser[res.Code]
	} else {
		data["user"] = models.GetUser(claims.Uuid)
		data["modules"] = models.GetRoleModules(uuid.FromStringOrNil(claims.RoleUuid))
		res.Data = data
	}
	c.JSON(http.StatusOK, res)
}

func AddUser(c *gin.Context) {
	name := c.PostForm("name")
	password := c.DefaultPostForm("password", "123456")
	avatar := c.PostForm("avatar")
	roleUuid := c.DefaultPostForm("roleUuid", "")
	valid := validation.Validation{}
	valid.Required(name, "name").Message("登录名为空")
	valid.Required(password, "password").Message("密码不能为空")
	valid.Required(avatar, "avatar").Message("头像不能为空")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		salts := models.AuthUserName(name)
		if salts == "" {
			salt := com.ToStr(time.Now().Unix())
			password = util.GeneratePwd(password, salt)
			user := &models.User{
				Name:     name,
				Password: password,
				Avatar:   avatar,
				RoleUuid: roleUuid,
				Salt:     salt,
			}
			b := models.AddUser(user)
			if b {
				code = e.SUCCESS
			} else {
				code = e.ERROR
			}
		} else {
			code = e.ERROR_EXIST_USER
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

func EditUser(c *gin.Context) {
	id := c.PostForm("id")
	name := c.PostForm("name")
	avatar := c.PostForm("avatar")
	roleUuid := c.PostForm("roleUuid")
	valid := validation.Validation{}
	valid.Required(id, "id").Message("用户ID不能为空")
	valid.Required(name, "name").Message("登录名为空")
	valid.Required(avatar, "avatar").Message("头像不能为空")
	valid.Required(roleUuid, "roleUuid").Message("角色ID不能为空")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		user := &models.User{
			Name:     name,
			Avatar:   avatar,
			RoleUuid: roleUuid,
		}
		b := models.EditUser(user, id)
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

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	valid := validation.Validation{}
	valid.Required(id, "id").Message("用户ID不能为空")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		b := models.DeleteUser(id)
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

func AuthUser(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	valid := validation.Validation{}
	valid.Required(name, "name").Message("登录名为空")
	valid.Required(password, "password").Message("密码不能为空")
	code := e.INVALID_PARAMS
	data := make(map[string]interface{})
	if !valid.HasErrors() {
		salt := models.AuthUserName(name)
		if salt != "" {
			password = util.GeneratePwd(password, salt)
			b, user := models.AuthUser(name, password)
			if b {
				modules := models.GetRoleModules(uuid.FromStringOrNil(user.RoleUuid))
				//验证成功
				token, _ := util.GenerateToken(user)
				code = e.SUCCESS
				data["token"] = token
				data["modules"] = modules
				data["user"] = user
			} else {
				code = e.ERROR_PASSWORD
			}
		} else {
			code = e.ERROR_NOT_EXIST_USER
		}
	} else {
		util.LoopLog(valid.Errors)
	}
	res := &util.Res{
		Code: code,
		Msg:  e.MsgUser[code],
	}
	if data != nil {
		res.Data = data
	}
	c.JSON(http.StatusOK, res)
}