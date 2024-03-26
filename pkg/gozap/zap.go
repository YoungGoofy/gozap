package gozap

import (
	"errors"
	"github.com/YoungTreezy/gozap/pkg/gozap/spiders"
)

var spider = spiders.NewSpiderScan(getDataFromConf())

func RunSpider() error {
	err := spider.GetConnectionId()
	if err != nil {
		return err
	}
	return nil
}

func GetStatus() (string, error) {
	if spider.SessionId == "" {
		return "", errors.New("any session not found")
	}
	if status, err := spider.GetStatus(); err != nil {
		return "", err
	} else {
		return status, nil
	}
}

func GetFullResult() (spiders.UrlsInScope, error) {
	if spider.SessionId == "" {
		return nil, errors.New("any session not found")
	}
	if result, err := spider.GetResult(); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}
