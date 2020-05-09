package svn

import (
	"os/exec"
)

type SvnInfo struct {
	PortName    string
	PortPath    string
	PortVersion string
	WorkDir     string
}

func New(port string, path string, version string, dir string) *SvnInfo {
	return &SvnInfo{
		PortName:    port,
		PortPath:    path,
		PortVersion: version,
		WorkDir:     dir,
	}
}

func svnCmd(args ...string) *exec.Cmd {
	return exec.Command("svn", args...)
}
