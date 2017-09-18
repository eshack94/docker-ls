package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mayflower/docker-ls/cli/util"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "docker-ls",
	Short: "browse docker registries",
	Long:  "Browse and examine repositories and tags in a docker registry",
}

var libraryFlags = &util.LibraryFlags{}

func init() {
	flags := rootCmd.PersistentFlags()
	libraryFlags.BindToFlags(flags)

	var configFile string
	flags.StringVarP(&configFile, "config", "c", "",
		"read config from specified file (default: look for config in home directory)",
	)

	cobra.OnInitialize(func() {
		if configFile != "" {
			viper.SetConfigFile(configFile)
		} else {
			home, err := homedir.Dir()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			viper.AddConfigPath(home)
			viper.SetConfigName(".docker-ls")
		}

		viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
		viper.SetEnvPrefix("DOCKER_LS")
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err == nil {
			fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		}
	})
}