package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type ESCmd struct {
	Url string `json:"url"`
}

func (es *ESCmd) Search(body any) ([]byte, error) {
	jsonBody, err := json.Marshal(body)

	if err != nil {
		log.Fatalf("error occuried when tryinh to parse json: %v", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", es.Url+"/"+"statistic9730"+"/_search", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}