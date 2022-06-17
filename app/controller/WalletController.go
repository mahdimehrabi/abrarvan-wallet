package controller

import (
	"challange/app/infrastracture"
	"challange/app/interfaces"
	"challange/app/models"
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
	if r.Method != http.MethodPost {
		http.Error(w, "not allowed", http.StatusMethodNotAllowed)
		return
	}
	user := new(models.User)
	err := user.FromJSON(r.Body)
	if err != nil {
		c.logger.Error(err.Error())
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	err = c.walletService.CreateUser(user)
	if err != nil {
		c.logger.Error(err.Error())
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, "Created user")
	w.WriteHeader(200)
	return
}
