// The MIT License (MIT)
//
// Copyright © 2019 Rangertaha <rangertaha@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// reconCmd represents the recon command
var reconCmd = &cobra.Command{
	Use:   "recon [objects]",
	Short: "Conduct reconnaissance of the local system",
	Long:  `Conduct reconnaissance of the local system and identifying high value objects to exfiltrate via routes`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("recon called")
	},
}

func init() {
	rootCmd.AddCommand(reconCmd)
	reconCmd.Flags().StringArrayP("type", "t", []string{"all"}, "Types of objects to search for")
	// reconCmd.Flags().BoolP("report", "r", false, "Print a report on the types of objects it found")
}
