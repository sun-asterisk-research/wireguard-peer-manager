package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/sun-asterisk-research/wireguard-peer-manager/config"
	"github.com/sun-asterisk-research/wireguard-peer-manager/server"
)

var cmd = &cobra.Command{
	Use:                   "wireguard-pm [OPTIONS] COMMAND [ARG...]",
	Short:                 "Wireguard peer manager",
	TraverseChildren:      true,
	DisableFlagsInUseLine: true,
	SilenceErrors:         true,
	PersistentPreRun:      preRun,
	Run: run,
}

func preRun(cmd *cobra.Command, _ []string) {
	if err := config.ResolveValues(); err != nil {
		logrus.Fatal("Cannot resolve config: ", err)
	}
}

func run(cmd *cobra.Command, _ []string) {
	if err := server.Start(*config.Values); err != nil {
		logrus.Fatal(err)
	}
}

func main() {
	flags := cmd.PersistentFlags()

	flags.StringP("host", "H", "0.0.0.0", "Host to listen on")
	flags.IntP("port", "p", 9000, "Port to listen on")
	flags.StringP("device", "d", "wg0", "Device to manage peers for")
	flags.StringP("bearer-token-auth", "a", "", "Expected bearer token for auth")

	config.ReadFlags(flags)

	cmd.Execute()
}
