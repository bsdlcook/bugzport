package poudriere

import (
	"fmt"
	"strings"
)

type PathT struct {
	LogDir     string
	CacheDir   string
	ImageDir   string
	PackageDir string
	WorkDir    string
}

type JailT struct {
	Name    string
	Version string
	Arch    string
	Method  string
	Mount   string
	FS      string
	Path    *PathT
}

func JailFromName(jail string, tree string) (*JailT, error) {
	info, err := readJail(jail)
	if err != nil {
		return &JailT{}, err
	}

	paths := &PathT{
		LogDir:     getPath(poudriereLogDir, info, tree),
		CacheDir:   getPath(poudriereCacheDir, info, tree),
		ImageDir:   getPath(poudriereImageDir, info, tree),
		PackageDir: getPath(poudrierePackageDir, info, tree),
		WorkDir:    getPath(poudriereWorkDir, info, tree),
	}

	return &JailT{
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
	out, err := poudriereCmd(false, "jail", "-j", jail, "-i").Output()

	if err != nil {
		return nil, fmt.Errorf("No such jail '%s' found in Poudriere. Is the the jail name correct?", jail)
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

func getPath(path string, info map[string]string, tree string) string {
	return fmt.Sprintf("%s/%s-%s", path, info["name"], tree)
}
