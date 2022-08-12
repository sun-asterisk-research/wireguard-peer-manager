package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// ReadFlags parses command flag for configs
// Flags override configs provided by environment variables & config file
func ReadFlags(flags *pflag.FlagSet) {
	viper.BindPFlag("Host", flags.Lookup("host"))
	viper.BindPFlag("Port", flags.Lookup("port"))
	viper.BindPFlag("Device", flags.Lookup("device"))
	viper.BindPFlag("PeerCIDRs", flags.Lookup("peer-cidr"))
	viper.BindPFlag("BearerTokenAuth", flags.Lookup("bearer-token-auth"))
}
