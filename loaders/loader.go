package loaders

import "github.com/wizcas/arse/data"

// Loader reads a config file and parse it into an ArseFile instance.
// Implement this interface to support various config file format.
type Loader interface {
	Load(filename string) (*data.ArseFile, error)
}

func SmartLoad(filename string) (*data.ArseFile, error) {
	var loader Loader
	loader = yaml
}
