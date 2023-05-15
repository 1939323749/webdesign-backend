package post

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mygo/router/api/v1/get"
	"net/http"
	"time"
)

type Contribution struct {
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
		return
	}

	// 连接到 MongoDB
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer client.Disconnect(ctx)

	// 获取 bilibili_contributions 集合
	collection := client.Database("test").Collection("bilibili_contributions")

	cover, description, err := get.Bilibiliparser(body.Url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse the URL"})
		return
	}
	// 创建新的 Contribution
	contribution := Contribution{
		Category:    body.Category,
		Cover:       cover,
		Description: description,
		ID:          0, // You need to extract this from the request body
		Vote:        0, // You need to extract this from the request body
	}

	// 将 Contribution 插入到集合中
	_, err = collection.InsertOne(ctx, contribution)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
