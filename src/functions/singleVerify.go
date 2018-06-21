package functions

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

// SingleProjectMode special setting for get sub folder
var SingleProjectMode = false

// HandleSingle handles a single project
func HandleSingle(list *Work) {

	argsWithoutProg := os.Args[1:]
	fullPath, _ := os.Getwd()

	fmt.Println("\nRunning in single project mode!")
	SingleProjectMode = true

	// Check if its the same dir, will not work if different drive. Maybe not needed
	if !strings.Contains(argsWithoutProg[0], fullPath) {
		fmt.Println("Bad path! Place zip or folder in the same location as the exe file")
		fmt.Println("\nProgram path:", argsWithoutProg[0])
		fmt.Println("file path: ", fullPath)
		SleepMs(time.Duration(Config.KeepOpenInSeconds * 1000))
		return
	}

	relativePath := argsWithoutProg[0][len(fullPath)+1:]

	// Checks if its a zip file and opens it
	if IsZip(relativePath) {
		newPath := fmt.Sprintf("./%s", relativePath[:len(relativePath)-4])
		Unzip(fmt.Sprintf("./%s", relativePath), newPath)
		relativePath = newPath
		if Config.DeleteUnzipedFolder {
			defer deleteUnzippedFolder(relativePath)
		}
	}

	list.Projects = append(list.Projects, Project{
		Done:       false,
		FolderName: relativePath,
		HTMLs:      []HTMLVerify{},
		CSSs:       []CSSVerify{},
	})

	// syncing
	var wg = sync.WaitGroup{}
	var wgUI = sync.WaitGroup{}

	wg.Add(1)
	wgUI.Add(1)
	go DoProject(list, 0, &wg, &wgUI)

	wgUI.Wait()           // Wait until all html files has been found and loaded
	go UpdateScreen(list) // the ui

	wg.Wait()
	list.Complete = true
}

func deleteUnzippedFolder(path string) {
	fmt.Println("Delete:", path)
	os.RemoveAll(path)
}
