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

func NewLog() *Log {
	l := &Log{}

	return l
}

func (l *Log) Save() {
	j, err := json.Marshal(l)
	if err != nil {
		return
	}

	url := "http://localhost:8081/log"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
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
