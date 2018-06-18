package functions

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// WalkHTML returns all paths to html doc in the path given in Project
func WalkHTML(project *Project) {
	skip := true

	// startPath is where the Walk will begin to find all html files
	startPath := fmt.Sprintf("./%s/%s", Config.FolderName, project.FolderName)

	// For when the user drags a folder or zip onto the exe file for single verify
	if SingleProjectMode {
		startPath = fmt.Sprintf("./%s", project.FolderName)
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
				// Lägg till hela sökvägen till html filen i HTMLs arrayen
				project.HTMLs = append(project.HTMLs, HTMLVerify{
					Path:        path,
					AllVerified: false,
					StrictVerify: StrictVerify{
						Verified: false,
					},
					HTML5Verify: HTML5Verify{
						Verified: false,
					},
				})
			}

		}
		return nil
	})

	if err != nil {
		log.Println(err.Error())
	}
}
