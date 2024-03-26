package gozap

import (
	"errors"
	"github.com/YoungTreezy/gozap/pkg/gozap/spiders"
)

type SpiderScan struct {
	url       string
	apiKey    string
	sessionId string
}

func NewSpiderScan(url, apiKey string) *SpiderScan {
	return &SpiderScan{url: url, apiKey: apiKey, sessionId: ""}
}

func (s *SpiderScan) GetConnectionId() error {
	id, err := spiders.GetConnectionId(s.apiKey, s.url)
	s.sessionId = id
	if err != nil {
		return err
	}
	return nil
}

func (s *SpiderScan) GetStatus() (string, error) {
	if s.sessionId == "" {
		return "", errors.New("any session not found")
	}
	if status, err := spiders.GetStatus(s.apiKey, s.sessionId); err != nil {
		return "", err
	} else {
		return status, nil
	}
}

func (s *SpiderScan) GetFullResult() (spiders.UrlsInScope, error) {
	if s.sessionId == "" {
		return nil, errors.New("any session not found")
	}
	if result, err := spiders.GetResult(s.apiKey, s.sessionId); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}
