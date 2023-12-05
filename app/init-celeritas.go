package main

import (
	"github.com/jongyunha/celeritas"
	"log"
	"os"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	cel := &celeritas.Celeritas{}
	err = cel.New(path)
	if err != nil {
		log.Fatal(nil)
	}

	cel.AppName = "app"
	cel.InfoLog.Println("Debug is set to", cel.Debug)
	return &application{App: cel}
}
