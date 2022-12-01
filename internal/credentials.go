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
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

const credsFileName = ".az-pipelines-creds"

type Credentials struct {
	Username     string
	Token        string
	Organization string
	Project      string
}

func CredsFileAbsolute() string {
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}
	return user.HomeDir + "/" + credsFileName
}

func WriteCredentials(creds Credentials) error {
	_, err := os.Create(CredsFileAbsolute())
	if nil != err {
		log.Fatalf("Could not create credentials file. Error: %s", err)
	}

	content, _ := json.Marshal(creds)
	return ioutil.WriteFile(CredsFileAbsolute(), content, 0644)
}

func ReadCredentials() (Credentials, error) {
	raw, err := ioutil.ReadFile(CredsFileAbsolute())

	if nil != err {
		log.Fatalf("Could not read credentials file. Error: %s", err)
	}

	var creds Credentials
	err = json.Unmarshal(raw, &creds)

	return creds, err
}
