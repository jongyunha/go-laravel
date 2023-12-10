package celeritas

import (
	"github.com/go-chi/chi/v5"
	"github.com/jongyunha/celeritas/render"
	"log"
)

type Celeritas struct {
	AppName  string
	Debug    bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	RootPath string
	Routes   *chi.Mux
	Render   *render.Render
	config   config
}

type initPaths struct {
	rootPath    string
	folderNames []string
}

type config struct {
	port     string
	renderer string
}
