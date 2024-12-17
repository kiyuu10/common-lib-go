package types

import (
	"fmt"
)

type String string

var _ fmt.Stringer = String("")

func (s String) String() string {
	return string(s)
}

func NewString(s string) String {
	return String(s)
}
