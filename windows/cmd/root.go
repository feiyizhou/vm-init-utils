package cmd

import (
	"github.com/spf13/cobra"
)

func NewWinRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(NewResetNetworkCMD())
	return rootCmd
}
