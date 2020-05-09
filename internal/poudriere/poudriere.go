package poudriere

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	poudrierePrefix  string = "/usr/local/poudriere"
	poudriereDataDir string = poudrierePrefix + "/data"

	poudriereLogDir     string = poudriereDataDir + "/logs/bulk"
	poudriereCacheDir   string = poudriereDataDir + "/cache"
	poudriereImageDir   string = poudriereDataDir + "/images"
	poudrierePackageDir string = poudriereDataDir + "/packages"
	poudriereWorkDir    string = poudriereDataDir + "/wrkdirs"
)

func poudriereCmd(output bool, args ...string) *exec.Cmd {
	cmd := exec.Command("poudriere", args...)

	if output {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		fmt.Println(cmd.Stdout)
	}

	return cmd
}

func poudriereVersion() string {
	out, _ := poudriereCmd(false, "version").Output()
	return strings.Trim(string(out), "\n")
}
