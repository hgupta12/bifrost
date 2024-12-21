package session

import "github.com/pion/webrtc/v4"

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