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
