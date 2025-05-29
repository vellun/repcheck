package depender

import (
	"encoding/json"
	"github.com/vellun/repcheck/pkg/structs"
	"log"
	"os/exec"

	"golang.org/x/mod/modfile"
)

type Depender interface {
	GetDeps() ([]structs.DepInfo, bool)
}

type goDepender struct {
	modFile    *modfile.File
	noIndirect bool
}

func New(modFile *modfile.File, noIndirect bool) Depender {
	return &goDepender{
		modFile:    modFile,
		noIndirect: noIndirect,
	}
}

func (d *goDepender) GetDeps() ([]structs.DepInfo, bool) {
	var deps []structs.DepInfo
	updatesFound := false

	for _, req := range d.modFile.Require {
		update, err := checkModuleUpdate(req.Mod.Path, req.Mod.Version)
		if err != nil {
			log.Printf("Error occured for dep %s: %v\n", req.Mod.Path, err)
			continue
		}
		if update == "" {
			continue
		}

		if d.noIndirect && req.Indirect {
			continue
		}

		updatesFound = true

		deps = append(deps, structs.DepInfo{
			Path:          req.Mod.Path,
			CurVersion:    req.Mod.Version,
			UpdateVersion: update,
			IsIndirect:    req.Indirect,
		})
	}

	return deps, updatesFound
}

func checkModuleUpdate(modulePath, currentVersion string) (string, error) {
	cmd := exec.Command("go", "list", "-u", "-m", "-json", modulePath)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	type Module struct {
		Path    string
		Version string
		Update  struct {
			Version string
		}
	}
	var module Module
	if err := json.Unmarshal(output, &module); err != nil {
		return "", err
	}

	if module.Update.Version != "" && module.Update.Version != currentVersion {
		return module.Update.Version, nil
	}
	return "", nil
}
