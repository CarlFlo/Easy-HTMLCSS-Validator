package functions

import (
	"fmt"
	"os"
	"time"
)

// HandleSingle handles a single project
func HandleSingle() {

	fmt.Println("\nRunning in single project mode!")

	argsWithoutProg := os.Args[1:]
	fmt.Println("Path:\n", argsWithoutProg[0])

	time.Sleep(time.Second * 120)
}
