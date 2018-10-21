package cmd

import "github.com/spf13/cobra"

var (
	Version  string
	Revision string
)

func init() {
	RootCmd.AddCommand(newVersionCmd())
}

func newVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print version",
	}
	return cmd
}
