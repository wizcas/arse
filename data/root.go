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
