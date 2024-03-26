package spiders

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetResult(apiKey, sessionId string) (UrlsInScope, error) {
	var result Result
	resp, err := http.Get(
		fmt.Sprintf("http://localhost:8080/JSON/spider/view/fullResults/?apikey=%s&scanId=%s", apiKey, sessionId))
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
	// TODO: remove FullResults[0].UrlsInScope
	outputStruct := result.FullResults[0].UrlsInScope
	return outputStruct, nil
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

	var result StatusResult
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

func decode(item string) (string, error) {
	var data map[string]string
	if err := json.Unmarshal([]byte(item), &data); err != nil {
		return "", errors.New(fmt.Sprintf("error: %s", err))
	}
	val := data["scan"]
	return val, nil
}
