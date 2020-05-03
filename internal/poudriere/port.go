package poudriere

import (
	"os/exec"
	"strings"
)

const (
	Github int = iota
	Gitlab
)

const (
	UseGithub string = "USE_GITHUB"
	UseGitlab string = "USE_GITLAB"

	GithubAccount string = "GH_ACCOUNT"
	GithubProject string = "GH_PROJECT"

	GitlabAccount string = "GL_ACCOUNT"
	GitlabProject string = "GL_PROJECT"
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

func makeVar(dir string, value string) string {
	cmd, _ := exec.Command("make", "-V", value, "-C", dir).Output()
	return strings.Trim(string(cmd), "\n")
}

func repoInfo(dir string) *RepoT {
	if makeVar(dir, UseGithub) != "" {
		return &RepoT{
			makeVar(dir, GithubAccount),
			makeVar(dir, GithubProject),
			Github,
		}
	} else if makeVar(dir, UseGitlab) != "" {
		return &RepoT{
			makeVar(dir, GitlabAccount),
			makeVar(dir, GitlabProject),
			Gitlab,
		}
	}

	return &RepoT{}
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
