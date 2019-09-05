package loader

import (
	"errors"
	"strings"

	"github.com/wizcas/arse/data"
)

type baseLoader struct {
	filename string
}

// Loader reads a config file and parse it into an ArseFile instance.
// Implement this interface to support various config file format.
type Loader interface {
	Load() (*data.ArseFile, error)
}

// SmartLoad tries to find the right file loader by filename & file extensions,
// and parse & load the file into an ArseFile.
func SmartLoad(filename string) (*data.ArseFile, error) {
	loader := getLoader(filename)
	if loader == nil {
		return nil, errors.New("unknown file type")
	}
	return loader.Load()
}

func getLoader(filename string) Loader {
	if strings.HasSuffix(filename, "Makefile") {
		return &makefileLoader{&baseLoader{filename}}
	}
	if strings.HasSuffix(filename, ".yml") {
		return &yamlLoader{&baseLoader{filename}}
	}
	return nil
}
