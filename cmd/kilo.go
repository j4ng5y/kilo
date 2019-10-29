package cmd

import (
	"github.com/spf13/cobra"
	"k8s.io/klog"
)

const (
	kiloVersion = "0.1.0"
)

var kiloCmd = &cobra.Command{
	Use:     "kilo",
	Short:   "",
	Long:    "",
	Example: "",
	Version: kiloVersion,
	Run: func(ccmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := kiloCmd.Execute(); err != nil {
		klog.Fatal(err)
	}
}
