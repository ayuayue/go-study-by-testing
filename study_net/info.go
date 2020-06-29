package main

import (
	"errors"
	"fmt"
	"net"
	"os"
)

func main() {
	//获取 计算机接口
	// interfaces, err := net.Interfaces()
	// if err != nil {
	//     os.Exit(1)
	// }
	//循环所有接口
	// for k,v := range interfaces {
	//     fmt.Printf("========%d=========\n",k)
	//硬件信息
	// fmt.Println(k,v)
	//MAC 地址
	// fmt.Println(k,v.HardwareAddr)

	// }
	//获取 ip 地址
	//ip
	ips, err := net.InterfaceAddrs()
	if err != nil {
		os.Exit(1)
	}
	for k, ip := range ips {
		fmt.Printf("========%d=========\n", k)
		if ipnet, ok := ip.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(net.ParseIP(ipnet.IP.String()))
				fmt.Println(IPString2Long(ipnet.IP.String()))
			}
		}
	}
}
func IPString2Long(ip string) (uint, error) {
	b := net.ParseIP(ip).To4()
	if b == nil {
		return 0, errors.New("invalid ipv4 format")
	}

	return uint(b[3]) | uint(b[2])<<8 | uint(b[1])<<16 | uint(b[0])<<24, nil
}
