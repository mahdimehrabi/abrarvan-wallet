package infrastracture

import (
	"bytes"
	"challange/app/interfaces"
	"encoding/json"
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

func BytesJsonToMap(bytes []byte) (map[string]interface{}, error) {
	mp := make(map[string]interface{})
	err := json.Unmarshal(bytes, &mp)
	return mp, err
}

func MapToJsonBytesBuffer(mp map[string]interface{}) *bytes.Buffer {
	j, err := json.Marshal(mp)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(j)
}
