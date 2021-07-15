package version

import (
	"github.com/spf13/cobra"

	"github.com/speedflow/speedflow/pkg/version"
)

var output = ""

// New returns a command to print version
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print Speedflow version",
		Run:   run,
	}

	cmd.Flags().StringVarP(&output, "output", "o", "", "One of '', 'yaml' or 'json'.")

	return cmd
}

// run returns the command
func run(cmd *cobra.Command, args []string) {
	version.Print(cmd.OutOrStdout(), output)
}
