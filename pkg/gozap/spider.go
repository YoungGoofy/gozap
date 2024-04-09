package gozap

import (
	"errors"
	"fmt"
	"github.com/YoungGoofy/gozap/pkg/gozap/spiders"
	"log"
	"net/http"
)

type (
	Spider struct {
		scanner   MainScan
		sessionId string
	}
	UrlsInScope []spiders.UrlsInScope
)

func NewSpider(scanner MainScan) *Spider {
	//sessionId, err := GetSpiderSessionCount()
	//if err != nil {
	//	log.Println(err)
	//	return nil
	//}
	return &Spider{scanner: scanner /*, sessionId: sessionId*/}
}

func (s *Spider) GetSessionId() error {
	id, err := spiders.GetConnectionId(s.scanner.apiKey, s.scanner.url)
	if err != nil {
		return err
	}
	//if err = PostSessionCount(id, "spider"); err != nil {
	//	return err
	//}
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
	/** This method return results in runtime*/
	var lastUrl UrlsInScope
	maxCount := 0
	minCount := 0
	for {
		select {
		case <-done: // Проверяем, получили ли сигнал о завершении
			return
		default:
			// Отправляем запрос на сервер
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
			// Проверяем, что получены непустые данные
			if len(urlsInScope) > 0 {
				// Сохраняем последний элемент
				lastUrl = urlsInScope[minCount : maxCount-1]
				// Отправляем последний элемент по каналу
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
