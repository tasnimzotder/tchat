package utils

import (
	"encoding/json"
	"errors"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter, errStr string, err error) {
	var errMsg error

	if err != nil {
		errMsg = err
	} else {
		errMsg = errors.New(errStr)
	}

	_ = json.NewEncoder(w).Encode(struct {
		Error string `json:"error"`
	}{
		Error: errMsg.Error(),
	})
}
