## Docs
### Models
##### StatusResult
Model for saving status of working response.
```go
type StatusResult struct {  
    Status string `json:"status"`  
}
```

##### UrlsInScope
Model for saving detail information about urls from web-site
```go
type UrlsInScope struct {  
    Processed          string `json:"processed"`  
    StatusReason       string `json:"statusReason"`  
    Method             string `json:"method"`  
    ReasonNotProcessed string `json:"reasonNotProcessed"`  
    MessageID          string `json:"messageId"`  
    URL                string `json:"url"`  
    StatusCode         string `json:"statusCode"`  
}
```

##### ScanProgress
Model for saving info about scanners in active Scanner
```go
type ScanProgress struct {  
    ScanProgress []interface{}  
}
```

##### Alert
Model for saving detail info about alerts
```go
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
```

***
### Funcs
##### NewMainScan
Create new main scanner
```go
NewMainScan(url, apiKey string) *MainScan
```

##### NewActiveScanner
Create new active scanner
```go
NewActiveScanner(s MainScan) *ActiveScanner
```
##### NewSpider
Create new spider
```go
NewSpider(scanner MainScan) *Spider
```

***
## Types 
#### MainScan
```go
type MainScan struct {  
    url    string  
    apiKey string  
}
```

##### MainScan.AddUrl
Add url for MainScan
```go
AddUrl(url string)
```

##### MainScan.AddApiKey
Add api key for MainScan
```go
AddApiKey(apiKey string)
```

##### MainScan.CountOfAlerts
Get count of alerts
```go
CountOfAlerts() (string, error)
```

##### MainScan.GetAlert
Get one alert by id
```go
GetAlert(alertId string) (models.AlertDetail, error)
```

##### MainScan.GetAlerts
Get some alerts
```go
GetAlerts(start, count string) (models.ListOfAlerts, error)
```

***
#### ActiveScanner
```go
type ActiveScanner struct {  
    scanner   MainScan  
    sessionId string  
}
```

##### ActiveScanner.StartActiveScan
Start new session of active scanning
```go
StartActiveScan() error
```

##### ActiveScanner.StopScan
Stop scanning all scanners
```go
StopScan() error
```

##### ActiveScanner.PauseScan
Pause scanning for all scanners
```go
PauseScan() error
```

##### ActiveScanner.ResumeScan
Resume scanning for all scanners
```go
ResumeScan() error
```

##### ActiveScanner.GetStatus
Get status for active scanning
```go
GetStatus() (string, error)
```

##### ActiveScanner.GetAlertIds
Get ids of alerts after active scanning
```go
GetAlertIds() ([]string, error)
```

##### ActiveScanner.ScanProgress
Get detailed info about different scanners
```go
ScanProgress() ([]interface{}, error)
```

##### ActiveScanner.SkipScanner
Skip scanning any scanner
```go
SkipScanner(pluginId string) (string, error)
```

***
#### Spider
```go
Spider struct {  
    scanner   MainScan  
    sessionId string  
}
```

##### Spider.StartPassiveScan
Start new session for passive scanning
```go
StartPassiveScan() error
```

##### Spider.GetStatus
Get status of running spider
```go
GetStatus() (string, error)
```

##### Spider.GetResult
Get urls after scraping service
```go
GetResult() (UrlsInScope, error)
```

##### Spider.AsyncGetResult
Async method for getting urls after scraping service
```go
AsyncGetResult(ch chan<- UrlsInScope, errCh chan<- error, statusCh chan string, done <-chan struct{})
```

##### Spider.StopScan
Stop spider
```go
StopScan() error
```

##### Spider.PauseScan
Pause spider
```go
PauseScan() error
```

##### Spider.ResumeScan
Resume spider
```go
ResumeScan() error
```

***
### Example
```go
package main  
  
import (  
    "fmt"  
	"github.com/YoungGoofy/gozap/pkg/gozap"
	"github.com/YoungGoofy/gozap/pkg/gozap/utils"
	"log"    
	"sync"
	)  
  
func main() {  
    scan := gozap.NewMainScan(utils.GetDataFromConf())  
    spider := gozap.NewSpider(*scan)  
    dataCh := make(chan gozap.UrlsInScope)  
    errCh := make(chan error)  
    done := make(chan struct{})  
    statusCh := make(chan string)  
    var wg sync.WaitGroup  
  
    if err := spider.StartPassiveScan(); err != nil {  
       log.Fatal(err)  
    }  
  
    wg.Add(1)  
    go func() {  
       defer wg.Done()  
       spider.AsyncGetResult(dataCh, errCh, statusCh, done)  
    }()  
  
    var count int  
    for {  
       select {  
       case urls := <-dataCh:  
          for _, url := range urls {  
             fmt.Println("Data: ", url)  
             count += 1  
          }  
       case <-errCh:  
          fmt.Println("Error: ", <-errCh)  
       case status := <-statusCh:  
          if status == "100" {  
             close(done)  
          }  
       case <-done:  
          wg.Wait()  
          close(dataCh)  
          close(errCh)  
          fmt.Println(count)  
          return  
       }  
    }  
}
```
