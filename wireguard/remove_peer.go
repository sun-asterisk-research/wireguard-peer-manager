package wireguard

import (
	"golang.zx2c4.com/wireguard/wgctrl"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

// RemovePeer remove peer with matching public key from the specified device
func RemovePeer(deviceName, pubkey string) error {
	client, err := wgctrl.New()
	if err != nil {
		return err
	}

	defer client.Close()

	wgKey, err := wgtypes.ParseKey(pubkey)
	if err != nil {
		return err
	}

	return client.ConfigureDevice(deviceName, wgtypes.Config{
		Peers: []wgtypes.PeerConfig{
			{
				PublicKey: wgKey,
				Remove:    true,
			},
		},
	})
}
