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
	buffer.WriteString("\nWarning! Config file needs to be present in the folder if user drags folder/zip on top of exe file from antoher directory or else default settings will be used\n")
	buffer.WriteString("\ncores: How many cores the program can use (def=2 max: number of pc cores)")
	buffer.WriteString("\nfolderName: The folder name the program will use for bulk validation")
	buffer.WriteString("\noutputFilename: ")
	buffer.WriteString("\ndispConfigOnStart: will display config on start with a small delay for reading it (3.5 seconds).")
	buffer.WriteString("\nmakeHelpTxt: If this txt file is to be created (def=true)")
	buffer.WriteString("\ndeleteUnzipedFolder: If true then .zip folders extracted data is deleted after use")
	buffer.WriteString("\nvalidateWithHTML5_verySlow: Requires sava installed and in Path and vnu.jar in the same folder as exe")
	buffer.WriteString("\nDrawUI: True if you want to show the progress in the cmd")
	buffer.WriteString("\nUpdateUiMs: (min: 50 max: 1000) How long the UI will wait in milliseconds before updating the progress on screen")
	buffer.WriteString("\n\nYou can both drag and drop a zip file or folder with html files inside them onto the exe file to validate just those files")

	return buffer.Bytes()
}
