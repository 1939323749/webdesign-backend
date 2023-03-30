package get

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/valyala/fastjson"
	"io/ioutil"
	"net/http"
)

func Bilibiliparser(c *gin.Context) {
	bvid := c.Query("bvid")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.bilibili.com/x/web-interface/view?bvid="+bvid, nil)
	if err != nil {
		fmt.Println(err)
	}
	responce, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	res, err := ioutil.ReadAll(responce.Body)
	if err != nil {
		fmt.Println(err)
	}

	title := fastjson.MustParseBytes(res).Get("data").Get("title")
	pic := fastjson.MustParseBytes(res).Get("data").Get("pic")

	client.CloseIdleConnections()
	c.JSON(http.StatusOK, gin.H{
		"title": string(title.GetStringBytes()),
		"pic":   string(pic.GetStringBytes()),
	})
}
