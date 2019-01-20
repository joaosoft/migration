package services

import (
	"fmt"

	"github.com/joaosoft/logger"
	"github.com/joaosoft/manager"
)

// AppConfig ...
type AppConfig struct {
	Migration *MigrationConfig `json:"migration"`
}

// MigrationConfig ...
type MigrationConfig struct {
	Host string           `json:"host"`
	Path string           `json:"path"`
	Db   manager.DBConfig `json:"db"`
	Log  struct {
		Level string `json:"level"`
	} `json:"log"`
}

// NewConfig ...
func NewConfig() (*MigrationConfig, error) {
	appConfig := &AppConfig{}
	if _, err := manager.NewSimpleConfig(fmt.Sprintf("/config/app.%s.json", GetEnv()), appConfig); err != nil {
		logger.Error(err.Error())

		return &MigrationConfig{}, err
	}

	return appConfig.Migration, nil
}
