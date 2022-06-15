package infrastracture

import (
	"challange/app/interfaces"
	"fmt"
	"net/http"
)

func BadRequestResponse(w http.ResponseWriter) {
	fmt.Fprint(w, "bad request")
	w.WriteHeader(http.StatusBadRequest)
}

func ErrorResponse(err error, logger interfaces.Logger, w http.ResponseWriter) {
	logger.Error(err.Error())
	fmt.Fprint(w, "bad request")
	w.WriteHeader(http.StatusBadRequest)
}

func SuccessResponse(w http.ResponseWriter, data string) {
	fmt.Fprint(w, data)
	w.WriteHeader(http.StatusOK)
}
