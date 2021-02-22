package main

import (
	"fmt"
	"net"
)

//获取本机的所有网卡
//isFiler 是否过滤没有分配IP的网卡
func GetAllIfaceNames(isFilter bool) ([]string, error) {
	ifaceNames := []string{}
	ifaces, err := net.Interfaces()
	if err != nil {
		return ifaceNames, err
	}
	for _, iface := range ifaces {
		if isFilter {
			IP, err := getIPv4ByIface(&iface)
			if err != nil || fmt.Sprintf("%v", IP) == "127.0.0.1" {
				continue
			}
		}
		ifaceNames = append(ifaceNames, iface.Name)
	}

	return ifaceNames, nil
}

func IsContain(hosts []HostItem, host HostItem) bool {
	for _, eachHost := range hosts {
		if eachHost.IP.String() == host.IP.String() && eachHost.MAC.String() == host.MAC.String() {
			return true
		}
	}
	return false
}
