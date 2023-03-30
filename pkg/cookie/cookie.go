package cookie

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cookie(c *gin.Context) {

	cookie, err := c.Cookie("gin_cookie")

	if err != nil {
		cookie = "NotSet"
		c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "cookie value:" + cookie,
	})
	fmt.Printf("Cookie value: %s \n", cookie)
}
