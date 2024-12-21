package lib

import (
	"bufio"
	"fmt"
	"os"
)

func ReadSDP() (string, error) {
	fmt.Printf("Paste remote SDP: ")
	r := bufio.NewReader(os.Stdin)

	data, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}

	return string(data), nil
}
