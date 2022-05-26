package helper

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
)

func CreatePost(url string, payload map[string]interface{}) error {
	payloadStr, _ := json.Marshal(payload)

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payloadStr))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated {
		var result map[string]interface{}
		decoder := json.NewDecoder(response.Body)
		err = decoder.Decode(&result)
		if err != nil {
			return err
		}

		var reason string
		if response.StatusCode == http.StatusBadRequest {
			reason = strings.Join(result["message"].([]string), " ")
		} else {
			reason = result["reason"].(string)
		}
		panic(reason)
	}

	return nil
}
