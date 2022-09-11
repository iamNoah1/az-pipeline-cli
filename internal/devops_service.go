package internal

import (
	"encoding/base64"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

func InvokeDevOpsAPI(url string, token string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	q := req.URL.Query()
	q.Add("api-version", "6.0")

	req.URL.RawQuery = q.Encode()

	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(token)))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Errored when sending request to the server")
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		message := ""
		if resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
			message = "Login is not valid, try to login again"
		}
		return nil, errors.New(message)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return responseBody, nil
}
