package functions

import (
	"fmt"
	"io/ioutil"
	"log"
	"sync"
)

// DoProject går igenom ett project
func DoProject(list *Work, index int, wg *sync.WaitGroup, wgUI *sync.WaitGroup) {
	defer wg.Done() // defer is done at the end before return

	WalkHTML(&list.Projects[index])
	wgUI.Done()

	if len(list.Projects[index].HTMLs) == 0 {
		log.Printf("%s has no html files!", list.Projects[index].FolderName)
		return
	}

	// Will iterate thru all html pages that were found
	for i := 0; i < len(list.Projects[index].HTMLs); i++ {

		html, err := ioutil.ReadFile(list.Projects[index].HTMLs[i].Path)
		if err != nil {
			log.Printf("%s - %v", list.Projects[index].FolderName, err.Error())
			return
		}

		// Post not working for some reason
		// ValidateHTML5Web(string(html), &list.Projects[index].HTMLs[i])

		if Config.ValidateWithHTML5_verySlow {
			err = ValidateHTML5(string(html), &list.Projects[index].HTMLs[i])
			if err != nil {
				log.Printf("[HTML5] Could not validate %s: %v", list.Projects[index].FolderName, err.Error())
			}
		} else {
			list.Projects[index].HTMLs[i].HTML5Verify.ErrorValidating = fmt.Errorf("Did not verify with HTML5 (Turned off in config. Warning is very slow and demanding)")
		}

		/*
			ValidateCSS(string(css), &list.Projects[index].HTMLs[i])
		*/

		err = ValidateHTMLStrict(string(html), &list.Projects[index].HTMLs[i])
		if err != nil {
			log.Printf("[XHTML 1.0 Strict] Could not validate %s: %v", list.Projects[index].FolderName, err.Error())
		}
		// Do html5 verify here

		//log.Println("Validated:", list.Projects[index].HTMLs[i].Path)
	}
	// Mark as done
	list.Projects[index].Done = true
}
