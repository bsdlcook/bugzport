package poudriere

import "fmt"

func GenerateReport(p *PortT) string {
	report := fmt.Sprintf(`
Amended:
 * Bumped DISTVERSION to %s and updated distinfo.

Changelog:
 * %s

 Tested:
  * portlint: OK (looks fine).
  * testport: OK (poudriere: %s).`, p.Version, changelog(p), poudriereVersion())

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
