package routers

import (
	"blog/controller/space"
	"github.com/gin-gonic/gin"
)

func initSpaceRouters(api *gin.RouterGroup) {
	api.POST("/singleUpload", space.UploadSingle)

}
