package main

import (
	"app/handlers"
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
	myHandlers := &handlers.Handlers{
		App: cel,
	}
	app := &application{
		App:      cel,
		Handlers: myHandlers,
	}

	app.App.Routes = app.routes()
	return app
}
