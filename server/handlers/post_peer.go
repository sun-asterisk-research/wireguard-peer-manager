package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/sun-asterisk-research/wireguard-peer-manager/config"
	"github.com/sun-asterisk-research/wireguard-peer-manager/types"
	"github.com/sun-asterisk-research/wireguard-peer-manager/wireguard"
)

type postPeerResponse struct {
	AllowedIPs []string `json:"allowed_ips"`
}

func PostPeer(w http.ResponseWriter, r *http.Request) {
	var peer types.Peer

	if err := json.NewDecoder(r.Body).Decode(&peer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}

	configuredPeer, err := wireguard.CreateOrUpdatePeer(config.Values.Device, peer)
	if err != nil {
		logrus.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response := postPeerResponse{
		AllowedIPs: configuredPeer.AllowedIps,
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		logrus.Error(err)
	}
}
