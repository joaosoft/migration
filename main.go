package main

import (
	"database/sql"
	"flag"
	"fmt"
	"migration/services"
	"os"
)

func main() {
	var cmdMigrate string
	var cmdNumber int

	flag.StringVar(&cmdMigrate, string(services.CmdMigrate), string(services.OptionUp), "Runs the specified command. Valid options are: `"+string(services.OptionUp)+"`, `"+string(services.OptionDown)+"`.")
	flag.IntVar(&cmdNumber, string(services.CmdNumber), 0, "Runs the specified command.")
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

	m.AddTag("custom", CustomHandler)
	if executed, err := m.Execute(services.MigrationOption(cmdMigrate), cmdNumber); err != nil {
		panic(err)
		os.Exit(1)
	} else {
		fmt.Printf("EXECUTED: %d", executed)
	}
	os.Exit(0)
}

func CustomHandler(option services.MigrationOption, tx *sql.Tx, data string) error {
	fmt.Printf("\nexecuting with option '%s' and data '%s", option, data)
	return nil
}
