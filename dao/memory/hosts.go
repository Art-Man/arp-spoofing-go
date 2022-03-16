package memory

import (
	"github.com/ShangRui-hash/arp-spoofing-go/models"
)

var (
	hosts = make(map[string]*models.Host, 0)
)

func AddHost(host *models.Host) {
	hosts[host.IP] = host
}

func GetHost(ip string) *models.Host {
	return hosts[ip]
}

func GetAllIP() []string {
	//获取hosts的所有key
	ips := make([]string, 0)
	for ip := range hosts {
		ips = append(ips, ip)
	}
	return ips
}

func GetAllHosts() []*models.Host {
	hostList := make([]*models.Host, 0)
	for _, host := range hosts {
		hostList = append(hostList, host)
	}
	return hostList
}

func ClearHosts() {
	hosts = make(map[string]*models.Host, 0)
	return
}
