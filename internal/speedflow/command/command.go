package command

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	cmdFlows "github.com/speedflow/speedflow/internal/speedflow/command/flows"
	cmdRoot "github.com/speedflow/speedflow/internal/speedflow/command/root"
	cmdVersion "github.com/speedflow/speedflow/internal/speedflow/command/version"
	"github.com/speedflow/speedflow/internal/speedflow/speedflow"
	ver "github.com/speedflow/speedflow/pkg/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Exitable is used to exit on error
	// Useful for tests
	Exitable = true

	cfgPathFile string
	cfgSubPath  = ".config/speedflow/"
	cfgFile     = "speedflow.yml"

	spPathFile string
	sfFile     = ".speedflow.yml"
)

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgPathFile != "" {
		viper.SetConfigFile(cfgPathFile)
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
func Execute() {
	defer func() {
		if r := recover(); r != nil {
			// TODO: Improve error message color
			fmt.Println("Internal Speedflow error")
			// TODO: Add logger at debug level
			// TODO: Add "tips" option
			// TODO: Get URL from outside
			fmt.Println("➡ Please report here: https://github.com/speedflow/speedflow/issues/new?labels=bug")
			os.Exit(1)
		}
	}()

	_, err := ExecuteC(os.Stdin, os.Stdout, os.Stderr)
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
		Run:     cmdRoot.Run,
		PreRun:  prerun,
	}

	cmd.SetIn(in)
	cmd.SetOut(out)
	cmd.SetErr(errIO)
	cmd.SetArgs(args)

	// Add flags
	flags(cmd)

	// Add subcommands
	cmd.AddCommand(cmdVersion.New())
	cmd.AddCommand(cmdFlows.New())

	return cmd, cmd.Execute()
}

func flags(cmd *cobra.Command) {
	// Current directory
	currentDir, err := os.Getwd()
	cobra.CheckErr(err)

	cmd.Flags().StringVarP(&spPathFile, "file", "f", filepath.Join(currentDir, sfFile), "Speedflow file")
	cmd.Flags().StringVarP(&cfgPathFile, "config", "c", filepath.Join(cfgPath(), cfgFile), "Configuration file")
	cmd.PersistentFlags().StringP("output", "o", "", "One of '', 'yaml' or 'json'.")
}

func prerun(cmd *cobra.Command, args []string) {
	err := speedflow.Load(spPathFile)

	var errFile *speedflow.ErrFile
	if errors.As(err, &errFile) {
		// TODO: Improve error message color
		cmd.PrintErrln("Unable to open the Speedflow file")
		// TODO: Add "tips" option
		cmd.PrintErrln("➡ Create an new Speedflow file? speedflow init")
		exit(1)
	}

	var errParse *speedflow.ErrParse
	if errors.As(err, &errParse) {
		// TODO: Improve error message with color
		// TODO: Add some parsing info error
		cmd.PrintErrln("Unable to read and parse the Speedflow file")
		// TODO: Add "tips" option
		cmd.PrintErrln("➡ Call the doctor? speedflow doctor")
		cmd.PrintErrf("%v\n", err)
		exit(1)
	}
}

func cfgPath() string {
	// Home directory
	homeDir, err := os.UserHomeDir()
	cobra.CheckErr(err)

	return filepath.Join(homeDir, cfgSubPath)
}

func exit(c int) {
	if Exitable {
		os.Exit(c)
	}
}
