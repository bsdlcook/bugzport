package poudriere

import (
	"os/exec"
	"strings"
)

const (
	poudrierePrefix  string = "/usr/local/poudriere"
	poudriereDataDir string = poudrierePrefix + "/data"

	poudriereLogDir     string = poudriereDataDir + "/logs"
	poudriereCacheDir   string = poudriereDataDir + "/cache"
	poudriereImageDir   string = poudriereDataDir + "/images"
	poudrierePackageDir string = poudriereDataDir + "/packages"
	poudriereWorkDir    string = poudriereDataDir + "/workdirs"
)

func poudriereCmd(args ...string) *exec.Cmd {
	return exec.Command("poudriere", args...)
}

func poudriereVersion() string {
	out, _ := poudriereCmd("version").Output()
	return strings.Trim(string(out), "\n")
}
