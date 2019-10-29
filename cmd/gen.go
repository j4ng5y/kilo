package cmd

import (
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"k8s.io/klog"
	"log"
	"os"
	"path"
)

var generateCmd = &cobra.Command{
	Use:     "gen",
	Short:   "generate a configuration file",
	Long:    "",
	Example: "",
	Run:     genFunc,
}

func init() {
	kiloCmd.AddCommand(generateCmd)
}

func genFunc(ccmd *cobra.Command, args []string) {
	if err := os.MkdirAll(varConfigPath, 0660); err != nil {
		klog.Warningf("unable to create %s due to error: %v, trying next directory...", varConfigPath, err)
	} else {
		if err := writeDefaultConfig(path.Join(varConfigPath, "config.yaml")); err != nil {
			klog.Fatalf("unable to write %s due to error: %v, failing...", path.Join(varConfigPath, "config.yaml"), err)
		}
		return
	}
	if err := os.MkdirAll(etcConfigPath, 0660); err != nil {
		klog.Warningf("unable to create %s due to error: %v, trying next directory...", etcConfigPath, err)
	} else {
		if err := writeDefaultConfig(path.Join(etcConfigPath, "config.yaml")); err != nil {
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
		if err := writeDefaultConfig(path.Join(d, "kilo", "config.yaml")); err != nil {
			log.Fatalf("unable to write %s due to error: %v, failing...", path.Join(d, "kilo", "config.yaml"), err)
		}
		return
	}
}

func writeDefaultConfig(fileName string) error {
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0660)
	if err != nil {
		return err
	}

	_, err = f.Write([]byte(defaultConfig))
	if err != nil {
		return err
	}

	return nil
}

const defaultConfig = `---
kilo:
  state:
    spec:
      etcd:
        startup_timeout: 60
  local_cluster:
    spec:
      install:
        type: single-master
        skip_download: false
        symlink: force
        version: latest
        bin_dir: /usr/local/bin
        systemd_dir: /etc/systemd/system
        name: kilo-k3s
      server:
        bind_address: localhost
        https_listen_port: 6443
        http_listen_port: 0
        data_dir: $HOME/.kilo/k3s
        log: true
        k8s_cluster:
          cidr: 10.42.0.0/16
          secret: 5u93r53cr37k1l053cr3t15453cr3t
          service_cidr: 10.43.0.0/16
          dns: 10.43.255.254
          domain: kilo.k3s.local
          disable_default_scheduler: false
          disable_default_cloud_controller_manager: false
          disable_default_network_policy_controller: false
          disable_flannel: false
          additional_apiserver_args:
            - NA
          additional_controller_args:
            - NA
          additional_scheduler_args:
            - NA
          additional_kubelet_args:
            - NA
          additional_kube_proxy_args:
            - NA
        storage:
          endpoint: sqlite
    contexts:
      - name: jgregory-cluster.home
  remote_clusters:
    contexts:
      - name: jgregory-cluster.us.dev.lite.granular.ag
`
