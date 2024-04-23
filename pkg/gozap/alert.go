package gozap

import (
	"github.com/YoungGoofy/gozap/pkg/gozap/alerts"
	"github.com/YoungGoofy/gozap/pkg/models"
)

func (s *MainScan) CountOfAlerts() (string, error) {
	count, err := alerts.CountOfAlerts(s.apiKey, s.url)
	if err != nil {
		return "", err
	}
	return count, err
}

func (s *MainScan) GetAlert(alertId string) (models.AlertDetail, error) {
	alert, err := alerts.GetAlert(s.apiKey, alertId)
	if err != nil {
		return models.AlertDetail{}, err
	}
	return alert, nil
}

func (s *MainScan) GetAlerts(start, count string) (models.ListOfAlerts, error) {
	alert, err := alerts.GetAlerts(s.apiKey, s.url, start, count)
	if err != nil {
		return models.ListOfAlerts{}, err
	}
	return alert, nil
}
