package flows

import (
	"fmt"

	"github.com/spf13/cobra"
)

var output = ""

// New returns a command to list flows
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "flows",
		Short: "List flows",
		Run:   run,
	}

	cmd.Flags().StringVarP(&output, "output", "o", "", "One of '', 'yaml' or 'json'.")

	return cmd
}

// run returns the command
func run(cmd *cobra.Command, args []string) {
	fmt.Println("TODO: To implement")
}
