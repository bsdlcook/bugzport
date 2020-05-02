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

var MakeVar = func(dir string, value string) string {
	cmd, _ := exec.Command("make", "-V", value, "-C", dir).Output()
	return strings.Trim(string(cmd), "\n")
}

var RepoInfo = func(dir string) *RepoT {
	if MakeVar(dir, UseGithub) != "" {
		return &RepoT{
			MakeVar(dir, GithubAccount),
			MakeVar(dir, GithubProject),
			Github,
		}
	} else if MakeVar(dir, UseGitlab) != "" {
		return &RepoT{
			MakeVar(dir, GitlabAccount),
			MakeVar(dir, GitlabProject),
			Gitlab,
		}
	}

	return &RepoT{}
}

func PortFromName(dir string) *PortT {
	return &PortT{
		Name:        MakeVar(dir, "PORTNAME"),
		Version:     MakeVar(dir, "PORTVERSION"),
		DistVersion: MakeVar(dir, "DISTVERSIONFULL"),
		Category:    MakeVar(dir, "CATEGORIES"),
		Maintainer:  MakeVar(dir, "MAINTAINER"),
		Repo:        RepoInfo(dir),
	}
}
