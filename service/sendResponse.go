package service

import (
	"encoding/json"
	"net/http"
)

func SendResponse(r http.ResponseWriter, pl interface{}) {
	js, err := json.Marshal(pl)
	if err != nil {
		http.Error(r, err.Error(), http.StatusInternalServerError)
	}
	r.Header().Set("Content-Type", "application/json")
	_, err = r.Write(js)
	if err != nil {
		panic(err.Error())
	}
}
