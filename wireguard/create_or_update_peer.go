package wireguard

import (
	"fmt"
	"net"
	"time"

	"github.com/sun-asterisk-research/wireguard-peer-manager/config"
	"github.com/sun-asterisk-research/wireguard-peer-manager/types"
	"golang.zx2c4.com/wireguard/wgctrl"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

// CreateOrUpdatePeer create a new peer or update an existing peer for the specified device
func CreateOrUpdatePeer(deviceName string, peer types.Peer) (types.Peer, error) {
	client, err := wgctrl.New()
	if err != nil {
		return peer, err
	}

	defer client.Close()

	pubkey, err := wgtypes.ParseKey(peer.PublicKey)
	if err != nil {
		return peer, err
	}

	var allowedIPs []net.IPNet

	// Assign an IP if non were specified
	if len(peer.AllowedIps) != 0 {
		allowedIPs, err = parseAllowedIPs(peer.AllowedIps)
		if err != nil {
			return peer, err
		}
	} else {
		ip, err := getFreeIP(client, deviceName, config.Values.PeerCIDRs)
		if err != nil {
			return peer, err
		}

		var mask net.IPMask
		if ip.To4() != nil {
			mask = net.CIDRMask(32, 8*net.IPv4len)
		} else {
			mask = net.CIDRMask(64, 8*net.IPv6len)
		}

		allowedIPs = []net.IPNet{
			{
				IP:   ip,
				Mask: mask,
			},
		}

		for _, allowedIP := range allowedIPs {
			peer.AllowedIps = append(peer.AllowedIps, allowedIP.String())
		}
	}

	var psk *wgtypes.Key
	if peer.PresharedKey != "" {
		if key, err := wgtypes.ParseKey(peer.PresharedKey); err == nil {
			psk = &key
		} else {
			return peer, err
		}
	}

	var persistentKeepaliveInterval *time.Duration
	if peer.PersistentKeepalive != 0 {
		if duration, err := time.ParseDuration(fmt.Sprintf("%ds", peer.PersistentKeepalive)); err == nil {
			persistentKeepaliveInterval = &duration
		} else {
			return peer, err
		}
	}

	peerConf := wgtypes.PeerConfig{
		PublicKey:                   pubkey,
		PresharedKey:                psk,
		AllowedIPs:                  allowedIPs,
		ReplaceAllowedIPs:           true,
		PersistentKeepaliveInterval: persistentKeepaliveInterval,
	}

	conf := wgtypes.Config{
		Peers: []wgtypes.PeerConfig{
			peerConf,
		},
	}

	return peer, client.ConfigureDevice(deviceName, conf)
}
