package functions

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"time"
)

// CheckPath checks if the path exsists
func CheckPath() {

	// Check if folder exists and creates it if not
	if _, err := os.Stat(Config.FolderName); os.IsNotExist(err) {

		// Check os and use correct settings
		switch currentOS := runtime.GOOS; currentOS {
		case "linux":
			os.Mkdir(Config.FolderName, 0777) // linux (Raspberry pi)
		case "windows":
			os.Mkdir(Config.FolderName, 0007) // Windows
		default:
			log.Println(fmt.Sprintf("Currently running on %s. No permission setup for it. Using windows permission setting for folder as default", currentOS))
			os.Mkdir(Config.FolderName, 0007)
		}

		log.Println(fmt.Sprintf("Folder %s has been created. Please insert project folders into it and restart the software", Config.FolderName))
		time.Sleep(time.Second * 2)
		log.Println("Program will close in 8 seconds")
		time.Sleep(time.Second * 8)
		os.Exit(1)
	}
}

// PopulateProjectArr populates the list with html for that folder
func PopulateProjectArr(list *Work) {

	files, err := ioutil.ReadDir(fmt.Sprintf("./%s/", Config.FolderName))
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(fmt.Sprintf("Totalt %d st mappar i %s:\n", len(files), Config.FolderName))

	for _, f := range files {

		if !f.IsDir() { // skips everything that isnt a folder
			continue
		}

		// Created a new project object for each folder
		list.Projects = append(list.Projects, Project{
			Done:       false,
			FolderName: f.Name(),
			HTMLs:      []HTMLVerify{},
			CSSs:       []CSSVerify{},
		})
	}
}
