package security

import (
	"fmt"
	"testing"
)

func ExampleHash() {
	tag := "hello tag"
	sha512 := Hash(tag, plainText)
	md5 := Md5(tag, plainText)
	smd5 := SimpleMd5(plainText)
	fmt.Println(ToHexString(sha512))
	fmt.Println(ToHexString(md5))
	fmt.Println(ToHexString(smd5))
	// Output:
	// 8de93ec7b111745d112d6ae96b16b3bad87bad03a734462bb7e89d2c47cb91da
	// d9d29b0dbca7306b5f4b32e03662f0a9
	// 3f37b87408cec44b69ec95d80fd3f5ba
}

func TestPassword(t *testing.T) {
	cipher, err := HashPassword(0, plainText)
	if err != nil {
		t.Error("HashPassword()", err)
		return
	}
	err = CheckPasswordHash(cipher, plainText)
	if err != nil {
		t.Error("CheckPasswordHash()", err)
	}
}
