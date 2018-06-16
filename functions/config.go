package functions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

// Config håller i alla variabler som blir inlästa
var Config *configStruct

type configStruct struct {
	Cores               int           `json:"cores"` // Hur många cores som go rutines får använda
	FolderName          string        `json:"folderName"`
	DispConfigOnStart   bool          `json:"dispConfigOnStart"`
	MakeHelpTxt         bool          `json:"makeHelpTxt"`
	DeleteUnzipedFolder bool          `json:"deleteUnzipedFolder"`
	DrawUI              bool          `json:"drawUI"`
	UpdateUiMs          time.Duration `json:"updateUiMs"`
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

	err = json.Unmarshal(file, &Config)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	log.Println("Success!")
	if Config.DispConfigOnStart {
		fmt.Println(string(file))
	}

	return nil
}

// Fösköker skapa config filen
func createConfig() error {
	log.Println("Creating config...")

	jsonData := []byte(`{
		"cores": 2,
		"folderName": "putProjectFoldersInHere",
		"dispConfigOnStart": true,
		"makeHelpTxt": true,
		"deleteUnzipedFolder": true,
		"DrawUI": true,
		"UpdateUiMs": 50}`)

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
