package src

import (
	md "exporter/models"
	"os"

	"gopkg.in/yaml.v2"
)

func ReadConf() md.ConfigExprt {
	f, err := os.Open("config.yml")
	if err != nil {
		logger("READ_CONF_FILE_ERR", err)
	}
	defer f.Close()
	var cfg md.ConfigExprt
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		logger("PARSE_CONF_FILE_ERR", err)
	}
	return cfg
}
