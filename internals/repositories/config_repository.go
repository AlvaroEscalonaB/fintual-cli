package internals

import (
	"fintual-cli/internals/models"

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

var ConfigRepository = Config{filePath: filePath}
