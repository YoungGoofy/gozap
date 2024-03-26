package gozap

import (
	"errors"
	"github.com/YoungTreezy/gozap/pkg/gozap/spiders"
)

type Scan struct {
	url       string
	apiKey    string
	sessionId string
}

func NewScan(url, apiKey string) *Scan {
	return &Scan{url: url, apiKey: apiKey, sessionId: "4"}
}

func (s *Scan) GetConnectionId() error {
	id, err := spiders.GetConnectionId(s.apiKey, s.url)
	s.sessionId = id
	if err != nil {
		return err
	}
	return nil
}

func (s *Scan) GetStatus() (string, error) {
	if s.sessionId == "" {
		return "", errors.New("any session not found")
	}
	if status, err := spiders.GetStatus(s.apiKey, s.sessionId); err != nil {
		return "", err
	} else {
		return status, nil
	}
}

func (s *Scan) GetFullResult() (spiders.UrlsInScope, error) {
	if s.sessionId == "" {
		return nil, errors.New("any session not found")
	}
	if result, err := spiders.GetResult(s.apiKey, s.sessionId); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}
