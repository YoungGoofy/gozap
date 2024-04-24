package gozap

type MainScan struct {
	url    string
	apiKey string
}

func NewMainScan(args ...string) *MainScan {
	if len(args) == 2 {
		return &MainScan{url: args[0], apiKey: args[1]}
	}
	return &MainScan{}
}

func (s *MainScan) AddUrl(url string) {
	s.url = url
}

func (s *MainScan) AddApiKey(apiKey string) {
	s.apiKey = apiKey
}
