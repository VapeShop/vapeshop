package cmd

import (
	"github.com/spf13/cobra"
)

var VapeshopCmd = &cobra.Command{
	Use:   "vape",
	Short: "Vapeshop is a reliable package manager for smart contracts",
	Long:  `Vapeshop is a reliable package manager for smart contracts`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}
