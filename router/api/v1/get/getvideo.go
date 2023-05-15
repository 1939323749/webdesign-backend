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

func GetVideo(c *gin.Context) {
	cli, er := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if er != nil {
		fmt.Println(er)
	}

	ctx := context.TODO()
	err := cli.Connect(ctx)
	if err != nil {
		fmt.Println(er)
	}

	defer cli.Disconnect(ctx)

	db := cli.Database("local")
	collection := db.Collection("video")

	// 执行查询操作，获取所有视频文档
	cursor, err := collection.Find(ctx, bson.M{})
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
		id := video["id"]
		url := video["cover"]
		category := video["category"]
		description := video["description"]

		formattedVideo := gin.H{
			"id":          id,
			"url":         url,
			"category":    category,
			"description": description,
		}

		formattedVideos = append(formattedVideos, formattedVideo)
	}

	c.JSON(http.StatusOK, formattedVideos)
}
