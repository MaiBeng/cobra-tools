package pkg

import (
	"encoding/json"
	"net/http"
	"strings"
)

func PostJson(url string, bodyStr interface{}, response interface{}) error {

	method := "POST"

	body, err := json.Marshal(bodyStr)
	if err != nil {
		return err
	}

	payload := strings.NewReader(string(body))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if response != nil {
		err = json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			return err
		}
	}

	return nil
}
