package put

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
)

type Updaterating struct {
	Bvid        string    `bson:"bvid"`
	Category    string    `bson:"category"`
	Cover       string    `bson:"cover"`
	Description string    `bson:"description"`
	ID          int       `bson:"id"`
	Time        time.Time `bson:"time"`
	Rating      float32   `bson:"rating"`
	RatingCount int       `bson:"ratingCount"`
}

func UpdateRating(c *gin.Context) {
	var update Updaterating
	if err := c.BindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	// Connect to MongoDB
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	defer client.Disconnect(ctx)

	// Get the video collection
	videoCollection := client.Database("local").Collection("video")

	// Get current video's rating and rating count
	var currentVideo Updaterating
	err = videoCollection.FindOne(ctx, bson.M{"bvid": update.Bvid}).Decode(&currentVideo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	// Calculate new rating
	newRating := (currentVideo.Rating*float32(currentVideo.RatingCount) + update.Rating) / float32(currentVideo.RatingCount+1)

	// Update the rating and rating count in the video collection
	filter := bson.M{"bvid": update.Bvid}                                                                  // Filter by bvid
	updateBson := bson.M{"$set": bson.M{"rating": newRating, "ratingCount": currentVideo.RatingCount + 1}} // Set the new rating and increment rating count

	_, err = videoCollection.UpdateOne(ctx, filter, updateBson)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
