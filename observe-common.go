package common

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type Config struct {
	ApiUrl    string
	ExtraPath string
	Customer  string
	Token     string
}

func SendPayload(payload interface{}, config Config) (string, error) {
	json_payload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req_url := "https://" + config.ApiUrl + "/" + config.ExtraPath
	bearer_auth := "Bearer " + config.Customer + " " + config.Token

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

func ReadConfig(config_path string) (Config, error) {
	config := Config{}
	config_file, err := os.Open(config_path)
	if err != nil {
		return config, err
	}
	defer config_file.Close()

	decoder := json.NewDecoder(config_file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
