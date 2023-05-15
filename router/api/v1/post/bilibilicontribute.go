package post

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mygo/utils"
	"net/http"
	"time"
)

type Contribution struct {
	Bvid        string `bson:"bvid"`
	Category    string `bson:"category"`
	Cover       string `bson:"cover"`
	Description string `bson:"description"`
	ID          int    `bson:"id"`
	Vote        int    `bson:"vote"`
}
type Body struct {
	Url      string `json:"url"`
	Category string `json:"category"`
}

func Bilibilicontribute(c *gin.Context) {
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	// 连接到 MongoDB
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	defer client.Disconnect(ctx)

	// 获取 video 和 counters 集合
	videoCollection := client.Database("local").Collection("video")
	counterCollection := client.Database("local").Collection("counters")

	bvid, description, cover, err := utils.Bilibiliparser(body.Url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse the URL"})
		fmt.Println(err)
		return
	}

	// 获取当前 videoId
	var counter struct {
		ID  string `bson:"_id"`
		Seq int    `bson:"sequence_value"`
	}
	err = counterCollection.FindOne(ctx, bson.M{"_id": "videoId"}).Decode(&counter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		fmt.Println("find")
		return
	}

	// 创建新的 Contribution
	contribution := Contribution{
		Bvid:        bvid,
		Category:    body.Category,
		Cover:       cover,
		Description: description,
		ID:          counter.Seq,
		Vote:        0,
	}

	// 将 Contribution 插入到 video 集合中
	_, err = videoCollection.InsertOne(ctx, contribution)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	// 更新 videoId
	_, err = counterCollection.UpdateOne(ctx, bson.M{"_id": "videoId"}, bson.M{"$inc": bson.M{"sequence_value": 1}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
