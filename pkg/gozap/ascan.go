package gozap

import (
	"errors"
	"fmt"
	"github.com/YoungGoofy/gozap/pkg/gozap/ascan"
	"github.com/YoungGoofy/gozap/pkg/models"
	"log"
	"net/http"
)

type (
	ActiveScanner struct {
		scanner   MainScan
		sessionId string
	}
)

func NewActiveScanner(s MainScan) *ActiveScanner {
	return &ActiveScanner{scanner: s}
}

func (as *ActiveScanner) StartActiveScan() error {
	sessionId, err := ascan.GetSessionId(as.scanner.apiKey, as.scanner.url)
	if err != nil {
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

func (as *ActiveScanner) ScanProgress() (models.HostProgress, error) {
	if as.sessionId == "" {
		return models.HostProgress{}, errors.New("any session not found")
	}
	scanProgress, err := ascan.ScanProgress(as.scanner.apiKey, as.sessionId)
	if err != nil {
		return models.HostProgress{}, err
	}
	r := convertInterfaceToString(scanProgress.ScanProgress[1].(map[string]interface{})["HostProcess"].([]interface{}))
	return r, nil
}

func convertInterfaceToString(ifaces []interface{}) models.HostProgress {
	batch := make([]models.Plugin, 0, len(ifaces))
	for _, itemI := range ifaces {
		var plugin models.Plugin
		plugin.PluginName = itemI.(map[string]interface{})["Plugin"].([]interface{})[0].(string)
		plugin.PluginID = itemI.(map[string]interface{})["Plugin"].([]interface{})[1].(string)
		plugin.PluginStatus = itemI.(map[string]interface{})["Plugin"].([]interface{})[3].(string)
		batch = append(batch, plugin)
	}
	return models.HostProgress{Plugins: batch}
}

func (as *ActiveScanner) SkipScanner(pluginId string) (string, error) {
	if as.sessionId == "" {
		return "", errors.New("any session not found")
	}
	if status, err := ascan.SkipScanner(as.scanner.apiKey, as.sessionId, pluginId); err != nil {
		return "", err
	} else {
		return status, nil
	}
}
