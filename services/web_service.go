package services

import (
	"fmt"
	"sync"

	"github.com/joaosoft/logger"
	"github.com/joaosoft/manager"
)

type WebService struct {
	config        *MigrationConfig
	isLogExternal bool
	pm            *manager.Manager
	mux           sync.Mutex
	logger        logger.ILogger
}

// NewWebService ...
func NewWebService(options ...WebServiceOption) (*WebService, error) {
	service := &WebService{
		pm:     manager.NewManager(manager.WithRunInBackground(false)),
		logger: logger.NewLogDefault("migration", logger.InfoLevel),
		config: &MigrationConfig{},
	}

	if service.isLogExternal {
		service.pm.Reconfigure(manager.WithLogger(service.logger))
	}

	// load configuration File
	appConfig := &AppConfig{}
	if simpleConfig, err := manager.NewSimpleConfig(fmt.Sprintf("/config/app.%s.json", GetEnv()), appConfig); err != nil {
		service.logger.Error(err.Error())
	} else if appConfig.Migration != nil {
		service.pm.AddConfig("config_app", simpleConfig)
		level, _ := logger.ParseLevel(appConfig.Migration.Log.Level)
		service.logger.Debugf("setting log level to %s", level)
		service.logger.Reconfigure(logger.WithLevel(level))
		service.config = appConfig.Migration
	}

	service.Reconfigure(options...)

	if service.config.Host == "" {
		service.config.Host = DefaultURL
	}

	simpleDB := manager.NewSimpleDB(&appConfig.Migration.Db)
	if err := service.pm.AddDB("db_postgres", simpleDB); err != nil {
		service.logger.Error(err.Error())
		return nil, err
	}

	web := manager.NewSimpleWebEcho(service.config.Host)
	controller := NewController(service.logger, NewInteractor(service.logger, NewStoragePostgres(service.logger, simpleDB)))
	controller.RegisterRoutes(web)

	service.pm.AddWeb("api_web", web)

	return service, nil
}

// Start ...
func (m *WebService) Start() error {
	return m.pm.Start()
}

// Stop ...
func (m *WebService) Stop() error {
	return m.pm.Stop()
}
