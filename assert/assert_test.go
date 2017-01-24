package assert

import (
	"testing"
)

func TestTrue(t *testing.T) {
	a := 3 != 3
	True(t, a, "How can a=%T(%v) is not true!", a, a)
}

type II int

func TestEqual(t *testing.T) {
	a := 3
	b := int32(3)
	c := II(3)
	Equal(t, c, a, "You've messed up a=%T(%v), c=%T(%v)", a, a, c, c)
	Equal(t, b, a, "You've messed up a=%T(%v), b=%T(%v)", a, a, b, b)
	Equal(t, 2, a, "You've messed up a=%T(%v)", a, a)
}
