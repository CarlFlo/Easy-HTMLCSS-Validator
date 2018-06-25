package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/signal"
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

/*
	folder dir
	help txt dir
*/

func main() {

	list := &functions.Work{
		Projects: []functions.Project{},
		Complete: false,
		Timing: functions.Timing{
			StartTime: time.Now(),
		},
		GracefulStop: false,
	}

	// Graceful stop - catches (ctrl + c)
	var gracefulStop = make(chan os.Signal, 1)
	signal.Notify(gracefulStop, os.Interrupt, os.Kill)
	go func() {
		sig := <-gracefulStop

		// Program is done. No need to preforme the code under
		if list.Complete {
			os.Exit(1)
			return
		}

		// set list to abort mode
		list.GracefulStop = true
		if !functions.Config.GracefulStop { //	If false then just stop the program on ctrl+c
			os.Exit(1)
		}
		functions.Clear()
		fmt.Println(fmt.Sprintf("Caught signal: %+v\nWait for 3 second to finish processing", sig))
		functions.SleepMs(3000)
		// Output result to file
		jsonDataJSON, _ := json.MarshalIndent(list, "", "   ")
		ioutil.WriteFile("interrupted-output.js", jsonDataJSON, 0644)
		os.Exit(1)
	}()

	functions.DeleteOldOutput()

	// Args Test
	//os.Args = append(os.Args, "D:\\Dropbox\\Kod\\Go\\Projects\\W3 validator - golang\\bin\\projectZip.zip")

	// For when user drags in a folder or .zip file onto the exe file for single verify
	if len(os.Args) > 1 {
		functions.HandleSingle(list)
		functions.SaveResult(list)
		// So the windows wont close.
		functions.SleepMs(time.Duration(functions.Config.KeepOpenInSeconds * 1000))
		return
	}

	functions.CheckPath()
	// todo: Unzip all .zip folders
	functions.PopulateProjectArr(list)

	if len(list.Projects) == 0 {
		fmt.Println("Nothing to verify")
		functions.SleepMs(4000)
		os.Exit(404)
	}

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

	// Do not preforme code if program was aborted mid run
	if !list.GracefulStop {

		list.Complete = true
		functions.SaveResult(list)
	}

	functions.SleepMs(time.Duration(functions.Config.KeepOpenInSeconds * 1000)) // So the windows wont close.
}
