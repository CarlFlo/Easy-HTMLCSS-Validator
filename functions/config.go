package functions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// Config håller i alla variabler som blir inlästa
var Config *configStruct

type configStruct struct {
	Cores       int    `json:"cores"` // Hur många cores som go rutines får använda
	FolderName  string `json:"folderName"`
	MakeHelpTxt bool   `json:"makeHelpTxt"`
}

// ReadConfig försöker läsa configen
func ReadConfig() {

	err := loadConfig()
	if err != nil {
		err = createConfig()
		if err != nil {
			panic("Could not create config file")
		}
		err := loadConfig()
		if err != nil {
			panic("Could not load newly created config file")
		}
	}
}

// Försöker läsa in configen
func loadConfig() error {
	log.Println("Reading config...")

	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		log.Println(err.Error())
		return err
	}

	log.Println("Success!")
	fmt.Println(string(file))

	err = json.Unmarshal(file, &Config)

	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

// Fösköker skapa config filen
func createConfig() error {
	log.Println("Creating config...")

	jsonData := []byte(`{
		"cores": 1,
		"folderName": "putProjectFoldersInHere",
		"makeHelpTxt": true}`)

	configStruct := configStruct{}

	err := json.Unmarshal(jsonData, &configStruct)
	if err != nil {
		log.Println("Error! Could not create file...")
		return err
	}

	jsonDataJSON, _ := json.MarshalIndent(configStruct, "", "   ")
	err = ioutil.WriteFile("config.json", jsonDataJSON, 0644)
	return nil
}
