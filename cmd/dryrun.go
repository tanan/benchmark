package cmd

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(newRunCmd())
}

func newDryRunCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dryrun",
		Short: "Manage Run resources",
		RunE:  dryRunCmd,
	}
	return cmd
}

func dryRunCmd(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("resource name is required.")
	} else if args[0] == "help" {
		return cmd.Help()
	}

	return nil
}
