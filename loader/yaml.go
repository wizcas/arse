package loader

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/wizcas/arse/data"
	y "gopkg.in/yaml.v2"
)

type yamlLoader struct {
	*baseLoader
}

// Load a YAML file from <filepath> into a parsed Arse configuration
func (l *yamlLoader) Load() (*data.ArseFile, error) {
	fmt.Printf("Loading file: %s...\n", l.filename)
	content, err := ioutil.ReadFile(l.filename)
	if err != nil {
		return nil, err
	}

	filemap := make(map[string]interface{})
	if err = y.Unmarshal(content, &filemap); err != nil {
		return nil, err
	}

	actions, err := loadActions(filemap)
	if err != nil {
		return nil, err
	}

	return &data.ArseFile{
		Actions: actions,
	}, nil
}

func loadActions(filemap map[string]interface{}) (data.ActionMap, error) {
	mapActions := asYm(filemap["actions"])
	if mapActions == nil {
		return nil, fmt.Errorf("actions not found")
	}

	collection := make(data.ActionMap)
	for k, v := range mapActions {
		name := k.(string)
		mapProps := asYm(v)
		if mapProps == nil {
			fmt.Printf("[WARN] %s is not a valid action\n", k)
			continue
		}

		collection[strings.ToLower(name)] = &data.Action{
			Name:        name,
			Description: mapProps.readString("description", ""),
			Script:      mapProps.readString("script", ""),
		}
	}
	return collection, nil
}
