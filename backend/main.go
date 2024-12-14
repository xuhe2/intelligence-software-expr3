package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/xuhe2/intelligence-software-expr3/controllers"
	"github.com/xuhe2/intelligence-software-expr3/initializers"
)

func main() {
	r := gin.Default()
	// CORS
	// 配置 CORS 中间件
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // 允许访问的域名
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/",controllers.Root)

	r.Run(":3001") // 监听并在 0.0.0.0:3001 上启动服务
}

func init(){
	initializers.LoadEnvVar(".env")
	initializers.ConnectToDB()
}
