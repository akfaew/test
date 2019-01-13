package test

import "testing"

func Test_Misc(t *testing.T) {
	True(t, true)
	False(t, false)
	DeepEqual(t, "string of text", "string of text")
	DeepEqual(t, []string{"a", "b"}, []string{"a", "b"})
	NoError(t, nil)
	Len(t, []string{"a", "b"}, 2)
	Len(t, []int{}, 0)
}
