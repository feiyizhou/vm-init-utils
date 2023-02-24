package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"runtime"
	"vm-init-utils/applications"
	"vm-init-utils/options"
	"vm-init-utils/utils"
	"vm-init-utils/validators"
)

func NewResetNetworkCMD() *cobra.Command {
	cleanFlagSet := pflag.NewFlagSet("changeDNS", pflag.ContinueOnError)
	networkFlags := options.NewNetworkFlags()
	cmd := &cobra.Command{
		Use:                "set-ip",
		Short:              "set-ip",
		Long:               "set-ip",
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cleanFlagSet.Parse(args); err != nil {
				return err
			}
			cmdArr := cleanFlagSet.Args()
			if len(cmdArr) > 0 {
				return fmt.Errorf("unknown command %s", cmdArr[0])
			}
			// short-circuit on help
			help, err := cleanFlagSet.GetBool("help")
			if err != nil {
				return errors.New(`"help" flag is non-bool, programmer error, please correct`)
			}
			if help {
				return cmd.Help()
			}
			os, err := utils.GetOSType()
			if len(cmdArr) > 0 {
				return fmt.Errorf("Failed get os type, err: %v ", err)
			}
			networkFlags.OsType = os
			if err = validators.ValidateFlagSet(*networkFlags); err != nil {
				return err
			}
			applications.NewResetNetworkPipeline(runtime.GOOS).ResetNetwork(networkFlags)
			return nil
		},
	}
	networkFlags.AddFlags(cleanFlagSet)
	cleanFlagSet.BoolP("help", "h", false, fmt.Sprintf("help for %s", cmd.Name()))
	// ugly, but necessary, because Cobra's default UsageFunc and HelpFunc pollute the flagset with global flags
	const usageFmt = "Usage:\n  %s\n\nFlags:\n%s"
	cmd.SetUsageFunc(func(cmd *cobra.Command) error {
		_, _ = fmt.Fprintf(cmd.OutOrStderr(), usageFmt, cmd.UseLine(), cleanFlagSet.FlagUsagesWrapped(2))
		return nil
	})
	cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		_, _ = fmt.Fprintf(cmd.OutOrStdout(), "%s\n\n"+usageFmt, cmd.Long, cmd.UseLine(), cleanFlagSet.FlagUsagesWrapped(2))
	})
	return cmd
}
