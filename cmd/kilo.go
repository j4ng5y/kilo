package main

import (
	"github.com/spf13/cobra"
	"k8s.io/klog"
)

const (
	kiloVersion = "0.1.0"
	etcConfigPath = "/etc/kilo"
	varConfigPath = "/var/lib/kilo"
)

var configFile string

func kiloFunc(ccmd *cobra.Command, args []string) {}

func main() {
	var (
		kiloCmd = &cobra.Command{
			Use:                        "kilo",
			Short:                      "",
			Long:                       "",
			Example:                    "",
			Version:                    kiloVersion,
			PersistentPreRun:           nil,
			PreRun:                     nil,
			Run:                        kiloFunc,
			PostRun:                    nil,
			PersistentPostRun:          nil,
		}

		initializeCmd = &cobra.Command{
			Use:                        "init",
			Short:                      "initialize kilo",
			Long:                       "",
			Example:                    "",
			PersistentPreRun:           nil,
			PreRun:                     nil,
			Run:                        initFunc,
			PostRun:                    nil,
			PersistentPostRun:          nil,
		}

		generateCmd = &cobra.Command{
			Use:                        "generate",
			Short:                      "generate a configuration file",
			Long:                       "",
			Example:                    "",
			PersistentPreRun:           nil,
			PreRun:                     nil,
			Run:                        genFunc,
			PostRun:                    nil,
			PersistentPostRun:          nil,
		}

		commitCmd = &cobra.Command{
			Use:                        "commit",
			Short:                      "commit local cluster changes",
			Long:                       "",
			Example:                    "",
			PersistentPreRun:           nil,
			PreRun:                     nil,
			Run:                        commitFunc,
			PostRun:                    nil,
			PersistentPostRun:          nil,
		}

		pushCmd = &cobra.Command{
			Use:                        "push",
			Short:                      "push local cluster changes to remote cluster",
			Long:                       "",
			Example:                    "",
			PersistentPreRun:           nil,
			PreRun:                     nil,
			Run:                        pushFunc,
			PostRun:                    nil,
			PersistentPostRun:          nil,
		}

		pullCmd = &cobra.Command{
			Use:                        "pull",
			Short:                      "pull remote cluster changes to local cluster",
			Long:                       "",
			Example:                    "",
			PersistentPreRun:           nil,
			PreRun:                     nil,
			Run:                        pullFunc,
			PostRun:                    nil,
			PersistentPostRun:          nil,
		}
	)

	kiloCmd.AddCommand(initializeCmd, generateCmd, pushCmd, pullCmd, commitCmd)

	initializeCmd.Flags().StringVarP(&configFile, "config-file", "f", "", "The config file to use")

	if err := kiloCmd.Execute(); err != nil {
		klog.Fatal(err.Error())
	}
}