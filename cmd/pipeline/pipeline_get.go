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

	"github.com/iamNoah1/az-pipeline-cli/internal"
	"github.com/spf13/cobra"
)

// pipelineGetCmd represents the pipelineGet command
var pipelineGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets the definition of a pipeline",
	Long:  `Gets the definition of a pipeline`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := internal.GetLogger()

		runId, err := cmd.Flags().GetString("runId")
		if err != nil {
			log.Fatal(err)
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

		// Don't ask why, but 1 is the pipeline definition
		responseBody, err = internal.InvokeDevOpsAPI(http.MethodGet, fmt.Sprintf("https://dev.azure.com/%s/%s/_apis/build/builds/%s/logs/1", creds.Organization, project, runId), creds.Token, nil)
		if err != nil {
			logger.Fatal(err)
		}
		fmt.Println(string(responseBody))
	},
}

func init() {
	pipelinesCmd.AddCommand(pipelineGetCmd)

	pipelineGetCmd.PersistentFlags().StringP("pipelineId", "i", "", "The id of the pipeline to show the runs for")
	pipelineGetCmd.MarkPersistentFlagRequired("pipelineId")

	pipelineGetCmd.PersistentFlags().StringP("runId", "r", "", "The id of the pipeline run to show logs for")
	pipelineGetCmd.MarkPersistentFlagRequired("runId")
}
