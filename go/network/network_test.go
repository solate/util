package network

import "testing"

func TestGetLocalIp(t *testing.T) {
	localIp, err := GetLocalIp()
	if err != nil {
		t.Error(err)
	}
	t.Log(localIp)
}
