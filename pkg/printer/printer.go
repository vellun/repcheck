package printer

import (
	"encoding/json"
	"fmt"
	"github.com/vellun/repcheck/pkg/structs"
	"log"
)

type Printer interface {
	Print(info structs.ModuleInfo, deps []structs.DepInfo, isUpdates bool)
}

type DefaultPrinter struct{}
type JSONPrinter struct{}

func (c *DefaultPrinter) Print(info structs.ModuleInfo, deps []structs.DepInfo, isUpdates bool) {
	fmt.Printf("Module: %s\n", info.Name)
	fmt.Printf("Go version: %s\n", info.GoVersion)

	if !isUpdates {
		fmt.Println("\nNo deps to update")
		return
	}

	fmt.Println("\nDeps to update:")

	for _, dep := range deps {
		res := fmt.Sprintf("%s: %s -> %s", dep.Path, dep.CurVersion, dep.UpdateVersion)
		if dep.IsIndirect {
			res += " // indirect"
		}

		fmt.Println(res)
	}
}

func (j *JSONPrinter) Print(info structs.ModuleInfo, deps []structs.DepInfo, isUpdates bool) {
	output := struct {
		Module structs.ModuleInfo `json:"module"`
		Deps   []structs.DepInfo  `json:"deps"`
	}{
		Module: info,
		Deps:   deps,
	}

	if !isUpdates {
		fmt.Println("\nNo deps to update")
		return
	}

	data, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		log.Fatalf("JSON formatting error: %v\n", err)
	}
	fmt.Println(string(data))
}
