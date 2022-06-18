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
	memoryDB.MockGetFn = func(s string) (string, error) {
		return `{"code":"hello","credit":50000,"consumerCount":1000}`, nil
	}
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
		"mobile": "09120401761",
		"code":   "hello",
	}
	req, _ := http.NewRequest("POST",
		"/use-code",
		infrastracture.MapToJsonBytesBuffer(data))
	w := httptest.NewRecorder()
	walletController.UseCode(w, req)
	body := w.Body.String()
	if !strings.Contains(body, "Congratulation") || w.Code != http.StatusOK {
		t.Error("using charge code failed")
	}
}
