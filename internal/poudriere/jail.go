package poudriere

import (
	"fmt"
	"strings"
)

type JailT struct {
	Name    string
	Version string
	Arch    string
	Method  string
	Mount   string
	FS      string
}

func JailFromName(jail string) (*JailT, error) {
	info, err := readJail(jail)
	if err != nil {
		return &JailT{}, err
	}

	return &JailT{
		Name:    info["name"],
		Version: info["version"],
		Arch:    info["arch"],
		Method:  info["method"],
		Mount:   info["mount"],
		FS:      info["fs"],
	}, nil
}

func readJail(jail string) (map[string]string, error) {
	out, err := poudriereCmd("jail", "-j", jail, "-i").Output()

	if err != nil {
		return nil, fmt.Errorf("No such jail '%s' found in Poudriere. Is the the jail name correct?", jail)
	}

	info := make(map[string]string)

	for _, line := range strings.Split(string(out), "\n") {
		line = strings.TrimPrefix(line, "Jail ")
		if len(line) < 1 {
			continue
		}

		value := strings.Split(line, ":")

		info[value[0]] = strings.TrimSpace(value[1])
	}

	return info, nil
}
