package jail

const (
	poudrierePrefix  string = "/usr/local/poudriere"
	poudriereDataDir string = poudrierePrefix + "/data"

	jailLogDir     string = poudriereDataDir + "/logs/bulk"
	jailCacheDir   string = poudriereDataDir + "/cache"
	jailImageDir   string = poudriereDataDir + "/images"
	jailPackageDir string = poudriereDataDir + "/packages"
	jailWorkDir    string = poudriereDataDir + "/wrkdirs"
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
