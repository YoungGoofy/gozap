package alerts

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
