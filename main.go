package main

import (
	"fmt"
	"os"

	"github.com/VapeShop/vapeshop/cmd"
)

func main() {
	if err := cmd.VapeshopCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
