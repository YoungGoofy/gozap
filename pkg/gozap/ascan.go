package gozap

import (
	"errors"
	"fmt"
	"github.com/YoungGoofy/gozap/pkg/gozap/ascan"
	"github.com/YoungGoofy/gozap/pkg/gozap/utils"
	"log"
	"net/http"
)

type ActiveScanner struct {
	scanner   Scan
	sessionId string
}

func NewActiveScanner(s Scan) *ActiveScanner {
	sessionId := "0"
	return &ActiveScanner{scanner: s, sessionId: sessionId}
}

func (as *ActiveScanner) GetSessionId() error {
	sessionId, err := ascan.GetSessionId(as.scanner.apiKey, as.scanner.url)
	if err != nil {
		return err
	}
	if err = utils.PostSessionCount(sessionId, "ascan"); err != nil {
		return err
	}
	as.sessionId = sessionId
	return nil
}

func (as *ActiveScanner) StopScan() error {
	if as.sessionId == "" {
		return errors.New("any session not found")
	}
	if status := ascan.EditScan(as.scanner.apiKey, as.sessionId, "stop"); status == http.StatusOK {
		log.Printf("\nGood stop\nStatus: %d", status)
		return nil
	} else {
		return errors.New(fmt.Sprintf("\nBad stop\nStatus: %d", status))
	}
}

func (as *ActiveScanner) PauseScan() error {
	if as.sessionId == "" {
		return errors.New("any session not found")
	}
	if status := ascan.EditScan(as.scanner.apiKey, as.sessionId, "pause"); status == http.StatusOK {
		log.Printf("\nGood pause\nStatus: %d", status)
		return nil
	} else {
		return errors.New(fmt.Sprintf("\nBad pause\nStatus: %d", status))
	}
}

func (as *ActiveScanner) ResumeScan() error {
	if as.sessionId == "" {
		return errors.New("any session not found")
	}
	if status := ascan.EditScan(as.scanner.apiKey, as.sessionId, "resume"); status == http.StatusOK {
		log.Printf("\nGood resume\nStatus: %d", status)
		return nil
	} else {
		return errors.New(fmt.Sprintf("\nBad resume\nStatus: %d", status))
	}
}

func (as *ActiveScanner) GetStatus() (string, error) {
	if as.sessionId == "" {
		return "", errors.New("any session not found")
	}
	if status, err := ascan.GetStatus(as.scanner.apiKey, as.sessionId); err != nil {
		return "", err
	} else {
		return status, nil
	}
}

func (as *ActiveScanner) GetAlertIds() ([]string, error) {
	if as.sessionId == "" {
		return nil, errors.New("any session not found")
	}
	ids, err := ascan.GetAlertsId(as.scanner.apiKey, as.sessionId)
	if err != nil {
		return nil, err
	}
	return ids, nil
}
