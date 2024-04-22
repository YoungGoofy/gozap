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
