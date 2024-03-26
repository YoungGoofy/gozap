package gozap

import (
	"fmt"
	"github.com/pelletier/go-toml"
)

func getDataFromConf() (string, string) {
	conf, err := toml.LoadFile("configs/config.toml")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	key := conf.Get("api.key").(string)
	url := conf.Get("api.urls.url").(string)
	return url, key
}
