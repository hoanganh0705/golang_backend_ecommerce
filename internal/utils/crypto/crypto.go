package crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

// tabnine: test | explain | document | ask
func GetHash(key string) string {

	// creare a sha256 hasher
	hash := sha256.New()

	// move the key that need to be hashed into the hasher and change it to bytes
	hash.Write([]byte(key))

	// get the hashed bytes
	hashBytes := hash.Sum(nil)

	// convert bytes to hex string
	return hex.EncodeToString(hashBytes)
}
