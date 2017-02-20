/*
 * Copyright (c) 2017 Zhao DAI <daidodo@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or any
 * later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see accompanying file LICENSE.txt
 * or <http://www.gnu.org/licenses/>.
 */

package assert

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type I interface {
	Fun()
}

type A struct {
	a interface{}
	b I
}

func (A) Fun() {}

type B struct {
}

type Bool bool
type Int int
type Uint uint
type Uintptr uintptr
type Float float64
type Complex complex128
type Str string
type Chan chan int
type Func func(int) bool
type Ptr *int
type UPtr unsafe.Pointer
type If I
type Array [3]interface{}
type Slice []interface{}
type Map map[interface{}]interface{}
type Struct A

func (Int) Fun() {}

//func (Ptr) String() string     { return "String of Ptr" }
//func (UPtr) String() string   { return "String of UPtr" }
//func (If) String() string      { return "String of If" }
func (Bool) String() string    { return "String of Bool" }
func (Int) String() string     { return "String of Int" }
func (Uint) String() string    { return "String of Uint" }
func (Uintptr) String() string { return "String of Uintptr" }
func (Float) String() string   { return "String of Float" }
func (Complex) String() string { return "String of Complex" }
func (Str) String() string     { return "String of Str" }
func (Chan) String() string    { return "String of Chan" }
func (Func) String() string    { return "String of Func" }
func (Array) String() string   { return "String of Array" }
func (Slice) String() string   { return "String of Slice" }
func (Map) String() string     { return "String of Map" }
func (Struct) String() string  { return "String of Struct" }

//func (Ptr) GoString() string     { return "GoString of Ptr" }
//func (UPtr) GoString() string   { return "GoString of UPtr" }
//func (If) GoString() string      { return "GoString of If" }
func (Bool) GoString() string    { return "GoString of Bool" }
func (Int) GoString() string     { return "GoString of Int" }
func (Uint) GoString() string    { return "GoString of Uint" }
func (Uintptr) GoString() string { return "GoString of Uintptr" }
func (Float) GoString() string   { return "GoString of Float" }
func (Complex) GoString() string { return "GoString of Complex" }
func (Str) GoString() string     { return "GoString of Str" }
func (Chan) GoString() string    { return "GoString of Chan" }
func (Func) GoString() string    { return "GoString of Func" }
func (Array) GoString() string   { return "GoString of Array" }
func (Slice) GoString() string   { return "GoString of Slice" }
func (Map) GoString() string     { return "GoString of Map" }
func (Struct) GoString() string  { return "GoString of Struct" }

//func (Ptr) Error() string     { return "Error of Ptr" }
//func (UPtr) Error() string   { return "Error of UPtr" }
//func (If) Error() string      { return "Error of If" }
func (Bool) Error() string    { return "Error of Bool" }
func (Int) Error() string     { return "Error of Int" }
func (Uint) Error() string    { return "Error of Uint" }
func (Uintptr) Error() string { return "Error of Uintptr" }
func (Float) Error() string   { return "Error of Float" }
func (Complex) Error() string { return "Error of Complex" }
func (Str) Error() string     { return "Error of Str" }
func (Chan) Error() string    { return "Error of Chan" }
func (Func) Error() string    { return "Error of Func" }
func (Array) Error() string   { return "Error of Array" }
func (Slice) Error() string   { return "Error of Slice" }
func (Map) Error() string     { return "Error of Map" }
func (Struct) Error() string  { return "Error of Struct" }

func H(s string) string {
	if len(s) < 1 {
		return ""
	}
	return kRED + s + kEND
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

type PInt uintptr
type PStr uintptr

func (p PInt) String() string {
	return "String of PInt"
}

func (p PStr) GoString() string {
	return "Go String of PStr"
}

func (p PStr) String() string {
	return "String of PStr"
}

func TestWriteKey(t *testing.T) {
	a := &A{}
	cs := []struct {
		e, e2 string
		h, h2 string
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
		{e: "(0+200.2i)", v: reflect.ValueOf(complex64(200.2i)), h: "(" + H("0+200.2i") + ")"},
		{e: "(100.1+0i)", v: reflect.ValueOf(complex64(100.1)), h: "(" + H("100.1+0i") + ")"},
		{e: "(100.1+200.2i)", v: reflect.ValueOf(complex64(100.1 + 200.2i)), h: "(" + H("100.1+200.2i") + ")"},
		{e: "(100.1+200.2i)", v: reflect.ValueOf(complex128(100.1 + 200.2i)), h: "(" + H("100.1+200.2i") + ")"},
		{e: `"abc"`, v: reflect.ValueOf(string("abc"))},
		{e: "<nil>", v: reflect.ValueOf(chan int(nil))},
		{e: "%v", v: reflect.ValueOf(make(chan int)), p: true},
		{e: "<nil>", v: reflect.ValueOf((func() int)(nil))},
		{e: "%v", v: reflect.ValueOf(func() {}), p: true},
		{e: "<nil>", v: reflect.ValueOf((*int)(nil))},
		{e: "%v", v: reflect.ValueOf(new(int)), p: true},
		{e: fmt.Sprintf("%p", a), v: reflect.ValueOf(a)},
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
		{e: "array", v: reflect.ValueOf(reflect.Array)},
		{e: "17", v: reflect.ValueOf(struct{ a reflect.Kind }{reflect.Array}).Field(0)},
		{e: "String of PInt", v: reflect.ValueOf(PInt(100))},
		{e: "0x64", v: reflect.ValueOf(struct{ a PInt }{100}).Field(0)},
		{e: "String of PStr", v: reflect.ValueOf(PStr(100))},
		{e: "0x64", v: reflect.ValueOf(struct{ a PStr }{100}).Field(0)},
		{e: "String of Bool", v: reflect.ValueOf(Bool(true))},
		{e: "String of Int", v: reflect.ValueOf(Int(100))},
		{e: "String of Uint", v: reflect.ValueOf(Uint(100))},
		{e: "String of Uintptr", v: reflect.ValueOf(Uintptr(100))},
		{e: "100", v: reflect.ValueOf(Float(100))},
		{e: "(100+0i)", v: reflect.ValueOf(Complex(100)), h: "(" + H("100+0i") + ")"},
		{e: `"100"`, v: reflect.ValueOf(Str("100"))},
	}
	for i, c := range cs {
		var d tValueDiffer
		d.writeKey(0, c.v, false)
		d.writeKey(1, c.v, true)
		r1, r2 := c.e, c.e2
		if c.p {
			r1 = fmt.Sprintf(r1, c.v)
		}
		if r2 != "" && c.p {
			r2 = fmt.Sprintf(r2, c.v)
		}
		h1, h2 := c.h, c.h2
		if h1 == "" {
			h1 = H(r1)
		}
		if h2 == "" {
			h2 = H(r2)
		}
		if r2 != "" && r2 == d.String(0) {
			Equal(t, r2, d.String(0), "i=%v\n%v", i, d.String(0))
		} else {
			Equal(t, r1, d.String(0), "i=%v\n%v", i, d.String(0))
		}
		if h2 != "" && h2 == d.String(1) {
			Equal(t, h2, d.String(1), "i=%v\n%v", i, d.String(1))
		} else {
			Equal(t, h1, d.String(1), "i=%v\n%v", i, d.String(1))
		}
	}
}

func TestWriteElem(t *testing.T) {
	pa := new(int)
	cs := []struct {
		e, e2 string
		h, h2 string
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
		{e: "(0+200.2i)", v: reflect.ValueOf(complex64(200.2i)), h: "(" + H("0+200.2i") + ")"},
		{e: "(100.1+0i)", v: reflect.ValueOf(complex64(100.1)), h: "(" + H("100.1+0i") + ")"},
		{e: "(100.1+200.2i)", v: reflect.ValueOf(complex64(100.1 + 200.2i)), h: "(" + H("100.1+200.2i") + ")"},
		{e: "(100.1+200.2i)", v: reflect.ValueOf(complex128(100.1 + 200.2i)), h: "(" + H("100.1+200.2i") + ")"},
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
		{v: reflect.ValueOf([5][]int{1: []int{}, 2: []int{100}, 3: []int{1, 2, 3}}),
			e: "[5][]int{\n\t<nil>, [],\n\t[100],\n\t[1 2 3],\n\t<nil>\n}",
			h: "\x1b[41m[5][]int{\x1b[0m\n\t\x1b[41m<nil>, [],\x1b[0m\n\t\x1b[41m[100],\x1b[0m\n\t\x1b[41m[1 2 3],\x1b[0m\n\t\x1b[41m<nil>\x1b[0m\n\x1b[41m}\x1b[0m"},
		{v: reflect.ValueOf([11][]int{1: []int{}, 2: []int{100}, 3: []int{1, 2, 3}}),
			e: "[11][]int{\n\t0:<nil>,\n\t1:[],\n\t2:[100],\n\t3:[1 2 3],\n\t4:<nil>,\n\t5:<nil>,\n\t6:<nil>,\n\t7:<nil>,\n\t8:<nil>,\n\t9:<nil>,\n\t10:<nil>\n}",
			h: "\x1b[41m[11][]int{\x1b[0m\n\t\x1b[41m0:<nil>,\x1b[0m\n\t\x1b[41m1:[],\x1b[0m\n\t\x1b[41m2:[100],\x1b[0m\n\t\x1b[41m3:[1 2 3],\x1b[0m\n\t\x1b[41m4:<nil>,\x1b[0m\n\t\x1b[41m5:<nil>,\x1b[0m\n\t\x1b[41m6:<nil>,\x1b[0m\n\t\x1b[41m7:<nil>,\x1b[0m\n\t\x1b[41m8:<nil>,\x1b[0m\n\t\x1b[41m9:<nil>,\x1b[0m\n\t\x1b[41m10:<nil>\x1b[0m\n\x1b[41m}\x1b[0m"},
		{e: "<nil>", v: reflect.ValueOf([]int(nil))},
		{e: "[]", v: reflect.ValueOf([]chan int{})},
		{e: fmt.Sprintf("[]*int{<nil>, %v, <nil>}", pa), v: reflect.ValueOf([]*int{1: pa, 2: nil})},
		{e: "[]uint{0:0, 1:0, 2:0, 3:0, 4:0, 5:0, 6:0, 7:0, 8:0, 9:0, 10:0}", v: reflect.ValueOf(make([]uint, 11))},
		{v: reflect.ValueOf([][]int{1: []int{}, 2: []int{100}, 3: []int{1, 2, 3}, 4: nil}),
			e: "[][]int{\n\t<nil>, [],\n\t[100],\n\t[1 2 3],\n\t<nil>\n}",
			h: "\x1b[41m[][]int{\x1b[0m\n\t\x1b[41m<nil>, [],\x1b[0m\n\t\x1b[41m[100],\x1b[0m\n\t\x1b[41m[1 2 3],\x1b[0m\n\t\x1b[41m<nil>\x1b[0m\n\x1b[41m}\x1b[0m"},
		{v: reflect.ValueOf([][]int{1: []int{}, 2: []int{100}, 3: []int{1, 2, 3}, 10: nil}),
			e: "[][]int{\n\t0:<nil>,\n\t1:[],\n\t2:[100],\n\t3:[1 2 3],\n\t4:<nil>,\n\t5:<nil>,\n\t6:<nil>,\n\t7:<nil>,\n\t8:<nil>,\n\t9:<nil>,\n\t10:<nil>\n}",
			h: "\x1b[41m[][]int{\x1b[0m\n\t\x1b[41m0:<nil>,\x1b[0m\n\t\x1b[41m1:[],\x1b[0m\n\t\x1b[41m2:[100],\x1b[0m\n\t\x1b[41m3:[1 2 3],\x1b[0m\n\t\x1b[41m4:<nil>,\x1b[0m\n\t\x1b[41m5:<nil>,\x1b[0m\n\t\x1b[41m6:<nil>,\x1b[0m\n\t\x1b[41m7:<nil>,\x1b[0m\n\t\x1b[41m8:<nil>,\x1b[0m\n\t\x1b[41m9:<nil>,\x1b[0m\n\t\x1b[41m10:<nil>\x1b[0m\n\x1b[41m}\x1b[0m"},
		{e: "<nil>", v: reflect.ValueOf(map[bool]int(nil))},
		{e: "map[]", v: reflect.ValueOf(map[bool]int{})},
		{e: "map[true:100 false:200]", e2: "map[false:200 true:100]", v: reflect.ValueOf(map[bool]int{true: 100, false: 200})},
		{e: "map[bool]chan int{true:<nil>}", v: reflect.ValueOf(map[bool]chan int{true: nil})},
		{e: fmt.Sprintf("map[*int]int{%v:10}", pa), v: reflect.ValueOf(map[*int]int{pa: 10})},
		{v: reflect.ValueOf(map[bool][]int{true: []int{1, 2, 3}, false: []int{100, 200}}),
			e:  "map[bool][]int{\n\ttrue:[1 2 3],\n\tfalse:[100 200]\n}",
			e2: "map[bool][]int{\n\tfalse:[100 200],\n\ttrue:[1 2 3]\n}",
			h:  "\x1b[41mmap[bool][]int{\x1b[0m\n\t\x1b[41mtrue:[1 2 3],\x1b[0m\n\t\x1b[41mfalse:[100 200]\x1b[0m\n\x1b[41m}\x1b[0m",
			h2: "\x1b[41mmap[bool][]int{\x1b[0m\n\t\x1b[41mfalse:[100 200],\x1b[0m\n\t\x1b[41mtrue:[1 2 3]\x1b[0m\n\x1b[41m}\x1b[0m"},
		{v: reflect.ValueOf(map[[3]int]bool{[3]int{1, 2, 3}: true, [3]int{100, 200, 300}: false}),
			e:  "map[[3]int]bool{\n\t[1 2 3]:true,\n\t[100 200 300]:false\n}",
			e2: "map[[3]int]bool{\n\t[100 200 300]:false,\n\t[1 2 3]:true\n}",
			h:  "\x1b[41mmap[[3]int]bool{\x1b[0m\n\t\x1b[41m[1 2 3]:true,\x1b[0m\n\t\x1b[41m[100 200 300]:false\x1b[0m\n\x1b[41m}\x1b[0m",
			h2: "\x1b[41mmap[[3]int]bool{\x1b[0m\n\t\x1b[41m[100 200 300]:false,\x1b[0m\n\t\x1b[41m[1 2 3]:true\x1b[0m\n\x1b[41m}\x1b[0m",
		},
		{e: "{}", v: reflect.ValueOf(struct{}{})},
		{e: `{a:0 b:"" c:<nil>}`, v: reflect.ValueOf(struct {
			a int
			b string
			c []uint
		}{})},
		{v: reflect.ValueOf(struct {
			b string
			c []uint
			a int
		}{c: []uint{1, 2, 3}}),
			e: "struct{\n\tb:\"\",\n\tc:[1 2 3],\n\ta:0\n}",
			h: "\x1b[41mstruct{\x1b[0m\n\t\x1b[41mb:\"\",\x1b[0m\n\t\x1b[41mc:[1 2 3],\x1b[0m\n\t\x1b[41ma:0\x1b[0m\n\x1b[41m}\x1b[0m"},
		{e: "array", v: reflect.ValueOf(reflect.Array)},
		{e: "17", v: reflect.ValueOf(struct{ a reflect.Kind }{reflect.Array}).Field(0)},
		{e: "String of PInt", v: reflect.ValueOf(PInt(100))},
		{e: "0x64", v: reflect.ValueOf(struct{ a PInt }{100}).Field(0)},
		{e: "String of PStr", v: reflect.ValueOf(PStr(100))},
		{e: "0x64", v: reflect.ValueOf(struct{ a PStr }{100}).Field(0)},
		{e: "String of Bool", v: reflect.ValueOf(Bool(true))},
		{e: "String of Int", v: reflect.ValueOf(Int(100))},
		{e: "String of Uint", v: reflect.ValueOf(Uint(100))},
		{e: "String of Uintptr", v: reflect.ValueOf(Uintptr(100))},
		{e: "100", v: reflect.ValueOf(Float(100))},
		{e: "(100+0i)", v: reflect.ValueOf(Complex(100)), h: "(" + H("100+0i") + ")"},
		{e: `"100"`, v: reflect.ValueOf(Str("100"))},
	}
	for i, c := range cs {
		var d tValueDiffer
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
		h1, h2 := c.h, c.h2
		if h1 == "" {
			h1 = H(r1)
		}
		if h2 == "" {
			h2 = H(r2)
		}
		if h2 != "" && h2 == d.String(1) {
			Equal(t, h2, d.String(1), "i=%v, r=\n%v", i, d.String(1))
		} else {
			Equal(t, h1, d.String(1), "i=%v, r=\n%v", i, d.String(1))
		}
	}
}

func TestWriteValueAfterType(t *testing.T) {
	a := &A{}
	pa := new(int)
	cs := []struct {
		e, e2 string
		h, h2 string
		v     reflect.Value
		p     bool
	}{
		{e: "", v: reflect.ValueOf(nil)},
		{e: "(true)", v: reflect.ValueOf(true), h: "(" + H("true") + ")"},
		{e: "(100)", v: reflect.ValueOf(int(100)), h: "(" + H("100") + ")"},
		{e: "(100)", v: reflect.ValueOf(int8(100)), h: "(" + H("100") + ")"},
		{e: "(100)", v: reflect.ValueOf(int16(100)), h: "(" + H("100") + ")"},
		{e: "(100)", v: reflect.ValueOf(int32(100)), h: "(" + H("100") + ")"},
		{e: "(100)", v: reflect.ValueOf(int64(100)), h: "(" + H("100") + ")"},
		{e: "(100)", v: reflect.ValueOf(uint(100)), h: "(" + H("100") + ")"},
		{e: "(100)", v: reflect.ValueOf(uint8(100)), h: "(" + H("100") + ")"},
		{e: "(100)", v: reflect.ValueOf(uint16(100)), h: "(" + H("100") + ")"},
		{e: "(100)", v: reflect.ValueOf(uint32(100)), h: "(" + H("100") + ")"},
		{e: "(100)", v: reflect.ValueOf(uint64(100)), h: "(" + H("100") + ")"},
		{e: "(0x64)", v: reflect.ValueOf(uintptr(100)), h: "(" + H("0x64") + ")"},
		{e: "(100)", v: reflect.ValueOf(float32(100)), h: "(" + H("100") + ")"},
		{e: "(100)", v: reflect.ValueOf(float64(100)), h: "(" + H("100") + ")"},
		{e: "(0+200.2i)", v: reflect.ValueOf(complex64(200.2i)), h: "(" + H("0+200.2i") + ")"},
		{e: "(100.1+0i)", v: reflect.ValueOf(complex64(100.1)), h: "(" + H("100.1+0i") + ")"},
		{e: "(100.1+200.2i)", v: reflect.ValueOf(complex64(100.1 + 200.2i)), h: "(" + H("100.1+200.2i") + ")"},
		{e: "(100.1+200.2i)", v: reflect.ValueOf(complex128(100.1 + 200.2i)), h: "(" + H("100.1+200.2i") + ")"},
		{e: `("abc")`, v: reflect.ValueOf(string("abc")), h: "(" + H(`"abc"`) + ")"},
		{e: "(nil)", v: reflect.ValueOf(chan int(nil)), h: "(" + H("nil") + ")"},
		{e: "(%v)", v: reflect.ValueOf(make(chan int)), p: true, h: "(" + H("%v") + ")"},
		{e: "(nil)", v: reflect.ValueOf((func() int)(nil)), h: "(" + H("nil") + ")"},
		{e: "(%v)", v: reflect.ValueOf(func() {}), p: true, h: "(" + H("%v") + ")"},
		{e: "(nil)", v: reflect.ValueOf((*int)(nil)), h: "(" + H("nil") + ")"},
		{e: "(%v)", v: reflect.ValueOf(new(int)), p: true, h: "(" + H("%v") + ")"},
		{e: fmt.Sprintf("(%p)", a), v: reflect.ValueOf(a), h: fmt.Sprintf("("+H("%p")+")", a)},
		{e: "(nil)", v: reflect.ValueOf(unsafe.Pointer(nil)), h: "(" + H("nil") + ")"},
		{e: "(%v)", v: reflect.ValueOf(unsafe.Pointer(new(int))), p: true, h: "(" + H("%v") + ")"},
		{v: reflect.ValueOf(struct{ a interface{} }{}).Field(0)},
		{e: "(nil)", v: reflect.ValueOf(struct{ a I }{}).Field(0), h: "(" + H("nil") + ")"},
		{e: "(100)", v: reflect.ValueOf(struct{ a interface{} }{100}).Field(0), h: "(" + H("100") + ")"},
		{e: "{}", v: reflect.ValueOf([...]chan int{})},
		{e: fmt.Sprintf("{<nil>, %v, <nil>}", pa), v: reflect.ValueOf([3]*int{1: pa})},
		{e: "{0:0, 1:0, 2:0, 3:0, 4:0, 5:0, 6:0, 7:0, 8:0, 9:0, 10:0}", v: reflect.ValueOf([11]uint{})},
		{e: `{
	<nil>, [],
	[100],
	[1 2 3],
	<nil>
}`, v: reflect.ValueOf([5][]int{1: []int{}, 2: []int{100}, 3: []int{1, 2, 3}}),
			h: "\x1b[41m{\x1b[0m\n\t\x1b[41m<nil>, [],\x1b[0m\n\t\x1b[41m[100],\x1b[0m\n\t\x1b[41m[1 2 3],\x1b[0m\n\t\x1b[41m<nil>\x1b[0m\n\x1b[41m}\x1b[0m"},
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
}`, v: reflect.ValueOf([11][]int{1: []int{}, 2: []int{100}, 3: []int{1, 2, 3}}),
			h: "\x1b[41m{\x1b[0m\n\t\x1b[41m0:<nil>,\x1b[0m\n\t\x1b[41m1:[],\x1b[0m\n\t\x1b[41m2:[100],\x1b[0m\n\t\x1b[41m3:[1 2 3],\x1b[0m\n\t\x1b[41m4:<nil>,\x1b[0m\n\t\x1b[41m5:<nil>,\x1b[0m\n\t\x1b[41m6:<nil>,\x1b[0m\n\t\x1b[41m7:<nil>,\x1b[0m\n\t\x1b[41m8:<nil>,\x1b[0m\n\t\x1b[41m9:<nil>,\x1b[0m\n\t\x1b[41m10:<nil>\x1b[0m\n\x1b[41m}\x1b[0m"},
		{e: "(nil)", v: reflect.ValueOf([]int(nil)), h: "(" + H("nil") + ")"},
		{e: "{}", v: reflect.ValueOf([]chan int{})},
		{e: fmt.Sprintf("{<nil>, %v, <nil>}", pa), v: reflect.ValueOf([]*int{1: pa, 2: nil})},
		{e: "{0:0, 1:0, 2:0, 3:0, 4:0, 5:0, 6:0, 7:0, 8:0, 9:0, 10:0}", v: reflect.ValueOf(make([]uint, 11))},
		{e: `{
	<nil>, [],
	[100],
	[1 2 3],
	<nil>
}`, v: reflect.ValueOf([][]int{1: []int{}, 2: []int{100}, 3: []int{1, 2, 3}, 4: nil}),
			h: "\x1b[41m{\x1b[0m\n\t\x1b[41m<nil>, [],\x1b[0m\n\t\x1b[41m[100],\x1b[0m\n\t\x1b[41m[1 2 3],\x1b[0m\n\t\x1b[41m<nil>\x1b[0m\n\x1b[41m}\x1b[0m"},
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
}`, v: reflect.ValueOf([][]int{1: []int{}, 2: []int{100}, 3: []int{1, 2, 3}, 10: nil}),
			h: "\x1b[41m{\x1b[0m\n\t\x1b[41m0:<nil>,\x1b[0m\n\t\x1b[41m1:[],\x1b[0m\n\t\x1b[41m2:[100],\x1b[0m\n\t\x1b[41m3:[1 2 3],\x1b[0m\n\t\x1b[41m4:<nil>,\x1b[0m\n\t\x1b[41m5:<nil>,\x1b[0m\n\t\x1b[41m6:<nil>,\x1b[0m\n\t\x1b[41m7:<nil>,\x1b[0m\n\t\x1b[41m8:<nil>,\x1b[0m\n\t\x1b[41m9:<nil>,\x1b[0m\n\t\x1b[41m10:<nil>\x1b[0m\n\x1b[41m}\x1b[0m"},
		{e: "(nil)", v: reflect.ValueOf(map[bool]int(nil)), h: "(" + H("nil") + ")"},
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
}`, v: reflect.ValueOf(map[bool][]int{true: []int{1, 2, 3}, false: []int{100, 200}}),
			h:  "\x1b[41m{\x1b[0m\n\t\x1b[41mtrue:[1 2 3],\x1b[0m\n\t\x1b[41mfalse:[100 200]\x1b[0m\n\x1b[41m}\x1b[0m",
			h2: "\x1b[41m{\x1b[0m\n\t\x1b[41mfalse:[100 200],\x1b[0m\n\t\x1b[41mtrue:[1 2 3]\x1b[0m\n\x1b[41m}\x1b[0m"},
		{e: `{
	[1 2 3]:true,
	[100 200 300]:false
}`, e2: `{
	[100 200 300]:false,
	[1 2 3]:true
}`, v: reflect.ValueOf(map[[3]int]bool{[3]int{1, 2, 3}: true, [3]int{100, 200, 300}: false}),
			h:  "\x1b[41m{\x1b[0m\n\t\x1b[41m[1 2 3]:true,\x1b[0m\n\t\x1b[41m[100 200 300]:false\x1b[0m\n\x1b[41m}\x1b[0m",
			h2: "\x1b[41m{\x1b[0m\n\t\x1b[41m[100 200 300]:false,\x1b[0m\n\t\x1b[41m[1 2 3]:true\x1b[0m\n\x1b[41m}\x1b[0m"},
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
		}{c: []uint{1, 2, 3}}),
			h: "\x1b[41m{\x1b[0m\n\t\x1b[41mb:\"\",\x1b[0m\n\t\x1b[41mc:[1 2 3],\x1b[0m\n\t\x1b[41ma:0\x1b[0m\n\x1b[41m}\x1b[0m"},
		{e: "(String of Bool)", v: reflect.ValueOf(Bool(true)), h: "(" + H("String of Bool") + ")"},
		{e: "(String of Int)", v: reflect.ValueOf(Int(100)), h: "(" + H("String of Int") + ")"},
		{e: "(String of Uint)", v: reflect.ValueOf(Uint(100)), h: "(" + H("String of Uint") + ")"},
		{e: "(String of Uintptr)", v: reflect.ValueOf(Uintptr(100)), h: "(" + H("String of Uintptr") + ")"},
		{e: "(100)", v: reflect.ValueOf(Float(100)), h: "(" + H("100") + ")"},
		{e: "(100+0i)", v: reflect.ValueOf(Complex(100)), h: "(" + H("100+0i") + ")"},
		{e: `("100")`, v: reflect.ValueOf(Str("100")), h: "(" + H(`"100"`) + ")"},
	}
	for i, c := range cs {
		var d tValueDiffer
		d.writeValueAfterType(0, c.v, false)
		d.writeValueAfterType(1, c.v, true)
		r1, r2 := c.e, c.e2
		if c.p {
			r1 = fmt.Sprintf(r1, c.v)
		}
		if r2 != "" && c.p {
			r2 = fmt.Sprintf(r2, c.v)
		}
		h1, h2 := c.h, c.h2
		if h1 == "" {
			h1 = H(r1)
		} else if c.p {
			h1 = fmt.Sprintf(h1, c.v)
		}
		if h2 == "" {
			h2 = H(r2)
		} else if c.p {
			h2 = fmt.Sprintf(h2, c.v)
		}
		if r2 != "" && r2 == d.String(0) {
			Equal(t, r2, d.String(0), "i=%v, r=\n%v", i, d.String(0))
		} else {
			Equal(t, r1, d.String(0), "i=%v, r=\n%v", i, d.String(0))
		}
		if h2 != "" && h2 == d.String(1) {
			Equal(t, h2, d.String(1), "i=%v, hr=\n%v", i, d.String(1))
		} else {
			Equal(t, h1, d.String(1), "i=%v, hr=\n%v", i, d.String(1))
		}
	}
}

func TestWriteTypeHeadChan(t *testing.T) {
	cs := []struct {
		e         string
		v         reflect.Type
		h, hd, he bool
	}{
		{"chan int", reflect.TypeOf(make(chan int)), false, false, false},
		{"chan " + H("int"), reflect.TypeOf(make(chan int)), false, false, true},
		{"chan int", reflect.TypeOf(make(chan int)), false, true, false},
		{"chan " + H("int"), reflect.TypeOf(make(chan int)), false, true, true},
		{H("chan") + " int", reflect.TypeOf(make(chan int)), true, false, false},
		{H("chan int"), reflect.TypeOf(make(chan int)), true, false, true},
		{H("chan") + " int", reflect.TypeOf(make(chan int)), true, true, false},
		{H("chan int"), reflect.TypeOf(make(chan int)), true, true, true},
		{"<-chan int", reflect.TypeOf(make(<-chan int)), false, false, false},
		{"<-chan " + H("int"), reflect.TypeOf(make(<-chan int)), false, false, true},
		{H("<-") + "chan int", reflect.TypeOf(make(<-chan int)), false, true, false},
		{H("<-") + "chan " + H("int"), reflect.TypeOf(make(<-chan int)), false, true, true},
		{H("<-chan") + " int", reflect.TypeOf(make(<-chan int)), true, false, false},
		{H("<-chan int"), reflect.TypeOf(make(<-chan int)), true, false, true},
		{H("<-chan") + " int", reflect.TypeOf(make(<-chan int)), true, true, false},
		{H("<-chan int"), reflect.TypeOf(make(<-chan int)), true, true, true},
		{"chan<- int", reflect.TypeOf(make(chan<- int)), false, false, false},
		{"chan<- " + H("int"), reflect.TypeOf(make(chan<- int)), false, false, true},
		{"chan" + H("<-") + " int", reflect.TypeOf(make(chan<- int)), false, true, false},
		{"chan" + H("<- int"), reflect.TypeOf(make(chan<- int)), false, true, true},
		{H("chan<-") + " int", reflect.TypeOf(make(chan<- int)), true, false, false},
		{H("chan<- int"), reflect.TypeOf(make(chan<- int)), true, false, true},
		{H("chan<-") + " int", reflect.TypeOf(make(chan<- int)), true, true, false},
		{H("chan<- int"), reflect.TypeOf(make(chan<- int)), true, true, true},
	}
	for i, c := range cs {
		var d tValueDiffer
		d.writeTypeHeadChan(0, c.v, c.h, c.hd)
		d.bufi(0).Write(c.he, "int")
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
		{v: reflect.TypeOf(make(chan Array))},
		{v: reflect.TypeOf(make(<-chan Array))},
		{v: reflect.TypeOf(make(chan<- Array))},
		{v: reflect.TypeOf(func() {})},
		{v: reflect.TypeOf(func(string) {})},
		{v: reflect.TypeOf(func(int, string) {})},
		{v: reflect.TypeOf(func(int, string) float32 { return 0 })},
		{v: reflect.TypeOf(func(int, string) (bool, float32, string) { return true, 0, "a" })},
		{v: reflect.TypeOf(func(Map, Slice) (Uintptr, Float, Str) { return 1, 0, "a" })},
		{v: reflect.TypeOf(new(int))},
		{v: reflect.TypeOf(new(chan int))},
		{v: reflect.TypeOf(new(UPtr))},
		{v: reflect.TypeOf(new(func(int, string)))},
		{v: reflect.TypeOf(new(func(int, string) (bool, int8, uint)))},
		{v: reflect.TypeOf(new(func(Map, Struct) (Uintptr, Float, Chan)))},
		{v: reflect.TypeOf(unsafe.Pointer(new(int)))},
		{v: reflect.ValueOf(struct{ a interface{} }{}).Field(0).Type()},
		{v: reflect.ValueOf(struct{ a I }{}).Field(0).Type()},
		{v: reflect.ValueOf(struct {
			a func(Map, Slice) (Uintptr, Float, Str)
		}{}).Field(0).Type()},
		{v: reflect.TypeOf([...]int{})},
		{v: reflect.TypeOf([...]int{1, 2, 3})},
		{v: reflect.TypeOf([10]Map{})},
		{v: reflect.TypeOf([]int{})},
		{v: reflect.TypeOf([]Array{})},
		{v: reflect.TypeOf(map[bool]int{})},
		{v: reflect.TypeOf(map[Complex]Struct{})},
		{v: reflect.TypeOf(struct{ a int }{}), e: "struct"},
		{v: reflect.TypeOf(A{})},
		{v: reflect.TypeOf(Int(100))},
		{v: reflect.TypeOf(Uint(100))},
		{v: reflect.TypeOf(Uintptr(100))},
		{v: reflect.TypeOf(Float(100))},
		{v: reflect.TypeOf(Complex(100))},
		{v: reflect.TypeOf(Str("100"))},
		{v: reflect.TypeOf(Chan(make(Chan)))},
		{v: reflect.TypeOf(Func(func(int) bool { return false }))},
		{v: reflect.TypeOf(Ptr(new(int)))},
		{v: reflect.TypeOf(UPtr(new(int)))},
		{v: reflect.ValueOf(struct{ a If }{}).Field(0).Type()},
		{v: reflect.TypeOf(Array{})},
		{v: reflect.TypeOf(Slice{})},
		{v: reflect.TypeOf(Map{})},
		{v: reflect.TypeOf(Struct{})},
	}
	for i, c := range cs {
		var d tValueDiffer
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

func TestWriteTypeBeforeValue(t *testing.T) {
	cs := []struct {
		e, h string
		v    reflect.Value
	}{
		{e: "<nil>", v: reflect.ValueOf(nil)},
		{v: reflect.ValueOf(true)},
		{v: reflect.ValueOf(int(100))},
		{v: reflect.ValueOf(int8(100))},
		{v: reflect.ValueOf(int16(100))},
		{v: reflect.ValueOf(int32(100))},
		{v: reflect.ValueOf(int64(100))},
		{v: reflect.ValueOf(uint(100))},
		{v: reflect.ValueOf(uint8(100))},
		{v: reflect.ValueOf(uint16(100))},
		{v: reflect.ValueOf(uint32(100))},
		{v: reflect.ValueOf(uint64(100))},
		{v: reflect.ValueOf(uintptr(100))},
		{v: reflect.ValueOf(float32(100))},
		{v: reflect.ValueOf(float64(100))},
		{v: reflect.ValueOf(complex64(100.1 + 200.2i))},
		{v: reflect.ValueOf(complex128(100.1 + 200.2i))},
		{v: reflect.ValueOf(string("abc"))},
		{e: "(chan int)", h: "(" + H("chan int") + ")", v: reflect.ValueOf(chan int(nil))},
		{e: "(chan int)", h: "(" + H("chan int") + ")", v: reflect.ValueOf(make(chan int))},
		{e: "(func() int)", h: "(" + H("func() int") + ")", v: reflect.ValueOf((func() int)(nil))},
		{e: "(func() int)", h: "(" + H("func() int") + ")", v: reflect.ValueOf(func() int { return 0 })},
		{e: "(*int)", h: "(" + H("*int") + ")", v: reflect.ValueOf((*int)(nil))},
		{e: "(*int)", h: "(" + H("*int") + ")", v: reflect.ValueOf(new(int))},
		{v: reflect.ValueOf(unsafe.Pointer(nil)), e: "(unsafe.Pointer)", h: "(" + H("unsafe.Pointer") + ")"},
		{v: reflect.ValueOf(unsafe.Pointer(new(int))), e: "(unsafe.Pointer)", h: "(" + H("unsafe.Pointer") + ")"},
		{e: "<nil>", v: reflect.ValueOf(struct{ a interface{} }{}).Field(0)},
		{v: reflect.ValueOf(struct{ a I }{}).Field(0)},
		{e: "int", v: reflect.ValueOf(struct{ a interface{} }{100}).Field(0)},
		{v: reflect.ValueOf([...]chan int{})},
		{v: reflect.ValueOf([3]*int{1: new(int)})},
		{v: reflect.ValueOf([11]uint{})},
		{v: reflect.ValueOf([5][]int{1: []int{}, 2: []int{100}, 3: []int{1, 2, 3}})},
		{v: reflect.ValueOf([11][]int{1: []int{}, 2: []int{100}, 3: []int{1, 2, 3}})},
		{v: reflect.ValueOf([]int(nil))},
		{v: reflect.ValueOf([]chan int{})},
		{v: reflect.ValueOf([]*int{1: new(int), 2: nil})},
		{v: reflect.ValueOf(make([]uint, 11))},
		{v: reflect.ValueOf([][]int{1: []int{}, 2: []int{100}, 3: []int{1, 2, 3}, 4: nil})},
		{v: reflect.ValueOf([][]int{1: []int{}, 2: []int{100}, 3: []int{1, 2, 3}, 10: nil})},
		{v: reflect.ValueOf(map[bool]int(nil))},
		{v: reflect.ValueOf(map[bool]int{})},
		{v: reflect.ValueOf(map[bool]int{true: 100, false: 200})},
		{v: reflect.ValueOf(map[bool]chan int{true: nil})},
		{v: reflect.ValueOf(map[*int]int{new(int): 10})},
		{v: reflect.ValueOf(map[bool][]int{true: []int{1, 2, 3}, false: []int{100, 200}})},
		{v: reflect.ValueOf(map[[3]int]bool{[3]int{1, 2, 3}: true, [3]int{100, 200, 300}: false})},
		{e: "struct", v: reflect.ValueOf(struct{}{})},
		{e: "struct", v: reflect.ValueOf(struct {
			a int
			b string
			c []uint
		}{})},
		{e: "struct", v: reflect.ValueOf(struct {
			b string
			c []uint
			a int
		}{c: []uint{1, 2, 3}})},
		{v: reflect.ValueOf(A{})},
		{v: reflect.ValueOf(Int(100))},
		{v: reflect.ValueOf(Uint(100))},
		{v: reflect.ValueOf(Uintptr(100))},
		{v: reflect.ValueOf(Float(100))},
		{v: reflect.ValueOf(Complex(100))},
		{v: reflect.ValueOf(Str("100"))},
		{v: reflect.ValueOf(Chan(make(Chan))), e: "(assert.Chan)", h: "(" + H("assert.Chan") + ")"},
		{v: reflect.ValueOf(Func(func(int) bool { return false })), e: "(assert.Func)", h: "(" + H("assert.Func") + ")"},
		{v: reflect.ValueOf(Ptr(new(int))), e: "(assert.Ptr)", h: "(" + H("assert.Ptr") + ")"},
		{v: reflect.ValueOf(UPtr(new(int))), e: "(assert.UPtr)", h: "(" + H("assert.UPtr") + ")"},
		{v: reflect.ValueOf(struct{ a If }{}).Field(0)},
		{v: reflect.ValueOf(Array{})},
		{v: reflect.ValueOf(Slice{})},
		{v: reflect.ValueOf(Map{})},
		{v: reflect.ValueOf(Struct{})},
	}
	for i, c := range cs {
		var d tValueDiffer
		d.writeTypeBeforeValue(0, c.v, false)
		d.writeTypeBeforeValue(1, c.v, true)
		e, h := c.e, c.h
		if e == "" {
			e = c.v.Type().String()
		}
		if h == "" {
			h = H(e)
		}
		Equal(t, e, d.String(0), "i=%v, r=\n%v", i, d.String(0))
		Equal(t, h, d.String(1), "i=%v, r=\n%v", i, d.String(1))
	}
}

func TestWriteTypeValue(t *testing.T) {
	a := &A{}
	pa := new(int)
	cs := []struct {
		e, e2 string
		v     reflect.Value
		p     bool
	}{
		{e: "<nil>", v: reflect.ValueOf(nil)},
		{v: reflect.ValueOf(true)},
		{v: reflect.ValueOf(int(100))},
		{v: reflect.ValueOf(int8(100))},
		{v: reflect.ValueOf(int16(100))},
		{v: reflect.ValueOf(int32(100))},
		{v: reflect.ValueOf(int64(100))},
		{v: reflect.ValueOf(uint(100))},
		{v: reflect.ValueOf(uint8(100))},
		{v: reflect.ValueOf(uint16(100))},
		{v: reflect.ValueOf(uint32(100))},
		{v: reflect.ValueOf(uint64(100))},
		{e: "uintptr(0x64)", v: reflect.ValueOf(uintptr(100))},
		{v: reflect.ValueOf(float32(100))},
		{v: reflect.ValueOf(float64(100))},
		{e: "complex64(100.1+200.2i)", v: reflect.ValueOf(complex64(100.1 + 200.2i))},
		{e: "complex128(100.1+200.2i)", v: reflect.ValueOf(complex128(100.1 + 200.2i))},
		{e: `string("abc")`, v: reflect.ValueOf(string("abc"))},
		{e: "(chan int)(nil)", v: reflect.ValueOf(chan int(nil))},
		{e: "(chan int)(%v)", v: reflect.ValueOf(make(chan int)), p: true},
		{e: "(func() int)(nil)", v: reflect.ValueOf((func() int)(nil))},
		{e: "(func() int)(%v)", v: reflect.ValueOf(func() int { return 0 }), p: true},
		{e: "(*int)(nil)", v: reflect.ValueOf((*int)(nil))},
		{e: "(*int)(%v)", v: reflect.ValueOf(new(int)), p: true},
		{e: "&[3]int{1, 2, 3}", v: reflect.ValueOf(&[...]int{1, 2, 3})},
		{e: "(*[3]int)(nil)", v: reflect.ValueOf((*[3]int)(nil))},
		{e: "&[]int{1, 2, 3}", v: reflect.ValueOf(&[]int{1, 2, 3})},
		{e: "(*[]int)(nil)", v: reflect.ValueOf((*[]int)(nil))},
		{e: "&map[bool]int{true:100}", v: reflect.ValueOf(&map[bool]int{true: 100})},
		{e: "(*map[bool]int)(nil)", v: reflect.ValueOf((*map[bool]int)(nil))},
		{e: "&assert.A{a:<nil>, b:<nil>}", v: reflect.ValueOf(&A{})},
		{e: "(*assert.A)(nil)", v: reflect.ValueOf((*A)(nil))},
		{e: fmt.Sprintf("&assert.A{a:<nil>, b:%p}", a), v: reflect.ValueOf(&A{b: a})},
		{e: "(unsafe.Pointer)(nil)", v: reflect.ValueOf(unsafe.Pointer(nil))},
		{e: "(unsafe.Pointer)(%v)", v: reflect.ValueOf(unsafe.Pointer(new(int))), p: true},
		{e: "<nil>", v: reflect.ValueOf(struct{ a interface{} }{}).Field(0)},
		{e: "assert.I(nil)", v: reflect.ValueOf(struct{ a I }{}).Field(0)},
		{e: "int(100)", v: reflect.ValueOf(struct{ a interface{} }{100}).Field(0)},
		{e: "[0]chan int{}", v: reflect.ValueOf([...]chan int{})},
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
		{e: "map[bool]int(nil)", v: reflect.ValueOf(map[bool]int(nil))},
		{e: "map[bool]int{}", v: reflect.ValueOf(map[bool]int{})},
		{e: "map[bool]int{true:100, false:200}", e2: "map[bool]int{false:200, true:100}", v: reflect.ValueOf(map[bool]int{true: 100, false: 200})},
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
		{e: "struct{}", v: reflect.ValueOf(struct{}{})},
		{e: `struct{a:0, b:"", c:<nil>}`, v: reflect.ValueOf(struct {
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
		{e: "assert.A{a:<nil>, b:<nil>}", v: reflect.ValueOf(A{})},
		{v: reflect.ValueOf(Int(100)), e: "assert.Int(String of Int)"},
		{v: reflect.ValueOf(Uint(100)), e: "assert.Uint(String of Uint)"},
		{v: reflect.ValueOf(Uintptr(100)), e: "assert.Uintptr(String of Uintptr)"},
		{v: reflect.ValueOf(Float(100))},
		{v: reflect.ValueOf(Complex(100.1 + 200.2i)), e: "assert.Complex(100.1+200.2i)"},
		{v: reflect.ValueOf(Str("100")), e: `assert.Str("100")`},
		{v: reflect.ValueOf(Chan(nil)), e: "(assert.Chan)(nil)"},
		{v: reflect.ValueOf(Chan(make(Chan))), e: "(assert.Chan)(%v)", p: true},
		{v: reflect.ValueOf(Func(nil)), e: "(assert.Func)(nil)"},
		{v: reflect.ValueOf(Func(func(int) bool { return false })), e: "(assert.Func)(%v)", p: true},
		{v: reflect.ValueOf(Ptr(nil)), e: "(assert.Ptr)(nil)"},
		{v: reflect.ValueOf(Ptr(new(int))), e: "(assert.Ptr)(%v)", p: true},
		{v: reflect.ValueOf(UPtr(nil)), e: "(assert.UPtr)(nil)"},
		{v: reflect.ValueOf(UPtr(new(int))), e: "(assert.UPtr)(%v)", p: true},
		{v: reflect.ValueOf(struct{ a If }{}).Field(0), e: "assert.If(nil)"},
		{v: reflect.ValueOf(Array{}), e: "assert.Array{<nil>, <nil>, <nil>}"},
		{v: reflect.ValueOf(&Array{}), e: "&assert.Array{<nil>, <nil>, <nil>}"},
		{v: reflect.ValueOf((*Array)(nil)), e: "(*assert.Array)(nil)"},
		{v: reflect.ValueOf(Slice(nil)), e: "assert.Slice(nil)"},
		{v: reflect.ValueOf(Slice{}), e: "assert.Slice{}"},
		{v: reflect.ValueOf(&Slice{}), e: "&assert.Slice{}"},
		{v: reflect.ValueOf(make(Slice, 3)), e: "assert.Slice{<nil>, <nil>, <nil>}"},
		{v: reflect.ValueOf(&Slice{a, 2, 3}), e: fmt.Sprintf("&assert.Slice{%p, 2, 3}", a)},
		{v: reflect.ValueOf((*Slice)(nil)), e: "(*assert.Slice)(nil)"},
		{v: reflect.ValueOf(Map(nil)), e: "assert.Map(nil)"},
		{v: reflect.ValueOf(Map{}), e: "assert.Map{}"},
		{v: reflect.ValueOf(&Map{}), e: "&assert.Map{}"},
		{v: reflect.ValueOf(Map{10: true}), e: "assert.Map{10:true}"},
		{v: reflect.ValueOf(&Map{10: true}), e: "&assert.Map{10:true}"},
		{v: reflect.ValueOf(Map{10: true, 20: false}), e: "assert.Map{10:true, 20:false}", e2: "assert.Map{20:false, 10:true}"},
		{v: reflect.ValueOf((*Map)(nil)), e: "(*assert.Map)(nil)"},
		{v: reflect.ValueOf(Struct{}), e: "assert.Struct{a:<nil>, b:<nil>}"},
		{v: reflect.ValueOf(&Struct{}), e: "&assert.Struct{a:<nil>, b:<nil>}"},
		{v: reflect.ValueOf((*Struct)(nil)), e: "(*assert.Struct)(nil)"},
		{v: reflect.ValueOf(reflect.Array), e: "reflect.Kind(array)"},
		{v: reflect.ValueOf(errors.New("abc")), e: `&errors.errorString{s:"abc"}`},
	}
	for i, c := range cs {
		var d tValueDiffer
		d.writeTypeValue(0, c.v, false, false)
		e := c.e
		if c.p {
			e = fmt.Sprintf(e, c.v)
		}
		if e == "" {
			e = fmt.Sprintf("%v(%v)", c.v.Type(), c.v)
		}
		if c.e2 == "" || e == d.String(0) {
			Equal(t, e, d.String(0), "i=%v, r=\n%v", i, d.String(0))
		} else {
			Equal(t, c.e2, d.String(0), "i=%v, r=\n%v", i, d.String(0))
		}
	}
}
