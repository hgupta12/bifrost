package lib

import (
	"os"

	"github.com/pion/webrtc/v4"
)

type Document struct {
	File *os.File
	Name string
	DC *webrtc.DataChannel
}