package gozap

type ActiveScanner struct {
	scan      Scan
	sessionId string
}

func NewActiveScanner(s Scan) *ActiveScanner {
	sessionId := "0"
	return &ActiveScanner{scan: s, sessionId: sessionId}
}

func (as *ActiveScanner) GetSessionId() error {
	return nil
}

func (as *ActiveScanner) StopScan() error {
	return nil
}

func (as *ActiveScanner) PauseScan() error {
	return nil
}

func (as *ActiveScanner) ResumeScan() error {
	return nil
}

func (as *ActiveScanner) GetStatus() (string, error) {
	return "", nil
}
