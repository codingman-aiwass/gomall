package convert

import (
	"fmt"
	"net"
	"strings"
)

// 假设 c.RocketMQ.NameServers 为 []string{"namesrv:9876"}
func ResolveNameServer(nameServer string) (string, error) {
	// 分离主机名和端口（这里简单假设格式正确）
	var host string
	var port string
	arr := strings.Split(nameServer, ":")
	host, port = arr[0], arr[1]

	ips, err := net.LookupIP(host)
	if err != nil {
		return "", err
	}
	if len(ips) == 0 {
		return "", fmt.Errorf("no IP found for host %s", host)
	}
	return fmt.Sprintf("%s:%s", ips[0].String(), port), nil
}
