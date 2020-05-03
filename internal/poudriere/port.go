package poudriere

import (
	"os/exec"
	"strings"
)

const (
	github int = iota
	gitlab

	githubUse string = "USE_GITHUB"
	gitlabUse string = "USE_GITLAB"

	githubAccount string = "GH_ACCOUNT"
	githubProject string = "GH_PROJECT"

	gitlabAccount string = "GL_ACCOUNT"
	gitlabProject string = "GL_PROJECT"
)

type RepoT struct {
	Account string
	Project string
	Type    int
}

type PortT struct {
	Name        string
	Version     string
	DistVersion string
	Category    string
	Maintainer  string
	Repo        *RepoT
}

func PortFromName(dir string) *PortT {
	return &PortT{
		Name:        makeVar(dir, "PORTNAME"),
		Version:     makeVar(dir, "PORTVERSION"),
		DistVersion: makeVar(dir, "DISTVERSIONFULL"),
		Category:    makeVar(dir, "CATEGORIES"),
		Maintainer:  makeVar(dir, "MAINTAINER"),
		Repo:        repoInfo(dir),
	}
}

func makeVar(dir string, value string) string {
	cmd, _ := exec.Command("make", "-V", value, "-C", dir).Output()
	return strings.Trim(string(cmd), "\n")
}

func repoInfo(dir string) *RepoT {
	if makeVar(dir, githubUse) != "" {
		return &RepoT{
			makeVar(dir, githubAccount),
			makeVar(dir, githubProject),
			github,
		}
	} else if makeVar(dir, gitlabUse) != "" {
		return &RepoT{
			makeVar(dir, gitlabAccount),
			makeVar(dir, gitlabProject),
			gitlab,
		}
	}

	return &RepoT{}
}
