package poudriere

import (
	"os/exec"
	"strings"
)

var PoudriereCmd = func(args ...string) *exec.Cmd {
	return exec.Command("poudriere", args...)
}

var PoudriereVersion = func() string {
	out, _ := PoudriereCmd("version").Output()
	return strings.Trim(string(out), "\n")
}
