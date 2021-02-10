/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/martijnxd/myhue/functions"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all lights",
	Long: `myhue list.

MYHue is a cli app to interact with a philips hue bridge.`,
	Run: func(cmd *cobra.Command, args []string) {
		token, bridge := functions.ConnectHUE()
		w, _ := cmd.Flags().GetBool("wide")
		r, _ := cmd.Flags().GetBool("reachable")

		functions.ListAll(r, w, token, bridge)
	},
}

func init() {
	var w, r bool
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(&w, "wide", "w", false, "wide list")
	listCmd.Flags().BoolVarP(&r, "reachable", "r", false, "display reachable lights only")

}
