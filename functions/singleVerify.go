package functions

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// SingleProjectMode special setting for get sub folder
var SingleProjectMode = false

// HandleSingle handles a single project
func HandleSingle() {

	var mutex = &sync.Mutex{}
	var wg = sync.WaitGroup{}

	argsWithoutProg := os.Args[1:]
	pwd, _ := os.Getwd()

	relativePath := argsWithoutProg[0][len(pwd)+1:]

	// Checks if its a zip file and opens it
	if IsZip(relativePath) {
		newPath := fmt.Sprintf("./%s", relativePath[:len(relativePath)-4])
		Unzip(fmt.Sprintf("./%s", relativePath), newPath)
		relativePath = newPath
		defer deleteUnzippedFolder(relativePath)
	}

	fmt.Println("\nRunning in single project mode!")
	SingleProjectMode = true

	DoProject(relativePath, &wg, mutex)
}

func deleteUnzippedFolder(path string) {
	fmt.Println("Delete:", path)
	os.RemoveAll(path)
	time.Sleep(time.Second * 120)
}
