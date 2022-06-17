package controller

import (
	"challange/app/infrastracture"
	"challange/app/interfaces"
	"challange/app/services"
	"fmt"
	"net/http"
)

type WalletController struct {
	logger        interfaces.Logger
	walletService services.WalletService
}

func NewWalletController(
	logger infrastracture.ArvanLogger,
	walletService services.WalletService,
) WalletController {
	return WalletController{
		logger:        &logger,
		walletService: walletService,
	}
}

func (c *WalletController) UseCode(w http.ResponseWriter, r *http.Request) {
	var jsonBytes []byte
	_, err := r.Body.Read(jsonBytes)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}

	err = c.walletService.CreateUser(jsonBytes)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}
	fmt.Fprint(w, "Hello World!")
	w.WriteHeader(200)
}
