package types

// Peer is the configuration for a wireguard peer
type Peer struct {
	// Base64 encoded public key
	PublicKey string `json:"public_key"`
	// Base64 encoded preshared key
	PresharedKey string `json:"preshared_key,omitempty"`
	// Peer's allowed ips, it might be any of IPv4 or IPv6 addresses in CIDR notation
	AllowedIps []string `json:"allowed_ips"`
	// Peer's persisten keepalive interval in seconds
	PersistentKeepalive int `json:"persistent_keepalive"`
	// Peer's endpoint in host:port format
	Endpoint string `json:"endpoint"`
}
