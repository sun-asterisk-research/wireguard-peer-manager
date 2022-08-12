package config

import "github.com/spf13/viper"

func ReadEnvs() {
	viper.BindEnv("Host", "WGPM_HOST")
	viper.BindEnv("Port", "WGPM_PORT")
	viper.BindEnv("Device", "WGPM_DEVICE")
	viper.BindEnv("PeerCIDRs", "WGPM_PEER_CIDRS")
	viper.BindEnv("BearerTokenAuth", "WGPM_BEARER_TOKEN_AUTH")
}
