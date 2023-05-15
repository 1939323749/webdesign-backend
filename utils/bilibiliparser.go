package utils

import (
	"github.com/valyala/fastjson"
	"io/ioutil"
	"net/http"
	"strings"
)

func Bilibiliparser(url string) (string, string, string, error) {
	bvid := extractBvidFromUrl(url)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.bilibili.com/x/web-interface/view?bvid="+bvid, nil)
	if err != nil {
		return "", "", "", err
	}
	response, err := client.Do(req)
	if err != nil {
		return "", "", "", err
	}
	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", "", "", err
	}

	title := fastjson.MustParseBytes(res).Get("data").Get("title").GetStringBytes()
	pic := fastjson.MustParseBytes(res).Get("data").Get("pic").GetStringBytes()

	return string(bvid), string(title), string(pic), nil
}

func extractBvidFromUrl(url string) string {
	// 示例解析逻辑：从"https://www.bilibili.com/video/BV11s4y1Q7Yf/?spm_id_from=333.934.0.0&vd_source=7d38dd52ce866e656fec20561e6ad46d"提取"BV11s4y1Q7Yf"
	bvid := ""

	// 提取URL中的BV号
	startIndex := strings.Index(url, "/BV")
	if startIndex != -1 {
		endIndex := strings.Index(url[startIndex+1:], "/")
		if endIndex != -1 {
			bvid = url[startIndex+1 : startIndex+endIndex+1]
		} else {
			bvid = url[startIndex+1:]
		}
	}

	return bvid
}
