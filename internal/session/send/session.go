package send

import (
	"github.com/hgupt12/bifrost/internal/lib"
	"github.com/pion/webrtc/v4"
)

type Session struct {
	PeerConnection *webrtc.PeerConnection
	controlChannel *webrtc.DataChannel
	channels []*lib.Document
	numberOfFiles int
}

func NewSession(numberOfFiles int) *Session {
	return &Session{
		channels: make([]*lib.Document, numberOfFiles),
		numberOfFiles: numberOfFiles,
	}
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