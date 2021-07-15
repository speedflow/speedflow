package command

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	cmdList "github.com/speedflow/speedflow/internal/speedflow/command/list"
	cmdVersion "github.com/speedflow/speedflow/internal/speedflow/command/version"
	"github.com/speedflow/speedflow/internal/speedflow/speedflow"
	ver "github.com/speedflow/speedflow/pkg/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfg string

	cfgSubPath = ".config/speedflow/"
	cfgFile    = "speedflow.yml"

	sfFile = ".speedflow.yml"
)

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfg != "" {
		viper.SetConfigFile(cfg)
	} else {
		viper.AddConfigPath(cfgPath())
		viper.SetConfigType("yml")
		viper.SetConfigName(cfgFile)
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

// ExecuteC executes command with arguments and returns command
func ExecuteC(in io.Reader, out, errIO io.Writer, args ...string) (*cobra.Command, error) {
	// Cobra initialization
	cobra.OnInitialize(initConfig)

	// Create root command
	cmd := &cobra.Command{
		Use:     "speedflow",
		Short:   "Increase your flow productivity with style",
		Version: ver.Version,
	}
	cmd.Run = run(cmd, in, out, errIO)

	cmd.SetIn(in)
	cmd.SetOut(out)
	cmd.SetErr(errIO)
	cmd.SetArgs(args)

	// Add subcommands
	cmd.AddCommand(cmdVersion.New(in, out, errIO))

	// Add flags
	flags(cmd)

	return cmd, cmd.Execute()
}

func flags(cmd *cobra.Command) {
	// Current directory
	currentDir, err := os.Getwd()
	cobra.CheckErr(err)

	cmd.Flags().BoolP("list", "l", false, "List flows")
	cmd.Flags().StringP("file", "f", filepath.Join(currentDir, sfFile), "Speedflow file")
	cmd.Flags().StringP("config", "c", filepath.Join(cfgPath(), cfgFile), "Configuration file")
}

func run(rootCmd *cobra.Command, in io.Reader, out, errIO io.Writer) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		// Load speedflow file
		f, err := cmd.Flags().GetString("file")
		cobra.CheckErr(err)
		err = speedflow.Load(f)
		cobra.CheckErr(err)

		// Display list
		if showList, _ := rootCmd.Flags().GetBool("list"); showList {
			err := cmdList.New(in, out, errIO).Execute()
			cobra.CheckErr(err)
			return
		}

		// Execute command
		// TODO: implement
		fmt.Fprintln(out, "Hello World!")
	}
}

func cfgPath() string {
	// Home directory
	homeDir, err := os.UserHomeDir()
	cobra.CheckErr(err)

	return filepath.Join(homeDir, cfgSubPath)
}
