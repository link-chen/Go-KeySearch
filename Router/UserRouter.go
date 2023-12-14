package Router

import (
	"WebBack/Controller"
	"github.com/gin-gonic/gin"
)

func InitUserRouter() *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type")

		if c.Request.Method == "OPTIONS" {
			// 对 OPTIONS 请求进行特殊处理或者直接返回 200 状态码
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	})

	SimpleUser := r.Group("/user")
	{
		SimpleUser.POST("/Login", Controller.UserLogin)
		SimpleUser.POST("/Register", Controller.NewCount)
		SimpleUser.GET("/GetClient", Controller.GetClient)

		SimpleUser.GET("/DownLoad", Controller.DownLoadFile)
		SimpleUser.GET("/GetFilesName", Controller.GetAllFilesName)
		SimpleUser.GET("/Test", Controller.TestJson)
		//SimpleUser.GET("/KeyWord", Controller.Search)
		SimpleUser.POST("/DownLoadByIndex", Controller.DownLoadByIndex)

		SimpleUser.POST("/KeyWord", Controller.Search)

		SimpleUser.POST("/UpFile", Controller.ReceiveFile)
	}
	SuperUser := r.Group("/SuperUser")
	{
		SuperUser.GET("/Login", func(c *gin.Context) {
			c.JSON(200, "SuperAdmin-Ok")
		})
	}

	return r
}
