package controllers

import (
	//"log"
	"encoding/json"
	"net/http"
)

type DecodeRequest struct {
	String    string `json:"string"`
	Algorithm string `json:"algorithm"`
	ApiKey    string `json:"api_key"`
}

type DecodeResponse struct {
	String string `json:"string"`
	Error  string `json:"error"` // omitempty?
}

func POSTDecode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	encoder := json.NewEncoder(w)

	defer r.Body.Close()
	var req DecodeRequest
	var res DecodeResponse

	if err := decoder.Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		res.Error = err.Error()
		encoder.Encode(res)

		return
	}

	if req.String == "" {
		http.Error(w, "Empty string", http.StatusInternalServerError)
		res.Error = "Empty string"
		encoder.Encode(res)

		return
	}

	res.String = req.String
	encoder.Encode(res)
}
