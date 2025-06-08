package armory

import "crypto/sha256"

func Hash256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}
