package poudriere

import (
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

func JailFromName(jail string) *JailT {
	info := readJail(jail)

	return &JailT{
		Name:    info["name"],
		Version: info["version"],
		Arch:    info["arch"],
		Method:  info["method"],
		Mount:   info["mount"],
		FS:      info["fs"],
	}
}

func readJail(jail string) map[string]string {
	out, _ := poudriereCmd("jail", "-j", jail, "-i").Output()
	info := make(map[string]string)

	for _, line := range strings.Split(string(out), "\n") {
		line = strings.TrimPrefix(line, "Jail ")
		if len(line) < 1 {
			continue
		}

		value := strings.Split(line, ":")

		info[value[0]] = strings.TrimSpace(value[1])
	}

	return info
}
