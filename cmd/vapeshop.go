package cmd

import (
	"github.com/spf13/cobra"
)

func init() {

	VapeshopCmd.AddCommand(exhaleCmd)
	VapeshopCmd.AddCommand(versionCmd)
}

var VapeshopCmd = &cobra.Command{
	Use:   "vape",
	Short: "Vapeshop is a reliable package manager for smart contracts",
	Long:  `Vapeshop is a reliable package manager for smart contracts`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() (err error) {
	initializeVapeShopFlags()
	err = VapeshopCmd.Execute()
	return
}

func initializeVapeShopFlags() {

}
