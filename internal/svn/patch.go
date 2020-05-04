package svn

import (
	"fmt"
	"os"
)

const (
	defaultReportDir string = ".report"
)

func (s *SVNInfo) WritePatch() {
	patchDir := fmt.Sprintf("%s%s/%s-%s/", s.WorkDir, defaultReportDir, s.PortName, s.PortVersion)
	patchFile := fmt.Sprintf("%s-%s.diff", s.PortName, s.PortVersion)

	file, _ := os.Create(patchDir + patchFile)
	defer file.Close()

	file.WriteString(s.generatePatch())
}

func (s *SVNInfo) generatePatch() string {
	diff := svnCmd("diff", s.PortPath)
	diff.Dir = s.WorkDir
	out, _ := diff.Output()

	return string(out)
}
