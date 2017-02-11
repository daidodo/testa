package assert

import (
	"reflect"
	"testing"
	"unsafe"
)

type I interface {
	Fun()
}

type A struct {
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
