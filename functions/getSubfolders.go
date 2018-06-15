package functions

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// WalkHTML retunerar alla sökvägar till html dokument i folderName sökvägen
func WalkHTML(folderName string) []string {

	htmlPaths := []string{}

	skip := true

	startPath := fmt.Sprintf("./%s/%s", Config.FolderName, folderName)

	// For when the user drags a folder or zip onto the exe file for single verify
	if SingleProjectMode {
		startPath = fmt.Sprintf("./%s", folderName)
	}

	err := filepath.Walk(startPath, func(path string, info os.FileInfo, err error) error {

		if skip {
			skip = false
			return nil
		}

		if err != nil {
			log.Println(err.Error())
			return err
		}
		if !info.IsDir() {
			if strings.EqualFold(info.Name()[len(info.Name())-4:], "html") {
				// Lägg till hela sökvägen från folderName i htmlPaths
				htmlPaths = append(htmlPaths, path)
			}

		}
		//fmt.Printf("Visited file: %s\n", path)
		return nil

	})

	if err != nil {
		log.Println(err.Error())
	}

	return htmlPaths

	/* Test
	fmt.Println(len(htmlPaths))

	for i := 0; i < len(htmlPaths); i++ {
		fmt.Println(htmlPaths[i])

		txt, err := ioutil.ReadFile(htmlPaths[i])
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		fmt.Println(string(txt))

	}
	*/
}
