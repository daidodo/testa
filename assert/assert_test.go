package assert

import (
	"testing"
)

func TestTrue(t *testing.T) {
	True(t, true)
	a := false
	True(t, a)
}

func TestFalse(t *testing.T) {
	False(t, false)
	a := 3 == 3
	False(t, a, "TestFalse: How can a is true!")
}

func TestEqualBoolTrue(t *testing.T) {
	a := true
	Equal(t, true, a)
	Equal(t, false, a, a)
}

func TestEqualBoolFalse(t *testing.T) {
	a := false
	Equal(t, false, a)
	Equal(t, true, a, a, true)
}

func TestEqualInt(t *testing.T) {
	var a int
	a = 100
	Equal(t, 100, a)
	Equal(t, 101, a, "TestEqualInt: a=%v is not 101", a)
}

func TestEqualIntRune(t *testing.T) {
	var a int
	a = 'A'
	Equal(t, 'A', a)
	Equal(t, 'B', a, "TestEqualIntRune: a=%v is not 'B'", a)
}

//func TestEqualInt8(t *testing.T) {
//    a := int8(100)
//    Equal(t, int8(100), a)
//    Equal(t, 101, a, "TestEqualInt8: a=%v is not 101", a)
//}

//type II int

//func TestEqual(t *testing.T) {
//    a := 3
//    b := int32(3)
//    c := II(3)
//    Equal(t, c, a, "You've messed up a=%T(%v), c=%T(%v)", a, a, c, c)
//    Equal(t, b, a, "You've messed up a=%T(%v), b=%T(%v)", a, a, b, b)
//    Equal(t, 2, a, "You've messed up a=%T(%v)", a, a)
//}
