package command

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

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
	}
	cmd.Run = run

	cmd.SetIn(in)
	cmd.SetOut(out)
	cmd.SetErr(errIO)
	cmd.SetArgs(args)

	// Add subcommands
	cmd.AddCommand(cmdVersion.New())

	// Add flags
	flags(cmd)

	return cmd, cmd.Execute()
}

func flags(cmd *cobra.Command) {
	// Current directory
	currentDir, err := os.Getwd()
	cobra.CheckErr(err)

	cmd.Flags().BoolP("list", "l", false, "List flows")
	cmd.Flags().StringVarP(&spPathFile, "file", "f", filepath.Join(currentDir, sfFile), "Speedflow file")
	cmd.Flags().StringVarP(&cfgPathFile, "config", "c", filepath.Join(cfgPath(), cfgFile), "Configuration file")
}

func run(cmd *cobra.Command, args []string) {
	// Load speedflow file
	if err := speedflow.Load(spPathFile); err != nil {
		cmd.PrintErrf("Error: %s\n", err)
		exit(1)
		return
	}

	// Display list
	if showList, _ := cmd.Flags().GetBool("list"); showList {
		cmd.Println("Flow     Name        ")
		cmd.Println("default  Default flow")
		return
	}

	// Execute command
	// TODO: implement
	cmd.Println("Hello World!")
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
