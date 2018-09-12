// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// gmailCmd represents the gmail command
var gmailCmd = &cobra.Command{
	Use:   "gmail",
	Short: "GMail file exfiltration",
	Long: `GMail file exfiltration`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gmail called")
	},
}

func init() {
	routeCmd.AddCommand(gmailCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gmailCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gmailCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	gmailCmd.Flags().StringP("username", "u", "", "Username to use with email")
	gmailCmd.Flags().StringP("password", "p", "", "Password to use with email")
	gmailCmd.Flags().StringP("server", "s", "", "Server address")
	gmailCmd.Flags().IntP("port", "", 587, "Server port to use")
}
