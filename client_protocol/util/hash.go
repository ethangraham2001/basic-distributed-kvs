package util

import (
	"crypto/md5"
)

// HashKey hashes the key with the MD5 algorithm, and returns the
// least-signifant four bytes concatenated.
func HashKey(key string) uint32 {
	keyHash := md5.Sum([]byte(key))

	var leastSignificantFourBytes uint32

	for i := 0; i <= 4; i++ {
		leastSignificantFourBytes += uint32(keyHash[i] << i)
	}

	return leastSignificantFourBytes
}
