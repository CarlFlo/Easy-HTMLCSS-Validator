package functions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

// SaveResult saves list to file to be opened by html/js after
func SaveResult(list *Work) {

	// For timing
	list.Timing.EndTime = time.Now()
	list.Timing.Duration = list.Timing.EndTime.Sub(list.Timing.StartTime)

	// Saving file
	jsonDataJSON, _ := json.MarshalIndent(list, "", "   ")

	ioutil.WriteFile(fmt.Sprintf("%s/output.js", getExeDir()), []byte(fmt.Sprintf("var OUTPUT_RESULT = %s", string(jsonDataJSON))), 0644)

	openSite()
}

func openSite() {
	if Config.OpenResultWeb {
		// Displays result in a custom local webpage using html and javascript
		OpenResultPage()
	}
}
