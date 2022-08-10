package wireguard

import (
	"fmt"
	"time"

	"github.com/sun-asterisk-research/wireguard-peer-manager/types"
	"golang.zx2c4.com/wireguard/wgctrl"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

// CreateOrUpdatePeer create a new peer or update an existing peer for the specified device
func CreateOrUpdatePeer(deviceName string, peer types.Peer) error {
	client, err := wgctrl.New()
	if err != nil {
		return err
	}

	defer client.Close()

	pubkey, err := wgtypes.ParseKey(peer.PublicKey)
	if err != nil {
		return err
	}

	allowedIPs, err := parseAllowedIPs(peer.AllowedIps)
	if err != nil {
		return err
	}

	var psk *wgtypes.Key
	if peer.PresharedKey != "" {
		if key, err := wgtypes.ParseKey(peer.PresharedKey); err == nil {
			psk = &key
		} else {
			return err
		}
	}

	var persistentKeepaliveInterval *time.Duration
	if peer.PersistentKeepalive != 0 {
		if duration, err := time.ParseDuration(fmt.Sprintf("%ds", peer.PersistentKeepalive)); err == nil {
			persistentKeepaliveInterval = &duration
		} else {
			return err
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

	return client.ConfigureDevice(deviceName, conf)
}
