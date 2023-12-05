package celeritas

type Celeritas struct {
	AppName string
	Debug   bool
	Version string
}

type initPaths struct {
	rootPath    string
	folderNames []string
}
