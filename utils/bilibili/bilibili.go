package bilibili

import "github.com/go-rod/rod"

func Bilibili(str string) {
	page := rod.New().MustConnect().MustPage("https://search.bilibili.com/all?keyword=" + str)
	page.MustWaitLoad().MustScreenshot("./asset/" + str + ".png")
}
