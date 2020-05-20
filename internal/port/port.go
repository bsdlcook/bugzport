package port

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func (p *Port) FullName() string {
	return fmt.Sprintf("%s/%s", p.Category, p.Name)
}

func FromName(dir string) (*Port, error) {
	valid := isPort(dir)

	if valid != nil {
		return &Port{}, valid
	}

	return &Port{
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

func makeVar(dir, value string) string {
	cmd, _ := exec.Command("make", "-V", value, "-C", dir).Output()
	return strings.Trim(string(cmd), "\n")
}

func repoInfo(dir string) *Repo {
	if makeVar(dir, githubUse) != "" {
		return &Repo{
			makeVar(dir, githubAccount),
			makeVar(dir, githubProject),
			Github,
		}
	} else if makeVar(dir, gitlabUse) != "" {
		return &Repo{
			makeVar(dir, gitlabAccount),
			makeVar(dir, gitlabProject),
			Gitlab,
		}
	}

	return &Repo{}
}

func usesInfo(dir string) int {
	uses := makeVar(dir, "USES")

	for _, use := range strings.Fields(uses) {
		switch use {
		case gomodUse:
			return Gomod
		case cargoUse:
			return Cargo
		}
	}

	return -1
}
