package command

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	cmdVersion "github.com/speedflow/speedflow/internal/speedflow/command/version"
	ver "github.com/speedflow/speedflow/pkg/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string

	cmd = &cobra.Command{
		Use:     "speedflow",
		Short:   "Increase your flow productivity with style",
		Version: ver.Version,
	}
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
func Execute(in io.Reader, out, err io.Writer) {
	cobra.OnInitialize(initConfig)

	cmd.SetIn(in)
	cmd.SetOut(out)
	cmd.SetErr(err)

	cmd.AddCommand(cmdVersion.New(in, out, err))

	cobra.CheckErr(cmd.Execute())
}
