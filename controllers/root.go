package controllers

import (
	"encoding/json"
	"github.com/shootnix/golatin-service/models"
	"log"
	"net/http"
)

type DecodeRequest struct {
	String string `json:"string"`
	Table  string `json:"algorithm"`
}

type DecodeResponse struct {
	String string `json:"string"`
	Error  string `json:"error"` // omitempty?
}

func POSTTranslit(w http.ResponseWriter, r *http.Request) {
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

	var err_msg string
	if req.String == "" {
		err_msg = "Empty string"
		http.Error(w, err_msg, http.StatusInternalServerError)
		res.Error = err_msg
		encoder.Encode(res)

		return
	}

	if req.Table == "" {
		err_msg = "Empty table"
		http.Error(w, err_msg, http.StatusInternalServerError)
		res.Error = err_msg
		encoder.Encode(res)

		return
	}

	/*
		apiUser, err := models.FindApiUser(req.ApiKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			res.Error = err.Error()
			encoder.Encode(res)

			return
		}

		log.Println("User with api_key `" + apiUser.ApiKey + "` found")
		if apiUser.CheckDisabled() {
			err_msg := "User is disabled"
			http.Error(w, err_msg, http.StatusInternalServerError)
			res.Error = err_msg
			encoder.Encode(res)

			return
		}

		if apiUser.CheckAllowedRequests() == 0 {
			err_msg = "Requests for this user isn't allowed"
			http.Error(w, err_msg, http.StatusInternalServerError)
			res.Error = err_msg
			encoder.Encode(res)

			return
		}
	*/

	transliter := models.NewTransliter(req.Table)
	resultString, err := transliter.GoLatin(req.String)
	if err != nil {
		err_msg = err.Error()
		http.Error(w, err_msg, http.StatusInternalServerError)
		res.Error = err_msg
		encoder.Encode(res)

		return
	}

	/*
		if err = apiUser.RegisterRequest(); err != nil {
			err_msg = err.Error()
			http.Error(w, err_msg, http.StatusInternalServerError)
			res.Error = err_msg
			encoder.Encode(res)

			return
		}
	*/

	res.String = resultString
	encoder.Encode(res)
}
