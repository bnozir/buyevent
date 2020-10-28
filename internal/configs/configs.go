package configs

import (
	"encoding/json"
	"io/ioutil"
)

// ApplicationConfigs application configurations
type ApplicationConfigs struct {
	CustomersDataFilePath         string `json:"customersDataFilePath"`
	NotificationsTypeDataFilePath string `json:"notificationsTypeDataFilePath"`
	NotificationsDataFilePath     string `json:"notificationsDataFilePath"`
	InfoLogFilePath               string `json:"infoLogFilePath"`
	ErrorLogFilePath              string `json:"errorLogFilePath"`
}

// Configs ApplicationConfigs instant
var Configs ApplicationConfigs

// Reads and set configurations from config-file
func init() {
	configFile, err := ioutil.ReadFile("./config/configs.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(configFile, &Configs)
	if err != nil {
		panic(err)
	}
}
