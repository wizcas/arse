package data

import (
	"fmt"
	"strings"
)

// Action is a set of user defined commands run by Arse
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

// ActionMap contains the parsed Actions that are used for runtime lookup and execution.
type ActionMap map[string]*Action

// ActionList consists of parsed Actions and can be sorted by their names.
type ActionList []*Action

func (a ActionList) Len() int      { return len(a) }
func (a ActionList) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ActionList) Less(i, j int) bool {
	return strings.Compare(
		strings.ToLower(a[i].Name),
		strings.ToLower(a[j].Name),
	) < 0
}
