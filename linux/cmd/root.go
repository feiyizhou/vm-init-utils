package cmd

import "github.com/spf13/cobra"

func NewLinuxRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(
		//hostnameChangeCmd,
		NewResetNetworkCMD(),
		//pwdChangeCmd,
	)
	return rootCmd
}
