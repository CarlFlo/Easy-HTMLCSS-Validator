package functions

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

// Clear will clear the screen
func Clear() {

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
