package security

import "encoding/hex"

// ToHexString ...
func ToHexString(data []byte) string {
	return hex.EncodeToString(data)

}

// DeHexString ...
func DeHexString(content string) ([]byte, error) {
	data, err := hex.DecodeString(content)
	if err != nil {
		return nil, err
	}
	return data, nil
}
