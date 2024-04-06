package gozap

type Scan struct {
	url    string
	apiKey string
}

func NewScan(url, apiKey string) *Scan {
	return &Scan{url: url, apiKey: apiKey}
}

func (s *Scan) AddUrl(url string) {
	s.url = url
}

func (s *Scan) AddApiKey(apiKey string) {
	s.apiKey = apiKey
}
