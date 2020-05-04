package svn

import (
	"os/exec"
)

type SVNInfo struct {
	PortName    string
	PortPath    string
	PortVersion string
	WorkDir     string
}

func New(port string, path string, version string, dir string) *SVNInfo {
	return &SVNInfo{
		PortName:    port,
		PortPath:    path,
		PortVersion: version,
		WorkDir:     dir,
	}
}

func svnCmd(args ...string) *exec.Cmd {
	return exec.Command("svn", args...)
}
