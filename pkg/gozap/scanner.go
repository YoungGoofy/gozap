package gozap

type Scan struct {
	url    string
	apiKey string
}

func NewScan(url, apiKey string) *Scan {
	return &Scan{url: url, apiKey: apiKey}
}
