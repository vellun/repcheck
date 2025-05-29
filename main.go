package main

import (
	"github.com/vellun/repcheck/pkg/cloner"
	"github.com/vellun/repcheck/pkg/depender"
	"github.com/vellun/repcheck/pkg/parser"
	"github.com/vellun/repcheck/pkg/printer"
	"github.com/vellun/repcheck/pkg/structs"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var jsonFlag bool
var noIndirectFlag bool

var rootCmd = &cobra.Command{
	Use:   "repcheck [repo url]",
	Short: "go deps checker",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repoURL := args[0]
		pwd, _ := os.Getwd()

		var cloner cloner.GitCloner
		repoDir, err := cloner.Clone(repoURL, pwd)
		defer os.RemoveAll(repoDir)

		if err != nil {
			log.Fatal(err)
		}

		var parser parser.GoModParser
		modFile, err := parser.Parse(repoDir)

		if err != nil {
			log.Fatal(err)
		}

		if err := os.Chdir(repoDir); err != nil {
			log.Fatalf("Ошибка: %v\n", err)
		}

		var pr printer.Printer
		pr = &printer.DefaultPrinter{}
		if jsonFlag {
			pr = &printer.JSONPrinter{}
		}

		info := structs.NewModuleInfo(
			modFile.Module.Mod.Path,
			modFile.Go.Version,
		)

		depender := depender.New(modFile, noIndirectFlag)
		deps, isUpdates := depender.GetDeps()

		pr.Print(info, deps, isUpdates)

	},
}

func Execute() {
	rootCmd.Flags().BoolVar(&jsonFlag, "json", false, "Print deps in json format")
	rootCmd.Flags().BoolVar(&noIndirectFlag, "noindirect", false, "Dont display indirect deps")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	Execute()
}
