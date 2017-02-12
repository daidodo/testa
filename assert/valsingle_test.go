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
	True(t, isReference(reflect.ValueOf(struct{ a I }{}).Field(0).Type()))
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
	True(t, isNonTrivial(reflect.ValueOf(struct{ a I }{}).Field(0).Type()))
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
	False(t, isComposite(reflect.ValueOf(struct{ a I }{}).Field(0).Type()))
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
		if r2 != "" && r2 == d.String(0) {
			Equal(t, r2, d.String(0), "i=%v", i)
		} else {
			Equal(t, r1, d.String(0), "i=%v", i)
		}
		if r2 != "" && H(r2) == d.String(1) {
			Equal(t, H(r2), d.String(1), "i=%v", i)
		} else {
			Equal(t, H(r1), d.String(1), "i=%v", i)
		}
	}
}

func TestWriteElem(t *testing.T) {
	pa := new(int)
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
		{e: "[]", v: reflect.ValueOf([...]chan int{})},
		{e: fmt.Sprintf("[3]*int{<nil>, %v, <nil>}", pa), v: reflect.ValueOf([3]*int{1: pa})},
		{e: "[11]uint{0:0, 1:0, 2:0, 3:0, 4:0, 5:0, 6:0, 7:0, 8:0, 9:0, 10:0}", v: reflect.ValueOf([11]uint{})},
		{e: `[5][]int{
	<nil>, [],
	[100],
	[1 2 3],
	<nil>
}`, v: reflect.ValueOf([5][]int{1: []int{}, 2: []int{100}, 3: []int{1, 2, 3}})},
		{e: `[11][]int{
	0:<nil>,
	1:[],
	2:[100],
	3:[1 2 3],
	4:<nil>,
	5:<nil>,
	6:<nil>,
	7:<nil>,
	8:<nil>,
	9:<nil>,
	10:<nil>
}`, v: reflect.ValueOf([11][]int{1: []int{}, 2: []int{100}, 3: []int{1, 2, 3}})},
		{e: "<nil>", v: reflect.ValueOf([]int(nil))},
		{e: "[]", v: reflect.ValueOf([]chan int{})},
		{e: fmt.Sprintf("[]*int{<nil>, %v, <nil>}", pa), v: reflect.ValueOf([]*int{1: pa, 2: nil})},
		{e: "[]uint{0:0, 1:0, 2:0, 3:0, 4:0, 5:0, 6:0, 7:0, 8:0, 9:0, 10:0}", v: reflect.ValueOf(make([]uint, 11))},
		{e: `[][]int{
	<nil>, [],
	[100],
	[1 2 3],
	<nil>
}`, v: reflect.ValueOf([][]int{1: []int{}, 2: []int{100}, 3: []int{1, 2, 3}, 4: nil})},
		{e: `[][]int{
	0:<nil>,
	1:[],
	2:[100],
	3:[1 2 3],
	4:<nil>,
	5:<nil>,
	6:<nil>,
	7:<nil>,
	8:<nil>,
	9:<nil>,
	10:<nil>
}`, v: reflect.ValueOf([][]int{1: []int{}, 2: []int{100}, 3: []int{1, 2, 3}, 10: nil})},
		{e: "<nil>", v: reflect.ValueOf(map[bool]int(nil))},
		{e: "map[]", v: reflect.ValueOf(map[bool]int{})},
		{e: "map[true:100 false:200]", e2: "map[false:200 true:100]", v: reflect.ValueOf(map[bool]int{true: 100, false: 200})},
		{e: "map[bool]chan int{true:<nil>}", v: reflect.ValueOf(map[bool]chan int{true: nil})},
		{e: fmt.Sprintf("map[*int]int{%v:10}", pa), v: reflect.ValueOf(map[*int]int{pa: 10})},
		{e: `map[bool][]int{
	true:[1 2 3],
	false:[100 200]
}`, e2: `map[bool][]int{
	false:[100 200],
	true:[1 2 3]
}`, v: reflect.ValueOf(map[bool][]int{true: []int{1, 2, 3}, false: []int{100, 200}})},
		{e: `map[[3]int]bool{
	[1 2 3]:true,
	[100 200 300]:false
}`, e2: `map[[3]int]bool{
	[100 200 300]:false,
	[1 2 3]:true
}`, v: reflect.ValueOf(map[[3]int]bool{[3]int{1, 2, 3}: true, [3]int{100, 200, 300}: false})},
		{e: "{}", v: reflect.ValueOf(struct{}{})},
		{e: `{a:0 b:"" c:<nil>}`, v: reflect.ValueOf(struct {
			a int
			b string
			c []uint
		}{})},
		{e: `struct{
	b:"",
	c:[1 2 3],
	a:0
}`, v: reflect.ValueOf(struct {
			b string
			c []uint
			a int
		}{c: []uint{1, 2, 3}})},
	}
	for i, c := range cs {
		var d ValueDiffer
		d.writeElem(0, c.v, false)
		d.writeElem(1, c.v, true)
		r1 := c.e
		if c.p {
			r1 = fmt.Sprintf(r1, c.v)
		}
		r2 := c.e2
		if r2 != "" && c.p {
			r2 = fmt.Sprintf(r2, c.v)
		}
		if r2 != "" && r2 == d.String(0) {
			Equal(t, r2, d.String(0), "i=%v, r=\n%v", i, d.String(0))
		} else {
			Equal(t, r1, d.String(0), "i=%v, r=\n%v", i, d.String(0))
		}
		if r2 != "" && H(r2) == d.String(1) {
			Equal(t, H(r2), d.String(1), "i=%v, r=\n%v", i, d.String(1))
		} else {
			Equal(t, H(r1), d.String(1), "i=%v, r=\n%v", i, d.String(1))
		}
	}
}

func TestWriteValueAfterType(t *testing.T) {
	pa := new(int)
	cs := []struct {
		e, e2 string
		v     reflect.Value
		p     bool
	}{
		{e: "", v: reflect.ValueOf(nil)},
		{e: "(true)", v: reflect.ValueOf(true)},
		{e: "(100)", v: reflect.ValueOf(int(100))},
		{e: "(100)", v: reflect.ValueOf(int8(100))},
		{e: "(100)", v: reflect.ValueOf(int16(100))},
		{e: "(100)", v: reflect.ValueOf(int32(100))},
		{e: "(100)", v: reflect.ValueOf(int64(100))},
		{e: "(100)", v: reflect.ValueOf(uint(100))},
		{e: "(100)", v: reflect.ValueOf(uint8(100))},
		{e: "(100)", v: reflect.ValueOf(uint16(100))},
		{e: "(100)", v: reflect.ValueOf(uint32(100))},
		{e: "(100)", v: reflect.ValueOf(uint64(100))},
		{e: "(0x64)", v: reflect.ValueOf(uintptr(100))},
		{e: "(100)", v: reflect.ValueOf(float32(100))},
		{e: "(100)", v: reflect.ValueOf(float64(100))},
		{e: "(100.1+200.2i)", v: reflect.ValueOf(complex64(100.1 + 200.2i))},
		{e: "(100.1+200.2i)", v: reflect.ValueOf(complex128(100.1 + 200.2i))},
		{e: `("abc")`, v: reflect.ValueOf(string("abc"))},
		{e: "(nil)", v: reflect.ValueOf(chan int(nil))},
		{e: "(%v)", v: reflect.ValueOf(make(chan int)), p: true},
		{e: "(nil)", v: reflect.ValueOf((func() int)(nil))},
		{e: "(%v)", v: reflect.ValueOf(func() {}), p: true},
		{e: "(nil)", v: reflect.ValueOf((*int)(nil))},
		{e: "(%v)", v: reflect.ValueOf(new(int)), p: true},
		{e: "(nil)", v: reflect.ValueOf(unsafe.Pointer(nil))},
		{e: "(%v)", v: reflect.ValueOf(unsafe.Pointer(new(int))), p: true},
		{e: "", v: reflect.ValueOf(struct{ a interface{} }{}).Field(0)},
		{e: "(nil)", v: reflect.ValueOf(struct{ a I }{}).Field(0)},
		{e: "(100)", v: reflect.ValueOf(struct{ a interface{} }{100}).Field(0)},
		{e: "{}", v: reflect.ValueOf([...]chan int{})},
		{e: fmt.Sprintf("{<nil>, %v, <nil>}", pa), v: reflect.ValueOf([3]*int{1: pa})},
		{e: "{0:0, 1:0, 2:0, 3:0, 4:0, 5:0, 6:0, 7:0, 8:0, 9:0, 10:0}", v: reflect.ValueOf([11]uint{})},
		{e: `{
	<nil>, [],
	[100],
	[1 2 3],
	<nil>
}`, v: reflect.ValueOf([5][]int{1: []int{}, 2: []int{100}, 3: []int{1, 2, 3}})},
		{e: `{
	0:<nil>,
	1:[],
	2:[100],
	3:[1 2 3],
	4:<nil>,
	5:<nil>,
	6:<nil>,
	7:<nil>,
	8:<nil>,
	9:<nil>,
	10:<nil>
}`, v: reflect.ValueOf([11][]int{1: []int{}, 2: []int{100}, 3: []int{1, 2, 3}})},
		{e: "(nil)", v: reflect.ValueOf([]int(nil))},
		{e: "{}", v: reflect.ValueOf([]chan int{})},
		{e: fmt.Sprintf("{<nil>, %v, <nil>}", pa), v: reflect.ValueOf([]*int{1: pa, 2: nil})},
		{e: "{0:0, 1:0, 2:0, 3:0, 4:0, 5:0, 6:0, 7:0, 8:0, 9:0, 10:0}", v: reflect.ValueOf(make([]uint, 11))},
		{e: `{
	<nil>, [],
	[100],
	[1 2 3],
	<nil>
}`, v: reflect.ValueOf([][]int{1: []int{}, 2: []int{100}, 3: []int{1, 2, 3}, 4: nil})},
		{e: `{
	0:<nil>,
	1:[],
	2:[100],
	3:[1 2 3],
	4:<nil>,
	5:<nil>,
	6:<nil>,
	7:<nil>,
	8:<nil>,
	9:<nil>,
	10:<nil>
}`, v: reflect.ValueOf([][]int{1: []int{}, 2: []int{100}, 3: []int{1, 2, 3}, 10: nil})},
		{e: "(nil)", v: reflect.ValueOf(map[bool]int(nil))},
		{e: "{}", v: reflect.ValueOf(map[bool]int{})},
		{e: "{true:100, false:200}", e2: "{false:200, true:100}", v: reflect.ValueOf(map[bool]int{true: 100, false: 200})},
		{e: "{true:<nil>}", v: reflect.ValueOf(map[bool]chan int{true: nil})},
		{e: fmt.Sprintf("{%v:10}", pa), v: reflect.ValueOf(map[*int]int{pa: 10})},
		{e: `{
	true:[1 2 3],
	false:[100 200]
}`, e2: `{
	false:[100 200],
	true:[1 2 3]
}`, v: reflect.ValueOf(map[bool][]int{true: []int{1, 2, 3}, false: []int{100, 200}})},
		{e: `{
	[1 2 3]:true,
	[100 200 300]:false
}`, e2: `{
	[100 200 300]:false,
	[1 2 3]:true
}`, v: reflect.ValueOf(map[[3]int]bool{[3]int{1, 2, 3}: true, [3]int{100, 200, 300}: false})},
		{e: "{}", v: reflect.ValueOf(struct{}{})},
		{e: `{a:0, b:"", c:<nil>}`, v: reflect.ValueOf(struct {
			a int
			b string
			c []uint
		}{})},
		{e: `{
	b:"",
	c:[1 2 3],
	a:0
}`, v: reflect.ValueOf(struct {
			b string
			c []uint
			a int
		}{c: []uint{1, 2, 3}})},
	}
	for i, c := range cs {
		var d ValueDiffer
		d.writeValueAfterType(0, c.v)
		r1 := c.e
		if c.p {
			r1 = fmt.Sprintf(r1, c.v)
		}
		r2 := c.e2
		if r2 != "" && c.p {
			r2 = fmt.Sprintf(r2, c.v)
		}
		if r2 != "" && r2 == d.String(0) {
			Equal(t, r2, d.String(0), "i=%v, r=\n%v", i, d.String(0))
		} else {
			Equal(t, r1, d.String(0), "i=%v, r=\n%v", i, d.String(0))
		}
	}
}

func TestWriteTypeHeadChan(t *testing.T) {
	cs := []struct {
		e         string
		v         reflect.Type
		h, hd, he bool
	}{
		{"chan ", reflect.TypeOf(make(chan int)), false, false, false},
		{"chan ", reflect.TypeOf(make(chan int)), false, false, true},
		{"chan ", reflect.TypeOf(make(chan int)), false, true, false},
		{"chan ", reflect.TypeOf(make(chan int)), false, true, true},
		{H("chan "), reflect.TypeOf(make(chan int)), true, false, false},
		{H("chan "), reflect.TypeOf(make(chan int)), true, false, true},
		{H("chan "), reflect.TypeOf(make(chan int)), true, true, false},
		{H("chan "), reflect.TypeOf(make(chan int)), true, true, true},
		{"<-chan ", reflect.TypeOf(make(<-chan int)), false, false, false},
		{"<-chan ", reflect.TypeOf(make(<-chan int)), false, false, true},
		{H("<-") + "chan ", reflect.TypeOf(make(<-chan int)), false, true, false},
		{H("<-") + "chan ", reflect.TypeOf(make(<-chan int)), false, true, true},
		{H("<-chan "), reflect.TypeOf(make(<-chan int)), true, false, false},
		{H("<-chan "), reflect.TypeOf(make(<-chan int)), true, false, true},
		{H("<-chan "), reflect.TypeOf(make(<-chan int)), true, true, false},
		{H("<-chan "), reflect.TypeOf(make(<-chan int)), true, true, true},
		{"chan<- ", reflect.TypeOf(make(chan<- int)), false, false, false},
		{"chan<- ", reflect.TypeOf(make(chan<- int)), false, false, true},
		{"chan" + H("<-") + " ", reflect.TypeOf(make(chan<- int)), false, true, false},
		{"chan" + H("<- "), reflect.TypeOf(make(chan<- int)), false, true, true},
		{H("chan<- "), reflect.TypeOf(make(chan<- int)), true, false, false},
		{H("chan<- "), reflect.TypeOf(make(chan<- int)), true, false, true},
		{H("chan<- "), reflect.TypeOf(make(chan<- int)), true, true, false},
		{H("chan<- "), reflect.TypeOf(make(chan<- int)), true, true, true},
	}
	for i, c := range cs {
		var d ValueDiffer
		d.writeTypeHeadChan(0, c.v, c.h, c.hd, c.he)
		Equal(t, c.e, d.String(0), "i=%v, r=\n%v", i, d.String(0))
	}
}

func TestWriteType(t *testing.T) {
	cs := []struct {
		e string
		v reflect.Type
	}{
		{v: reflect.TypeOf(true)},
		{v: reflect.TypeOf(int(100))},
		{v: reflect.TypeOf(int8(100))},
		{v: reflect.TypeOf(int16(100))},
		{v: reflect.TypeOf(int32(100))},
		{v: reflect.TypeOf(int64(100))},
		{v: reflect.TypeOf(uint(100))},
		{v: reflect.TypeOf(uint8(100))},
		{v: reflect.TypeOf(uint16(100))},
		{v: reflect.TypeOf(uint32(100))},
		{v: reflect.TypeOf(uint64(100))},
		{v: reflect.TypeOf(uintptr(100))},
		{v: reflect.TypeOf(float32(100))},
		{v: reflect.TypeOf(float64(100))},
		{v: reflect.TypeOf(complex64(100))},
		{v: reflect.TypeOf(complex128(100))},
		{v: reflect.TypeOf(string("abc"))},
		{v: reflect.TypeOf(make(chan int))},
		{v: reflect.TypeOf(make(<-chan int))},
		{v: reflect.TypeOf(make(chan<- int))},
		{v: reflect.TypeOf(func() {})},
		{v: reflect.TypeOf(func(string) {})},
		{v: reflect.TypeOf(func(int, string) {})},
		{v: reflect.TypeOf(func(int, string) float32 { return 0 })},
		{v: reflect.TypeOf(func(int, string) (bool, float32, string) { return true, 0, "a" })},
		{v: reflect.TypeOf(new(int))},
		{v: reflect.TypeOf(new(chan int))},
		{v: reflect.TypeOf(new(func(int, string)))},
		{v: reflect.TypeOf(new(func(int, string) (bool, int8, uint)))},
		{v: reflect.TypeOf(unsafe.Pointer(new(int)))},
		{v: reflect.ValueOf(struct{ a interface{} }{}).Field(0).Type()},
		{v: reflect.ValueOf(struct{ a I }{}).Field(0).Type()},
		{v: reflect.TypeOf([...]int{})},
		{v: reflect.TypeOf([...]int{1, 2, 3})},
		{v: reflect.TypeOf([]int{})},
		{v: reflect.TypeOf(map[bool]int{})},
		{v: reflect.TypeOf(struct{ a int }{}), e: "struct"},
		{v: reflect.TypeOf(A{})},
	}
	for i, c := range cs {
		var d ValueDiffer
		d.writeType(0, c.v, false)
		d.writeType(1, c.v, true)
		e := c.e
		if e == "" {
			e = c.v.String()
		}
		Equal(t, e, d.String(0), "i=%v, r=\n%v", i, d.String(0))
		Equal(t, H(e), d.String(1), "i=%v, r=\n%v", i, d.String(1))
	}
}
