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
	"github.com/spf13/cobra"
)

// routeCmd represents the route command
var routeCmd = &cobra.Command{
	Use:   "route",
	Short: "Exfiltration routes",
	Long: `Exfiltration routes`,
	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("Exfiltration routes...")
	//},
}

func init() {
	rootCmd.AddCommand(routeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	routeCmd.PersistentFlags().IntP("timing", "t", 0, "Control the timing for sending files")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rshellCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	routeCmd.Flags().BoolP("all", "a", false, "Use all exfiltration routes with default configurations")
	routeCmd.Flags().BoolP("recieve", "r", false, "Recieve data from exfiltration routes")
}
