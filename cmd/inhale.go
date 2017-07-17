// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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

	"github.com/VapeShop/vapeshop/definitions"

	"github.com/spf13/cobra"
)

// inhaleCmd represents the inhale command
var inhaleCmd = &cobra.Command{
	Use:   "inhale",
	Short: "Install a package of smart contracts onto your system",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("inhale called")
	},
}

func buildInhaleCommand() {
	VapeshopCmd.AddCommand(inhaleCmd)
	addInhaleFlags()
}

func addInhaleFlags() {
	inhaleCmd.Flags().StringSliceVar(p, "flavors", []string{}, "Choose from a variety of smart contract bindings to generate for the smart contracts in this package. Options include"+fmt.Sprintf("%v", definitions.Flavors))
}
