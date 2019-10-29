package main

import (
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/klog"
	"log"
	"os"
	"path"
)

func genFunc(ccmd *cobra.Command, args []string) {
	if err := os.MkdirAll(varConfigPath, 0660); err != nil {
		klog.Warningf("unable to create %s due to error: %v, trying next directory...", varConfigPath, err)
	} else {
		if err := viper.WriteConfigAs(path.Join(varConfigPath, "config.yaml")); err != nil {
			klog.Fatalf("unable to write %s due to error: %v, failing...", path.Join(varConfigPath, "config.yaml"), err)
		}
		return
	}
	if err := os.MkdirAll(etcConfigPath, 0660); err != nil {
		klog.Warningf("unable to create %s due to error: %v, trying next directory...", etcConfigPath, err)
	} else {
		if err := viper.WriteConfigAs(path.Join(etcConfigPath, "config.yaml")); err != nil {
			klog.Fatalf("unable to write %s due to error: %v, failing...", path.Join(etcConfigPath, "config.yaml"), err)
		}
		return
	}
	d, err := homedir.Dir()
	if err != nil {
		klog.Fatalf("unable to determine the users home directory, error: %v", err)
	}
	if err := os.MkdirAll(path.Join(d, "kilo"), 0660); err != nil {
		log.Fatalf("unable to create %s due to error: %v, failing...", path.Join(d, "kilo"), err)
	} else {
		if err := viper.WriteConfigAs(path.Join(d, "kilo", "config.yaml")); err != nil {
			log.Fatalf("unable to write %s due to error: %v, failing...", path.Join(d, "kilo", "config.yaml"), err)
		}
		return
	}
}