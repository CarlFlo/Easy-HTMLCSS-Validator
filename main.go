package main

import (
	"math/rand"
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
	rand.Seed(time.Now().UTC().UnixNano())
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
		Timing: functions.Timing{
			StartTime: time.Now(),
		},
	}

	/* Fix singleVerify */

	// For when user drags in a folder or .zip file onto the exe file for single verify
	if len(os.Args) > 1 {
		// Fix
		functions.HandleSingle(list)
		return
	}

	functions.CheckPath()
	// Unzip all .zip folders
	functions.PopulateProjectArr(list)

	var wg = sync.WaitGroup{}
	var wgUI = sync.WaitGroup{}

	for i := 0; i < len(list.Projects); i++ {
		wg.Add(1)
		wgUI.Add(1)
		go functions.DoProject(list, i, &wg, &wgUI)
	}
	wgUI.Wait()                     // Wait until all html files has been found and loaded
	go functions.UpdateScreen(list) // the ui

	wg.Wait()
	list.Timing.EndTime = time.Now()
	list.Complete = true

	functions.SleepMs(9000000) // So the windows wont close. Change this later
}
