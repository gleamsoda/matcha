package password

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestHash(t *testing.T) {
	password := "password"
	hashedPassword, err := Hash(password)
	assert.NoError(t, err)
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	assert.NoError(t, err)
}

func TestVerify(t *testing.T) {
	password := "password"
	hashedPassword, _ := Hash(password)

	err := Verify(password, string(hashedPassword))
	assert.NoError(t, err)

	err = Verify("wrongpassword", string(hashedPassword))
	assert.Error(t, err)
}
