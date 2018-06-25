package functions

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"
)

// Get full dir of exe
func getExeDir() string {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		log.Fatal(err)
	}

	return dir
}

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

// ExitGoroutine will stop the goroutine that calls this function
func ExitGoroutine() {
	runtime.Goexit()
}

// SleepMs will sleep for n milliseconds
func SleepMs(n time.Duration) {
	time.Sleep(time.Millisecond * n)
}

// SetCmdSize sets the cmd size //
func SetCmdSize(cols, lines string) {

	cmd := exec.Command("mode", "con:", fmt.Sprintf("cols=%s", cols), fmt.Sprintf("lines=%s", lines))
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// RandomString returns a random string of n length
func RandomString(n int) string {

	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
