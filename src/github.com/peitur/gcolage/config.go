package gcolage

import (
	"encoding/json"
	"log"
)

type FileCollectorConfig struct {
	ConfigPath string `json:"config_path"`
	TargetPath string `json:"store"`
	Checksum   string `json:"checksum"`
}

type PipCollectorConfig struct {
	ConfigPath string
	TargetPath string
	Checksum   string
}

func ReadConfigFile(filename string) FileCollectorConfig {
	buffer, err := ReadFileRaw(filename)
	if err != nil {
		log.Fatal(err)
	}

	var conf FileCollectorConfig
	err = json.Unmarshal([]byte(buffer), &conf)
	if err != nil {
		log.Fatal(err)
	}

	return conf
}
