package routers

import (
	"blog/controller/blogs"
	"github.com/gin-gonic/gin"
)

func initBlogRouters(api *gin.RouterGroup) {
	//tag
	api.GET("/tags", blogs.Tags)
	api.GET("/tags/all", blogs.AllTags)
	api.POST("/tag", blogs.AddTag)
	api.PUT("/tags", blogs.EditTag)
	api.DELETE("/tag/:id", blogs.DelTag)

	//category
	api.GET("/categories", blogs.Categories)
	api.GET("/categories/all", blogs.AllCategories)
	api.POST("/category", blogs.AddCategory)
	api.PUT("/category", blogs.EditCategory)
	api.DELETE("/category/:id", blogs.DelCategory)

	//article
	api.GET("/articles", blogs.Articles)
	api.POST("/article", blogs.AddArticle)
	api.PUT("/article", blogs.EditArticle)
	api.DELETE("/article/:id", blogs.DelArticle)

}
