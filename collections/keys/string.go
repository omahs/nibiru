package keys

import (
	"fmt"
)

// String converts any member of the string typeset into a StringKey
// NOTE(mercilex): this exists to avoid type errors in which bytes are being
// converted to a StringKey which is not correct behavior.
func String[T ~string](v T) StringKey {
	return StringKey(v)
}

type StringKey string

func (s StringKey) KeyBytes() []byte {
	if err := noNullTermination(s); err != nil {
		panic(fmt.Errorf("invalid StringKey: %w", err))
	}
	return append([]byte(s), 0x00) // null terminate it for safe prefixing
}

func (s StringKey) FromKeyBytes(b []byte) (int, Key) {
	l := len(b)
	if l < 2 {
		panic("invalid StringKey bytes")
	}
	for i, c := range b {
		if c == 0 {
			return i, StringKey(b[:i])
		}
	}
	panic(fmt.Errorf("StringKey is not null terminated: %s", s))
}

func (s StringKey) String() string {
	return string(s)
}
