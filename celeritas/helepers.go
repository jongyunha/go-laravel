package celeritas

import "os"

func (c *Celeritas) CreateDirIfNotExist(path string) error {
	const mode = 0755
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.Mkdir(path, mode)
	}

	return nil
}
