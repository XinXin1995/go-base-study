package main

import (
	"blog/pkg/set"
	"blog/routers"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := routers.InitRouter()
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	_ = router.Run(fmt.Sprintf(":%d", set.HTTPPort))
	//server := &http.Server{
	//	Addr:           fmt.Sprintf("127.0.0.1:%d", setting.HTTPPort),
	//	Handler:        router,
	//	ReadTimeout:    setting.ReadTimeout,
	//	WriteTimeout:   setting.WriteTimeout,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//err := server.ListenAndServe()
	//if err != nil {
	//	fmt.Println("服务启动失败：", err)
	//}
}
