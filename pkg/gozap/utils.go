package gozap

import (
	"errors"
	"fmt"
	"github.com/pelletier/go-toml"
)

//type Scanner interface {
//	Stop()
//	Pause()
//	Resume()
//	GetConnectionId() error
//	GetStatus() (string, error)
//}

func GetDataFromConf() (string, string) {
	conf, err := toml.LoadFile("configs/config.toml")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	key := conf.Get("api.key").(string)
	url := conf.Get("api.urls.url").(string)
	return url, key
}

func GetSpiderSessionCount() (string, error) {
	conf, err := toml.LoadFile("pkg/configs/config.toml")
	if err != nil {
		fmt.Println("Error: ", err)
		return "", err
	}
	count := conf.Get("session.spider.count").(string)
	return count, nil
}

func PostSpiderSessionCount(id string) error {
	conf, err := toml.LoadFile("pkg/configs/config.toml")
	if err != nil {
		return errors.New(fmt.Sprintf("bad connect to file: %s", err))
	}
	conf.Set("session.spider.count", id)
	return nil
}
