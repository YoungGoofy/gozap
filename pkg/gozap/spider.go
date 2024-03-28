package gozap

import (
	"errors"
	"fmt"
	"github.com/YoungTreezy/gozap/pkg/gozap/spiders"
	"log"
	"net/http"
)

type Spider struct {
	scanner   Scan
	sessionId string
}

func NewSpider(scanner Scan) *Spider {
	sessionId, err := GetSessionCount()
	if err != nil {
		log.Println(err)
		return nil
	}
	return &Spider{scanner: scanner, sessionId: sessionId}
}

func (s *Spider) GetConnectionId() error {
	id, err := spiders.GetConnectionId(s.scanner.apiKey, s.scanner.url)
	if err = PostSessionCount(id); err != nil {
		return err
	}
	s.sessionId = id
	if err != nil {
		return err
	}
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

func (s *Spider) GetResult() (spiders.UrlsInScope, error) {
	if s.sessionId == "" {
		return nil, errors.New("any session not found")
	}
	if result, err := spiders.GetResult(s.scanner.apiKey, s.sessionId); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

func (s *Spider) StopScan() error {
	if s.sessionId == "" {
		return errors.New("any session not found")
	}
	if status := spiders.EditScan(s.scanner.apiKey, s.sessionId, "stop"); status == http.StatusOK {
		log.Printf("Status: %d", status)
		return nil
	} else {
		return errors.New(fmt.Sprintf("Status: %d", status))
	}
}

func (s *Spider) PauseScan() error {
	if s.sessionId == "" {
		return errors.New("any session not found")
	}
	if status := spiders.EditScan(s.scanner.apiKey, s.sessionId, "pause"); status == http.StatusOK {
		log.Printf("Status: %d", status)
		return nil
	} else {
		return errors.New(fmt.Sprintf("Status: %d", status))
	}
}

func (s *Spider) ResumeScan() error {
	if s.sessionId == "" {
		return errors.New("any session not found")
	}
	if status := spiders.EditScan(s.scanner.apiKey, s.sessionId, "resume"); status == http.StatusOK {
		log.Printf("Status: %d", status)
		return nil
	} else {
		return errors.New(fmt.Sprintf("Status: %d", status))
	}
}
