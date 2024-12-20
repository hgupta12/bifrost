package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Get the current version of Bifrost",
	Long: "Get the current version of Bifrost",
	Run: func(cmd *cobra.Command, args []string){
		fmt.Println("v0.1")
	},
}