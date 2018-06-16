package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"

	"./functions"
)

func init() {

	//functions.SleepMs(4000)
	//functions.SetCmdSize("150", "25")
	functions.Clear()
	functions.ReadConfig()
	if functions.Config.MakeHelpTxt {
		functions.MakeHelpFile()
	}
	runtime.GOMAXPROCS(functions.Config.Cores)
}

func main() {

	list := &functions.Work{
		Mutex:    sync.Mutex{},
		Projects: []functions.Project{},
		Complete: false,
	}

	/* Fix singleVerify */

	// For when user drags in a folder or .zip file onto the exe file for single verify
	if len(os.Args) > 1 {

		functions.HandleSingle(list)
		return
	}

	functions.CheckPath()
	// Unzip all .zip folders
	functions.PopulateProjectArr(list)

	var wg = sync.WaitGroup{}
	var wgUI = sync.WaitGroup{}
	startTime := time.Now()

	for i := 0; i < len(list.Projects); i++ {
		wg.Add(1)
		wgUI.Add(1)
		go functions.DoProject(list, i, &wg, &wgUI)
	}
	wgUI.Wait()                     // Wait until all html files has been found and loaded
	go functions.UpdateScreen(list) // the ui

	wg.Wait()

	list.Complete = true

	fmt.Println("(It's done even if it says that one isn't 100%)")
	fmt.Println(fmt.Sprintf("Validation took: %v\n", time.Now().Sub(startTime)))

	functions.ShowResult(list)

	functions.SleepMs(9000000)
}
