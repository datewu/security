package security

import "encoding/hex"

// ToString ...
func ToString(data []byte) string {
	return hex.EncodeToString(data)

}

// ToBytes ...
func ToBytes(content string) ([]byte, error) {
	data, err := hex.DecodeString(content)
	if err != nil {
		return nil, err
	}
	return data, nil
}
