package functions

import (
	"fmt"
	"os"
	"os/exec"
)

// OpenResultPage will open a html page in the browser to show the result
func OpenResultPage() {

	cmd := "cmd"
	args := []string{"/c", "start", "./web/showResult.html"}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Println("Failed to open local webpage:")
		fmt.Fprintln(os.Stderr, err)
	}

	fmt.Println("Successfully started webpage")
}
