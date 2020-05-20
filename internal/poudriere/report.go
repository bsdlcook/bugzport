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

	file.WriteString(generateReport(j.Port))
}

func generateReport(p *port.Port) string {
	report := fmt.Sprintf(`%s: Update to %s

Amended:
 * Bumped DISTVERSION to %s and updated distinfo.
 %s

Changelog:
 * %s

Tested:
 * portlint: OK (looks fine).
 * testport: OK (poudriere: %s).`, p.FullName(), p.Version, p.Version, uses(p), changelog(p), poudriereVersion())

	return report
}

func changelog(p *port.Port) string {
	switch p.Repo.Type {
	case port.Github:
		return fmt.Sprintf("https://github.com/%s/%s/releases/%s", p.Repo.Account, p.Repo.Project, p.DistVersion)
	case port.Gitlab:
		return fmt.Sprintf("https://gitlab.com/%s/%s/-/tags/%s", p.Repo.Account, p.Repo.Project, p.DistVersion)
	default:
		return "[change me]"
	}
}

func uses(p *port.Port) string {
	switch p.Uses {
	case port.Gomod:
		return "* Updated *_TUPLE dependency list."
	case port.Cargo:
		return "* Updated CARGO_CRATES dependency list."
	default:
		return ""
	}
}
