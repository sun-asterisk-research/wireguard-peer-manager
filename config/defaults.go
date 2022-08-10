package config

import "github.com/spf13/viper"

func setDefaults() {
	viper.SetDefault("Host", "0.0.0.0")
	viper.SetDefault("Port", 9000)
	viper.SetDefault("Device", "wg0")
}
