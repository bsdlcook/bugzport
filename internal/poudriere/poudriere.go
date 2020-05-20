package poudriere

import (
	"os/exec"
	"strings"
)

func poudriereCmd(args []string) *exec.Cmd {
	return exec.Command("poudriere", args...)
}

func poudriereVersion() string {
	out, _ := poudriereCmd([]string{"version"}).Output()
	return strings.Trim(string(out), "\n")
}
