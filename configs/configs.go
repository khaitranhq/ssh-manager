package configs

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Connection struct {
	Host     string `json:"host"`
	KeyPath  string `json:"key_path"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type Group struct {
	GroupName   string       `json:"group_name"`
	Connections []Connection `json:"connections"`
}

func LoadConfigs() []Group {
	if _, err := os.Stat(CONFIG_FILE_PATH); err != nil {
		emptyGroupArr := []Group{}
		file, _ := json.MarshalIndent(emptyGroupArr, "", " ")
		ioutil.WriteFile(CONFIG_FILE_PATH, file, 0644)

		return emptyGroupArr
	}

	return []Group{}
}
