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
package pipeline

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/iamNoah1/az-pipeline-cli/internal"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

// listPipelinesCmd represents the listPipelines command
var listPipelinesCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all pipelines of a given project",
	Long:  `Lists all pipelines of a given project`,
	Run: func(cmd *cobra.Command, args []string) {
		project, err := cmd.Flags().GetString("project")
		if err != nil {
			log.Fatal(err)
		}

		creds, err := internal.ReadCredentials()
		if err != nil {
			log.Fatal(err)
		}

		responseBody, err := internal.InvokeDevOpsAPI(http.MethodGet, fmt.Sprintf("https://dev.azure.com/%s/%s/_apis/pipelines", creds.Organization, project), creds.Token, nil)
		if err != nil {
			log.Fatal(err)
		}

		var responseJson internal.PipelineResponse
		json.Unmarshal([]byte(responseBody), &responseJson)
		//fmt.Print(string(responseBody))

		printPipelines(responseJson)
	},
}

func printPipelines(pipelines internal.PipelineResponse) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Name"})

	for _, pipeline := range pipelines.Value {
		t.AppendRow(table.Row{pipeline.Id, pipeline.Name})
	}

	t.Render()
}

func init() {
	pipelinesCmd.AddCommand(listPipelinesCmd)
}
