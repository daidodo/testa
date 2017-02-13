package assert

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestWriteTypeDiffValues(t *testing.T) {
	a := func() int { return 1 }
	cs := []struct {
		v1, v2         reflect.Value
		s1, s2         string
		n1, n2, om, cf bool
	}{
		{v1: reflect.ValueOf(true), v2: reflect.ValueOf(false)},
		{v1: reflect.ValueOf(int(100)), v2: reflect.ValueOf(int(101))},
		{v1: reflect.ValueOf(int8(100)), v2: reflect.ValueOf(int8(101))},
		{v1: reflect.ValueOf(int16(100)), v2: reflect.ValueOf(int16(101))},
		{v1: reflect.ValueOf(int32(100)), v2: reflect.ValueOf(int32(101))},
		{v1: reflect.ValueOf(int64(100)), v2: reflect.ValueOf(int64(101))},
		{v1: reflect.ValueOf(uint(100)), v2: reflect.ValueOf(uint(101))},
		{v1: reflect.ValueOf(uint8(100)), v2: reflect.ValueOf(uint8(101))},
		{v1: reflect.ValueOf(uint16(100)), v2: reflect.ValueOf(uint16(101))},
		{v1: reflect.ValueOf(uint32(100)), v2: reflect.ValueOf(uint32(101))},
		{v1: reflect.ValueOf(uint64(100)), v2: reflect.ValueOf(uint64(101))},
		{v1: reflect.ValueOf(uintptr(100)), v2: reflect.ValueOf(uintptr(101)), s1: H("0x64"), s2: H("0x65")},
		{v1: reflect.ValueOf(float32(100.25)), v2: reflect.ValueOf(float32(101.25))},
		{v1: reflect.ValueOf(float64(100.25)), v2: reflect.ValueOf(float64(101.25))},
		{v1: reflect.ValueOf(complex64(100.25 + 200.5i)), v2: reflect.ValueOf(complex64(101.25 + 200.5i)), s1: "(" + H("100.25") + "+200.5)", s2: "(" + H("101.25") + "+200.5)"},
		{v1: reflect.ValueOf(complex64(100.25 + 200.5i)), v2: reflect.ValueOf(complex64(100.25 + 201.5i)), s1: "(100.25+" + H("200.5") + ")", s2: "(100.25+" + H("201.5") + ")"},
		{v1: reflect.ValueOf(complex64(100.25 + 200.5i)), v2: reflect.ValueOf(complex64(101.25 + 201.5i)), s1: "(\x1b[41m100.25\x1b[0m+\x1b[41m200.5\x1b[0m)", s2: "(\x1b[41m101.25\x1b[0m+\x1b[41m201.5\x1b[0m)"},
		{v1: reflect.ValueOf(complex128(100.25 + 200.5i)), v2: reflect.ValueOf(complex128(101.25 + 200.5i)), s1: "(" + H("100.25") + "+200.5)", s2: "(" + H("101.25") + "+200.5)"},
		{v1: reflect.ValueOf(complex128(100.25 + 200.5i)), v2: reflect.ValueOf(complex128(100.25 + 201.5i)), s1: "(100.25+" + H("200.5") + ")", s2: "(100.25+" + H("201.5") + ")"},
		{v1: reflect.ValueOf(complex128(100.25 + 200.5i)), v2: reflect.ValueOf(complex128(101.25 + 201.5i)), s1: "(\x1b[41m100.25\x1b[0m+\x1b[41m200.5\x1b[0m)", s2: "(\x1b[41m101.25\x1b[0m+\x1b[41m201.5\x1b[0m)"},
		{v1: reflect.ValueOf(chan int(nil)), v2: reflect.ValueOf(make(chan int))},
		{v1: reflect.ValueOf(make(chan int)), v2: reflect.ValueOf(make(chan int))},
		{v1: reflect.ValueOf((func() int)(nil)), v2: reflect.ValueOf(func() int { return 1 })},
		{v1: reflect.ValueOf(func() int { return 2 }), v2: reflect.ValueOf(func() int { return 1 })},
		{v1: reflect.ValueOf(a), v2: reflect.ValueOf(a), cf: true},
		{v1: reflect.ValueOf((*int)(nil)), v2: reflect.ValueOf(new(int))},
		{v1: reflect.ValueOf(new(int)), v2: reflect.ValueOf(new(int))},
		{v1: reflect.ValueOf(unsafe.Pointer(nil)), v2: reflect.ValueOf(unsafe.Pointer(new(int)))},
		{v1: reflect.ValueOf(unsafe.Pointer(new(int))), v2: reflect.ValueOf(unsafe.Pointer(new(int)))},
		{v1: reflect.ValueOf(A{}).Field(0), v2: reflect.ValueOf(A{a: 100}).Field(0), s2: "int(100)"},
		{v1: reflect.ValueOf(A{}).Field(1), v2: reflect.ValueOf(A{b: A{}}).Field(1), s2: "assert.A{a:<nil>, b:<nil>}"},
		{v1: reflect.ValueOf(A{b: A{a: 100}}).Field(1), v2: reflect.ValueOf(A{b: A{a: 100}}).Field(1), s2: "assert.A{a:<nil>, b:<nil>}"},
	}
	for i, c := range cs {
		f := func(v1, v2 reflect.Value, s1, s2 string) {
			var d ValueDiffer
			d.writeTypeDiffValues(v1, v2)
			if s1 == "" {
				s1 = fmt.Sprintf(H("%v"), v1)
			}
			if s2 == "" {
				s2 = fmt.Sprintf(H("%v"), v2)
			}
			Caller(1).Equal(t, s1, d.String(0), "i=%v, r1=\n%v\n%v", i, d.String(0), d.String(1))
			Caller(1).Equal(t, s2, d.String(1), "i=%v, r2=\n%v\n%v", i, d.String(0), d.String(1))
			Caller(1).Equal(t, c.n1, d.Attrs[NewLine], "i=%v, NewLine0: Attrs=%v", i, d.Attrs)
			Caller(1).Equal(t, c.n2, d.Attrs[NewLine+1], "i=%v, NewLine1: Attrs=%v", i, d.Attrs)
			Caller(1).Equal(t, c.om, d.Attrs[OmitSame], "i=%v, OmitSame: Attrs=%v", i, d.Attrs)
			Caller(1).Equal(t, c.cf, d.Attrs[CompFunc], "i=%v, CompFunc: Attrs=%v", i, d.Attrs)
		}
		f(c.v1, c.v2, c.s1, c.s2)
		f(c.v2, c.v1, c.s2, c.s1)
	}
}

func TestWriteTypeDiffValuesString(t *testing.T) {
	eq := func(x1, x2 string, s1, s2 string) {
		var d ValueDiffer
		d.writeDiffValuesString(reflect.ValueOf(x1), reflect.ValueOf(x2))
		Caller(1).Equal(t, s1, d.String(0), "%s\n%s", d.String(0), d.String(1))
		Caller(1).Equal(t, s2, d.String(1), "%s\n%s", d.String(0), d.String(1))
	}
	eq("abcaadef", "accaa", "\"a\x1b[41mb\x1b[0mcaa\x1b[41mdef\x1b[0m\"", "\"a\x1b[41mc\x1b[0mcaa\"")
	eq("This is\x83这是 Chinese 中文！", "This is    这不是Chinase 汉文吗?", "\"This is\x1b[41m\\x83\x1b[0m这\x1b[41m是 \x1b[0mChin\x1b[41me\x1b[0mse \x1b[41m中\x1b[0m文\x1b[41m！\x1b[0m\"", "\"This is\x1b[41m    \x1b[0m这\x1b[41m不是\x1b[0mChin\x1b[41ma\x1b[0mse \x1b[41m汉\x1b[0m文\x1b[41m吗?\x1b[0m\"")
}
