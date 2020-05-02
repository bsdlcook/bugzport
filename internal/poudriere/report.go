package poudriere

import "fmt"

func GenerateReport(port *PortT) string {
	report := "Amended:\n" +
		" * Bumped DISTVERSION to " + port.Version + " and updated distinfo.\n\n" +
		"Changelog:\n" +
		" * " + Changelog(port) + "\n\n" +
		"Tested:\n" +
		" * portlint: OK (looks fine).\n" +
		" * testport: OK (poudriere: " + PoudriereVersion() + ")."

	return report
}

func Changelog(port *PortT) string {
	switch port.Repo.Type {
	case Github:
		return fmt.Sprintf("https://github.com/%s/%s/releases/%s", port.Repo.Account, port.Repo.Project, port.DistVersion)
	case Gitlab:
		return fmt.Sprintf("https://gitlab.com/%s/%s/-/tags/%s", port.Repo.Account, port.Repo.Project, port.DistVersion)
	default:
		return "[change me]"
	}
}
