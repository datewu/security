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
	// 85aba1262bf609a3c618cb1686db45580d231a0c4f2f547aaaeb768d5e1b1c92
	// 563753d2208f028c66f169bd16ab0644
	// 8c04b86aa649a1835d23587afce0aadf
}

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
