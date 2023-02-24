package cmd

import "github.com/spf13/cobra"

func NewChangeCmd() *cobra.Command {
	rootCmd := &cobra.Command{}

	rootCmd.AddCommand(
		//hostnameChangeCmd,
		ipChangeCmd,
		//pwdChangeCmd,
	)

	return rootCmd
}
