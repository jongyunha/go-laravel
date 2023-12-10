package main

import (
	"app/handlers"
	"github.com/jongyunha/celeritas"
)

type application struct {
	App      *celeritas.Celeritas
	Handlers *handlers.Handlers
}

func main() {
	c := initApplication()

	c.App.ListenAndServe()
}
