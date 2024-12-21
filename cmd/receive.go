package cmd

import (
	"fmt"

	"github.com/hgupt12/bifrost/internal/lib"
	"github.com/hgupt12/bifrost/internal/session"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(receiveCmd)
}

var receiveCmd = &cobra.Command{
	Use: "receive",
	Short: "Receive files",
	Long: "Receive files",
	RunE: func(cmd *cobra.Command, args []string) error {
		input, err := lib.ReadSDP()

		if err != nil {
			return err
		}
		
		s := session.NewSession()
		answer, err := s.CreateRecieverConnection(input)
		if err != nil {
			return err
		}

		s.HandleState()
		
		fmt.Println(answer)
		select{}
	},
}