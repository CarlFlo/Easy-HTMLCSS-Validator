package functions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"runtime"
	"time"
)

// Config håller i alla variabler som blir inlästa
var Config *configStruct

type configStruct struct {
	Cores                      int           `json:"cores"` // Hur många cores som go rutines får använda
	FolderName                 string        `json:"folderName"`
	OutputFilename             string        `json:"outputFilename"`
	DispConfigOnStart          bool          `json:"dispConfigOnStart"`
	DisplayResult              bool          `json:"displayResult"`
	MakeHelpTxt                bool          `json:"makeHelpTxt"`
	DeleteUnzipedFolder        bool          `json:"deleteUnzipedFolder"`
	ValidateWithHTML5_verySlow bool          `json:validateWithHTML5_verySlow`
	DrawUI                     bool          `json:"drawUI"`
	UpdateUIMs                 time.Duration `json:"updateUIMs"`
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

// Tries to read config file
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
	checkConfigValues()
	if Config.DispConfigOnStart {
		fmt.Println(string(file)) // This value will not update if checkConfigValues changed anything
		SleepMs(3500)
	}

	return nil
}

// Tries to create config file
func createConfig() error {
	log.Println("Creating config...")

	// Default settings
	configStruct := configStruct{
		Cores:                      2,
		FolderName:                 "Projects",
		OutputFilename:             "output.json",
		DispConfigOnStart:          true,
		DisplayResult:              true,
		MakeHelpTxt:                true,
		DeleteUnzipedFolder:        true,
		ValidateWithHTML5_verySlow: false,
		DrawUI:     true,
		UpdateUIMs: 50,
	}

	jsonDataJSON, _ := json.MarshalIndent(configStruct, "", "   ")
	err := ioutil.WriteFile("config.json", jsonDataJSON, 0644)
	if err != nil {
		panic(err.Error())
	}
	return nil
}

// Checks so values in config file are ok
func checkConfigValues() {

	// The amount of cores that goroutines can use should not be over the number of cores available
	if Config.Cores > runtime.NumCPU() {
		Config.Cores = runtime.NumCPU()
		fmt.Println("\"cores\" in config was invalid and thus changed to:", Config.Cores)
	}
	// Minimum update rate is 50ms and max is 1000ms
	if Config.UpdateUIMs < 50 {
		Config.UpdateUIMs = 50
		fmt.Println("\"updateUIMs\" in config was invalid and thus changed to:", Config.UpdateUIMs)
	} else if Config.UpdateUIMs > 1000 {
		Config.UpdateUIMs = 1000
		fmt.Println("\"updateUIMs\" in config was invalid and thus changed to:", Config.UpdateUIMs)
	}

	// Check output name
	if len(Config.OutputFilename) == 0 {
		Config.OutputFilename = "output.json"
	}
}
