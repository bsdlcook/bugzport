package poudriere

import (
	"fmt"
	"os"

	"gitlab.com/lcook/bugzport/internal/port"
)

const (
	defaultReportDir  string = ".report"
	defaultReportFile string = "summary"
)

func WriteReport(j *Job) {
	reportDir := fmt.Sprintf("%s%s/%s-%s/", j.WorkDir, defaultReportDir, j.Port.Name, j.Port.Version)

	os.MkdirAll(reportDir, os.ModePerm)
	file, _ := os.Create(reportDir + defaultReportFile)

	defer file.Close()

	file.WriteString(generateReport(j))
}

func generateReport(j *Job) string {
	report := fmt.Sprintf(`%s: Update to %s

Changelog:

 * %s

QA:

 * portlint: OK (looks fine).
 * testport: OK (poudriere: %s, %s).

MFH: <Yes/No> (<reason>, <merge to quarterly branch>).`, j.Port.FullName(), j.Port.Version, changelog(j.Port), j.Jail.Version, j.Jail.Arch)

	return report
}

func changelog(p *port.Port) string {
	switch p.Repo.Type {
	case port.Github:
		return fmt.Sprintf("https://github.com/%s/%s/releases/%s", p.Repo.Account, p.Repo.Project, p.DistVersion)
	case port.Gitlab:
		return fmt.Sprintf("https://gitlab.com/%s/%s/-/tags/%s", p.Repo.Account, p.Repo.Project, p.DistVersion)
	default:
		return "<change me>"
	}
}
