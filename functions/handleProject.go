package functions

import (
	"io/ioutil"
	"log"
	"sync"
)

// DoProject går igenom ett project
func DoProject(folderName string, wg *sync.WaitGroup, mutex *sync.Mutex) {

	defer wg.Done() // defer is done at the end

	htmlPaths := WalkHTML(folderName)

	if len(htmlPaths) == 0 {
		log.Printf("%s has no html files!", folderName)
		return
	}

	for i := 0; i < len(htmlPaths); i++ {

		txt, err := ioutil.ReadFile(htmlPaths[i])
		if err != nil {
			log.Printf("%s - %v", folderName, err.Error())
			return
		}

		err = ValidateHTML(string(txt), htmlPaths[i], i)
		if err != nil {
			log.Printf("Could not validate %s: %v", folderName, err.Error())
		}
		log.Println("Validated:", htmlPaths[i])
	}
}
