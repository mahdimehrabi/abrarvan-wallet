package routes

import (
	"fmt"
	"net/http"
)

type WalletRoutes struct {
}

func NewWalletRoutes() WalletRoutes {
	return WalletRoutes{}
}

func (r WalletRoutes) AddRoutes(sm *http.ServeMux) {
	sm.HandleFunc("/wallet/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world")
		w.WriteHeader(200)
	})
}
