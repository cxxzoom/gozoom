package gocache

// ByteView save immutable []byte
type ByteView struct {
	b []byte
}

// Len return len of ByteView.b
func (b ByteView) Len() int {
	return len(b.b)
}

// CloneBytes return copy of B.b
func CloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}

// ByteSlice call CloneBytes
func (b ByteView) ByteSlice() []byte {
	return CloneBytes(b.b)
}

// String convert string and return
func (b ByteView) String() string {
	return string(b.b)
}
