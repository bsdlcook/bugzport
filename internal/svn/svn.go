package svn

import (
	"os/exec"

	"gitlab.com/lcook/bugzport/internal/port"
)

type SvnInfo struct {
	Port    *port.Port
	WorkDir string
}

func New(port *port.Port, dir string) *SvnInfo {
	return &SvnInfo{
		Port:    port,
		WorkDir: dir,
	}
}

func svnCmd(args ...string) *exec.Cmd {
	return exec.Command("svn", args...)
}
