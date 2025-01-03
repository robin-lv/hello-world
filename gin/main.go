package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// 定义一个路由
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	return r
}

// curl -X GET http://localhost:80/hello
func main() {
	r := setupRouter()
	r.Run(":80") // 启动服务，监听 8080 端口
}
