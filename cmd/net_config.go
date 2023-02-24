package cmd

import "github.com/spf13/cobra"

func NewNetConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "net-conf",
		Short: "net-conf",
		Long:  "net-conf",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	return cmd
}
