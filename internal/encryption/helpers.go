package encryption

import (
	"crypto/aes"
	"crypto/rand"
	"ftcli/config"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/chacha20"
)

// Generates the initilization vector (IV)
// Returns IV (16bytes)
func GenerateIV() ([]byte, error) {

	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}
	return iv, nil

}

// Returns a 16byte salt
func GenerateSalt() ([]byte, error) {

	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}
	return salt, nil

}

// Returns a 12byte nonce
func GenerateNonce() ([]byte, error) {

	// This mighbe be why i'm still getting high mem-use on the send side.
	// nonces are read from the header on the receive side.
	// This could be taking up unncessary memory. how to use this in a more memory efficient way?
	// OR OR OR
	// i should send just once nonce with a counter
	nonce := make([]byte, chacha20.NonceSize)
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}
	return nonce, nil

}

// Generates a master key using Argon2. See config for cost parameter specifics.
func GenerateMasterKey(salt []byte, password string) []byte {
	return argon2.Key(
		[]byte(password),
		salt,
		config.Time,
		config.Memory,
		config.Threads,
		config.KeyLength)
}
