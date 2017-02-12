package assert

import (
	"reflect"
	"testing"
)

func TestWriteTypeDiffValues(t *testing.T) {
	cs := []struct {
		v1, v2 interface{}
		s1, s2 string
	}{
		{true, false, H("true"), H("false")},
		{int(100), int(101), H("100"), H("101")},
		{int8(100), int8(101), H("100"), H("101")},
		{int16(100), int16(101), H("100"), H("101")},
		{int32(100), int32(101), H("100"), H("101")},
		{int64(100), int64(101), H("100"), H("101")},
		{uint(100), uint(101), H("100"), H("101")},
		{uint8(100), uint8(101), H("100"), H("101")},
		{uint16(100), uint16(101), H("100"), H("101")},
		{uint32(100), uint32(101), H("100"), H("101")},
		{uint64(100), uint64(101), H("100"), H("101")},
		{uintptr(100), uintptr(101), H("0x64"), H("0x65")},
		{float32(100.123), float32(101.123), H("100.123"), H("101.123")},
		{float64(100.123), float64(101.123), H("100.123"), H("101.123")},
		{complex64(100.25 + 200.5i), complex64(101.25 + 200.5i), "(\x1b[41m100.25\x1b[0m+200.5)", "(\x1b[41m101.25\x1b[0m+200.5)"},
		{complex64(100.25 + 200.5i), complex64(100.25 + 201.5i), "(100.25+\x1b[41m200.5\x1b[0m)", "(100.25+\x1b[41m201.5\x1b[0m)"},
		{complex64(100.25 + 200.5i), complex64(101.25 + 201.5i), "(\x1b[41m100.25\x1b[0m+\x1b[41m200.5\x1b[0m)", "(\x1b[41m101.25\x1b[0m+\x1b[41m201.5\x1b[0m)"},
		{complex128(100.25 + 200.5i), complex128(101.25 + 200.5i), "(\x1b[41m100.25\x1b[0m+200.5)", "(\x1b[41m101.25\x1b[0m+200.5)"},
		{complex128(100.25 + 200.5i), complex128(100.25 + 201.5i), "(100.25+\x1b[41m200.5\x1b[0m)", "(100.25+\x1b[41m201.5\x1b[0m)"},
		{complex128(100.25 + 200.5i), complex128(101.25 + 201.5i), "(\x1b[41m100.25\x1b[0m+\x1b[41m200.5\x1b[0m)", "(\x1b[41m101.25\x1b[0m+\x1b[41m201.5\x1b[0m)"},
		{nil, make(chan int), H("0x64"), H("0x65")},
	}
	for i, c := range cs {
		var d ValueDiffer
		d.writeTypeDiffValues(reflect.ValueOf(c.v1), reflect.ValueOf(c.v2))
		Equal(t, c.s1, d.String(0), "i=%v, r1=\n%v\n%v", i, d.String(0), d.String(1))
		Equal(t, c.s2, d.String(1), "i=%v, r2=\n%v\n%v", i, d.String(0), d.String(1))
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
