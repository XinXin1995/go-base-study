package routers

import (
	"blog/middleware"
	"blog/pkg/setting"
	"blog/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Static("/static", "./static")
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("api/v1")
	{
		apiv1.POST("/user/auth", v1.AuthUser)
		//user
		apiv1.GET("/users", v1.GetUsers)
		apiv1.POST("/user/register", v1.AddUser)
		apiv1.PUT("/user/edit", v1.EditUser)
		apiv1.DELETE("/user/:id", v1.DeleteUser)

		//role
		apiv1.GET("/roles", v1.GetRoles)
		apiv1.POST("/role/add", v1.AddRole)
		apiv1.PUT("/role/:id", v1.EditRole)
		apiv1.DELETE("/role/:id", v1.DeleteRole)
		apiv1.GET("/roles/all", v1.GetAllRoles)
		apiv1.GET("/role/moduleAdd", v1.AddRoelModules)
		apiv1.GET("/role/modules", v1.GetRoleModules)

		//module
		apiv1.POST("/module/add", v1.AddModule)
		apiv1.GET("/modules", v1.GetModules)
		apiv1.GET("/modules/all", v1.GetAllModules)
		apiv1.DELETE("/module/:id", v1.DeleteModule)
		apiv1.PUT("/module/:id", v1.EditModule)
		apiv1.POST("/module/api/add", v1.AddModuleApis)
		apiv1.GET("/module/apis", v1.ModuleApis)
		//api
		apiv1.POST("/api/add", v1.AddApi)
		apiv1.GET("/apis", v1.GetApis)
		apiv1.GET("/apis/all", v1.GetAllApis)
		apiv1.DELETE("/api/:id", v1.DeleteApi)
		apiv1.PUT("/api/:id", v1.EditApi)

		apiv1.POST("/singleUpload", v1.UploadSingle)
	}

	return r
}
