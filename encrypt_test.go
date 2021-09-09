package security

import (
	"testing"
)

var plainText = []byte("Hello world, i need a hash salt. 'i'm the salt!'")

func TestAESEncryptAndConvers(t *testing.T) {
	key := &[32]byte{0: 128, 11: 127, 14: 225, 31: 224}
	c1, err := AES256Encrypt(plainText, key)
	if err != nil {
		t.Error("Encrypt()", err)
	}
	hex1 := ToHexString(c1)
	hex11 := ToHexString(c1)
	if hex1 != hex11 {
		t.Error("ToString() failed")
	}
	c2, _ := AES256Encrypt(plainText, key)
	hex2 := ToHexString(c2)
	// Even the same key will produce different ciphertext
	if hex1 == hex2 {
		t.Error("Encrypt() failed! With the same key, ciphertext should be diffent")
	}

	p1, err := AES256Decrypt(c1, key)
	if err != nil {
		t.Error("Decrypt()", err)
	}
	if string(p1) != string(plainText) {
		t.Error("Decrypt() failed! Expect:", string(plainText), "gotten:", string(p1))
	}
	b, err := DeHexString(hex2)
	if err != nil {
		t.Error("ToBytes() failed")
	}
	p2, err := AES256Decrypt(b, key)
	if err != nil {
		t.Error("Decrypt()", err)
	}
	if string(p2) != string(plainText) {
		t.Error("Decrypt() failed, expect:", string(plainText), "gotten:", string(p2))
	}
}
