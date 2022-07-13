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
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/iamNoah1/az-pipeline-cli/internal"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get a list of all projects of your organization",
	Long:  `Get a list of all projects of your organization`,
	Run: func(cmd *cobra.Command, args []string) {
		creds, err := internal.ReadCredentials()
		if err != nil {
			log.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://dev.azure.com/%s/_apis/projects", creds.Organization), nil)
		if err != nil {
			log.Fatal(err)
		}

		q := req.URL.Query()
		q.Add("api-version", "6.0")

		req.URL.RawQuery = q.Encode()

		req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(creds.Token)))

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

		var responseJson internal.ProjectResponse
		json.Unmarshal([]byte(responseBody), &responseJson)

		for _, project := range responseJson.Value {
			fmt.Println(project)
		}
	},
}

func init() {
	projectCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
