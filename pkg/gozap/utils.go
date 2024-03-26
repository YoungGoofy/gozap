package gozap

import (
	"fmt"
	"github.com/pelletier/go-toml"
)

func GetDataFromConf() (string, string) {
	conf, err := toml.LoadFile("configs/config.toml")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	key := conf.Get("api.key").(string)
	url := conf.Get("api.urls.url").(string)
	return url, key
}

func GetSessionCount() string {
	conf, err := toml.LoadFile("configs/config.toml")
	if err != nil {
		fmt.Println("Error: ", err)
		return ""
	}
	count := conf.Get("session.count").(string)
	return count
}
