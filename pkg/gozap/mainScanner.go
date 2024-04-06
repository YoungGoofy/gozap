package gozap

type MainScan struct {
	url    string
	apiKey string
}

func NewMainScan(url, apiKey string) *MainScan {
	return &MainScan{url: url, apiKey: apiKey}
}

func (s *MainScan) AddUrl(url string) {
	s.url = url
}

func (s *MainScan) AddApiKey(apiKey string) {
	s.apiKey = apiKey
}
