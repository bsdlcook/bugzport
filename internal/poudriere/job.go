package poudriere

import (
	"fmt"
	"io/ioutil"
	"os"

	"gitlab.com/lcook/bugzport/internal/svn"
	"gitlab.com/lcook/bugzport/utils"
)

func (j *Job) Run() {
	svn := svn.New(j.Port.Name, j.Port.FullName(), j.Port.Version, j.WorkDir)
	build := buildStatus(j)
	buildFlags := []string{"-j", j.Jail.Name, "-p", j.Tree, j.Port.FullName()}

	switch {
	case j.Options.Interactive:
		buildFlags = append([]string{"-i"}, buildFlags...)
		fallthrough
	case j.Options.Config:
		buildFlags = append([]string{"-c"}, buildFlags...)
		fallthrough
	default:
		buildFlags = append([]string{"testport"}, buildFlags...)
	}

	cmd := poudriereCmd(buildFlags)

	if j.Options.Output || j.Options.Interactive || j.Options.Config {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if j.Options.Interactive {
			cmd.Stdin = os.Stdin
		}
	}

	utils.SpinStart(build)
	cmd.Run()
	utils.SpinStop(build)

	if j.Options.Report {
		WriteReport(j)
		svn.WritePatch()
		utils.CopyFile(portLogFile(j), reportLogFile(j))
	}
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

	if j.Options.Output || j.Options.Interactive || j.Options.Config {
		buildStatus.Writer = ioutil.Discard
	}

	return buildStatus
}

func portLogFile(j *Job) string {
	return fmt.Sprintf("%s/latest/logs/%s-%s.log", j.Jail.Path.LogDir, j.Port.Name, j.Port.Version)
}

func reportLogFile(j *Job) string {
	return fmt.Sprintf("%s%s/%s-%s/%s-%s.log", j.WorkDir, defaultReportDir, j.Port.Name, j.Port.Version, j.Port.Name, j.Port.Version)
}
