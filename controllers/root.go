package controllers

import (
	//"log"
	"net/http"
	"encoding/json"
)

type DecodeRequest struct {
	String string `json:"string"`
	Algorithm string `json:"algorithm"`
	ApiKey string `json:"api_key"`
}

type DecodeResponse struct {
	String string `json:"string"`
	Error string `json:"error"`  // omitempty?
}

func POSTDecode (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var req DecodeRequest
	var resp DecodeResponse
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		resp.Error = err.Error() 

		json.NewEncoder(w).Encode(resp)
		return
	}
	

	resp.String = req.String
	json.NewEncoder(w).Encode(resp)
}