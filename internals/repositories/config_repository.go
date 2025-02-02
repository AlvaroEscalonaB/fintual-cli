package internals

import (
	"fintual-cli/internals/models"
	"log"

	"gopkg.in/ini.v1"
)

const filePath = "config.ini"

type Config struct {
	filePath string
}

func (c Config) GetConfig() (*models.ConfigINI, error) {
	iniData, err := ini.Load(c.filePath)

	if err != nil {
		return nil, err
	}

	var configFromConfig models.ConfigINI

	err = iniData.MapTo(&configFromConfig)

	if err != nil {
		return nil, err
	}

	return &configFromConfig, nil
}

func (c Config) SetConfig(field string, value string) error {
	config, err := ini.Load(c.filePath)

	if err != nil {
		return err
	}

	userSection := config.Section("user")

	if field == "email" {
		userSection.Key("email").SetValue(value)
	}

	if field == "password" {
		userSection.Key("password").SetValue(value)
	}

	err = config.SaveTo("config.ini")

	if err != nil {
		log.Fatalf("Error saving the file %v", err)
	}

	return nil
}

var ConfigRepository = Config{filePath: filePath}
