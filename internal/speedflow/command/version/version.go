package version

import (
	"io"

	"github.com/spf13/cobra"

	"github.com/speedflow/speedflow/pkg/version"
)

var output = ""

// New returns a command to print version
func New(in io.Reader, out, err io.Writer) (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "version",
		Short: "Print Speedflow version",
		Run:   run(out),
	}

	cmd.Flags().StringVarP(&output, "output", "o", "", "One of '', 'yaml' or 'json'.")

	return
}

// run returns the command
func run(out io.Writer) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		version.Print(out, output)
	}
}
