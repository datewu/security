package security

import (
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

// BcryptPwd generates a bcrypt hash of the password using work factor 14.
func BcryptPwd(f int, password []byte) ([]byte, error) {
	if f == 0 {
		f = 14
	}
	return bcrypt.GenerateFromPassword(password, f)
}

// CheckBcryptPwd securely compares a bcrypt hashed password with its possible
// plaintext equivalent.  Returns nil on success, or an error on failure.
func CheckBcryptPwd(hashed, plainText []byte) error {
	return bcrypt.CompareHashAndPassword(hashed, plainText)
}

// EncodeSignatureJWT Encodes an ECDSA signature according to
// https://tools.ietf.org/html/rfc7515#appendix-A.3.1
func EncodeSignatureJWT(sig []byte) string {
	return base64.RawURLEncoding.EncodeToString(sig)
}

// DecodeSignatureJWT Decodes an ECDSA signature according to
// https://tools.ietf.org/html/rfc7515#appendix-A.3.1
func DecodeSignatureJWT(b64sig string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(b64sig)
}
