package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"vm-init-utils/linux/options"
	"vm-init-utils/linux/services"
	"vm-init-utils/linux/utils"
	"vm-init-utils/linux/validators"
)

func NewResetNetworkCMD() *cobra.Command {
	cleanFlagSet := pflag.NewFlagSet("changeIP", pflag.ContinueOnError)
	networkFlags := options.NewNetworkFlags()
	cmd := &cobra.Command{
		Use:                "set-ip",
		Short:              "set-ip",
		Long:               "set-ip",
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				err    error
				help   bool
				osInfo *utils.Data
			)
			if err = cleanFlagSet.Parse(args); err != nil {
				return utils.MadeErr(err, "Filed to parse flag")
			}
			cmdArr := cleanFlagSet.Args()
			if len(cmdArr) > 0 {
				return utils.MadeErr(nil, fmt.Sprintf("unknown command %v", cmdArr))
			}

			// short-circuit on help
			help, err = cleanFlagSet.GetBool("help")
			if err != nil {
				return utils.MadeErr(err, `"help" flag is non-bool, programmer error, please correct`)
			}
			if help {
				return cmd.Help()
			}

			osInfo, err = utils.OSInfo()
			if err != nil {
				return utils.MadeErr(err, "Failed to get os info")
			}
			if osInfo == nil {
				return utils.MadeErr(nil, "OS info is nil")
			}
			networkFlags.OsType = osInfo.ID

			if err = validators.ValidateFlagSet(*networkFlags); err != nil {
				return utils.MadeErr(err, "The param is invalid")
			}

			services.NewLinuxService().ReSetNetWork(networkFlags)
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
