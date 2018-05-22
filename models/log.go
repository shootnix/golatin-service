package models

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Log struct {
	Module      string `json:"module"`
	Action      string `json:"action"`
	Result      string `json:"result"`
	Description string `json:"desc"`
}

var LogServiceURL string

func NewLog() *Log {
	l := &Log{}

	return l
}

func (l *Log) Save() {
	j, err := json.Marshal(l)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", LogServiceURL, bytes.NewBuffer(j))
	if err != nil {
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	/*
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}
		if 200 != resp.StatusCode {
			return
		}
	*/
}
