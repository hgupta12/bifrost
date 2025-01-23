package send

import (
	"fmt"
	"os"

	"github.com/hgupt12/bifrost/internal/lib"
	"github.com/pion/webrtc/v4"
)

func (s *Session) CreateTransferChannels(paths []string) error {

	for i := 0; i < len(paths); i++ {
		file, err := os.Open(paths[i])

		if err != nil {
			return err
		}

		f, err := file.Stat()
		if err != nil {
			return err
		}

		s.channels[i] = &lib.Document{
			File: file,
			Name: f.Name(),
		}

		s.channels[i].DC, err = s.PeerConnection.CreateDataChannel(fmt.Sprintf("dc%d", i), nil)

		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Session) CreateControlChannel() error {
	channel, err := s.PeerConnection.CreateDataChannel("control", nil)

	s.controlChannel = channel

	s.controlChannel.OnOpen(func () {
		fmt.Println("Waiting for receiver's confirmation...")

		err := s.controlChannel.SendText(fmt.Sprintf("Do you want to receive %d files [Y/n]: ",s.numberOfFiles))
		if err != nil {
			panic(err)
		}
	})
	s.controlChannel.OnClose(func () {
		fmt.Println("Closed!!")
	})

	s.controlChannel.OnMessage(func(msg webrtc.DataChannelMessage) {
		if string(msg.Data) == "Y" {
			fmt.Println("Starting file transfer...")
			// TODO - add file sending logic
		}
	})

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
