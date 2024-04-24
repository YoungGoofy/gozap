package models

type UrlsInScope struct {
	Processed          string `json:"processed"`
	StatusReason       string `json:"statusReason"`
	Method             string `json:"method"`
	ReasonNotProcessed string `json:"reasonNotProcessed"`
	MessageID          string `json:"messageId"`
	URL                string `json:"url"`
	StatusCode         string `json:"statusCode"`
}

type Result struct {
	FullResults []struct {
		UrlsInScope    []UrlsInScope `json:"urlsInScope"`
		UrlsOutOfScope []string      `json:"urlsOutOfScope"`
		UrlsIoError    []interface{} `json:"urlsIoError"`
	} `json:"fullResults"`
}

type StatusResult struct {
	Status string `json:"status"`
}

type ScanProgress struct {
	ScanProgress []interface{}
}

type Alert struct {
	SourceId    string            `json:"sourceid"`
	Other       string            `json:"other"`
	Method      string            `json:"method"`
	Evidence    string            `json:"evidence"`
	PluginID    string            `json:"pluginId"`
	CweId       string            `json:"cweid"`
	Confidence  string            `json:"confidence"`
	WascId      string            `json:"wascid"`
	Description string            `json:"description"`
	MessageID   string            `json:"messageId"`
	InputVector string            `json:"inputVector"`
	URL         string            `json:"url"`
	Tags        map[string]string `json:"tags"`
	Reference   string            `json:"reference"`
	Solution    string            `json:"solution"`
	Alert       string            `json:"alert"`
	Param       string            `json:"param"`
	Attack      string            `json:"attack"`
	Name        string            `json:"name"`
	Risk        string            `json:"risk"`
	ID          string            `json:"id"`
	AlertRef    string            `json:"alertRef"`
}

type AlertDetail struct {
	Alert Alert `json:"alert"`
}

type ListOfAlerts struct {
	Alert []Alert `json:"alerts"`
}

type HostProgress struct {
	Plugins []Plugin
}
type Plugin struct {
	PluginName   string
	PluginID     string
	PluginStatus string
}
