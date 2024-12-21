package session

import (
	"github.com/hgupt12/bifrost/internal/lib"
	"github.com/pion/webrtc/v4"
)

func (s *Session) CreateRecieverConnection(offerString string) (string, error) {
	offer := webrtc.SessionDescription{}

	if err := lib.Decode(offerString,&offer); err != nil {
		return "", err
	}

	if err := s.CreateConnection(); err != nil {
		return "", err
	}

	if err := s.PeerConnection.SetRemoteDescription(offer); err != nil {
		return "", err
	} 

	answer, err := s.PeerConnection.CreateAnswer(nil)
	if err != nil {
		return "", err
	}

	if err = s.PeerConnection.SetLocalDescription(answer); err != nil {
		return "", err
	}

	gatherDone := webrtc.GatheringCompletePromise(s.PeerConnection)
	<-gatherDone
	encodedAnswer, err := lib.Encode(answer)
	if err != nil {
		return "", err
	}

	return encodedAnswer, nil
}

