package main

import (
	"errors"
	"fmt"
	"testing"

	"github.com/daidodo/testa/assert"
	assert2 "github.com/stretchr/testify/assert"
)

func TestZero3(t *testing.T) {
	var a chan int
	assert.NotZero(t, a)
	assert.Zero(t, a)
}

func TestZero2(t *testing.T) {
	a := int8(0)
	assert.Zero(t, a)
	assert.NotZero(t, a)
}

func TestZero(t *testing.T) {
	a := int8(8)
	assert.NotZero(t, a)
	assert.Zero(t, a)
}

func TestNotEqualType2(t *testing.T) {
	var a, b *map[string]int
	var c *map[string]uint
	assert.NotEqualType(t, c, a)
	assert.NotEqualType(t, b, a)
}

func TestNotEqualType(t *testing.T) {
	a := 100
	b := uint(100)
	assert.NotEqualType(t, b, a)
	assert.NotEqualType(t, 0, a)
}

func TestEqualType6(t *testing.T) {
	var a, b *map[string]int
	var c *map[string]uint
	assert.EqualType(t, b, a)
	assert.EqualType(t, c, a)
}

func TestEqualType5(t *testing.T) {
	var a, b map[string]int
	var c map[string]uint
	assert.EqualType(t, b, a)
	assert.EqualType(t, c, a)
}

func TestEqualType4(t *testing.T) {
	a := make(map[string]int)
	b := (*map[string]int)(nil)
	assert.EqualType(t, b, a)
}

func TestEqualType3(t *testing.T) {
	type Str string
	a := make(map[string]int)
	b := make(map[Str]bool)
	assert.EqualType(t, b, a)
}

func TestEqualType2(t *testing.T) {
	a := errors.New("abc")
	b := errors.New("abd")
	assert.EqualType(t, b, a)
	assert.EqualType(t, nil, a)
}

func TestEqualType(t *testing.T) {
	a := 100
	b := uint(100)
	assert.EqualType(t, 0, a)
	assert.EqualType(t, b, a)
}

func TestNotEqualValueError(t *testing.T) {
	e1 := errors.New("abc")
	e2 := errors.New("abc")
	e3 := errors.New("abd")
	assert.NotEqual(t, fmt.Sprintf("%p", e1), fmt.Sprintf("%p", e2))
	assert.NotEqualValue(t, e1, e3)
	assert.NotEqualValue(t, e1, e2)
}

func TestNotEqualValue(t *testing.T) {
	a := uint(100)
	assert.NotEqualValue(t, 101, a)
	assert.NotEqualValue(t, 100, a)
}

func TestEqualValuePointerInt2(t *testing.T) {
	a, b := new(int), new(uint)
	assert.Equal(t, a, b)
	assert.EqualValue(t, a, b)
	*a = 101
	assert.Equal(t, a, b)
	assert.EqualValue(t, a, b)
}

func TestEqualValuePointerInt(t *testing.T) {
	a, b := new(int), new(int)
	assert.Equal(t, a, b)
	assert.EqualValue(t, a, b)
	*a = 101
	assert.Equal(t, a, b)
	assert.EqualValue(t, a, b)
}

func TestEqualValueError(t *testing.T) {
	e1 := errors.New("abc")
	e2 := errors.New("abc")
	e3 := errors.New("abd")
	assert.NotEqual(t, fmt.Sprintf("%p", e1), fmt.Sprintf("%p", e2))
	assert.EqualValue(t, e1, e2)
	assert.EqualValue(t, e1, e3)
}

func TestEqualValBug(t *testing.T) {
	a := int8(0)
	b := int32(-1000000000)
	assert2.EqualValues(t, b, a) //This is a BUG!
	assert.EqualValue(t, b, a)
}

func TestEqualValue(t *testing.T) {
	a := uint(100)
	assert.EqualValue(t, 100, a)
	assert.EqualValue(t, 101, a)
}

func TestErrorPrintDiff(t *testing.T) {
	e1 := errors.New("abc")
	e2 := errors.New("abc")
	e3 := errors.New("abd")
	assert.Equal(t, e1, e2)
	assert.Equal(t, e1, e3)
}

func TestErrorPrint(t *testing.T) {
	e1 := errors.New("abc")
	assert.Equal(t, 100, e1)
}

func TestNotNil(t *testing.T) {
	assert.NotNil(t, make(chan int))
	var a chan int
	assert.NotEqual(t, nil, a)
	assert.NotNil(t, a)
}

func TestNil(t *testing.T) {
	var a chan int
	assert.Nil(t, a)
	assert.Nil(t, nil)
	assert.Equal(t, nil, a)
}

func TestNotEqual(t *testing.T) {
	a := 3
	assert.NotEqual(t, 3, a, "a should not be 3")
}

func TestFalse(t *testing.T) {
	assert.False(t, false)
	a := 3 == 3
	assert.False(t, a, "How can a is true!")
}

func TestTrue(t *testing.T) {
	assert.True(t, true)
	a := false
	assert.True(t, a)
}

func TestEqualChanRW(t *testing.T) {
	a := make(chan int)
	var b <-chan int = a
	assert.Equal(t, b, a, "Is a=%#v same as %#v", a, b)
}

func TestEqualChanDiffType(t *testing.T) {
	a := make(chan int)
	b := make(chan string)
	assert.Equal(t, b, a, "Is a=%#v same as %#v", a, b)
}

func TestEqualFuncNilDiffType(t *testing.T) {
	var a func(int) string
	var b func(string) int
	a = nil
	b = nil
	assert.Equal(t, b, a, "Is a=%#v same as %#v?", a, b)
}

func TestEqualInterfaceNil1(t *testing.T) {
	var a, b interface{}
	assert.Equal(t, b, a)
	i := 3
	a = &i
	assert.Equal(t, b, a, "Is a=%#v same as %v?", a, b)
}
func TestEqualStruct(t *testing.T) {
	a := struct {
		a interface{}
		b interface{}
	}{a: 1, b: "abc"}
	b := struct {
		a interface{}
		b interface{}
	}{a: 1, b: "abc"}
	assert.Equal(t, a, b)
	c := struct {
		a interface{}
		b interface{}
	}{a: 'A', b: "abcd"}
	assert.Equal(t, a, c)
}

func TestEqualMap2(t *testing.T) {
	a := map[int]string{}
	b := map[int]string{}
	assert.Equal(t, b, a)
	a[-1] = "aaggs"
	b[2] = "aaa"
	b[4] = "ert"
	b[5] = "crt"
	assert.Equal(t, b, a, "a=%#v is not b=%#v", a, b)
}

func TestEqualMap(t *testing.T) {
	a := map[int]string{1: "abc", 2: "cde", 3: "xyz"}
	b := map[int]string{1: "abc", 2: "cde", 3: "xyz"}
	assert.Equal(t, b, a)
	a[-1] = "aaggs"
	b[2] = "aaa"
	b[4] = "ert"
	b[5] = "crt"
	assert.Equal(t, b, a, "a=%#v is not b=%#v", a, b)
}

func TestEqualInterface(t *testing.T) {
	var a, b interface{}
	assert.Equal(t, b, a)
	a = int(3)
	b = a
	assert.Equal(t, b, a)
	b = int(4)
	assert.Equal(t, b, a, "Is a=%#v same as %v?", a, b)
}

func TestEqualFuncArrayLong(t *testing.T) {
	f := func(int) string { return "1" }
	g := func(int) string { return "2" }
	a := [11]func(int) string{f, g, nil, f, nil, f, g, f}
	b := a
	assert.Equal(t, b, a, "Is a same as b?")
}

func TestEqualFuncArray(t *testing.T) {
	f := func(int) string { return "1" }
	g := func(int) string { return "2" }
	a := [...]func(int) string{f, g, nil}
	b := a
	assert.Equal(t, b, a, "Is a=%#v same as %#v?", a, b)
}

func TestEqualFunc(t *testing.T) {
	a := func(int) string { return "1" }
	b := a
	assert.Equal(t, b, a, "Is a=%#v same as %v?", a, b)
}

func TestEqualFuncNil2(t *testing.T) {
	var a, b func(int) string
	a = nil
	b = nil
	assert.Equal(t, b, a)
	b = func(int) string { return "1" }
	assert.Equal(t, b, a, "Is a=%#v same as %v?", a, b)
}

func TestEqualFuncNil1(t *testing.T) {
	var a, b func(int) string
	a = func(int) string { return "1" }
	b = nil
	assert.Equal(t, b, a, "Is a=%#v same as %v?", a, b)
}

func TestEqualChanRO(t *testing.T) {
	a := make(<-chan int)
	b := a
	assert.Equal(t, b, a)
	c := make(chan int)
	assert.Equal(t, c, a, "Is a=%#v same as %#v", a, c)
}

func TestEqualChan(t *testing.T) {
	a := make(chan int)
	b := a
	assert.Equal(t, b, a)
	b = make(chan int)
	assert.Equal(t, b, a, "Is a=%#v same as %#v", a, b)
}

//func TestEqualArrayLong3Dense(t *testing.T) {
//    var a, b [11][5][11]int
//    for i := range a {
//        for j := range a[0] {
//            for k := range a[0][0] {
//                a[i][j][k] = 0 + i + j*j + k*k*k
//                b[i][j][k] = 0 + i*i + j*j*j + k
//            }
//        }
//    }
//    assert.Equal(t, b, a, "a is not b")
//}

func TestEqualArrayShort3Dense(t *testing.T) {
	var a, b [5][5][5]int
	for i := range a {
		for j := range a[0] {
			for k := range a[0][0] {
				a[i][j][k] = 1 + i + j*j + k*k*k
				b[i][j][k] = 1 + i*i + j*j*j + k
			}
		}
	}
	assert.Equal(t, b, a, "a is not b")
}

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
	assert.Equal(t, b, a)
	b[1][1][1]++
	b[len(b)/2][len(b[0])/2][len(b[0][0])/3]++
	b[len(b)-1][len(b[0])-1][len(b[0][0])-2]++
	b[len(b)-1][len(b[0])-1][len(b[0][0])-1]++
	assert.Equal(t, b, a, "a is not b")
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
	assert.Equal(t, b, a)
	b[1][1][1]++
	b[len(b)/2][len(b[0])/2][len(b[0][0])/3]++
	b[len(b)-1][len(b[0])-1][len(b[0][0])-2]++
	b[len(b)-1][len(b[0])-1][len(b[0][0])-1]++
	assert.Equal(t, b, a, "a is not b")
}

func TestEqualArrayLong2(t *testing.T) {
	var a [14][14]int
	for i := range a {
		for j := range a[0] {
			a[i][j] = 100 + i + j*j
		}
	}
	b := a
	assert.Equal(t, b, a)
	b[1][1]++
	b[len(b)/2][len(b[0])/2]++
	b[len(b)-1][len(b[0])-2]++
	b[len(b)-1][len(b[0])-1]++
	assert.Equal(t, b, a, "a is not b")
}

func TestEqualArrayShort2(t *testing.T) {
	var a [10][10]int
	for i := range a {
		for j := range a[0] {
			a[i][j] = 100 + i + j*j
		}
	}
	b := a
	assert.Equal(t, b, a)
	b[1][1]++
	b[len(b)/2][len(b[0])/2]++
	b[len(b)-1][len(b[0])-2]++
	b[len(b)-1][len(b[0])-1]++
	assert.Equal(t, b, a, "a is not b")
}

func TestEqualArrayLong1(t *testing.T) {
	var a [14]int
	for i := range a {
		a[i] = 100 + i*i
	}
	b := a
	assert.Equal(t, b, a)
	b[1]++
	b[len(b)/2]++
	b[len(b)-2]++
	b[len(b)-1]++
	assert.Equal(t, b, a, "a=%#v is not %#v", a, b)
}

func TestEqualArrayShort1(t *testing.T) {
	var a [10]int
	for i := range a {
		a[i] = 100 + i
	}
	b := a
	assert.Equal(t, b, a)
	b[1]++
	b[len(b)/2]++
	b[len(b)-2]++
	b[len(b)-1]++
	assert.Equal(t, b, a, "a=%#v is not %#v", a, b)
}

func TestEqualComplex128(t *testing.T) {
	a := complex(float64(100.123), float64(12.345))
	b := a
	assert.Equal(t, b, a)
	b = complex(real(a)+.01, imag(b)+.1)
	assert.Equal(t, b, a, "a=%T%#v is not %T%#v", a, a, b, b)
}

func TestEqualComplex128Imag(t *testing.T) {
	a := complex(float64(100.123), float64(12.345))
	b := a
	assert.Equal(t, b, a)
	b = complex(real(a), imag(b)+.1)
	assert.Equal(t, b, a, "a=%T%#v is not %T%#v", a, a, b, b)
}

func TestEqualComplex128Real(t *testing.T) {
	a := complex(float64(100.123), float64(12.345))
	b := a
	assert.Equal(t, b, a)
	b = complex(real(a)+.1, imag(b))
	assert.Equal(t, b, a, "a=%T%#v is not %T%#v", a, a, b, b)
}

func TestEqualComplex64(t *testing.T) {
	a := complex(float32(100.123), float32(12.345))
	b := a
	assert.Equal(t, b, a)
	b = complex(real(a)+.01, imag(b)+.1)
	assert.Equal(t, b, a, "a=%T%v is not %T%v", a, a, b, b)
}

func TestEqualComplex64Imag(t *testing.T) {
	a := complex(float32(100.123), float32(12.345))
	b := a
	assert.Equal(t, b, a)
	b = complex(real(a), imag(b)+.1)
	assert.Equal(t, b, a, "a=%T%v is not %T%v", a, a, b, b)
}

func TestEqualComplex64Real(t *testing.T) {
	a := complex(float32(100.123), float32(12.345))
	b := a
	assert.Equal(t, b, a)
	b = complex(real(a)+.1, imag(b))
	assert.Equal(t, b, a, "a=%T%v is not %T%v", a, a, b, b)
}

func TestEqualFloat64(t *testing.T) {
	a := float64(100.123)
	b := a
	assert.Equal(t, b, a)
	b++
	assert.Equal(t, b, a, "a=%T(%v) is not %T(%v)", a, a, b, b)
}

func TestEqualFloat32(t *testing.T) {
	a := float32(100.123)
	b := a
	assert.Equal(t, b, a)
	b++
	assert.Equal(t, b, a, "a=%T(%v) is not %T(%v)", a, a, b, b)
}

func TestEqualUIntptr(t *testing.T) {
	a := uintptr(100)
	b := a
	assert.Equal(t, b, a)
	b++
	assert.Equal(t, b, a, "a=%T(%#v) is not %T(%#v)", a, a, b, b)
}

func TestEqualUInt64(t *testing.T) {
	a := uint64(100)
	b := a
	assert.Equal(t, b, a)
	b++
	assert.Equal(t, b, a, "a=%T(%v) is not %T(%v)", a, a, b, b)
}

func TestEqualUInt32(t *testing.T) {
	a := uint32(100)
	b := a
	assert.Equal(t, b, a)
	b++
	assert.Equal(t, b, a, "a=%T(%v) is not %T(%v)", a, a, b, b)
}

func TestEqualUInt16(t *testing.T) {
	a := uint16(100)
	b := a
	assert.Equal(t, b, a)
	b++
	assert.Equal(t, b, a, "a=%T(%v) is not %T(%v)", a, a, b, b)
}

func TestEqualUInt8(t *testing.T) {
	a := uint8(100)
	b := a
	assert.Equal(t, b, a)
	b++
	assert.Equal(t, b, a, "a=%T(%v) is not %T(%v)", a, a, b, b)
}

func TestEqualUInt(t *testing.T) {
	a := uint(100)
	b := a
	assert.Equal(t, b, a)
	b++
	assert.Equal(t, b, a, "a=%T(%v) is not %T(%v)", a, a, b, b)
}

func TestEqualInt64(t *testing.T) {
	a := int64(100)
	b := a
	assert.Equal(t, b, a)
	b++
	assert.Equal(t, b, a, "a=%T(%v) is not %T(%v)", a, a, b, b)
}

func TestEqualInt32(t *testing.T) {
	a := int32(100)
	b := a
	assert.Equal(t, b, a)
	b++
	assert.Equal(t, b, a, "a=%T(%v) is not %T(%v)", a, a, b, b)
}

func TestEqualInt16(t *testing.T) {
	a := int16(100)
	b := a
	assert.Equal(t, b, a)
	b++
	assert.Equal(t, b, a, "a=%T(%v) is not %T(%v)", a, a, b, b)
}

func TestEqualInt8(t *testing.T) {
	a := int8(100)
	b := a
	assert.Equal(t, b, a)
	b++
	assert.Equal(t, b, a, "a=%T(%v) is not %T(%v)", a, a, b, b)
}

func TestEqualInt(t *testing.T) {
	a := int(100)
	b := a
	assert.Equal(t, b, a)
	b++
	assert.Equal(t, b, a, "a=%T(%v) is not %T(%v)", a, a, b, b)
}

func TestEqualBoolFalse(t *testing.T) {
	a := false
	assert.Equal(t, false, a)
	assert.Equal(t, true, a, a, true)
}

func TestEqualBoolTrue(t *testing.T) {
	a := true
	assert.Equal(t, true, a)
	assert.Equal(t, false, a, a)
}

func myTest(t *testing.T, e, a int) {
	assert.Caller(1).Equal(t, e, a)
}

func TestA(t *testing.T) {
	myTest(t, 1, 1)
	myTest(t, 10, 10)
	myTest(t, 100, 101)
}

func TestSlice2(t *testing.T) {
	a := []int{1, 21, 3, 4, 5}
	b := a
	assert.Equal(t, b, a)
	c := []int{1, 2, 3, 4, 6, 7, 8, 10, 11, 12, 13}
	assert.Equal(t, c, a, "a=%T(%v) is not %T(%v)", a, a, c, c)
}

func TestSlice(t *testing.T) {
	a := []int{1, 21, 3, 4, 5}
	b := a
	assert.Equal(t, b, a)
	c := []int{1, 2, 3, 4, 6, 7, 8, 10, 11, 12}
	assert.Equal(t, c, a, "a=%T(%v) is not %T(%v)", a, a, c, c)
}

func TestArray(t *testing.T) {
	a := [...]int{1, 21, 3, 4, 5}
	b := a
	assert.Equal(t, b, a)
	c := [...]int{1, 2, 3, 4, 6}
	assert.Equal(t, c, a, "a=%T(%v) is not %T(%v)", a, a, c, c)
}

func TestInt(t *testing.T) {
	var a int
	a = 'A'
	fmt.Printf("\t%v\n", a)
	assert.Equal(t, 'A', a, "a=%v is not 'A'", a)
}

func TestMap(t *testing.T) {
	m1 := make(map[int]string)
	m2 := make(map[uint]string)
	var s1, s2 []rune
	for i := 0; i < 3; i++ {
		s1 = append(s1, rune('a'+i))
		s2 = append(s2, rune('b'+i))
		m1[i] = string(s1)
		m2[uint(i)] = string(s2)
	}
	assert.Equal(t, m1, m2)
}
