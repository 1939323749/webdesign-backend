package put

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"mygo/pkg/setting"
	"net/http"
)

func Content(c *gin.Context) {
	cont := c.Query("cont")
	valid := validation.Validation{}
	e := valid.Required(cont, "cont").Error
	if e != nil {
		fmt.Println(e)
		fmt.Println(cont)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "fail",
		})
		return
	}
	setting.CONTENT = cont
	setting.SaveContent(cont)
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
