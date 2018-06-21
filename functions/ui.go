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

		// Kill this goroutines
		if list.GracefulStop {
			ExitGoroutine()
		}

		Clear()
		totalDone = 0
		for i := 0; i < len(list.Projects); i++ {
			done := 0
			if !list.Projects[i].Done { // Project is not done
				total := len(list.Projects[i].HTMLs)
				done = 0
				for j := 0; j < total; j++ {
					if list.Projects[i].HTMLs[j].AllVerified {
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

	if Config.DisplayResult {
		showResult(list)
	}
	// End message
	fmt.Println("\n\nVerify done!\nWindow will close in", Config.KeepOpenInSeconds, "seconds")
}

// ShowResult shows the result of the validation
func showResult(list *Work) {

	fmt.Println(fmt.Sprintf("Validation took: %v\n", list.Timing.EndTime.Sub(list.Timing.StartTime)))

	/*	Lists Folder name and the html files in the project */
	for _, project := range list.Projects {
		fmt.Println("\nFolder name:", project.FolderName)

		// Show css result
		fmt.Println("CSS:\n")

		for val, htmlFile := range project.HTMLs {
			// Filepath
			fmt.Println(fmt.Sprintf("File [%d]: %s", val+1, htmlFile.Path))

			// HTML5
			fmt.Println("\n[HTML5]:")
			if htmlFile.HTML5Verify.ErrorValidating != nil {
				fmt.Println("Error validating:", htmlFile.HTML5Verify.ErrorValidating.Error())
			} else if htmlFile.HTML5Verify.HasWarningsOrErrors {
				fmt.Println("NOT OK") // Currently never run. Its ok or error. Fix later
			} else {
				fmt.Println("OK")
			}

			// XHTML 1.0 Strict
			fmt.Println(fmt.Sprintf("\n[XHTML 1.0 Strict]: %s", htmlFile.StrictVerify.Result))
			// Strict Warnings
			fmt.Println(" [Warnings]:")
			for k, v := range htmlFile.StrictVerify.Warnings {
				fmt.Println(fmt.Sprintf("  [%d]: %s", k, v))
			}
			// Strict info
			fmt.Println("\n [Info]:")
			for k, v := range htmlFile.StrictVerify.Infos {
				fmt.Println(fmt.Sprintf("  [%d]: %s", k, v))
			}
			fmt.Println("\n [Errors]:")
			for k, errStrArr := range htmlFile.StrictVerify.Errors { // todo look at this
				fmt.Println(fmt.Sprintf("  Group %d: %s\n", k, errStrArr.ErrorType))
				for nr, val := range errStrArr.ErrorStrings {
					fmt.Println(fmt.Sprintf("   [%d]:\n    Line:%s\n    Error:%s\n    Text:%s\n", nr, val.Line, val.Error, val.TextFromHTML))
				}
			}
			fmt.Println("")
		}
	}
}
