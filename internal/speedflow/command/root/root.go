package root

import (
	"fmt"

	"github.com/spf13/cobra"
)

// run returns the command
func Run(cmd *cobra.Command, args []string) {
	// TODO: Future usage of output
	o, _ := cmd.PersistentFlags().GetString("output")
	fmt.Printf("TODO: To implement root command (output: %s)\n", o)
}
