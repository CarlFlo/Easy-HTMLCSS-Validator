package functions

import (
	"io/ioutil"
	"log"
	"os"
)

var helpTxt = `CONFIG
\ncores: 
\nfolderName: 
\nreqHTMLFolder: 
\nmakeHelpTxt: 
`

// MakeHelpFile creates a helper file
func MakeHelpFile() {

	if _, err := os.Stat("/help.txt"); os.IsNotExist(err) {

		err = ioutil.WriteFile("help.txt", []byte(helpTxt), 0644)
		if err != nil {
			log.Println(err.Error())
		}
	}
}
