package poudriere

import (
	"fmt"
	"os"
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

func generateReport(p *PortT) string {
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

func changelog(p *PortT) string {
	switch p.Repo.Type {
	case github:
		return fmt.Sprintf("https://github.com/%s/%s/releases/%s", p.Repo.Account, p.Repo.Project, p.DistVersion)
	case gitlab:
		return fmt.Sprintf("https://gitlab.com/%s/%s/-/tags/%s", p.Repo.Account, p.Repo.Project, p.DistVersion)
	default:
		return "[change me]"
	}
}

func uses(p *PortT) string {
	switch p.Uses {
	case gomod:
		return "* Updated *_TUPLE dependeny list."
	case cargo:
		return "* Updated CARGO_CRATES dependeny list."
	default:
		return ""
	}
}
