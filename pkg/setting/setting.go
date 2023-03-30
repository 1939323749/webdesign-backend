package setting

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	PORT    uint16
	RUNMODE string
	CONTENT string
)

func Init() {
	viper.SetConfigFile("./conf/conf.json")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	PORT = viper.GetUint16("port")
	RUNMODE = viper.GetString("runmode")
	CONTENT = viper.GetString("content")
}
func SaveContent(string2 string) {
	viper.SetConfigFile("./conf/conf.json")
	viper.Set("CONTENT", string2)
	err := viper.WriteConfig()
	if err != nil {
		fmt.Println(err)
	}
}
