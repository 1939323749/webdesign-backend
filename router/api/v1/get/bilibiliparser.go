package get

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/valyala/fastjson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"net/http"
	"strings"
)

func Bilibiliparser(c *gin.Context) (string, string, error) {
	url := c.Query("url")
	bvid := extractBvidFromUrl(url)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.bilibili.com/x/web-interface/view?bvid="+bvid, nil)
	if err != nil {
		fmt.Println(err)
	}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	res, err := ioutil.ReadAll(response.Body)
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

	cli, er := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if er != nil {
		fmt.Println(er)
	}

	ctx := context.TODO()
	err = cli.Connect(ctx)
	if err != nil {
		fmt.Println(er)
	}

	defer cli.Disconnect(ctx)

	db := cli.Database("local")
	collection := db.Collection("video")

	video := bson.D{
		{Key: "id", Value: 1},
		{Key: "cover", Value: string(pic.GetStringBytes())},
		{Key: "description", Value: string(title.GetStringBytes())},
		{Key: "vote", Value: 0},
		{Key: "category", Value: "Your Category"},
	}

	_, err = collection.InsertOne(ctx, video)
	if err != nil {
		print(err)
	}
	return string(title.GetStringBytes()), string(pic.GetStringBytes()), nil
}

func extractBvidFromUrl(url string) string {
	// 示例解析逻辑：从"https://www.bilibili.com/video/BV11s4y1Q7Yf/?spm_id_from=333.934.0.0&vd_source=7d38dd52ce866e656fec20561e6ad46d"提取"BV11s4y1Q7Yf"
	bvid := ""

	// 提取URL中的BV号
	startIndex := strings.Index(url, "/BV")
	if startIndex != -1 {
		endIndex := strings.Index(url[startIndex+1:], "/")
		if endIndex != -1 {
			bvid = url[startIndex+1 : startIndex+endIndex+1]
		} else {
			bvid = url[startIndex+1:]
		}
	}

	return bvid
}
