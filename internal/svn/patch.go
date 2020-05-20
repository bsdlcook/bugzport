package svn

import (
	"fmt"
	"os"
)

const (
	defaultReportDir string = ".report"
)

func (s *SvnInfo) WritePatch() {
	patchDir := fmt.Sprintf("%s%s/%s-%s/", s.WorkDir, defaultReportDir, s.Port.Name, s.Port.Version)
	patchFile := fmt.Sprintf("%s-%s.diff", s.Port.Name, s.Port.Version)

	file, _ := os.Create(patchDir + patchFile)
	defer file.Close()

	file.WriteString(s.generatePatch())
}

func (s *SvnInfo) generatePatch() string {
	diff := svnCmd("diff", s.Port.FullName())
	diff.Dir = s.WorkDir
	out, _ := diff.Output()

	return string(out)
}
