package scanner

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/YoungTreezy/gozap/pkg/gozap/spiders"
	"io/ioutil"
	"net/http"
)

type SpiderScan struct {
	url       string
	apiKey    string
	SessionId string
}

func NewSpiderScan(url, apiKey string) *SpiderScan {
	return &SpiderScan{url: url, apiKey: apiKey, SessionId: ""}
}

func (s *SpiderScan) GetResult() (spiders.UrlsInScope, error) {
	var result spiders.Result
	resp, err := http.Get(
		fmt.Sprintf("http://localhost:8080/JSON/spider/view/fullResults/?apikey=%s&scanId=%s", s.apiKey, s.SessionId))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error: %s", err))
	}
	if err = json.Unmarshal([]byte(string(body)), &result); err != nil {
		return nil, errors.New(fmt.Sprintf("error: %s", err))
	}
	outputStruct := result.FullResults[0].UrlsInScope
	return outputStruct, nil
}

func (s *SpiderScan) GetStatus() (string, error) {
	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/JSON/spider/view/status/?apikey=%s&scanId=%s", s.apiKey, s.SessionId))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result spiders.StatusResult
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	return result.Status, nil
}

func (s *SpiderScan) GetConnectionId() error {
	resp, err := http.Get(
		fmt.Sprintf("http://localhost:8080/JSON/spider/action/scan/?apikey=%s&url=%s&contextName=&recurse=", s.apiKey, s.url))
	if err != nil {
		return errors.New(fmt.Sprintf("error: %s", err))
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New(fmt.Sprintf("error: %s", err))
	}
	fmt.Println("Answer: ", string(body))
	id, err := decode(string(body))
	if err != nil {
		return errors.New(fmt.Sprintf("error: %s", err))
	}
	s.SessionId = id
	return nil
}

func decode(item string) (string, error) {
	var data map[string]string
	if err := json.Unmarshal([]byte(item), &data); err != nil {
		return "", errors.New(fmt.Sprintf("error: %s", err))
	}
	val := data["scan"]
	return val, nil
}
