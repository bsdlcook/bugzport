package poudriere

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
)

type Job struct {
	Jail *JailT
	Port *PortT
	Tree string
}

var StatusBegin = func(j Job) *spinner.Spinner {
	infoMsg := "Builder environment:\n" +
		"\n\tVersion\t\t-- " + j.Jail.Version +
		"\n\tArch\t\t-- " + j.Jail.Arch +
		"\n\tMethod\t\t-- " + j.Jail.Method +
		"\n\tMount\t\t-- " + j.Jail.Mount + "\n\n"
	statusMsg := fmt.Sprintf(" Building package %s/%s @ %s <%s>", j.Port.Category, j.Port.Name, j.Port.Version, j.Port.Maintainer)

	fmt.Print(infoMsg)
	buildStatus := spinner.New(spinner.CharSets[11], 120*time.Millisecond)
	buildStatus.Color("blue")
	buildStatus.Suffix = statusMsg
	buildStatus.Start()

	return buildStatus
}

func (j Job) Run() {
	status := StatusBegin(j)
	PoudriereCmd("testport", "-j", j.Jail.Name, "-p", j.Tree, fmt.Sprintf("%s/%s", j.Port.Category, j.Port.Name)).Run()
	status.Stop()
}
