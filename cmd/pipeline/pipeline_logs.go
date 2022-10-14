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

// pipelineLogsCmd represents the pipelineLogs command
var pipelineLogsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Gets the logs of a pipeline run",
	Long:  `Gets the logs of a pipeline run`,
	Run: func(cmd *cobra.Command, args []string) {
		project, err := cmd.Flags().GetString("project")
		if err != nil {
			log.Fatal(err)
		}

		pipelineId, err := cmd.Flags().GetString("pipelineId")
		if err != nil {
			log.Fatal(err)
		}

		runId, err := cmd.Flags().GetString("runId")
		if err != nil {
			log.Fatal(err)
		}

		creds, err := internal.ReadCredentials()
		if err != nil {
			log.Fatal(err)
		}

		responseBody, err := internal.InvokeDevOpsAPI(http.MethodGet, fmt.Sprintf("https://dev.azure.com/%s/%s/_apis/pipelines/%s/runs/%s/logs", creds.Organization, project, pipelineId, runId), creds.Token, nil)
		if err != nil {
			log.Fatal(err)
		}

		var responseJson internal.PipelineRunLogsResponse
		json.Unmarshal([]byte(responseBody), &responseJson)

		printPipelineRunLog(responseJson, creds, project)
	},
}

func printPipelineRunLog(pipelineRunLogResponse internal.PipelineRunLogsResponse, creds internal.Credentials, project string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Id", "Name", "State", "Result", "Created", "Finished"})
	//for _, pipelineRunLog := range pipelineRunLogResponse.Logs {
	responseBody, err := internal.InvokeDevOpsAPI(http.MethodGet, fmt.Sprintf("https://dev.azure.com/%s/%s/_build/results?buildId=380&view=logs", creds.Organization, project), creds.Token, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseBody))
	//}
}

func init() {
	pipelinesCmd.AddCommand(pipelineLogsCmd)

	pipelineLogsCmd.PersistentFlags().StringP("pipelineId", "i", "", "The id of the pipeline to show the runs for")
	pipelineLogsCmd.MarkPersistentFlagRequired("pipelineId")

	pipelineLogsCmd.PersistentFlags().StringP("runId", "r", "", "The id of the pipeline run to show logs for")
	pipelineLogsCmd.MarkPersistentFlagRequired("runId")
}
