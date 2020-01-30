package x16rv2

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestX16RV2(t *testing.T) {
	targetHex := "e3c18aa20b87fd4d56905e05c33d6cae7d3209653c21aae78f1f1b9cde818100"
	targetBytes, _ := hex.DecodeString(targetHex)
	headerHex := "04000000fedcba9876543210123456789abcdef0031d8f75ade0746ec80b7020000000000f33171b804978ce997aafd70e7daffc44fba61609538b77773e32a9b830a73ea732dd15baa220"
	headerBytes, _ := hex.DecodeString(headerHex)
	hash := X16RV2(headerBytes)
	for i, j := 0, len(hash)-1; i < j; i, j = i+1, j-1 {
		hash[i], hash[j] = hash[j], hash[i]
	}
	if !bytes.Equal(hash, targetBytes) {
		t.Fatalf("wrong hash result, got %s, should be %s", hex.EncodeToString(hash), targetHex)
	}
}
