# migration
[![Build Status](https://travis-ci.org/joaosoft/migration.svg?branch=master)](https://travis-ci.org/joaosoft/migration) | [![codecov](https://codecov.io/gh/joaosoft/migration/branch/master/graph/badge.svg)](https://codecov.io/gh/joaosoft/migration) | [![Go Report Card](https://goreportcard.com/badge/github.com/joaosoft/migration)](https://goreportcard.com/report/github.com/joaosoft/migration) | [![GoDoc](https://godoc.org/github.com/joaosoft/migration?status.svg)](https://godoc.org/github.com/joaosoft/migration)

A simple database migration tool to integrate in your projects

###### If i miss something or you have something interesting, please be part of this project. Let me know! My contact is at the end.

## With support for
* Migration Up and Down options
* Custom Tags with handler to Up and Down options

## Mode's
* Database and Rabbitmq ```-mode all``` (default)
* Database ```-mode database```
* Rabbitmq ```-mode rabbitmq```

## Dependecy Management 
>### Dep

Project dependencies are managed using Dep. Read more about [Dep](https://github.com/golang/dep).
* Install dependencies: `dep ensure`
* Update dependencies: `dep ensure -update`


>### Go
```
go get github.com/joaosoft/migration
```

## Usage 
This examples are available in the project at [migration/main/cmd/main.go](https://github.com/joaosoft/migration/tree/master/main/cmd/main.go)
> Migration commands
```
// migrate up all migrations
migration -migrate up -mode database

// migrate up 2 migrations
migration -migrate up -number 2

// migrate down one migration
migration -migrate down

// migrate down 2 migration
migration -migrate down -number 2

// migrate down all migration
migration -migrate down -number -1
```

> By code
```
import (
	"flag"
	"fmt"
	github.com/joaosoft/migration/services
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
```


> Configuration file (the vhost can be omitted and set on the migration file)
```
{
  "migration": {
    "host": "localhost:8001",
    "path": {
      "database": "schema/db/postgres/example",
      "rabbitmq": "schema/rabbitmq/example"
    },
    "db": {
      "schema": "migration",
      "driver": "postgres",
      "datasource": "postgres://user:password@localhost:7000/postgres?sslmode=disable&search_path=migration"
    },
    "rabbitmq": {
      "host": "localhost:15672",
      "vhost": "dev"
    },
    "log": {
      "level": "info"
    }
  },
  "manager": {
    "log": {
      "level": "info"
    }
  },
  "client": {
    "log": {
      "level": "info"
    }
  }
}
```

> Migration file example
```
-- migrate up
CREATE TABLE migration.test1();

-- custom up
teste do joao A
teste do joao B



-- migrate down
DROP TABLE migration.test1;

-- custom down
teste do joao 1
teste do joao 2
```

## Administration
> Get a migration (GET)
```
http://localhost:8001/api/v1/migrations/<migration_name>
```
> Get migrations (GET)
```
http://localhost:8001/api/v1/migrations
```
>>+ Create a migration (POST)
```
http://localhost:8001/api/v1/migrations
```
with body:
```
{
	"id_migration": "<migration_name",
	"executed_at": "2018-06-02 13:11:37"
}
```
> Delete a migration (DELETE)
```
http://localhost:8001/api/v1/migrations/<migration_name>
```
> Delete migrations (DELETE)
```
http://localhost:8001/api/v1/migrations
```


## Known issues

## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com
