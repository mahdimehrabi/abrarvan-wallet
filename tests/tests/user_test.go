package tests

import (
	"challange/app/controller"
	"challange/app/infrastracture"
	"challange/app/repository"
	"challange/app/services"
	"challange/tests/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateUser(t *testing.T) {
	logger := infrastracture.ArvanLogger{}
	db := mocks.NewDB()
	memoryDB := mocks.NewMemoryDB()
	walletRepository := repository.WalletRepository{
		Logger:   &logger,
		DB:       db,
		MemoryDB: memoryDB,
	}
	walletService := services.WalletService{
		WalletRepository: &walletRepository,
	}
	walletController := controller.WalletController{
		Logger:        &logger,
		WalletService: walletService,
	}

	data := map[string]interface{}{
		"mobile":         "09120401761",
		"credit":         4000,
		"receivedCharge": true,
	}
	req, _ := http.NewRequest("POST",
		"/use-code",
		infrastracture.MapToJsonBytesBuffer(data))
	w := httptest.NewRecorder()
	walletController.UseCode(w, req)
	body := w.Body.String()
	if !strings.Contains(body, "Created") || w.Code != http.StatusOK {
		t.Error("user not created")
	}
}
