package storage

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"time"
)

// GetID returns a random ID string of the format prefix_12timeChars12randomChars
func GetID(prefix string) (string, error) {
	b := make([]byte, 12)
	t := time.Now()
	binary.BigEndian.PutUint64(b[:], uint64(t.UnixNano()))
	if _, err := rand.Read(b[6:12]); err != nil {
		return "", err
	}

	return fmt.Sprintf("%s_%x", prefix, b), nil
}
