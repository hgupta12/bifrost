package receive

import (
	"fmt"

	"github.com/pion/webrtc/v4"
)

type Session struct {
	PeerConnection *webrtc.PeerConnection

}

func NewSession() *Session {
	return &Session{}
}

func (s *Session) CreateConnection() error {
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}
	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		panic(err)
	}

	s.PeerConnection = peerConnection

	return nil
}

func (s *Session) HandleState() {
	s.PeerConnection.OnICEConnectionStateChange(func(state webrtc.ICEConnectionState) {
		fmt.Printf("ICE Connection State has changed: %s\n", state.String())
		if state == webrtc.ICEConnectionStateConnected {
			fmt.Println("Connection successfully established!")
		}
		if state == webrtc.ICEConnectionStateFailed {
			fmt.Println("Connection failed!")
		}
	})

	s.PeerConnection.OnDataChannel(func (dc *webrtc.DataChannel) {
		dc.OnOpen(func() {
			fmt.Printf("Data Channel opened %s, %d\n", dc.Label(), dc.ID())
		})

		dc.OnClose(func () {
			fmt.Println("Data Channel closed")
		})

		dc.OnMessage(func(msg webrtc.DataChannelMessage) {

			if dc.Label() == "control" {
				fmt.Print(string(msg.Data))
				var consent []byte

				if _, err := fmt.Scanln(&consent); err != nil {
					panic(err)
				}
				fmt.Println(consent)
				err := dc.SendText(string(consent))

				if err != nil {
					panic(err)
				}
			}

		})
	})

}
