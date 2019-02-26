package services

import (
	"fmt"

	"github.com/joaosoft/manager"
)

// AppConfig ...
type AppConfig struct {
	Migration *MigrationConfig `json:"migration"`
}

// MigrationConfig ...
type MigrationConfig struct {
	Host string `json:"host"`
	Path struct {
		Database string `json:"database"`
		Rabbitmq string `json:"rabbitmq"`
	} `json:"path"`
	Db       manager.DBConfig `json:"db"`
	RabbitMq *struct {
		Host string `json:"host"`
	} `json:"rabbitmq"`
	Log struct {
		Level string `json:"level"`
	} `json:"log"`
}

// NewConfig ...
func NewConfig() (*AppConfig, manager.IConfig, error) {
	appConfig := &AppConfig{}
	simpleConfig, err := manager.NewSimpleConfig(fmt.Sprintf("/config/app.%s.json", GetEnv()), appConfig)

	return appConfig, simpleConfig, err
}
