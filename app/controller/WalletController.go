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
	Logger        interfaces.Logger
	WalletService services.WalletService
}

func NewWalletController(
	logger infrastracture.ArvanLogger,
	walletService services.WalletService,
) WalletController {
	return WalletController{
		Logger:        &logger,
		WalletService: walletService,
	}
}

func (c *WalletController) UseCode(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "not allowed", http.StatusMethodNotAllowed)
		return
	}
	reqData := new(models.UseCodeReq)
	err := reqData.FromJSON(r.Body)
	if err != nil {
		c.Logger.Error(err.Error())
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	//try to create user with mobile number
	c.WalletService.CreateUser(&models.User{
		Mobile: reqData.Mobile,
	})

	err = c.WalletService.TryDecreaseConsumerCount(reqData.Code)
	if err != nil {
		c.Logger.Error(err.Error())
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, "Created reqData")
	w.WriteHeader(200)
	return
}
