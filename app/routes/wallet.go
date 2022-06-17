package routes

import (
	"challange/app/controller"
	"net/http"
)

type WalletRoutes struct {
	walletController controller.WalletController
}

func NewWalletRoutes(walletController controller.WalletController) WalletRoutes {
	return WalletRoutes{
		walletController: walletController,
	}
}

func (r WalletRoutes) AddRoutes(sm *http.ServeMux) {
	sm.HandleFunc("/use-code", r.walletController.UseCode)
}
