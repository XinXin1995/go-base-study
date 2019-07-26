package routers

import (
	"blog/controller/blogs"
	"github.com/gin-gonic/gin"
)

func initBlogRouters(api *gin.RouterGroup) {
	//tag
	api.GET("/tags", blogs.Tags)
	api.POST("/tag", blogs.AddTag)
	api.PUT("/tag", blogs.EditTag)
	api.DELETE("/tag/:id", blogs.DelTag)

	//category
	api.GET("/categories", blogs.Categories)
	api.POST("/category", blogs.AddCategory)
	api.PUT("/category", blogs.EditCategory)
	api.DELETE("/category/:id", blogs.DelCategory)

}
