package parser

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sort"
	"strings"
)

func Load(filepath string) (*ArseFile, error) {
	fmt.Printf("Loading file: %s...\n", filepath)
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	filemap := make(map[string]interface{})
	if err = yaml.Unmarshal(content, &filemap); err != nil {
		return nil, err
	}

	actions, err := loadActions(filemap)
	if err != nil {
		return nil, err
	}

	return &ArseFile{
		actions,
	}, nil
}

func loadActions(filemap map[string]interface{}) (ActionCollection, error) {
	mapActions := asYm(filemap["actions"])
	if mapActions == nil {
		return nil, fmt.Errorf("actions not found\n")
	}
	actionList := sortableActions{}
	for k, v := range mapActions {
		name := k.(string)
		mapProps := asYm(v)
		if mapProps == nil {
			fmt.Printf("[WARN] %s is not a valid action\n", k)
			continue
		}

		actionList = append(actionList, &Action{
			Name:        name,
			Description: mapProps.readString("description", ""),
			Script:      mapProps.readString("script", ""),
		})
	}

	sort.Sort(actionList)
	collection := make(ActionCollection)
	for _, action := range actionList {
		collection[strings.ToLower(action.Name)] = action
	}
	return collection, nil
}
