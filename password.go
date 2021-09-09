package security

import (
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
func CheckBcryptPwd(hash, password []byte) error {
	return bcrypt.CompareHashAndPassword(hash, password)
}
