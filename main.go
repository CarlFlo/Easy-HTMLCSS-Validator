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
		fmt.Println("Drag and drop is not supported yet!")
		functions.SleepMs(10000)
		os.Exit(1)

		functions.HandleSingle(list)
		return
	}

	functions.CheckPath()
	// Unzip all .zip folders
	functions.PopulateProjectArr(list)

	var wg = sync.WaitGroup{}
	startTime := time.Now()
	go functions.UpdateScreen(list) // the ui

	for i := 0; i < len(list.Projects); i++ {
		wg.Add(1)
		go functions.DoProject(list, i, &wg)
	}

	wg.Wait()

	list.Complete = true

	functions.SleepMs(functions.Config.UpdateUiMs) // So the update ui wont overwrite
	functions.Clear()
	fmt.Println(fmt.Sprintf("Validation took: %v\n", time.Now().Sub(startTime)))

	/*	Lists Folder name and the html files in the project */
	for _, project := range list.Projects {
		fmt.Println("\nFolder name:", project.FolderName)
		for val, htmlFile := range project.HTMLs {
			// Filepath
			fmt.Println(fmt.Sprintf("[%d] File: %s", val+1, htmlFile.Path))
			// Errors
			fmt.Println("Errors:", htmlFile.TotalErrors)
			for i := 0; i < len(htmlFile.Errors); i++ {
				fmt.Println(htmlFile.Errors[i])
			}
			// Warnings
			fmt.Println("\nWarnings:", htmlFile.TotalWarnings)
			for i := 0; i < len(htmlFile.Warnings); i++ {
				fmt.Println(htmlFile.Errors[i])
			}
			fmt.Println("")
		}
	}
	fmt.Println("\n\nDone!")

	functions.SleepMs(9000000)
}
