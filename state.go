package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/rberg2/sawtooth-go-sdk/processor"
	"github.com/rberg2/sawtooth-go-sdk/signing"
)

// State ...
type State struct {
	Context *processor.Context
}

// NewStateInstance ...
func NewStateInstance(context *processor.Context) *State {
	return &State{context}
}

// ProofOfIntegrityHash create a proof of intergrity hash
func ProofOfIntegrityHash(data []byte) string {
	return fmt.Sprintf("%x", md5.Sum(data))
}

// VerifyPOIHash verify proof of integrity hash provided for some data is valid
func VerifyPOIHash(b []byte, check string) (string, bool) {
	hash := ProofOfIntegrityHash(b)
	return hash, hash == check
}

// VerifySig verify a signed message
func VerifySig(signature string, publicKey, message []byte, hexEnc bool) bool {
	ctx := signing.NewSecp256k1Context()
	pub := signing.NewSecp256k1PublicKey(publicKey)
	sig := []byte(signature)
	msg := encodeMsg(message, hexEnc)
	return ctx.Verify(sig, msg, pub)
}

func encodeMsg(message []byte, hexEnc bool) []byte {
	if !hexEnc {
		return message
	}
	encoded := BytesToHex(message)
	return []byte(encoded)
}

// StrToHex return hex encoded string
func StrToHex(str string) string {
	return strings.ToLower(hex.EncodeToString([]byte(str)))
}

// BytesToHex return hex encoded string
func BytesToHex(b []byte) string {
	return strings.ToLower(hex.EncodeToString(b))
}

// DecodeHexStr return hex decoded bytes
func DecodeHexStr(str string) ([]byte, error) {
	return hex.DecodeString(str)
}

// DecodeHexBytes return hex decoded bytes
func DecodeHexBytes(b []byte) ([]byte, error) {
	return hex.DecodeString(string(b))
}
