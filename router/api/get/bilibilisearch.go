package get

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mygo/utils/bilibili"
	"os"
)

func Bilibilisearch(c *gin.Context) {
	var str string = c.Query("k")
	userFile := "./asset/" + str + ".png"
	fl, err := os.Open(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		bilibili.Bilibili(str)
		return
	}
	defer fl.Close()
}
