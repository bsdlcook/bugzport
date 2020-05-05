package poudriere

import (
	"fmt"

	"gitlab.com/lcook/bugzport/internal/svn"
	"gitlab.com/lcook/bugzport/utils"
)

type Job struct {
	Jail    *JailT
	Port    *PortT
	Tree    string
	WorkDir string
}

func (j *Job) Run() {
	svn := svn.New(j.Port.Name, j.Port.FullName(), j.Port.Version, j.WorkDir)
	build := buildStatus(j)

	utils.SpinStart(build)
	poudriereCmd("testport", "-j", j.Jail.Name, "-p", j.Tree, fmt.Sprintf("%s/%s", j.Port.Category, j.Port.Name)).Run()
	utils.SpinStop(build)

	WriteReport(j)
	svn.WritePatch()

	utils.CopyFile(portLogFile(j), reportLogFile(j))
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

func portLogFile(j *Job) string {
	return fmt.Sprintf("%s/latest/logs/%s-%s.log", j.Jail.Path.LogDir, j.Port.Name, j.Port.Version)
}

func reportLogFile(j *Job) string {
	return fmt.Sprintf("%s%s/%s-%s/%s-%s.log", j.WorkDir, defaultReportDir, j.Port.Name, j.Port.Version, j.Port.Name, j.Port.Version)
}
