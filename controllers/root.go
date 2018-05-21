package controllers

import (
	"encoding/json"
	"errors"
	"github.com/shootnix/golatin-transliter-service/models"
	"net/http"
)

type TranslitRequest struct {
	String string `json:"string"`
	Table  string `json:"table"`
}

type TranslitResponse struct {
	String string `json:"string"`
	Error  string `json:"error"` // omitempty?
}

func POSTTranslit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	encoder := json.NewEncoder(w)

	l := models.NewLog()
	l.Action = "translit"
	l.Module = "transliter"
	l.Result = "FAIL"

	defer r.Body.Close()
	var req TranslitRequest
	var res TranslitResponse

	if err := decoder.Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		res.Error = err.Error()

		l.Description = err.Error()
		go l.Save()

		encoder.Encode(res)

		return
	}

	if err := req.Validate(); err != nil {
		errmsg := err.Error()
		http.Error(w, errmsg, http.StatusInternalServerError)
		res.Error = errmsg

		l.Description = errmsg
		go l.Save()

		encoder.Encode(res)

		return
	}

	t, err := models.NewTransliter(req.Table)
	if err != nil {
		errmsg := err.Error()
		http.Error(w, errmsg, http.StatusInternalServerError)
		res.Error = errmsg

		l.Description = errmsg
		go l.Save()

		encoder.Encode(res)

		return
	}

	resultString, err := t.GoLatin(req.String)
	if err != nil {
		errmsg := err.Error()
		http.Error(w, errmsg, http.StatusInternalServerError)
		res.Error = errmsg

		l.Description = errmsg
		go l.Save()

		encoder.Encode(res)

		return
	}

	l.Result = "SUCCESS"
	l.Description = req.String
	go l.Save()

	res.String = resultString
	encoder.Encode(res)
}

func (req TranslitRequest) Validate() error {
	errmsg := "Validation errors:"
	n_errs := 0

	if req.String == "" {
		errmsg = errmsg + " `string` required."
		n_errs = n_errs + 1
	}

	if req.Table == "" {
		errmsg = errmsg + " `table` required."
		n_errs = n_errs + 1
	}

	if n_errs > 0 {
		return errors.New(errmsg)
	}
	// else
	return nil
}
