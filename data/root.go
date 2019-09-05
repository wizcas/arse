package data

import (
	"fmt"
	"strings"
)

// ArseFile is the root runtime configuration parsed from an input config file.
type ArseFile struct {
	Actions ActionMap
}

func (f ArseFile) String() string {
	arrActionString := make([]string, len(f.Actions))
	index := 0
	for _, action := range f.Actions {
		arrActionString[index] = action.String()
		index++
	}
	return fmt.Sprintf(`
/////////////// ARSEFILE ///////////////

%d action(s)
%s

////////////////////////////////////////`, len(f.Actions), strings.Join(arrActionString, "\n"))
}

// Action finder by case-insensitive name. Returns nil if not found.
func (f *ArseFile) Action(name string) *Action {
	action, ok := f.Actions[name]
	if !ok {
		return nil
	}
	return action
}
