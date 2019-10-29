package cmd

import "github.com/spf13/cobra"

var commitCmd = &cobra.Command{
	Use:     "commit",
	Short:   "commit local cluster changes",
	Long:    "",
	Example: "",
	Run:     commitFunc,
}

func init() {
	kiloCmd.AddCommand(commitCmd)
}

func commitFunc(ccmd *cobra.Command, args []string) {

}
