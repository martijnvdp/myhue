/*
Copyright © 2021 M van der Ploeg <EMAIL ADDRESS>

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

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "set light 0 on 49",
	Long: `Set light on/off and specify brightness. For example:
	myhue set light 0 on 49 ,turn light 0 on with a brightness level of 49
MYHue is a cli app to interact with a philips hue bridge.`,
	Run: func(cmd *cobra.Command, args []string) {
		token, bridge := functions.ConnectHUE()
		l, err := cmd.Flags().GetInt("light")
		br, err := cmd.Flags().GetInt("bright")
		a, err := cmd.Flags().GetBool("on")
		if err == nil {
			functions.SetLight(&l, &a, &br, token, bridge)
		}
	},
}

func init() {
	var l, br int
	var on bool
	rootCmd.AddCommand(setCmd)
	setCmd.Flags().IntVarP(&l, "light", "l", 0, "light id")
	setCmd.Flags().IntVarP(&br, "bright", "b", 0, "brightness level")
	setCmd.Flags().BoolVarP(&on, "on", "o", false, "Light on or off")
}
