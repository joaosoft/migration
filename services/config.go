package services

import (
	"fmt"

	"github.com/joaosoft/manager"
)

// AppConfig ...
type AppConfig struct {
	Migration MigrationConfig `json:"migration"`
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
func NewConfig() (*AppConfig, manager.IConfig, error) {
	appConfig := &AppConfig{}
	simpleConfig, err := manager.NewSimpleConfig(fmt.Sprintf("/config/app.%s.json", GetEnv()), appConfig)

	if appConfig.Migration.Host == "" {
		appConfig.Migration.Host = DefaultURL
	}

	return appConfig, simpleConfig, err
}
