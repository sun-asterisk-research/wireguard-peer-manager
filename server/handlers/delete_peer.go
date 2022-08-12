package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/sun-asterisk-research/wireguard-peer-manager/config"
	"github.com/sun-asterisk-research/wireguard-peer-manager/types"
	"github.com/sun-asterisk-research/wireguard-peer-manager/wireguard"
)

func DeletePeer(w http.ResponseWriter, r *http.Request) {
	var peer types.Peer

	if err := json.NewDecoder(r.Body).Decode(&peer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}

	if err := wireguard.RemovePeer(config.Values.Device, peer.PublicKey); err != nil {
		logrus.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)
}
