package wireguard

import "net"

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
