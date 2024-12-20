package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "bifrost",
	Short: "Bifrost is a P2P file sharing tool.",
	Long: "Bifrost is a P2P file sharing tool.",
	Run: func(cmd *cobra.Command, args []string){
		fmt.Println("Bifrost v0.1")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}