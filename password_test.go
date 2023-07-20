package security

import "testing"

func TestPassword(t *testing.T) {
	cipher, err := BcryptPwd(0, plainText)
	if err != nil {
		t.Error("HashPassword()", err)
		return
	}
	err = CheckBcryptPwd(cipher, plainText)
	if err != nil {
		t.Error("CheckPasswordHash()", err)
	}
}
