package main

import (
	"fmt"
	"mygo/pkg/setting"
	"mygo/pkg/static"
	"mygo/router"
	"net/http"
	"strconv"
)

func main() {
	setting.Init()
	r := router.InitRouter()
	static.Init(r)
	fmt.Println(strconv.Itoa(int(setting.PORT)))
	http.ListenAndServe(":"+strconv.Itoa(int(setting.PORT)), r)
}
