# call idm api with Go

### idm api manual
http://www.internetdownloadmanager.com/support/idm_api.html

### Implementation of Python
https://stackoverflow.com/questions/22587681/use-idminternet-download-manager-api-with-python


### Example
```go
package main

import (
	"log"

	"github.com/greyh4t/idm"
)

func main() {
	lt, err := idm.NewIDMLinkTransmitter2()
	if err != nil {
		log.Fatal(err)
	}

	err = lt.SendLinkToIDM(idm.Link{
		URL:   "http://www.example.com/",
		Flags: idm.FlagSlience,
	})
	if err != nil {
		log.Fatal(err)
	}
}
```