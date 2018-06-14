package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"
	"time"

	"./functions"
)

func init() {

	functions.Clear()
	functions.ReadConfig()
	if functions.Config.MakeHelpTxt {
		functions.MakeHelpFile()
	}
	runtime.GOMAXPROCS(functions.Config.Cores)
}

func main() {

	// Användaren drog in en mapp eller zip fil på exe filen
	if len(os.Args) > 1 {
		functions.HandleSingle()
		return
	}

	functions.CheckPath()
	projectDirs := functions.GetProjectDirs()

	log.Println("Start")

	startTime := time.Now()
	iterate(projectDirs)
	fmt.Println(fmt.Sprintf("Duration: %v", time.Now().Sub(startTime)))

	fmt.Println("Done")
	time.Sleep(time.Second * 20)
}

func iterate(projectDirs []string) {

	var mutex = &sync.Mutex{}
	var wg = sync.WaitGroup{}

	for i := 0; i < len(projectDirs); i++ {
		wg.Add(1)
		go functions.DoProject(projectDirs[i], &wg, mutex)
	}

	wg.Wait()
}
