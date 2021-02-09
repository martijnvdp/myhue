/*
Copyright © 2021 M van der Ploeg

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

type Hueconfig struct {
	USER  string `mapstructure:"USER"`
	TOKEN string `mapstructure:"TOKEN"`
	IP    string `mapstructure:"IP"`
}

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "myhue",
	Short: "a command line tool to interact with the philips hue bridge api",
	Long: `A command line tool to interact with the philips hue bridge api. For example:

before first run push the button on the hue bridge before run:
set the following env vars:

"MYHUE_USER":"someusername"
"MYHUE_IP":"1.1.1.1" optional if not set it will searche for the bridge
"MYHUE_TOKEN":"token after first user creation"

list all lights: myhue list.
turn on light 1 with a brightness level of 55%:
myhue light 1 on 55.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.myhue.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	var hueconfig Hueconfig
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
		configName := "config.myhue"
		configType := "yml"
		viper.AddConfigPath(home)
		viper.SetConfigName(configName)
		viper.SetConfigType(configType)
		cfgFile = filepath.Join(home, configName+"."+configType)
	}
	viper.SetEnvPrefix("MYHUE_") //not working
	viper.AllowEmptyEnv(true)
	viper.AutomaticEnv()
	viper.Unmarshal(&hueconfig)
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		_, err := os.Stat(cfgFile)
		if !os.IsExist(err) {
			if _, err := os.Create(cfgFile); err != nil {
			}
		}
		if err := viper.SafeWriteConfig(); err != nil {
		}
	}
}
