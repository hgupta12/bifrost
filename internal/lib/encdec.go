package lib

import (
	"encoding/base64"
	"encoding/json"
)

func Encode(obj interface{}) (string, error) {
	b, err := json.Marshal(obj)

	if err != nil {
		return "", err
	}

	encodedOffer := base64.StdEncoding.EncodeToString(b)

	return encodedOffer, nil
}

func Decode(in string, obj interface{}) error {
	decodedBytes, err := base64.StdEncoding.DecodeString(in)

	if err != nil {
		return err
	}

	return json.Unmarshal(decodedBytes, obj)
} 