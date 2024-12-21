package session

import (
	"fmt"

	"github.com/hgupt12/bifrost/internal/lib"
	"github.com/pion/webrtc/v4"
)

func (s *Session) CreateDataChannel() error {
	_, err := s.PeerConnection.CreateDataChannel("control", nil)

	if err != nil {
		return err
	}

	return nil
}

func (s *Session) CreateOffer() (string, error) {
	offer, err := s.PeerConnection.CreateOffer(nil)
	if err != nil {
		return "", err
	}
	
	gatherDone := webrtc.GatheringCompletePromise(s.PeerConnection)
	
	if err = s.PeerConnection.SetLocalDescription(offer); err != nil {
		return "", err
	}
	<-gatherDone
	
	offer2 := s.PeerConnection.LocalDescription()
	encodedOffer, _ := lib.Encode(offer2)

	return encodedOffer, nil
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
}
