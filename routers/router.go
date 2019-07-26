package routers

import (
	"blog/controller/setting"
	"blog/middleware"
	"blog/pkg/set"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Static("/static", "./static")
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	gin.SetMode(set.RunMode)
	r.POST("/auth", setting.AuthUser)
	apiv1 := r.Group("api/v1")
	apiv1.Use(middleware.JWT())
	{
		initSettingRouters(apiv1)
		initSpaceRouters(apiv1)
		initBlogRouters(apiv1)
	}

	return r
}
