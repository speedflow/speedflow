package version

import (
	"github.com/spf13/cobra"

	"github.com/speedflow/speedflow/pkg/version"
)

// New returns a command to print version
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print Speedflow version",
		Run:   run,
	}

	return cmd
}

// run returns the command
func run(cmd *cobra.Command, args []string) {
	o, _ := cmd.Parent().PersistentFlags().GetString("output")
	version.Print(cmd.OutOrStdout(), o)
}
