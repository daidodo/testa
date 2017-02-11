package assert

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type I interface {
	Fun()
}

type A struct {
}

func H(s string) string {
	if len(s) < 1 {
		return ""
	}
	return kRED + s + kEND
}

func TestInterfaceName(t *testing.T) {
	type I interface {
		Fun()
	}
	v := reflect.ValueOf(struct {
		a interface{}
		b I
	}{})
	Equal(t, "", interfaceName(nil))
	Equal(t, "", interfaceName(v.Field(0).Type()))
	Equal(t, "assert.I", interfaceName(v.Field(1).Type()))
}

func TestStructName(t *testing.T) {
	v := reflect.ValueOf(struct{}{})
	Equal(t, "", structName(nil))
	Equal(t, "struct", structName(v.Type()))
	Equal(t, "assert.A", structName(reflect.TypeOf(A{})))
}

func TestIsReference(t *testing.T) {
	False(t, isReference(nil))
	False(t, isReference(reflect.TypeOf(true)))
	False(t, isReference(reflect.TypeOf(int(100))))
	False(t, isReference(reflect.TypeOf(int8(100))))
	False(t, isReference(reflect.TypeOf(int16(100))))
	False(t, isReference(reflect.TypeOf(int32(100))))
	False(t, isReference(reflect.TypeOf(int64(100))))
	False(t, isReference(reflect.TypeOf(uint(100))))
	False(t, isReference(reflect.TypeOf(uint8(100))))
	False(t, isReference(reflect.TypeOf(uint16(100))))
	False(t, isReference(reflect.TypeOf(uint32(100))))
	False(t, isReference(reflect.TypeOf(uint64(100))))
	False(t, isReference(reflect.TypeOf(uintptr(100))))
	False(t, isReference(reflect.TypeOf(float32(100))))
	False(t, isReference(reflect.TypeOf(float64(100))))
	False(t, isReference(reflect.TypeOf(complex64(100))))
	False(t, isReference(reflect.TypeOf(complex128(100))))
	False(t, isReference(reflect.TypeOf(string("abc"))))
	True(t, isReference(reflect.TypeOf(make(chan int))))
	True(t, isReference(reflect.TypeOf(func() {})))
	True(t, isReference(reflect.TypeOf(new(int))))
	True(t, isReference(reflect.TypeOf(unsafe.Pointer(new(int)))))
	True(t, isReference(reflect.ValueOf(struct{ a interface{} }{}).Field(0).Type()))
	True(t, isReference(reflect.TypeOf([...]int{1, 2, 3})))
	True(t, isReference(reflect.TypeOf([]int{1, 2, 3})))
	True(t, isReference(reflect.TypeOf(map[int]bool{100: true})))
	True(t, isReference(reflect.TypeOf(A{})))
}

func TestIsNonTrivial(t *testing.T) {
	False(t, isNonTrivial(nil))
	False(t, isNonTrivial(reflect.TypeOf(true)))
	False(t, isNonTrivial(reflect.TypeOf(int(100))))
	False(t, isNonTrivial(reflect.TypeOf(int8(100))))
	False(t, isNonTrivial(reflect.TypeOf(int16(100))))
	False(t, isNonTrivial(reflect.TypeOf(int32(100))))
	False(t, isNonTrivial(reflect.TypeOf(int64(100))))
	False(t, isNonTrivial(reflect.TypeOf(uint(100))))
	False(t, isNonTrivial(reflect.TypeOf(uint8(100))))
	False(t, isNonTrivial(reflect.TypeOf(uint16(100))))
	False(t, isNonTrivial(reflect.TypeOf(uint32(100))))
	False(t, isNonTrivial(reflect.TypeOf(uint64(100))))
	False(t, isNonTrivial(reflect.TypeOf(uintptr(100))))
	False(t, isNonTrivial(reflect.TypeOf(float32(100))))
	False(t, isNonTrivial(reflect.TypeOf(float64(100))))
	False(t, isNonTrivial(reflect.TypeOf(complex64(100))))
	False(t, isNonTrivial(reflect.TypeOf(complex128(100))))
	False(t, isNonTrivial(reflect.TypeOf(string("abc"))))
	False(t, isNonTrivial(reflect.TypeOf(make(chan int))))
	False(t, isNonTrivial(reflect.TypeOf(func() {})))
	False(t, isNonTrivial(reflect.TypeOf(new(int))))
	False(t, isNonTrivial(reflect.TypeOf(unsafe.Pointer(new(int)))))
	True(t, isNonTrivial(reflect.ValueOf(struct{ a interface{} }{}).Field(0).Type()))
	True(t, isNonTrivial(reflect.TypeOf([...]int{1, 2, 3})))
	True(t, isNonTrivial(reflect.TypeOf([]int{1, 2, 3})))
	True(t, isNonTrivial(reflect.TypeOf(map[int]bool{100: true})))
	True(t, isNonTrivial(reflect.TypeOf(A{})))
}

func TestIsComposite(t *testing.T) {
	False(t, isComposite(nil))
	False(t, isComposite(reflect.TypeOf(true)))
	False(t, isComposite(reflect.TypeOf(int(100))))
	False(t, isComposite(reflect.TypeOf(int8(100))))
	False(t, isComposite(reflect.TypeOf(int16(100))))
	False(t, isComposite(reflect.TypeOf(int32(100))))
	False(t, isComposite(reflect.TypeOf(int64(100))))
	False(t, isComposite(reflect.TypeOf(uint(100))))
	False(t, isComposite(reflect.TypeOf(uint8(100))))
	False(t, isComposite(reflect.TypeOf(uint16(100))))
	False(t, isComposite(reflect.TypeOf(uint32(100))))
	False(t, isComposite(reflect.TypeOf(uint64(100))))
	False(t, isComposite(reflect.TypeOf(uintptr(100))))
	False(t, isComposite(reflect.TypeOf(float32(100))))
	False(t, isComposite(reflect.TypeOf(float64(100))))
	False(t, isComposite(reflect.TypeOf(complex64(100))))
	False(t, isComposite(reflect.TypeOf(complex128(100))))
	False(t, isComposite(reflect.TypeOf(string("abc"))))
	False(t, isComposite(reflect.TypeOf(make(chan int))))
	False(t, isComposite(reflect.TypeOf(func() {})))
	False(t, isComposite(reflect.TypeOf(new(int))))
	False(t, isComposite(reflect.TypeOf(unsafe.Pointer(new(int)))))
	False(t, isComposite(reflect.ValueOf(struct{ a interface{} }{}).Field(0).Type()))
	True(t, isComposite(reflect.TypeOf([...]int{1, 2, 3})))
	True(t, isComposite(reflect.TypeOf([]int{1, 2, 3})))
	True(t, isComposite(reflect.TypeOf(map[int]bool{100: true})))
	True(t, isComposite(reflect.TypeOf(A{})))
}

func TestIsNonTrivialElem(t *testing.T) {
	False(t, isNonTrivialElem(reflect.ValueOf(nil)))
	False(t, isNonTrivialElem(reflect.ValueOf(100)))
	False(t, isNonTrivialElem(reflect.ValueOf(struct{ a interface{} }{}).Field(0)))
	True(t, isNonTrivialElem(reflect.ValueOf(struct{ a interface{} }{[]int{1, 2, 3}}).Field(0)))
	False(t, isNonTrivialElem(reflect.ValueOf([...]int{})))
	True(t, isNonTrivialElem(reflect.ValueOf([...]int{1, 2, 3})))
	False(t, isNonTrivialElem(reflect.ValueOf([]int(nil))))
	False(t, isNonTrivialElem(reflect.ValueOf([]int{})))
	True(t, isNonTrivialElem(reflect.ValueOf([]int{1, 2, 3})))
	False(t, isNonTrivialElem(reflect.ValueOf(map[bool]int(nil))))
	False(t, isNonTrivialElem(reflect.ValueOf(map[bool]int{})))
	True(t, isNonTrivialElem(reflect.ValueOf(map[bool]int{true: 1})))
	False(t, isNonTrivialElem(reflect.ValueOf(struct{}{})))
	True(t, isNonTrivialElem(reflect.ValueOf(struct{ a int }{})))
}

func TestAttrElemStruct(t *testing.T) {
	False(t, attrElemStruct(reflect.ValueOf(struct{}{})))
	False(t, attrElemStruct(reflect.ValueOf(struct {
		a int
		b string
		c []uint
	}{})))
	True(t, attrElemStruct(reflect.ValueOf(struct {
		b string
		c []uint
		a int
	}{c: []uint{1, 2, 3}})))
}

func TestAttrElemMap(t *testing.T) {
	cs := []struct {
		tp, ml bool
		m      interface{}
	}{
		{false, false, map[[3]int][]int{}},
		{true, false, map[bool]chan int{true: nil}},
		{true, false, map[*int]int{new(int): 10}},
		{true, true, map[bool][]int{true: []int{1, 2, 3}}},
		{true, true, map[[3]int]bool{[...]int{1, 2, 3}: true}},
	}
	for i, c := range cs {
		a1, a2 := attrElemMap(reflect.ValueOf(c.m))
		Equal(t, c.tp, a1, "i=%v", i)
		Equal(t, c.ml, a2, "i=%v", i)
	}
}

func TestAttrElemArray(t *testing.T) {
	cs := []struct {
		tp, id, ml bool
		m          interface{}
	}{
		{false, false, false, [...][]int{}},
		{true, false, false, [...]*int{new(int)}},
		{true, true, false, [11]int{}},
		{true, false, true, [3][]int{1: []int{1}}},
		{true, true, true, [11][]int{9: []int{1}}},
	}
	for i, c := range cs {
		a1, a2, a3 := attrElemArray(reflect.ValueOf(c.m))
		Equal(t, c.tp, a1, "i=%v", i)
		Equal(t, c.id, a2, "i=%v", i)
		Equal(t, c.ml, a3, "i=%v", i)
	}
}

func TestWriteKey(t *testing.T) {
	cs := []struct {
		e, e2 string
		v     reflect.Value
		p     bool
	}{
		{e: "<nil>", v: reflect.ValueOf(nil)},
		{e: "true", v: reflect.ValueOf(true)},
		{e: "100", v: reflect.ValueOf(int(100))},
		{e: "100", v: reflect.ValueOf(int8(100))},
		{e: "100", v: reflect.ValueOf(int16(100))},
		{e: "100", v: reflect.ValueOf(int32(100))},
		{e: "100", v: reflect.ValueOf(int64(100))},
		{e: "100", v: reflect.ValueOf(uint(100))},
		{e: "100", v: reflect.ValueOf(uint8(100))},
		{e: "100", v: reflect.ValueOf(uint16(100))},
		{e: "100", v: reflect.ValueOf(uint32(100))},
		{e: "100", v: reflect.ValueOf(uint64(100))},
		{e: "0x64", v: reflect.ValueOf(uintptr(100))},
		{e: "100", v: reflect.ValueOf(float32(100))},
		{e: "100", v: reflect.ValueOf(float64(100))},
		{e: "(100.1+200.2i)", v: reflect.ValueOf(complex64(100.1 + 200.2i))},
		{e: "(100.1+200.2i)", v: reflect.ValueOf(complex128(100.1 + 200.2i))},
		{e: `"abc"`, v: reflect.ValueOf(string("abc"))},
		{e: "<nil>", v: reflect.ValueOf(chan int(nil))},
		{e: "%v", v: reflect.ValueOf(make(chan int)), p: true},
		{e: "<nil>", v: reflect.ValueOf((func() int)(nil))},
		{e: "%v", v: reflect.ValueOf(func() {}), p: true},
		{e: "<nil>", v: reflect.ValueOf((*int)(nil))},
		{e: "%v", v: reflect.ValueOf(new(int)), p: true},
		{e: "<nil>", v: reflect.ValueOf(unsafe.Pointer(nil))},
		{e: "%v", v: reflect.ValueOf(unsafe.Pointer(new(int))), p: true},
		{e: "<nil>", v: reflect.ValueOf(struct{ a interface{} }{}).Field(0)},
		{e: "100", v: reflect.ValueOf(struct{ a interface{} }{100}).Field(0)},
		{e: "[1 2 3]", v: reflect.ValueOf([...]int{1, 2, 3})},
		{e: "<nil>", v: reflect.ValueOf([]int(nil))},
		{e: "[1 2 3]", v: reflect.ValueOf([]int{1, 2, 3})},
		{e: "<nil>", v: reflect.ValueOf(map[bool]int(nil))},
		{e: "map[true:100 false:200]", e2: "map[false:200 true:100]", v: reflect.ValueOf(map[bool]int{true: 100, false: 200})},
		{e: "{}", v: reflect.ValueOf(struct{}{})},
		{e: "{a:100 b:false c:0xc8}", v: reflect.ValueOf(struct {
			a int
			b bool
			c uintptr
		}{100, false, 200})},
	}
	for i, c := range cs {
		var d ValueDiffer
		d.writeKey(0, c.v, false)
		d.writeKey(1, c.v, true)
		r1 := c.e
		if c.p {
			r1 = fmt.Sprintf(r1, c.v)
		}
		r2 := c.e2
		if r2 != "" && c.p {
			r2 = fmt.Sprintf(r2, c.v)
		}
		if r2 == d.String(0) {
			Equal(t, r2, d.String(0), "i=%v", i)
		} else {
			Equal(t, r1, d.String(0), "i=%v", i)
		}
		if H(r2) == d.String(1) {
			Equal(t, H(r2), d.String(1), "i=%v", i)
		} else {
			Equal(t, H(r1), d.String(1), "i=%v", i)
		}
	}
}
