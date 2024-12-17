package types

import (
	"encoding/base64"
	"encoding/hex"
	"math/big"
	"strings"

	"github.com/kiyuu10/common-lib-go/erroy"
)

type RawBytes []byte

func (b RawBytes) MarshalBinary() ([]byte, error) {
	return b, nil
}

func (b *RawBytes) UnmarshalBinary(input []byte) error {
	*b = input
	return nil
}

func (b RawBytes) MarshalText() ([]byte, error) {
	return b.MarshalBinary()
}

func (b *RawBytes) UnmarshalText(input []byte) error {
	return b.UnmarshalBinary(input)
}

func (b RawBytes) Bytes() []byte {
	return b
}

func hexBytesTrimPrefix(input string) string {
	if strings.HasPrefix(input, "0x") {
		return input[2:]
	} else {
		return input
	}
}

type HexBytes []byte

func (b HexBytes) MarshalText() ([]byte, error) {
	return []byte(b.String()), nil
}

func (b *HexBytes) UnmarshalText(input []byte) error {
	var inputStr = hexBytesTrimPrefix(string(input))
	data, err := hex.DecodeString(inputStr)
	if err != nil {
		return erroy.WrapStack(err, "hex bytes: decode text")
	}
	*b = data
	return nil
}

func (b HexBytes) Bytes() []byte {
	return b
}

func (b HexBytes) String() string {
	return hex.EncodeToString(b)
}

func (b HexBytes) BigInt() *big.Int {
	return new(big.Int).SetBytes(b)
}

type HexNumberBytes struct {
	HexBytes
}

func (b *HexNumberBytes) UnmarshalText(input []byte) error {
	inputStr := hexBytesTrimPrefix(string(input))
	if len(inputStr)&1 != 0 {
		inputStr = "0" + inputStr
	}
	return b.HexBytes.UnmarshalText([]byte(inputStr))
}

type Base64Bytes []byte

func (b Base64Bytes) MarshalText() ([]byte, error) {
	return []byte(b.String()), nil
}

func (b *Base64Bytes) UnmarshalText(input []byte) error {
	data, err := base64.StdEncoding.DecodeString(string(input))
	if err != nil {
		return erroy.WrapStack(err, "base64 bytes: decode text")
	}
	*b = data
	return nil
}

func (b Base64Bytes) Bytes() []byte {
	return b
}

func (b Base64Bytes) String() string {
	return base64.StdEncoding.EncodeToString(b)
}

func (b Base64Bytes) BigInt() *big.Int {
	return new(big.Int).SetBytes(b)
}
