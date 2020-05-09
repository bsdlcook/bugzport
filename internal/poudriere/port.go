package poudriere

import (
	"fmt"
	"os"
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

const (
	gomod int = iota
	cargo

	gomodUse = "go:modules"
	cargoUse = "cargo"
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
	Uses        int
}

func (p *PortT) FullName() string {
	return fmt.Sprintf("%s/%s", p.Category, p.Name)
}

func PortFromName(dir string) (*PortT, error) {
	valid := isPort(dir)

	if valid != nil {
		return &PortT{}, valid
	}

	return &PortT{
		Name:        makeVar(dir, "PORTNAME"),
		Version:     makeVar(dir, "PORTVERSION"),
		DistVersion: makeVar(dir, "DISTVERSIONFULL"),
		Category:    makeVar(dir, "CATEGORIES"),
		Maintainer:  makeVar(dir, "MAINTAINER"),
		Repo:        repoInfo(dir),
		Uses:        usesInfo(dir),
	}, nil
}

func isPort(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return fmt.Errorf("'%s' is not a valid port to build", dir)
	}

	return nil
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

func usesInfo(dir string) int {
	uses := makeVar(dir, "USES")

	for _, use := range strings.Fields(uses) {
		switch use {
		case gomodUse:
			return gomod
		case cargoUse:
			return cargo
		}
	}

	return -1
}
