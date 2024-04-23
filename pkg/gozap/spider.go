package gozap

import (
	"errors"
	"fmt"
	"github.com/YoungGoofy/gozap/pkg/gozap/spiders"
	"github.com/YoungGoofy/gozap/pkg/models"
	"log"
	"net/http"
)

type (
	Spider struct {
		scanner   MainScan
		sessionId string
	}
	UrlsInScope []models.UrlsInScope
)

func NewSpider(scanner MainScan) *Spider {
	return &Spider{scanner: scanner}
}

func (s *Spider) StartPassiveScan() error {
	id, err := spiders.GetConnectionId(s.scanner.apiKey, s.scanner.url)
	if err != nil {
		return err
	}
	s.sessionId = id
	return nil
}

func (s *Spider) GetStatus() (string, error) {
	if s.sessionId == "" {
		return "", errors.New("any session not found")
	}
	if status, err := spiders.GetStatus(s.scanner.apiKey, s.sessionId); err != nil {
		return "", err
	} else {
		return status, nil
	}
}

func (s *Spider) GetResult() (UrlsInScope, error) {
	if result, err := spiders.GetResult(s.scanner.apiKey, s.sessionId); err != nil {
		return nil, err
	} else {
		urlsInScope := result.FullResults[0].UrlsInScope
		return urlsInScope, nil
	}
}

func (s *Spider) AsyncGetResult(ch chan<- UrlsInScope, errCh chan<- error, statusCh chan string, done <-chan struct{}) {
	var lastUrl UrlsInScope
	maxCount := 0
	minCount := 0
	for {
		select {
		case <-done:
			return
		default:
			result, err := spiders.GetResult(s.scanner.apiKey, s.sessionId)
			if err != nil {
				errCh <- err
			}
			urlsInScope := result.FullResults[0].UrlsInScope
			if len(urlsInScope) > maxCount {
				maxCount = len(urlsInScope)
			} else {
				continue
			}
			if len(urlsInScope) > 0 {
				lastUrl = urlsInScope[minCount : maxCount-1]
				ch <- lastUrl
			}
			minCount = maxCount - 1
		}
		status, err := spiders.GetStatus(s.scanner.apiKey, s.sessionId)
		if err != nil {
			errCh <- err
		}
		statusCh <- status
	}
}

func (s *Spider) StopScan() error {
	if s.sessionId == "" {
		return errors.New("any session not found")
	}
	if status := spiders.EditScan(s.scanner.apiKey, s.sessionId, "stop"); status == http.StatusOK {
		log.Printf("\nGood stop\nStatus: %d", status)
		return nil
	} else {
		return errors.New(fmt.Sprintf("\nBad stop\nStatus: %d", status))
	}
}

func (s *Spider) PauseScan() error {
	if s.sessionId == "" {
		return errors.New("any session not found")
	}
	if status := spiders.EditScan(s.scanner.apiKey, s.sessionId, "pause"); status == http.StatusOK {
		log.Printf("\nGood pause\nStatus: %d", status)
		return nil
	} else {
		return errors.New(fmt.Sprintf("\nBad pause\nStatus: %d", status))
	}
}

func (s *Spider) ResumeScan() error {
	if s.sessionId == "" {
		return errors.New("any session not found")
	}
	if status := spiders.EditScan(s.scanner.apiKey, s.sessionId, "resume"); status == http.StatusOK {
		log.Printf("\nGood resume\nStatus: %d", status)
		return nil
	} else {
		return errors.New(fmt.Sprintf("\nBad resume\nStatus: %d", status))
	}
}
