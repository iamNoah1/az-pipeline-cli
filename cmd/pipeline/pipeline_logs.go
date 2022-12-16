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
	"net/http"

	"github.com/iamNoah1/az-pipeline-cli/internal"
	"github.com/spf13/cobra"
)

var pipelineLogsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Gets the logs of a pipeline run",
	Long:  `Gets the logs of a pipeline run`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := internal.GetLogger()

		runId, err := cmd.Flags().GetString("runId")
		if err != nil {
			logger.Fatal(err)
		}

		creds, err := internal.ReadCredentials()
		if err != nil {
			logger.Fatal(err)
		}

		project := getProject(cmd, creds)

		responseBody, err := internal.InvokeDevOpsAPI(http.MethodGet, fmt.Sprintf("https://dev.azure.com/%s/%s/_apis/build/builds/%s/logs", creds.Organization, project, runId), creds.Token, nil)
		if err != nil {
			logger.Fatal(err)
		}

		var logsResponse internal.PipelineRunLogsResponse
		json.Unmarshal([]byte(responseBody), &logsResponse)

		// starting with 2 here, because 1 is the pipeline definition
		for i := 2; i < logsResponse.Count; i++ {
			responseBody, err := internal.InvokeDevOpsAPI(http.MethodGet, fmt.Sprintf("https://dev.azure.com/%s/%s/_apis/build/builds/%s/logs/%d", creds.Organization, project, runId, i), creds.Token, nil)
			if err != nil {
				logger.Fatal(err)
			}
			logger.Info(string(responseBody))
		}

	},
}

func init() {
	pipelinesCmd.AddCommand(pipelineLogsCmd)

	pipelineLogsCmd.PersistentFlags().StringP("pipelineId", "i", "", "The id of the pipeline to show the runs for")
	pipelineLogsCmd.MarkPersistentFlagRequired("pipelineId")

	pipelineLogsCmd.PersistentFlags().StringP("runId", "r", "", "The id of the pipeline run to show logs for")
	pipelineLogsCmd.MarkPersistentFlagRequired("runId")
}
