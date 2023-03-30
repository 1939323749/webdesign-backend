package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mygo/pkg/cookie"
	"mygo/pkg/setting"
	"mygo/router/api/get"
	"mygo/router/api/post"
	"mygo/router/api/put"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RUNMODE)
	apiv1 := r.Group("/api")
	apiv1.GET("/get/ping", get.Ping)
	apiv1.GET("/bilibili", func(c *gin.Context) {
		get.Bilibilisearch(c)
		var str string = c.Query("k")
		c.Request.URL.Path = "/img/" + str + ".png"
		fmt.Println()
		r.HandleContext(c)
	})
	apiv1.GET("/bp", get.Bilibiliparser)
	apiv1.POST("/post/upload", post.Upload)
	apiv1.PUT("/put/content", put.Content)
	r.GET("/cookie", cookie.Cookie)
	return r
}
