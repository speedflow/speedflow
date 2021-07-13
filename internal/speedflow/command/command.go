package command

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	cmdList "github.com/speedflow/speedflow/internal/speedflow/command/list"
	cmdVersion "github.com/speedflow/speedflow/internal/speedflow/command/version"
	ver "github.com/speedflow/speedflow/pkg/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
)

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		cobra.CheckErr(err)
		viper.AddConfigPath(dir)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".speedflow")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

// Execute executes command
func Execute(inIO io.Reader, outIO, errIO io.Writer) {
	_, err := ExecuteC(inIO, outIO, errIO)
	cobra.CheckErr(err)
}

func ExecuteC(in io.Reader, out, err io.Writer, args ...string) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:     "speedflow",
		Short:   "Increase your flow productivity with style",
		Version: ver.Version,
	}
	cmd.Run = run(cmd, in, out, err)

	cobra.OnInitialize(initConfig)

	cmd.SetIn(in)
	cmd.SetOut(out)
	cmd.SetErr(err)
	cmd.SetArgs(args)

	cmd.Flags().BoolP("list", "l", false, "List flows")

	cmd.AddCommand(cmdVersion.New(in, out, err))

	return cmd, cmd.Execute()
}

func run(rootCmd *cobra.Command, in io.Reader, out, errIO io.Writer) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if showList, _ := rootCmd.Flags().GetBool("list"); showList {
			err := cmdList.New(in, out, errIO).Execute()
			cobra.CheckErr(err)
			return
		}

		// TODO: implement
		fmt.Fprintln(out, "Hello World!")
	}
}
