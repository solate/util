package etcd

import (
	"testing"
)

func TestEtcdInit(t *testing.T) {
	etcd := NewEtcd([]string{"127.0.0.1:2379"}, 1000, "", "")
	err := etcd.Init()
	if err != nil {
		t.Errorf("conn etcd err: %v", err)
	}
}
