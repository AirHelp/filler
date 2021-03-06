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

	"github.com/AirHelp/filler/version"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the Filler verison information",
	Long:  "Prints the Filler verison information",
	Run: func(cmd *cobra.Command, args []string) {
		v := version.Get()
		fmt.Printf("Version: %s\n", v.Version)
		fmt.Printf("GitCommit: %s\n", v.GitCommit)
		fmt.Printf("GitTreeState: %s\n", v.GitTreeState)
		fmt.Printf("BuildDate: %s\n", v.BuildDate)
		fmt.Printf("GoVersion: %s\n", v.GoVersion)
		fmt.Printf("Compiler: %s\n", v.Compiler)
		fmt.Printf("Platform: %s\n", v.Platform)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
