package request

import (
	"io/ioutil"
	"net/http"
)

func NewRequest(method, url, token string) ([]byte, int, error) {
	client := http.Client{}

	req, reqErr := http.NewRequest(method, url, nil)
	if reqErr != nil {
		return nil, 0, reqErr
	}

	//this is the man of the day
	if token != "" {
		req.Header.Add("X-RapidAPI-Key", token)
		req.Header.Add("Content-Type", "application/json")
	}

	resp, respErr := client.Do(req)
	if respErr != nil {
		return nil, 0, respErr
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	return resp_body, resp.StatusCode, nil
}
