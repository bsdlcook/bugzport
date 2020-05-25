package port

const (
	Github int = iota + 1
	Gitlab

	githubUse string = "USE_GITHUB"
	gitlabUse string = "USE_GITLAB"

	githubAccount string = "GH_ACCOUNT"
	githubProject string = "GH_PROJECT"

	gitlabAccount string = "GL_ACCOUNT"
	gitlabProject string = "GL_PROJECT"
)

const (
	Gomod int = iota
	Cargo

	gomodUse = "go:modules"
	cargoUse = "cargo"
)

type Repo struct {
	Account string
	Project string
	Type    int
}

type Meta struct {
	LogName string
}

type Port struct {
	Name        string
	Version     string
	DistVersion string
	Category    string
	Maintainer  string
	Repo        *Repo
	Uses        int
	Meta        *Meta
}
