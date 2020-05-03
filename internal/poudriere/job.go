package poudriere

import (
	"fmt"

	"gitlab.com/lcook/bugzport/utils"
)

type Job struct {
	Jail *JailT
	Port *PortT
	Tree string
}

var buildStatus = func(j Job) utils.SpinMessage {
	infoMessage := "Builder environment:\n" +
		"\n\tVersion\t\t-- " + j.Jail.Version +
		"\n\tArch\t\t-- " + j.Jail.Arch +
		"\n\tMethod\t\t-- " + j.Jail.Method +
		"\n\tMount\t\t-- " + j.Jail.Mount + "\n\n"
	fmt.Print(infoMessage)

	buildMessage := fmt.Sprintf(" Building package %s/%s @ %s <%s>", j.Port.Category, j.Port.Name, j.Port.Version, j.Port.Maintainer)
	buildStatus := utils.Spinner(buildMessage)

	return buildStatus
}

func (j Job) Run() {
	build := buildStatus(j)

	utils.SpinStart(build)
	PoudriereCmd("testport", "-j", j.Jail.Name, "-p", j.Tree, fmt.Sprintf("%s/%s", j.Port.Category, j.Port.Name)).Run()
	utils.SpinStop(build)
}
