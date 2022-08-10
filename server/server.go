package server

import (
	"crypto/sha512"
	"crypto/subtle"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"github.com/sun-asterisk-research/wireguard-peer-manager/config"
	"github.com/sun-asterisk-research/wireguard-peer-manager/server/handlers"
)

func BearerTokenAuth(h http.HandlerFunc, token string) http.HandlerFunc {
	if token == "" {
		return h
	}

	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		expect := "Bearer " + token

		authSha := sha512.Sum512([]byte(auth))
		expectSha := sha512.Sum512([]byte(expect))

		if subtle.ConstantTimeCompare(authSha[:], expectSha[:]) == 1 {
			h(w, r)
		} else {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}

func Start(cfg config.Config) error {
	router := httprouter.New()

	router.HandlerFunc("POST", "/peers", BearerTokenAuth(handlers.PostPeer, cfg.BearerTokenAuth))
	router.HandlerFunc("POST", "/peers/del", BearerTokenAuth(handlers.DeletePeer, cfg.BearerTokenAuth))

	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	logrus.Infof("Listening on %s", addr)
	logrus.Infof("Managing device %s", cfg.Device)

	return http.ListenAndServe(addr, router)
}
