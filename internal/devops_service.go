package internal

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
)

func InvokeDevOpsAPI(url string, token string) []byte {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("api-version", "6.0")

	req.URL.RawQuery = q.Encode()

	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(token)))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Errored when sending request to the server")
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return responseBody
}
