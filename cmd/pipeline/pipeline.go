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
	"fmt"
	"log"

	"github.com/iamNoah1/az-pipeline-cli/internal"

	"github.com/iamNoah1/az-pipeline-cli/cmd"
	"github.com/spf13/cobra"
)

// pipelinesCmd represents the pipeline command
var pipelinesCmd = &cobra.Command{
	Use:   "pipeline",
	Short: "Umbrella command for everything related to pipelines",
	Long:  `Umbrella command for everything related to pipelines`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Error: must also specify a subcommand for the project resource")
	},
}

func getProject(cmd *cobra.Command, creds internal.Credentials) string {
	projectFromFlag, err := cmd.Flags().GetString("project")
	if err != nil {
		log.Fatal(err)
	}

	var project string
	if "" == projectFromFlag {
		if "" == creds.Project {
			log.Fatal("Project must be set either with 'project set' command or through flag")
		}
		project = creds.Project
	}
	return project
}

func Init() {
	cmd.RootCmd.AddCommand(pipelinesCmd)

	pipelinesCmd.PersistentFlags().StringP("project", "p", "", "The project, you want to list the pipelines for.")
	//pipelinesCmd.MarkPersistentFlagRequired("project")
}
