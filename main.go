package main

import (
	"flag"
	"fmt"
	"github.com/joaosoft/migration/services"
	"os"
)

func main() {
	var cmdMigrate string
	var cmdNumber int
	var cmdMode string
	var executorModes []services.ExecutorMode

	flag.StringVar(&cmdMigrate, string(services.CmdMigrate), string(services.OptionUp), "Runs the specified command. Valid options are: `"+string(services.OptionUp)+"`, `"+string(services.OptionDown)+"`.")
	flag.IntVar(&cmdNumber, string(services.CmdNumber), 0, "Runs the specified command.")
	flag.StringVar(&cmdMode, string(services.CmdMode), string(services.ExecutorModeAll), "Runs the specific mode. Valid options are: `"+string(services.ExecutorModeAll)+"`, `"+string(services.ExecutorModeDatabase)+"`, `"+string(services.ExecutorModeRabbitMq)+"`.")
	flag.Parse()

	m, err := services.NewCmdService()
	if err != nil {
		panic(err)
		os.Exit(1)
	}

	if err := m.Start(); err != nil {
		panic(err)
		os.Exit(1)
	}

	switch mode := services.ExecutorMode(cmdMode); {
	case mode == services.ExecutorModeAll:
		executorModes = append(executorModes,
			[]services.ExecutorMode{services.ExecutorModeDatabase, services.ExecutorModeRabbitMq}...,
		)
	default:
		executorModes = append(executorModes, mode)
	}

	m.AddTag("custom", CustomHandler)

	for _, mode := range executorModes {
		if _, err := m.Execute(services.MigrationOption(cmdMigrate), cmdNumber, mode); err != nil {
			panic(err)
			os.Exit(1)
		}
	}

	os.Exit(0)
}

func CustomHandler(option services.MigrationOption, conn services.Executor, data string) error {
	fmt.Printf("\nexecuting with option '%s' and data '%s", option, data)
	return nil
}
