package parser

import (
	"fmt"
	"strings"
)

type ArseFile struct {
	actions ActionCollection
}

func (f ArseFile) String() string {
	arrActionString := make([]string, len(f.actions))
	index := 0
	for _, action := range f.actions {
		arrActionString[index] = action.String()
		index++
	}
	return fmt.Sprintf(`
/////////////// ARSEFILE ///////////////

%d action(s)
%s

////////////////////////////////////////`, len(f.actions), strings.Join(arrActionString, "\n"))
}

type ActionCollection map[string]*Action

type Action struct {
	Name        string
	Description string
	Script      string
}

func (a Action) String() string {
	return fmt.Sprintf(`
========================================
Name:         %s
Description:  %s
........................................

%s
========================================`, a.Name, a.Description, a.Script)
}

type sortableActions []*Action

func (a sortableActions) Len() int      { return len(a) }
func (a sortableActions) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a sortableActions) Less(i, j int) bool {
	return strings.Compare(
		strings.ToLower(a[i].Name),
		strings.ToLower(a[j].Name),
	) < 0
}
