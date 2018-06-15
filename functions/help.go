package functions

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
)

// MakeHelpFile creates a helper file
func MakeHelpFile() {

	if _, err := os.Stat("./help.txt"); os.IsNotExist(err) {

		err = ioutil.WriteFile("help.txt", createString(), 0644)
		if err != nil {
			log.Println(err.Error())
		}
	}
}

// Ugly solution
func createString() []byte {

	var buffer bytes.Buffer

	buffer.WriteString("CONFIG")
	buffer.WriteString("\ncores: How many cores the program can use (def=1)")
	buffer.WriteString("\nfolderName: The folder name the program will use for bulk validation")
	buffer.WriteString("\nmakeHelpTxt: If this txt file is to be created (def=true)")
	buffer.WriteString("\n\n You can both drag and drop a zip file or folder with html files inside them onto the exe file to validate just those files")

	return buffer.Bytes()
}
