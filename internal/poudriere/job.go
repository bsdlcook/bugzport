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

func (j *Job) Run() {
	build := buildStatus(j)

	utils.SpinStart(build)
	poudriereCmd("testport", "-j", j.Jail.Name, "-p", j.Tree, fmt.Sprintf("%s/%s", j.Port.Category, j.Port.Name)).Run()
	utils.SpinStop(build)
}

func buildStatus(j *Job) utils.SpinMessage {
	fmt.Print(fmt.Sprintf(`Builder environment:

	Version		-- %s
	Arch		-- %s
	Method		-- %s
	Mount		-- %s

`, j.Jail.Version, j.Jail.Arch, j.Jail.Method, j.Jail.Mount))

	buildMessage := fmt.Sprintf(" Building package %s/%s @ %s <%s>", j.Port.Category, j.Port.Name, j.Port.Version, j.Port.Maintainer)
	buildStatus := utils.Spinner(buildMessage)

	return buildStatus
}
