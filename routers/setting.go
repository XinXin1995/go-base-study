package routers

import (
	set "blog/controller/setting"
	"github.com/gin-gonic/gin"
)

func initSettingRouters(api *gin.RouterGroup) {
	//user
	api.GET("/users", set.GetUsers)
	api.POST("/user", set.AddUser)
	api.PUT("/user", set.EditUser)
	api.DELETE("/user/:id", set.DeleteUser)
	api.GET("/user", set.GetUser)

	//role
	api.GET("/roles", set.GetRoles)
	api.POST("/role", set.AddRole)
	api.PUT("/role/:id", set.EditRole)
	api.DELETE("/role/:id", set.DeleteRole)
	api.GET("/roles/all", set.GetAllRoles)
	api.POST("/role/moduleAdd", set.AddRoleModules)
	api.GET("/role/modules", set.GetRoleModules)

	//module
	api.POST("/module", set.AddModule)
	api.GET("/modules", set.GetModules)
	api.GET("/modules/all", set.GetAllModules)
	api.DELETE("/module/:id", set.DeleteModule)
	api.PUT("/module/:id", set.EditModule)
	api.POST("/module/api", set.AddModuleApis)
	api.GET("/module/apis", set.ModuleApis)
	//api
	api.POST("/api", set.AddApi)
	api.GET("/apis", set.GetApis)
	api.GET("/apis/all", set.GetAllApis)
	api.DELETE("/api/:id", set.DeleteApi)
	api.PUT("/api/:id", set.EditApi)
}
