package celeritas

import "os"

func (c *Celeritas) CreateDirIfNotExist(path string) error {
	const mode = 0755
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.Mkdir(path, mode)
	}

	return nil
}

func (c *Celeritas) CreateFileIfNotExist(file string) error {
	const mode = 0755
	if _, err := os.Stat(file); os.IsNotExist(err) {
		file, err := os.Create(file)
		if err != nil {
			return err
		}

		defer func(file *os.File) {
			_ = file.Close()
		}(file)
	}
	return nil
}
