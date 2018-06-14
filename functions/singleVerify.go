package functions

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// DragAndDropMode special setting for get sub folder
var DragAndDropMode = false

// HandleSingle handles a single project
func HandleSingle() {

	var mutex = &sync.Mutex{}
	var wg = sync.WaitGroup{}

	argsWithoutProg := os.Args[1:]
	pwd, _ := os.Getwd()

	relativePath := argsWithoutProg[0][len(pwd)+1:]

	fmt.Println("\nRunning in single project mode!")
	DragAndDropMode = true

	DoProject(relativePath, &wg, mutex)

	time.Sleep(time.Second * 120)
}
