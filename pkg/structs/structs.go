package structs

type ModuleInfo struct {
	Name      string
	GoVersion string
}

type DepInfo struct {
	Path          string
	CurVersion    string
	UpdateVersion string
	IsIndirect    bool
}

func NewModuleInfo(name string, goVersion string) ModuleInfo {
	return ModuleInfo{
		Name:      name,
		GoVersion: goVersion,
	}
}
