package main

import (
	"github.com/joaosoft/migration/services"
)

func main() {
	m, err := services.NewWebService()
	if err != nil {
		panic(err)
	}

	if err := m.Start(); err != nil {
		panic(err)
	}
}
