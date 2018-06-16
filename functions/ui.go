package functions

import (
	"fmt"
)

// UpdateScreen updates the screen with the validation progress
func UpdateScreen(list *Work) {
	if !Config.DrawUI {
		return
	}

	// Calculate total html docs
	totalHTMLFiles := 0

	for i := 0; i < len(list.Projects); i++ {
		for j := 0; j < len(list.Projects[i].HTMLs); j++ {
			totalHTMLFiles++
		}
	}

	for !list.Complete {
		Clear()

		fmt.Println(fmt.Sprintf("Processing %d files", totalHTMLFiles))
		for i := 0; i < len(list.Projects); i++ {

			if !list.Projects[i].Done { // Project is not done
				total := len(list.Projects[i].HTMLs)
				done := 0
				for j := 0; j < total; j++ {
					if list.Projects[i].HTMLs[j].Verified {
						done++
					}
				}
				fmt.Println(fmt.Sprintf("%s - [%d/%d] %.2f%%", list.Projects[i].FolderName, done, total, (float64(done)/float64(total))*100))
			} else { // Done
				fmt.Println(fmt.Sprintf("%s - 100%%", list.Projects[i].FolderName))
			}
		}
		SleepMs(Config.UpdateUiMs)
	}

}

// ShowResult shows the result of the validation
func ShowResult(list *Work) {

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
}
