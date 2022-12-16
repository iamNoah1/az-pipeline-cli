/*
Copyright © 2022 Noah Ispas <noahispas.public@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package internal

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func InvokeDevOpsAPI(method string, url string, token string, body io.Reader) ([]byte, error) {
	logger := GetLogger()

	logger.DPanicf("Going to call '%s' with '%s' method", url, method)

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		logger.Fatal(err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	q := req.URL.Query()
	q.Add("api-version", "6.0")

	req.URL.RawQuery = q.Encode()

	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(token)))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Fatal("Errored when sending request to the server", err)
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		message := ""
		if resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
			responseBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				logger.Fatal(err)
				return nil, err
			}
			message = fmt.Sprintf("Login is not valid, try to login again. Error details: %s", string(responseBody))
		} else {
			responseBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				logger.Fatal(err)
				return nil, err
			}
			message = fmt.Sprintf("Azure DevOps API returned an error indicating http code. Error details: %s", string(responseBody))
		}
		logger.Error(message)
		return nil, errors.New(message)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Fatal(err)
		return nil, err
	}

	return responseBody, nil
}
