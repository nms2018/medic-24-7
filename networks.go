package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type IpSettings struct {
	IP4 string `json:"ip_4"`
	IP6 string `json:"ip_6"`
}

func externalIP() []byte {
	ipsetting := IpSettings{}

	ifaces, err := net.Interfaces()
	if err != nil {
		return nil
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return nil
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip4 := ip.To4()
			ip6 := ip.To16()

			if ip4 == nil {
				continue // not an ipv4 address
			}

			ipsetting.IP4 = ip4.String()
			ipsetting.IP6 = ip6.String()

			jsonResponse, jsonError := json.Marshal(ipsetting)

			if jsonError != nil {
				fmt.Println(jsonError)
				return nil
			}
			return jsonResponse
		}
	}
	fmt.Println("not connected")
	return nil
}
