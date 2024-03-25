package lru

import "testing"

type String string

func (s String) Len() int {
	return len(s)
}

func TestCache_Get(t *testing.T) {
	l := New(int64(0), nil)
	l.Add("key1", String("1234"))
	// TODO
}
