package security

import (
	"testing"
)

var plainText = []byte("hello admin")

func TestEncryptAndConvers(t *testing.T) {
	key := &[32]byte{0: 128, 11: 127, 14: 225, 31: 224}
	c1, err := Encrypt(plainText, key)
	if err != nil {
		t.Error("Encrypt()", err)
	}
	hex1 := ToString(c1)
	hex11 := ToString(c1)
	if hex1 != hex11 {
		t.Error("ToString() failed")
	}
	c2, _ := Encrypt(plainText, key)
	hex2 := ToString(c2)
	// Even the same key will produce different ciphertext
	if hex1 == hex2 {
		t.Error("Encrypt() failed with the same key, ciphertext should be diffent")
	}

	p1, err := Decrypt(c1, key)
	if string(p1) != string(plainText) {
		t.Error("Decrypt() failed, expect:", string(plainText), "gotten:", string(p1))
	}
	b, err := ToBytes(hex2)
	if err != nil {
		t.Error("ToBytes() failed")
	}
	p2, err := Decrypt(b, key)
	if string(p2) != string(plainText) {
		t.Error("Decrypt() failed, expect:", string(plainText), "gotten:", string(p2))
	}
}
