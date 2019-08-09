package test

import "testing"

func TestTest(t *testing.T) {
	True(t, true)
	False(t, false)
	EqualStr(t, "string of text", "string of text")
	EqualInt(t, 42, 42)
	DeepEqual(t, "string of text", "string of text")
	DeepEqual(t, []string{"a", "b"}, []string{"a", "b"})
	NoError(t, nil)
	Len(t, []string{"a", "b"}, 2)
	Len(t, []int{}, 0)
	Len(t, map[string]string{"a": "a", "b": "b"}, 2)
	Len(t, map[string]string{}, 0)

	IsNil(t, nil)
	IsNotNil(t, 5)
}
