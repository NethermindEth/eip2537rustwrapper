package main

import (
	"encoding/hex"
	"testing"
)

func Test(t *testing.T) {
	input := make([]byte, 256)
	res, err := G1Add(input)
	t.Log(hex.EncodeToString(res))
	t.Log(err)
}
