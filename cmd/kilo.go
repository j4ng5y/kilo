package main

import (
	"github.com/j4ng5y/go-lumber"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

const versionFile = "./"

func versionFunc(log *lumber.Lumber) string {
	f, err := os.Open("./kilo_version.txt")
	if err != nil {
		log.Fatal(err.Error(), 1)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err.Error(), 1)
	}

	return string(b)
}

func kiloFunc(ccmd *cobra.Command, args []string) {

}

func main() {
	var (
		log = lumber.New()
		kiloCmd = &cobra.Command{
			Use:                        "kilo",
			Short:                      "",
			Long:                       "",
			Example:                    "",
			Version:                    versionFunc(log),
			PersistentPreRun:           nil,
			PreRun:                     nil,
			Run:                        kiloFunc,
			PostRun:                    nil,
			PersistentPostRun:          nil,
		}
	)

	if err := kiloCmd.Execute(); err != nil {
		log.Fatal(err.Error(), 1)
	}
}