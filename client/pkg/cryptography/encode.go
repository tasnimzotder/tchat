package cryptography

import "encoding/base64"

func EncodeBase64(data []byte) string {
	encodedData := base64.StdEncoding.EncodeToString(data)
	return encodedData
}

func DecodeBase64(data string) ([]byte, error) {
	decodedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	return decodedData, nil
}
