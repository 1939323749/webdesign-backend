package router

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"mygo/pkg/cookie"
	"mygo/pkg/setting"
	get2 "mygo/router/api/v1/get"
	"mygo/router/api/v1/post"
	"mygo/router/api/v1/put"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RUNMODE)

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	apiv1 := r.Group("/api")
	apiv1.GET("/get/ping", get2.Ping)
	apiv1.GET("/bilibili", func(c *gin.Context) {
		get2.Bilibilisearch(c)
		var str string = c.Query("k")
		c.Request.URL.Path = "/img/" + str + ".png"
		fmt.Println()
		r.HandleContext(c)
	})
	apiv1.POST("/post/bili", post.Bilibilicontribute)
	apiv1.GET("/getvideo", get2.GetVideo)
	apiv1.POST("/post/upload", post.Upload)
	apiv1.PUT("/put/content", put.Content)
	r.GET("/cookie", cookie.Cookie)
	return r
}
