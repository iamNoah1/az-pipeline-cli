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

	"github.com/iamNoah1/az-pipeline-cli/internal"
	"github.com/spf13/cobra"
)

// pipelineRunsCmd represents the pipelineRuns command
var pipelineRunsCmd = &cobra.Command{
	Use:   "runs",
	Short: "List runs of a given pipeline",
	Long:  `List runs of a given pipeline`,
	Run: func(cmd *cobra.Command, args []string) {
		project, err := cmd.Flags().GetString("project")
		if err != nil {
			log.Fatal(err)
		}

		pipelineId, err := cmd.Flags().GetString("pipelineId")
		if err != nil {
			log.Fatal(err)
		}

		creds, err := internal.ReadCredentials()
		if err != nil {
			log.Fatal(err)
		}

		responseBody, err := internal.InvokeDevOpsAPI(fmt.Sprintf("https://dev.azure.com/%s/%s/_apis/pipelines/%s/runs", creds.Organization, project, pipelineId), creds.Token)
		if err != nil {
			log.Fatal(err)
		}

		var responseJson internal.PipelineRunResponse
		json.Unmarshal([]byte(responseBody), &responseJson)
		//fmt.Print(string(responseBody))

		for _, pipelineRun := range responseJson.Value {
			fmt.Printf("Name: %s, Created: %s, Finished: %s, State: %s, Result: %s ,\n", pipelineRun.Name, pipelineRun.Created, pipelineRun.Finished, pipelineRun.State, pipelineRun.Result)
		}
	},
}

func init() {
	pipelinesCmd.AddCommand(pipelineRunsCmd)

	pipelineRunsCmd.PersistentFlags().StringP("pipelineId", "i", "", "The id of the pipeline to show the runs for")
	pipelineRunsCmd.MarkPersistentFlagRequired("pipelineId")
}