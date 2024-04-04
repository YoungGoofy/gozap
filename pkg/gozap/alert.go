package gozap

import (
	"github.com/YoungGoofy/gozap/pkg/gozap/alerts"
)

func (s *Scan) CountOfAlerts() (string, error) {
	count, err := alerts.CountOfAlerts(s.apiKey, s.url)
	if err != nil {
		return "", err
	}
	return count, err
}

func (s *Scan) GetAlert(alertId string) (alerts.AlertDetail, error) {
	alert, err := alerts.GetAlert(s.apiKey, alertId)
	if err != nil {
		return alerts.AlertDetail{}, err
	}
	return alert, nil
}

func (s *Scan) GetAlerts(start, count string) (alerts.ListOfAlerts, error) {
	alert, err := alerts.GetAlerts(s.apiKey, s.url, start, count)
	if err != nil {
		return alerts.ListOfAlerts{}, err
	}
	return alert, nil
}
