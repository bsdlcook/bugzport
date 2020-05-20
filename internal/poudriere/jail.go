package poudriere

import (
	"fmt"
	"strings"
)

type Path struct {
	LogDir     string
	CacheDir   string
	ImageDir   string
	PackageDir string
	WorkDir    string
}

type Jail struct {
	Name    string
	Version string
	Arch    string
	Method  string
	Mount   string
	FS      string
	Path    *Path
}

func JailFromName(jail, tree string) (*Jail, error) {
	info, err := readJail(jail)
	if err != nil {
		return &Jail{}, err
	}

	paths := getPaths(info, tree)

	return &Jail{
		Name:    info["name"],
		Version: info["version"],
		Arch:    info["arch"],
		Method:  info["method"],
		Mount:   info["mount"],
		FS:      info["fs"],
		Path:    paths,
	}, nil
}

func readJail(jail string) (map[string]string, error) {
	out, err := poudriereCmd([]string{"jail", "-j", jail, "-i"}).Output()

	if err != nil {
		return nil, fmt.Errorf("no such jail '%s' found in Poudriere. Is the jail name correct?", jail)
	}

	info := make(map[string]string)

	for _, line := range strings.Split(string(out), "\n") {
		line = strings.TrimPrefix(line, "Jail ")
		if len(line) < 1 {
			continue
		}

		value := strings.Split(line, ":")

		info[value[0]] = strings.TrimSpace(value[1])
	}

	return info, nil
}

func getPaths(info map[string]string, tree string) *Path {
	return &Path{
		LogDir:     fmtPath(poudriereLogDir, info, tree),
		CacheDir:   fmtPath(poudriereCacheDir, info, tree),
		ImageDir:   fmtPath(poudriereImageDir, info, tree),
		PackageDir: fmtPath(poudrierePackageDir, info, tree),
		WorkDir:    fmtPath(poudriereWorkDir, info, tree),
	}
}

func fmtPath(path string, info map[string]string, tree string) string {
	return fmt.Sprintf("%s/%s-%s", path, info["name"], tree)
}
