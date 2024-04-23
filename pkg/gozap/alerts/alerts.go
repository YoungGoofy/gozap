package alerts

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/YoungGoofy/gozap/pkg/models"
	"io/ioutil"
	"net/http"
)

func CountOfAlerts(apiKey, baseUrl string) (string, error) {
	count := struct {
		NumberOfAlerts string `json:"numberOfAlerts"`
	}{}
	resp, err := http.Get(
		fmt.Sprintf("http://localhost:8080/JSON/alert/view/numberOfAlerts/?apikey=%s&baseurl=%s", apiKey, baseUrl))
	if err != nil {
		return "", errors.New("bad request in CountOfAlerts")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("no body in CountOfAlerts")
	}
	if err = json.Unmarshal(body, &count); err != nil {
		return "", errors.New("bad unmarshal in CountOfAlerts")
	}
	return count.NumberOfAlerts, nil
}

func GetAlert(apiKey, alertId string) (models.AlertDetail, error) {
	var alert models.AlertDetail
	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/JSON/alert/view/alert/?apikey=%s&id=%s", apiKey, alertId))
	if err != nil {
		return models.AlertDetail{}, errors.New("bad request in GetAlert")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.AlertDetail{}, errors.New("no body in GetAlert")
	}
	if err = json.Unmarshal(body, &alert); err != nil {
		return models.AlertDetail{}, errors.New(fmt.Sprintf("bad unmarshal in GetAlert: %s", err))
	}
	return alert, nil
}

func GetAlerts(apiKey, baseUrl string, start, count string) (models.ListOfAlerts, error) {
	var alerts models.ListOfAlerts
	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/JSON/alert/view/alerts/?apikey=%s&baseurl=%s&start=%s&count=%s", apiKey, baseUrl, start, count))
	if err != nil {
		return models.ListOfAlerts{}, errors.New("bad request in GetAlert")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.ListOfAlerts{}, errors.New("no body in GetAlert")
	}
	if err = json.Unmarshal(body, &alerts); err != nil {
		return models.ListOfAlerts{}, errors.New("bad unmarshal in GetAlert")
	}
	return alerts, nil

}
