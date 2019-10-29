package cmd

import "github.com/spf13/cobra"

var pushCmd = &cobra.Command{
	Use:     "push",
	Short:   "push local cluster changes to remote cluster",
	Long:    "",
	Example: "",
	Run:     pushFunc,
}

func init() {
	kiloCmd.AddCommand(pushCmd)
}

func pushFunc(ccmd *cobra.Command, args []string) {

}
