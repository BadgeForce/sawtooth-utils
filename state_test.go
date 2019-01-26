package utils_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	utils "github.com/BadgeForce/sawtooth-utils"

	"github.com/rberg2/sawtooth-go-sdk/signing"
)

var (
	poiData, _ = json.Marshal(map[string]interface{}{"hello": "world"})
	poiHash    = "fbc24bcc7a1794758fc1327fcfebdaf6"

	cryptoCtx = signing.NewSecp256k1Context()
	privKey   = cryptoCtx.NewRandomPrivateKey()
	pubKey    = cryptoCtx.GetPublicKey(privKey)

	hwHex     = "68656c6c6f776f726c64"
	strForHex = "helloworld"
)

// TestProofOfIntegrityHash ...
func TestProofOfIntegrityHash(t *testing.T) {
	if hash := utils.ProofOfIntegrityHash(poiData); hash != poiHash {
		t.Errorf("invalid proof of integrity hash expected %s got %s", poiHash, hash)
	}
}

// TestVerifyPOIHash ...
func TestVerifyPOIHash(t *testing.T) {
	if hash, valid := utils.VerifyPOIHash(poiData, poiHash); !valid {
		t.Errorf("expected hash to be valid got %v", valid)
	} else if hash != poiHash {
		t.Errorf("hash mismatch expected %s got %s", poiHash, hash)
	}

	b, _ := json.Marshal(map[string]interface{}{"hello": "universe"})

	if hash, valid := utils.VerifyPOIHash(b, poiHash); valid {
		t.Errorf("expected hash to be invalid got %v", valid)
	} else if hash == poiHash {
		t.Errorf("expected different hashes hash1(%s) != hash2(%s)", hash, poiHash)
	}
}

// TestVerifySig ...
func TestVerifySig(t *testing.T) {
	message := []byte(fmt.Sprintf("%x", time.Now().Unix()))
	sig := fmt.Sprintf("%s", cryptoCtx.Sign(message, privKey))

	if valid := utils.VerifySig(string(sig), pubKey.AsBytes(), message, false); !valid {
		t.Error("expected signature validation to pass")
	}

	if valid := utils.VerifySig(fmt.Sprintf("%s%s", sig[0:len(sig)-1], "O"), pubKey.AsBytes(), message, false); valid {
		t.Error("expected signature validation to fail when invalid signature used")
	}

	fakePub := cryptoCtx.GetPublicKey(cryptoCtx.NewRandomPrivateKey())
	if valid := utils.VerifySig(string(sig), fakePub.AsBytes(), message, false); valid {
		t.Error("expected signature validation to fail when invalid publickey used")
	}

	if valid := utils.VerifySig(string(sig), pubKey.AsBytes(), append(message, []byte("foo")[0]), false); valid {
		t.Error("expected signature validation to fail when invalid message used")
	}
}

// TestStrToHex ...
func TestStrToHex(t *testing.T) {
	if hex := utils.StrToHex(strForHex); hex != hwHex {
		t.Errorf("invalid hex encoded string expected %s got %s", hwHex, hex)
	}
}

// TestBytesToHex ...
func TestBytesToHex(t *testing.T) {
	if hex := utils.BytesToHex([]byte(strForHex)); hex != hwHex {
		t.Errorf("invalid hex encoded string expected %s got %s", hwHex, hex)
	}
}

// TestDecodeHexStr ...
func TestDecodeHexStr(t *testing.T) {
	if str, err := utils.DecodeHexStr(hwHex); err != nil {
		t.Errorf("expected error to be nil got %s", err)
	} else if string(str) != strForHex {
		t.Errorf("invalid hex edcoded string expected %s got %s", strForHex, str)
	}
}

// TestDecodeHexBytes ...
func TestDecodeHexBytes(t *testing.T) {
	if str, err := utils.DecodeHexBytes([]byte(hwHex)); err != nil {
		t.Errorf("expected error to be nil got %s", err)
	} else if string(str) != strForHex {
		t.Errorf("invalid hex edcoded string expected %s got %s", strForHex, str)
	}
}
