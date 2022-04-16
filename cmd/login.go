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
package cmd

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strings"

	"github.com/spf13/cobra"
)

const credsFileName = ".az-pipelines-creds"

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to Azure DevOps",
	Long:  `Loging in to Azure DevOps for this CLI means to grab username and PAT and store them for further commands`,

	Run: func(cmd *cobra.Command, args []string) {
		exists, err := exists(credsFileAbsolute())
		if nil != err {
			log.Fatal(err)
		}
		if exists {
			fmt.Println("logged in")
			return
		}

		fmt.Print("username: ")
		var username = readFromConsole()

		fmt.Print("PAT: ")
		var token = readFromConsole()

		creds := Credentials{username, token}
		err = writeCredentials(creds)

		if nil != err {
			log.Fatal(err)
		} else {
			fmt.Println("logged in")
		}
	},
}

func exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

type Credentials struct {
	Username, Token string
}

func credsFileAbsolute() string {
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}
	return user.HomeDir + "/" + credsFileName
}

func readFromConsole() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func writeCredentials(creds Credentials) error {
	_, err := os.Create(credsFileAbsolute())
	if nil != err {
		log.Fatalf("Could not create credentials file. Error: %s", err)
	}

	content, _ := json.Marshal(creds)
	return ioutil.WriteFile(credsFileAbsolute(), content, 0644)
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
