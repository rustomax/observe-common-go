package common

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func SendPayload(payload interface{}, ApiUrl, ExtraPath, Customer, Token string) (string, error) {
	json_payload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req_url := "https://" + ApiUrl + "/" + ExtraPath
	bearer_auth := "Bearer " + Customer + " " + Token

	req, _ := http.NewRequest(http.MethodPost, req_url, bytes.NewBuffer(json_payload))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", bearer_auth)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	result := string(bodyBytes)
	return result, nil
}
