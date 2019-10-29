package state

import (
	"go.etcd.io/etcd/embed"
	"k8s.io/klog"
	"time"
)

func InitState(timeout int) {
	cfg := embed.NewConfig()
	cfg.Dir = "default.etcd"
	e, err := embed.StartEtcd(cfg)
	if err != nil {
		klog.Fatal(err)
	}
	defer e.Close()

	select {
	case <-e.Server.ReadyNotify():
		klog.V(1).Info("etcd started")
	case <-time.After(time.Duration(timeout)*time.Second):
		e.Server.Stop()
		klog.Errorf("etcd did not start within %d seconds", timeout)
	}
	klog.Fatal(<-e.Err())
}
