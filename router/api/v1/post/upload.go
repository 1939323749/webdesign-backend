package post

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Upload(c *gin.Context) {
	file, _ := c.FormFile("file")

	err := c.SaveUploadedFile(file, "./asset/"+file.Filename)
	if err != nil {
		c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
	}
}
