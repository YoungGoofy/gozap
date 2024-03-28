package ascan

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetSessionId(apiKey, url string) (string, error) {
	var result map[string]string
	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/JSON/ascan/action/scan/?apikey=%s&url=%s&recurse=&inScopeOnly=&scanPolicyName=&method=&postData=&contextId=", apiKey, url))
	if err != nil {
		return "", errors.New(fmt.Sprintf("bad request: %v", err))
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return "", errors.New(fmt.Sprintf("bad unmarshal: %s", err))
	}
	return result["scan"], nil
}

func EditScan(apiKey, url, action string) int {
	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/JSON/ascan/action/%s/?apikey=%s&scanId=%s", action, apiKey, url))
	if err != nil {
		log.Fatal("Bad request")
		return -1
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusBadRequest {
		return http.StatusBadRequest
	}
	return http.StatusOK
}

func GetStatus(apiKey, sessionId string) (string, error) {
	var status map[string]string
	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/JSON/ascan/view/status/?apikey=%s&scanId=%s", apiKey, sessionId))
	if err != nil {
		return "", errors.New(fmt.Sprintf("bad request: %v", err))
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if err = json.Unmarshal(body, &status); err != nil {
		return "", errors.New(fmt.Sprintf("bad unmarshal: %s", err))
	}
	return status["status"], nil
}

func GetAlertsId(apiKey, sessionId string) ([]string, error) {
	var ids map[string][]string
	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/JSON/ascan/view/alertsIds/?apikey=%s&scanId=%s", apiKey, sessionId))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("bad request: %v", err))
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, &ids); err != nil {
		return nil, errors.New(fmt.Sprintf("bad unmarshal: %s", err))
	}
	return ids["alertsIds"], nil
}
