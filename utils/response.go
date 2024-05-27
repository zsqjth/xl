package utils

import (
	"encoding/json"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, code int, msg string, data interface{}) {
	response := map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
