package wireguard

import (
	"fmt"
	"net"
	"net/netip"

	"golang.zx2c4.com/wireguard/wgctrl"
)

func parseAllowedIPs(ips []string) ([]net.IPNet, error) {
	var ipnets []net.IPNet

	for _, cidr := range ips {
		_, parsed, err := net.ParseCIDR(cidr)
		if err != nil {
			return ipnets, err
		}

		ipnets = append(ipnets, *parsed)
	}

	return ipnets, nil
}

func getFreeIP(client *wgctrl.Client, deviceName string, cidrs []string) (net.IP, error) {
	device, err := client.Device(deviceName)
	if err != nil {
		return nil, err
	}

	var usedSubnets []net.IPNet
	for _, peer := range device.Peers {
		usedSubnets = append(usedSubnets, peer.AllowedIPs...)
	}

	for _, cidr := range cidrs {
		ip, err := getFreeIPFromCIDR(cidr, usedSubnets)
		if err != nil {
			return nil, err
		}

		if ip != nil {
			return ip, nil
		}
	}

	return nil, fmt.Errorf("no available IP in specified CIDRs")
}

func getFreeIPFromCIDR(cidr string, used []net.IPNet) (net.IP, error) {
	prefix, err := netip.ParsePrefix(cidr)
	if err != nil {
		return nil, err
	}

	addr := prefix.Addr()
	var ip net.IP
	for {
		addr = addr.Next()
		ip = net.ParseIP(addr.String())

		if !checkUsedIP(ip, used) {
			break
		}

		if !prefix.Contains(addr) {
			return nil, nil
		}
	}

	return ip, nil
}

func checkUsedIP(ip net.IP, used []net.IPNet) bool {
	for _, subnet := range used {
		if subnet.Contains(ip) {
			return true
		}
	}

	return false
}
