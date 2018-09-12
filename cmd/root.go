// Copyright © 2018 rangertaha rangertaha@gmail.com
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
	"os"
	"fmt"

	"github.com/spf13/cobra"
	//"github.com/rangertaha/exfil"
)


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "exfil [command]",
	Short: "Exfiltrate data from a system",
	Long: `Data exfiltration framework designed to triage and exfiltrate files by user selected priority`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	//// Basic options
	//rootCmd.PersistentFlags().StringArrayP("keyboards", "k", []string{"en1"},
	//	"Keyboards/layouts ID to use")
	////rootCmd.PersistentFlags().StringArrayP("languages", "l", []string{"all"},
	////	"Language ID to use for linguistic typos")
	//
	//// Processing
	//rootCmd.PersistentFlags().IntP("concurrency", "c", 50,
	//	"Number of concurrent workers")
	//rootCmd.PersistentFlags().StringArrayP("typos", "t", []string{"all"},
	//	"Types of typos to perform")
	//
	//// Post Processing options for retrieving additional data
	//rootCmd.PersistentFlags().StringArrayP("funcs", "x", []string{"idna"},
	//	"Extra functions for data or filtering")
	//
	//// Output options
	//rootCmd.PersistentFlags().StringP("file", "f", "", "Output filename")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Output additional details")
}

