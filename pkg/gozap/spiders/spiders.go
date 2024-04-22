package spiders

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/YoungGoofy/gozap/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
)

func GetResult(apiKey, sessionId string) (models.Result, error) {
	var result models.Result
	resp, err := http.Get(
		fmt.Sprintf("http://localhost:8080/JSON/spider/view/fullResults/?apikey=%s&scanId=%s", apiKey, sessionId))
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, errors.New(fmt.Sprintf("error: %s", err))
	}
	if err = json.Unmarshal([]byte(body), &result); err != nil {
		return result, errors.New(fmt.Sprintf("error: %s", err))
	}
	return result, nil
}

func GetStatus(apiKey, sessionId string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/JSON/spider/view/status/?apikey=%s&scanId=%s", apiKey, sessionId))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result models.StatusResult
	if err = json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	return result.Status, nil
}

func GetConnectionId(apiKey, url string) (string, error) {
	resp, err := http.Get(
		fmt.Sprintf("http://localhost:8080/JSON/spider/action/scan/?apikey=%s&url=%s&contextName=&recurse=", apiKey, url))
	if err != nil {
		return "", errors.New(fmt.Sprintf("error: %s", err))
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New(fmt.Sprintf("error: %s", err))
	}
	fmt.Println("Answer: ", string(body))
	id, err := decode(string(body))
	if err != nil {
		return "", errors.New(fmt.Sprintf("error: %s", err))
	}
	return id, nil
}

func EditScan(apiKey, sessionId, action string) int {
	resp, err := http.Get(
		fmt.Sprintf("http://localhost:8080/JSON/spider/action/%s/?apikey=%s&scanId=%s", action, apiKey, sessionId))
	if err != nil {
		log.Println(errors.New("bad request in StopScan"))
		return -1
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusBadRequest {
		return http.StatusBadRequest
	}
	return http.StatusOK
}

func decode(item string) (string, error) {
	var data map[string]string
	if err := json.Unmarshal([]byte(item), &data); err != nil {
		return "", errors.New(fmt.Sprintf("error: %s", err))
	}
	val := data["scan"]
	return val, nil
}
