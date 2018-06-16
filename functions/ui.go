package functions

import (
	"fmt"
)

// UpdateScreen updates the screen with the validation progress
func UpdateScreen(list *Work) {

	if !Config.DrawUI {
		return
	}

	for !list.Complete {
		Clear()
		fmt.Println("Progress...")
		for i := 0; i < len(list.Projects); i++ {

			if !list.Projects[i].Done { // not done

				total := len(list.Projects[i].HTMLs)
				done := 0

				for j := 0; j < total; j++ {
					if list.Projects[i].HTMLs[j].Verified {
						done++
					}
				}

				fmt.Println(fmt.Sprintf("%s - %.2f%%", list.Projects[i].FolderName, (float64(done)/float64(total))*100))

			} else { // Done
				fmt.Println(fmt.Sprintf("%s - 100%%", list.Projects[i].FolderName))
			}
		}

		SleepMs(Config.UpdateUiMs)
	}
}
