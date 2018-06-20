package functions

import (
	"io/ioutil"
	"log"
	"sync"
)

// DoProject går igenom ett project
func DoProject(list *Work, index int, wg *sync.WaitGroup, wgUI *sync.WaitGroup) {

	// Retreves all files in that path and its sub paths
	WalkHTML(&list.Projects[index])
	if !SingleProjectMode {
		defer wg.Done() // defer is done at the end before return
		wgUI.Done()
	}

	if len(list.Projects[index].HTMLs) == 0 {
		log.Printf("%s has no html files!", list.Projects[index].FolderName)
		return
	}

	// CSS
	for i := 0; i < len(list.Projects[index].CSSs); i++ {

		css, err := ioutil.ReadFile(list.Projects[index].CSSs[i].Path)
		if err != nil {
			log.Printf("%s - %v", list.Projects[index].FolderName, err.Error())
			return
		}

		err = validateCSS(string(css), &list.Projects[index].CSSs[i])
		list.Projects[index].CSSs[i].ErrorValidating = err
	}

	// HTML - Iterate
	for i := 0; i < len(list.Projects[index].HTMLs); i++ {

		if list.GracefulStop {
			ExitGoroutine()
		}

		html, err := ioutil.ReadFile(list.Projects[index].HTMLs[i].Path)
		if err != nil {
			log.Printf("%s - %v", list.Projects[index].FolderName, err.Error())
			return
		}

		// HTML5
		err = validateHTML5(string(html), &list.Projects[index].HTMLs[i])
		if err != nil {
			log.Printf("[HTML5] Could not validate %s: %v", list.Projects[index].FolderName, err.Error())
		}
		list.Projects[index].HTMLs[i].HTML5Verify.ErrorValidating = err

		// XHTML 1.0 Strict
		err = ValidateHTMLStrict(string(html), &list.Projects[index].HTMLs[i])
		if err != nil {
			log.Printf("[XHTML 1.0 Strict] Could not validate %s: %v", list.Projects[index].FolderName, err.Error())
		}
	}
	// Mark as done
	list.Projects[index].Done = true
}
