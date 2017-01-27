package assert

import (
	"testing"
)

func TestEqualArrayLong3(t *testing.T) {
	var a [15][15][15]int
	for i := range a {
		for j := range a[0] {
			for k := range a[0][0] {
				a[i][j][k] = 100 + i + j*j + k*k*k
			}
		}
	}
	b := a
	Equal(t, b, a)
	b[1][1][1]++
	b[len(b)/2][len(b[0])/2][len(b[0][0])/3]++
	b[len(b)-1][len(b[0])-1][len(b[0][0])-2]++
	b[len(b)-1][len(b[0])-1][len(b[0][0])-1]++
	Equal(t, b, a, "a is not b")
}

func TestEqualArrayShort3(t *testing.T) {
	var a [5][5][5]int
	for i := range a {
		for j := range a[0] {
			for k := range a[0][0] {
				a[i][j][k] = 100 + i + j*j + k*k*k
			}
		}
	}
	b := a
	Equal(t, b, a)
	b[1][1][1]++
	b[len(b)/2][len(b[0])/2][len(b[0][0])/3]++
	b[len(b)-1][len(b[0])-1][len(b[0][0])-2]++
	b[len(b)-1][len(b[0])-1][len(b[0][0])-1]++
	Equal(t, b, a, "a is not b")
}

func TestEqualArrayLong2(t *testing.T) {
	var a [14][14]int
	for i := range a {
		for j := range a[0] {
			a[i][j] = 100 + i + j*j
		}
	}
	b := a
	Equal(t, b, a)
	b[1][1]++
	b[len(b)/2][len(b[0])/2]++
	b[len(b)-1][len(b[0])-2]++
	b[len(b)-1][len(b[0])-1]++
	Equal(t, b, a, "a is not b")
}

func TestEqualArrayShort2(t *testing.T) {
	var a [10][10]int
	for i := range a {
		for j := range a[0] {
			a[i][j] = 100 + i + j*j
		}
	}
	b := a
	Equal(t, b, a)
	b[1][1]++
	b[len(b)/2][len(b[0])/2]++
	b[len(b)-1][len(b[0])-2]++
	b[len(b)-1][len(b[0])-1]++
	Equal(t, b, a, "a is not b")
}

func TestEqualArrayLong1(t *testing.T) {
	var a [14]int
	for i := range a {
		a[i] = 100 + i*i
	}
	b := a
	Equal(t, b, a)
	b[1]++
	b[len(b)/2]++
	b[len(b)-2]++
	b[len(b)-1]++
	Equal(t, b, a, "a=%#v is not %#v", a, b)
}

func TestEqualArrayShort1(t *testing.T) {
	var a [10]int
	for i := range a {
		a[i] = 100 + i
	}
	b := a
	Equal(t, b, a)
	b[1]++
	b[len(b)/2]++
	b[len(b)-2]++
	b[len(b)-1]++
	Equal(t, b, a, "a=%#v is not %#v", a, b)
}

func TestEqualComplex128(t *testing.T) {
	a := complex(float64(100.123), float64(12.345))
	b := a
	Equal(t, b, a)
	b = complex(real(a)+.01, imag(b)+.1)
	Equal(t, b, a, "a=%T%#v is not %T%#v", a, a, b, b)
}

func TestEqualComplex128Imag(t *testing.T) {
	a := complex(float64(100.123), float64(12.345))
	b := a
	Equal(t, b, a)
	b = complex(real(a), imag(b)+.1)
	Equal(t, b, a, "a=%T%#v is not %T%#v", a, a, b, b)
}

func TestEqualComplex128Real(t *testing.T) {
	a := complex(float64(100.123), float64(12.345))
	b := a
	Equal(t, b, a)
	b = complex(real(a)+.1, imag(b))
	Equal(t, b, a, "a=%T%#v is not %T%#v", a, a, b, b)
}

func TestEqualComplex64(t *testing.T) {
	a := complex(float32(100.123), float32(12.345))
	b := a
	Equal(t, b, a)
	b = complex(real(a)+.01, imag(b)+.1)
	Equal(t, b, a, "a=%T%v is not %T%v", a, a, b, b)
}

func TestEqualComplex64Imag(t *testing.T) {
	a := complex(float32(100.123), float32(12.345))
	b := a
	Equal(t, b, a)
	b = complex(real(a), imag(b)+.1)
	Equal(t, b, a, "a=%T%v is not %T%v", a, a, b, b)
}

func TestEqualComplex64Real(t *testing.T) {
	a := complex(float32(100.123), float32(12.345))
	b := a
	Equal(t, b, a)
	b = complex(real(a)+.1, imag(b))
	Equal(t, b, a, "a=%T%v is not %T%v", a, a, b, b)
}

func TestEqualFloat64(t *testing.T) {
	a := float64(100.123)
	b := a
	Equal(t, b, a)
	b++
	Equal(t, b, a, "a=%T(%v) is not %T(%v)", a, a, b, b)
}

func TestEqualFloat32(t *testing.T) {
	a := float32(100.123)
	b := a
	Equal(t, b, a)
	b++
	Equal(t, b, a, "a=%T(%v) is not %T(%v)", a, a, b, b)
}

func TestEqualUIntptr(t *testing.T) {
	a := uintptr(100)
	b := a
	Equal(t, b, a)
	b++
	Equal(t, b, a, "a=%T(%#v) is not %T(%#v)", a, a, b, b)
}

func TestEqualUInt64(t *testing.T) {
	a := uint64(100)
	b := a
	Equal(t, b, a)
	b++
	Equal(t, b, a, "a=%T(%v) is not %T(%v)", a, a, b, b)
}

func TestEqualUInt32(t *testing.T) {
	a := uint32(100)
	b := a
	Equal(t, b, a)
	b++
	Equal(t, b, a, "a=%T(%v) is not %T(%v)", a, a, b, b)
}

func TestEqualUInt16(t *testing.T) {
	a := uint16(100)
	b := a
	Equal(t, b, a)
	b++
	Equal(t, b, a, "a=%T(%v) is not %T(%v)", a, a, b, b)
}

func TestEqualUInt8(t *testing.T) {
	a := uint8(100)
	b := a
	Equal(t, b, a)
	b++
	Equal(t, b, a, "a=%T(%v) is not %T(%v)", a, a, b, b)
}

func TestEqualUInt(t *testing.T) {
	a := uint(100)
	b := a
	Equal(t, b, a)
	b++
	Equal(t, b, a, "a=%T(%v) is not %T(%v)", a, a, b, b)
}

func TestEqualInt64(t *testing.T) {
	a := int64(100)
	b := a
	Equal(t, b, a)
	b++
	Equal(t, b, a, "a=%T(%v) is not %T(%v)", a, a, b, b)
}

func TestEqualInt32(t *testing.T) {
	a := int32(100)
	b := a
	Equal(t, b, a)
	b++
	Equal(t, b, a, "a=%T(%v) is not %T(%v)", a, a, b, b)
}

func TestEqualInt16(t *testing.T) {
	a := int16(100)
	b := a
	Equal(t, b, a)
	b++
	Equal(t, b, a, "a=%T(%v) is not %T(%v)", a, a, b, b)
}

func TestEqualInt8(t *testing.T) {
	a := int8(100)
	b := a
	Equal(t, b, a)
	b++
	Equal(t, b, a, "a=%T(%v) is not %T(%v)", a, a, b, b)
}

func TestEqualInt(t *testing.T) {
	a := int(100)
	b := a
	Equal(t, b, a)
	b++
	Equal(t, b, a, "a=%T(%v) is not %T(%v)", a, a, b, b)
}

func TestEqualBoolFalse(t *testing.T) {
	a := false
	Equal(t, false, a)
	Equal(t, true, a, a, true)
}

func TestEqualBoolTrue(t *testing.T) {
	a := true
	Equal(t, true, a)
	Equal(t, false, a, a)
}

func TestFalse(t *testing.T) {
	False(t, false)
	a := 3 == 3
	False(t, a, "How can a is true!")
}

func TestTrue(t *testing.T) {
	True(t, true)
	a := false
	True(t, a)
}
