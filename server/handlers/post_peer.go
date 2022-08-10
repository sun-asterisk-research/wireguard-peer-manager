package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/sun-asterisk-research/wireguard-peer-manager/config"
	"github.com/sun-asterisk-research/wireguard-peer-manager/types"
	"github.com/sun-asterisk-research/wireguard-peer-manager/wireguard"
)

func PostPeer(w http.ResponseWriter, r *http.Request) {
	var peer types.Peer

	if err := json.NewDecoder(r.Body).Decode(&peer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}

	if err := wireguard.CreateOrUpdatePeer(config.Values.Device, peer); err != nil {
		logrus.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
