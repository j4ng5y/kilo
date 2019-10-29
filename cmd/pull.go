package cmd

import "github.com/spf13/cobra"

var pullCmd = &cobra.Command{
	Use:     "pull",
	Short:   "pull remote cluster changes to local cluster",
	Long:    "",
	Example: "",
	Run:     pullFunc,
}

func init() {
	kiloCmd.AddCommand(pullCmd)
}

func pullFunc(ccmd *cobra.Command, args []string) {

}
