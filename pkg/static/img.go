package static

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init(r *gin.Engine) {
	r.StaticFS("/img", http.Dir("./asset"))
}
