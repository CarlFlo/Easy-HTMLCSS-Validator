package functions

import (
	"fmt"
	"runtime"
)

// UpdateScreen updates the screen with the validation progress
func UpdateScreen(list *Work) {
	if !Config.DrawUI {
		return
	}

	// Calculate the total amount of html docs to process
	totalHTMLFiles := 0
	totalDone := 0
	for i := 0; i < len(list.Projects); i++ {
		totalHTMLFiles += len(list.Projects[i].HTMLs)
	}

	// So the UI will run once more when done to make sure nothing is at <100%
	another := false

	for !list.Complete && !another {
		if list.Complete {
			another = true
		}

		Clear()
		totalDone = 0
		for i := 0; i < len(list.Projects); i++ {
			done := 0
			if !list.Projects[i].Done { // Project is not done
				total := len(list.Projects[i].HTMLs)
				done = 0
				for j := 0; j < total; j++ {
					if list.Projects[i].HTMLs[j].Verified {
						done++
					}
				}
				fmt.Println(fmt.Sprintf("%s - [%d/%d] %.2f%%", list.Projects[i].FolderName, done, total, (float64(done)/float64(total))*100))
			} else { // Done

				done += len(list.Projects[i].HTMLs)
				fmt.Println(fmt.Sprintf("%s - 100%%", list.Projects[i].FolderName))
			}
			totalDone += done
		}
		fmt.Println(fmt.Sprintf("\nProcessing %d/%d [%.2f%%] html files\nActive goroutines: %d", totalDone, totalHTMLFiles, (float64(totalDone)/float64(totalHTMLFiles))*100, runtime.NumGoroutine()))

		SleepMs(Config.UpdateUIMs)
	}
	showResult(list)
}

// ShowResult shows the result of the validation
func showResult(list *Work) {

	fmt.Println(fmt.Sprintf("Validation took: %v\n", list.Timing.EndTime.Sub(list.Timing.StartTime)))

	/*	Lists Folder name and the html files in the project */
	for _, project := range list.Projects {
		fmt.Println("\nFolder name:", project.FolderName)
		for val, htmlFile := range project.HTMLs {
			// Filepath
			fmt.Println(fmt.Sprintf("[%d] File: %s", val+1, htmlFile.Path))
			// Errors
			fmt.Println("Errors:", htmlFile.TotalErrors)
			for k, v := range htmlFile.Errors {
				fmt.Println(fmt.Sprintf("%d: %s", k, v))
			}
			// Warnings
			fmt.Println("\nWarnings:", htmlFile.TotalWarnings)

			for k, v := range htmlFile.Warnings {
				fmt.Println(fmt.Sprintf("%d: %s", k, v))
			}

			fmt.Println("")
		}
	}
	fmt.Println("\n\nDone!")
}
