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
	pwd, _ := os.Getwd()

	// Check if its the same dir, will not work if different drive. Maybe not needed
	if !strings.Contains(argsWithoutProg[0], pwd) {
		fmt.Println("Bad path! Place zip or folder in the same location as the exe file")
		SleepMs(5000)
		return
	}

	relativePath := argsWithoutProg[0][len(pwd)+1:]

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
	})

	fmt.Println("\nRunning in single project mode!")
	SingleProjectMode = true
	go UpdateScreen(list)                                    // the ui
	DoProject(list, 0, &sync.WaitGroup{}, &sync.WaitGroup{}) // These WaitGroup are not needed/used when doing 1 project but func needs them
}

func deleteUnzippedFolder(path string) {
	fmt.Println("Delete:", path)
	os.RemoveAll(path)
	time.Sleep(time.Second * 120)
}
