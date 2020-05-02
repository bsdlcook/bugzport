package poudriere

import (
	"fmt"
	"github.com/briandowns/spinner"
	"time"
)

type Job struct {
	Jail *JailT
	Port *PortT
	Tree string
}

var StatusBegin = func(job Job) *spinner.Spinner {
	infoMsg := "Builder enviroment:\n" +
		"\n\tVersion\t\t-- " + job.Jail.Version +
		"\n\tArch\t\t-- " + job.Jail.Arch +
		"\n\tMethod\t\t-- " + job.Jail.Method +
		"\n\tMount\t\t-- " + job.Jail.Mount + "\n\n"
	statusMsg := fmt.Sprintf(" Building package %s/%s @ %s <%s>", job.Port.Category, job.Port.Name, job.Port.Version, job.Port.Maintainer)

	fmt.Print(infoMsg)
	buildStatus := spinner.New(spinner.CharSets[11], 120*time.Millisecond)
	buildStatus.Color("blue")
	buildStatus.Suffix = statusMsg
	buildStatus.Start()

	return buildStatus
}

func (job Job) Run() {
	status := StatusBegin(job)
	PoudriereCmd("testport", "-j", job.Jail.Name, "-p", job.Tree, fmt.Sprintf("%s/%s", job.Port.Category, job.Port.Name)).Run()
	status.Stop()
}
