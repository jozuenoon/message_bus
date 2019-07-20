package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/inconshreveable/log15"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var collectorHost string
var queryHost string

var logger log15.Logger

var signals chan os.Signal

var rootCmd = &cobra.Command{
	Use:   "mbcli",
	Short: "Message bus client, travel time discovery system.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mbcli.yaml)")
	rootCmd.PersistentFlags().StringVar(&collectorHost, "collector_host", "127.0.0.1:9000",
		"configure collector host")
	rootCmd.PersistentFlags().StringVar(&queryHost, "query_host", "127.0.0.1:8000",
		"configure query host")

	signals = make(chan os.Signal, 2)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	logger = log15.New()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".mbcli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".mbcli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
