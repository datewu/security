package security

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha512"
)

// Hash generates a hash of data using HMAC-SHA-512/256. The tag is intended to
// be a natural-language string describing the purpose of the hash, such as
// "hash file for lookup key" or "master secret to client secret".  It serves
// as an HMAC "key" and ensures that different purposes will have different
// hash output. This function is NOT suitable for hashing passwords.
func Hash(tag string, data []byte) []byte {
	h := hmac.New(sha512.New512_256, []byte(tag))
	h.Write(data)
	return h.Sum(nil)
}

// Md5 which you should not use in password
func Md5(tag string, data []byte) []byte {
	hash := hmac.New(md5.New, []byte(tag))
	hash.Write(data)
	return hash.Sum(nil)
}

// SimpleMd5 which does simple sum
func SimpleMd5(data []byte) []byte {
	sum := md5.Sum(data)
	return sum[:]
}
