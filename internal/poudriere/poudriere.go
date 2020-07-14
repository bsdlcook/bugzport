package poudriere

import (
	"os/exec"
)

func poudriereCmd(args []string) *exec.Cmd {
	return exec.Command("poudriere", args...)
}
