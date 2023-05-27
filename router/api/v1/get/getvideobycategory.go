package get

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

func Getvideobycategory(c *gin.Context) {

	var category string = c.Query("c")
	cli, er := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if er != nil {
		fmt.Println(er)
		return
	}

	ctx := context.TODO()
	err := cli.Connect(ctx)
	if err != nil {
		fmt.Println(er)
		return
	}

	defer cli.Disconnect(ctx)

	db := cli.Database("local")
	collection := db.Collection("video")

	// 执行查询操作，获取匹配类别的视频文档
	cursor, err := collection.Find(ctx, bson.M{"category": category})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cursor.Close(ctx)

	var videos []bson.M
	if err := cursor.All(ctx, &videos); err != nil {
		fmt.Println(err)
		return
	}

	// 格式化视频文档，并构建返回结果
	var formattedVideos []gin.H
	for _, video := range videos {
		bvid := video["bvid"]
		id := video["id"]
		url := video["cover"]
		category := video["category"]
		description := video["description"]
		vote := video["vote"]
		time := video["time"]
		rating := video["rating"]
		ratingCount := video["ratingCount"]

		formattedVideo := gin.H{
			"bvid":        bvid,
			"id":          id,
			"url":         url,
			"category":    category,
			"description": description,
			"vote":        vote,
			"time":        time,
			"rating":      rating,
			"ratingCount": ratingCount,
		}

		formattedVideos = append(formattedVideos, formattedVideo)
	}

	c.JSON(http.StatusOK, formattedVideos)
}
