package config

import "github.com/spf13/viper"

type Config struct {
	Host            string
	Port            int
	Device          string
	BearerTokenAuth string
}

func init() {
	setDefaults()
}

// Values is the resolved runner config
var Values *Config

// ResolveValues resolves the runner config & populate `Values` with the resolved config
func ResolveValues() error {
	return viper.Unmarshal(&Values)
}
