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
	"fmt"
	"log"

	"github.com/iamNoah1/az-pipeline-cli/internal"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to Azure DevOps",
	Long:  `Loging in to Azure DevOps for this CLI means to grab username and PAT and store them for further commands`,

	Run: func(cmd *cobra.Command, args []string) {
		exists, err := internal.FileExists(internal.CredsFileAbsolute())
		if nil != err {
			log.Fatal(err)
		}

		force, err := cmd.Flags().GetBool("force")
		if nil != err {
			log.Fatal(err)
		}

		if exists && !force {
			fmt.Println("logged in")
			return
		}

		fmt.Print("username: ")
		var username string = internal.ReadFromConsole()

		fmt.Print("organization: ")
		var organization string = internal.ReadFromConsole()

		fmt.Print("PAT: ")
		var token string = ":" + internal.ReadFromConsole()

		creds := internal.Credentials{Username: username, Token: token, Organization: organization}
		err = internal.WriteCredentials(creds)

		if nil != err {
			log.Fatal(err)
		} else {
			fmt.Println("logged in")
		}
	},
}

func init() {
	RootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	loginCmd.Flags().BoolP("force", "f", false, "Force login even if already logged in")
}
