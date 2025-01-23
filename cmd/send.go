package cmd

import (
	"fmt"

	"github.com/hgupt12/bifrost/internal/lib"
	session "github.com/hgupt12/bifrost/internal/session/send"
	"github.com/pion/webrtc/v4"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(sendCmd)
}

var sendCmd = &cobra.Command{
	Use:   "send [filepath...]",
	Short: "Send files",
	Long:  "Send files",
	RunE: func(cmd *cobra.Command, args []string) error {
		s := session.NewSession(len(args))
		fmt.Println(args)
		if err := s.CreateConnection(); err != nil {
			return err
		}
		if err := s.CreateControlChannel(); err != nil {
			return err
		}

		if err := s.CreateTransferChannels(args); err != nil {
			return err
		}		

		encodedOffer, err := s.CreateOffer()
		if err != nil {
			return err
		}
		fmt.Println(encodedOffer)
		s.HandleState()

		answer, err := lib.ReadSDP()
		if err != nil {
			return err
		}

		decodedAnswer := webrtc.SessionDescription{}
		if err = lib.Decode(answer, &decodedAnswer); err != nil {
			return err
		}

		if err = s.PeerConnection.SetRemoteDescription(decodedAnswer); err != nil {
			return err
		}

		select{}
	},
}
