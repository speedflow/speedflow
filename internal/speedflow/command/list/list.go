package list

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

var output = ""

// New returns a command to print list
func New(in io.Reader, out, err io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all flows",
		Run:   run(out),
	}

	cmd.SetIn(in)
	cmd.SetOut(out)
	cmd.SetErr(err)

	cmd.Flags().StringVarP(&output, "output", "o", "", "One of '', 'yaml' or 'json'.")

	return cmd
}

// run returns the command
func run(out io.Writer) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		// TODO: implement
		fmt.Fprintln(out, "Flow     Name        ")
		fmt.Fprintln(out, "default  Default flow")
	}
}
