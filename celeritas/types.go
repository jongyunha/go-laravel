package celeritas

import "log"

type Celeritas struct {
	AppName  string
	Debug    bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	RootPath string
}

type initPaths struct {
	rootPath    string
	folderNames []string
}
