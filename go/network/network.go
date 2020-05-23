package network

import (
	"errors"
	"net"
)

// 获取本机IP地址
func GetLocalIp() (ip string, err error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return
	}

	for _, item := range netInterfaces {
		if (item.Flags & net.FlagUp) != 0 {
			addrs, err := item.Addrs()
			if err != nil {
				return "", err
			}

			for _, addr := range addrs {
				if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
					if ipNet.IP.To4() != nil {
						ip = ipNet.IP.String()
						return ip, nil
					}
				}
			}
		}
	}
	err = errors.New("connected to the network?")
	return

}
