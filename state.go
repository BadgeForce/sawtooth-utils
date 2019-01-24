package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/hyperledger/sawtooth-sdk-go/signing"
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

// VerifyPOIHash ...
func VerifyPOIHash(b []byte, check string) (string, bool) {
	hash := ProofOfIntegrityHash(b)
	return hash, hash == check
}

// VerifySig ...
func VerifySig(signature, publicKey string, message []byte, hexEnc bool) bool {
	ctx := signing.NewSecp256k1Context()
	pub := signing.NewSecp256k1PublicKey([]byte(publicKey))
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

// HexToStr return hex decoded bytes
func HexToStr(str string) ([]byte, error) {
	return hex.DecodeString(str)
}

// HexToBytes return hex decoded bytes
func HexToBytes(b []byte) ([]byte, error) {
	return hex.DecodeString(string(b))
}
