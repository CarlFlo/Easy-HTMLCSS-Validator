package functions

import (
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

		err = ValidateHTML(string(html), &list.Projects[index].HTMLs[i])
		if err != nil {
			log.Printf("Could not validate %s: %v", list.Projects[index].FolderName, err.Error())
		}

		//log.Println("Validated:", list.Projects[index].HTMLs[i].Path)
	}
	// Mark as done
	list.Projects[index].Done = true
}
