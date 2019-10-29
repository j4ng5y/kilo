package main

import (
	"fmt"
	"github.com/j4ng5y/kilo/pkg/state"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/klog"
	"path"
)

func initConfig() error {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		d, err := homedir.Dir()
		if err != nil {
			return fmt.Errorf("error reading the home directory, error: %w", err)
		}

		viper.SetConfigName("config")
		viper.AddConfigPath(varConfigPath)
		viper.AddConfigPath(etcConfigPath)
		viper.AddConfigPath(path.Join(d, "kilo"))

	}

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading in the configuration file, error: %w", err)
	}
	return nil
}

func initFunc(ccmd *cobra.Command, args []string) {
	if err := initConfig(); err != nil {
		klog.Fatal(err.Error())
	}
	go state.InitState(viper.GetInt("kilo.state.etcd.startup_timeout"))
}