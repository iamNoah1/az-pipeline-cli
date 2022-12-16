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
package project

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/iamNoah1/az-pipeline-cli/internal"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get a list of all projects in your organization",
	Long:  `Get a list of all projects in your organization`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := internal.GetLogger()

		creds, err := internal.ReadCredentials()
		if err != nil {
			logger.Fatal(err)
		}

		responseBody, err := internal.InvokeDevOpsAPI(http.MethodGet, fmt.Sprintf("https://dev.azure.com/%s/_apis/projects", creds.Organization), creds.Token, nil)
		if err != nil {
			logger.Fatal(err)
		}

		var responseJson internal.ProjectResponse
		json.Unmarshal([]byte(responseBody), &responseJson)

		printProjects(responseJson)
	},
}

func printProjects(projects internal.ProjectResponse) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Id", "Name"})

	for _, project := range projects.Value {
		t.AppendRow(table.Row{project.Id, project.Name})
	}

	t.Render()
}

func init() {
	projectCmd.AddCommand(listCmd)
}
