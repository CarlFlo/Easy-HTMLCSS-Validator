package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"sync"
	"time"

	"./functions"
)

func main() {

	clear()
	functions.ReadConfig()
	if functions.Config.MakeHelpTxt {
		functions.MakeHelpFile()
	}
	runtime.GOMAXPROCS(functions.Config.Cores)

	checkPath()
	projectDirs := getProjectDirs()

	log.Println("Start")

	startTime := time.Now()
	iterate(projectDirs)
	fmt.Println(fmt.Sprintf("Duration: %v", time.Now().Sub(startTime)))

	fmt.Println("Done")
	time.Sleep(time.Second * 20)
	os.Exit(1)
}

func iterate(projectDirs []string) {

	var mutex = &sync.Mutex{}
	var wg = sync.WaitGroup{}

	for i := 0; i < len(projectDirs); i++ {
		wg.Add(1)
		go doProject(projectDirs[i], &wg, mutex)
	}

	wg.Wait()
}

// Går igenom ett project
func doProject(folderName string, wg *sync.WaitGroup, mutex *sync.Mutex) {

	defer wg.Done() // defer is done at the end

	htmlPaths := functions.WalkHTML(folderName)

	if len(htmlPaths) == 0 {
		log.Printf("%s has no html files!", folderName)
		return
	}

	for i := 0; i < len(htmlPaths); i++ {

		txt, err := ioutil.ReadFile(htmlPaths[i])
		if err != nil {
			log.Printf("%s - %v", folderName, err.Error())
			return
		}

		err = validateHTML(string(txt), htmlPaths[i], i)
		if err != nil {
			log.Printf("Could not validate %s: %v", folderName, err.Error())
		}
		log.Println("Validated:", htmlPaths[i])
	}
}

// Hämtar alla mappnamn i path
func getProjectDirs() []string {

	var dirs []string

	files, err := ioutil.ReadDir(fmt.Sprintf("./%s/", functions.Config.FolderName))
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(fmt.Sprintf("Totalt %d st mappar i %s:\n", len(files), functions.Config.FolderName))

	for _, f := range files {
		dirs = append(dirs, f.Name())
		fmt.Println(f.Name())
	}

	fmt.Println("------END------\n")
	return dirs
}

func validateHTML(html, path string, tmp int) error {

	resp, err := http.PostForm("https://validator.w3.org/check",
		url.Values{ // name
			"charset":  {"(detect automatically)"},
			"fragment": {html},               // HTML koden
			"doctype":  {"XHTML 1.0 Strict"}, // "HTML5" eller "XHTML 1.0 Strict"
			"group":    {"1"},                // Gruppera felen tillsammans
			"st":       {"1"},
			"outline":  {"1"}, // Får ut antalet h1-hx element längst ner
		})

	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	//functions.ParseHTML(resp.Body) // body istället

	ioutil.WriteFile(string(fmt.Sprintf("./verified/%v.html", tmp)), body, 0644) // debug

	return nil
}

// Clear the screen
func clear() {

	// Check os and use correct settings
	switch currentOS := runtime.GOOS; currentOS {
	case "linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		log.Println(fmt.Sprintf("Currently running on %s. No clear setting for this type", currentOS))
	}
}

func checkPath() {

	// Check if folder exists and creates it if not
	if _, err := os.Stat(functions.Config.FolderName); os.IsNotExist(err) {

		// Check os and use correct settings
		switch currentOS := runtime.GOOS; currentOS {
		case "linux":
			os.Mkdir(functions.Config.FolderName, 0777) // linux (Raspberry pi)
		case "windows":
			os.Mkdir(functions.Config.FolderName, 0007) // Windows
		default:
			log.Println(fmt.Sprintf("Currently running on %s. No permission setup for it. Using windows permission setting for folder as default", currentOS))
			os.Mkdir(functions.Config.FolderName, 0007)
		}

		log.Println(fmt.Sprintf("Folder %s has been created. Please insert project folders into it and restart the software", functions.Config.FolderName))
		time.Sleep(time.Second * 2)
		log.Println("Program will close in 8 seconds")
		time.Sleep(time.Second * 8)
		os.Exit(1)
	}
}
