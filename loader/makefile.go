package loader

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/wizcas/arse/data"
)

type makefileLoader struct {
	*baseLoader
}

var reTargetName *regexp.Regexp

func init() {
	compiled, err := regexp.Compile(`^([a-zA-Z]+[a-zA-Z\._-]+):`)
	if err != nil {
		log.Fatalf("failed to compile Makefile regexp: %v", err)
	}
	reTargetName = compiled
}

func (l *makefileLoader) Load() (*data.ArseFile, error) {

	f, err := os.Open(l.filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	usageBuffer := []string{}
	actions := make(data.ActionMap)
	_ = actions
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "## ") {
			usageBuffer = append(usageBuffer, line[3:])
			continue
		}
		matches := reTargetName.FindStringSubmatch(line)
		// Detect if this line defines a Makefile target.
		// If so, extract its name and register as an action.
		// Targets without usage text are considered as internal
		// and therefore won't be registered.
		if matches != nil && len(matches) >= 2 && len(usageBuffer) > 0 {
			// fmt.Printf("%t > %v\n", matches, matches)
			name := matches[1]
			actions[strings.ToLower(name)] = &data.Action{
				Name:        name,
				Description: strings.Join(usageBuffer, "\n"),
				Script:      fmt.Sprintf("make -f %s %s", l.filename, name),
			}
		}
		usageBuffer = []string{}
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return &data.ArseFile{
		Actions: actions,
	}, nil
}
