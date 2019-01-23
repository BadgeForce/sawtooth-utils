package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/rberg2/sawtooth-go-sdk/processor"
)

// State ...
type State struct {
	Context *processor.Context
}

// NewStateInstance ...
func NewStateInstance(context *processor.Context) *State {
	return &State{context}
}

// ProofOfIntegrityHash ...
func ProofOfIntegrityHash(data []byte) string {
	return fmt.Sprintf("%x", md5.Sum(data))
}

// ToHex return hex encoded string
func ToHex(str string) string {
	return strings.ToLower(hex.EncodeToString([]byte(str)))
}

// FromHex return hex decoded string
func FromHex(str string) ([]byte, error) {
	return hex.DecodeString(str)
}
