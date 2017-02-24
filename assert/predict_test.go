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
	"reflect"
	"testing"
	"unsafe"
)

func TestValueEqual(t *testing.T) {
	a := make(chan int)
	b := func() int { return 0 }
	c := new(int)
	*c = 100
	cs := []struct {
		v1, v2 reflect.Value
		e, ci  bool
		k, k2  reflect.Kind
	}{
		{v1: reflect.ValueOf(nil), v2: reflect.ValueOf(nil), e: true},
		{v1: reflect.ValueOf(nil), v2: reflect.ValueOf(100)},
		{v1: reflect.ValueOf(100.5), v2: reflect.ValueOf(100), ci: true},
		{v1: reflect.ValueOf(100), v2: reflect.ValueOf(100), ci: true},
		{v1: reflect.ValueOf(101), v2: reflect.ValueOf(100), ci: true},
		{v1: reflect.ValueOf(struct{ a bool }{true}).Field(0), v2: reflect.ValueOf(struct{ a bool }{false}).Field(0), k: reflect.Bool},
		{v1: reflect.ValueOf(struct{ a bool }{true}).Field(0), v2: reflect.ValueOf(struct{ a bool }{true}).Field(0), k: reflect.Bool, e: true},
		{v1: reflect.ValueOf(struct{ a int }{100}).Field(0), v2: reflect.ValueOf(struct{ a int }{101}).Field(0), k: reflect.Int},
		{v1: reflect.ValueOf(struct{ a int }{101}).Field(0), v2: reflect.ValueOf(struct{ a int }{101}).Field(0), k: reflect.Int, e: true},
		{v1: reflect.ValueOf(struct{ a int8 }{100}).Field(0), v2: reflect.ValueOf(struct{ a int8 }{101}).Field(0), k: reflect.Int8},
		{v1: reflect.ValueOf(struct{ a int8 }{101}).Field(0), v2: reflect.ValueOf(struct{ a int8 }{101}).Field(0), k: reflect.Int8, e: true},
		{v1: reflect.ValueOf(struct{ a int16 }{100}).Field(0), v2: reflect.ValueOf(struct{ a int16 }{101}).Field(0), k: reflect.Int16},
		{v1: reflect.ValueOf(struct{ a int16 }{101}).Field(0), v2: reflect.ValueOf(struct{ a int16 }{101}).Field(0), k: reflect.Int16, e: true},
		{v1: reflect.ValueOf(struct{ a int32 }{100}).Field(0), v2: reflect.ValueOf(struct{ a int32 }{101}).Field(0), k: reflect.Int32},
		{v1: reflect.ValueOf(struct{ a int32 }{101}).Field(0), v2: reflect.ValueOf(struct{ a int32 }{101}).Field(0), k: reflect.Int32, e: true},
		{v1: reflect.ValueOf(struct{ a int64 }{100}).Field(0), v2: reflect.ValueOf(struct{ a int64 }{101}).Field(0), k: reflect.Int64},
		{v1: reflect.ValueOf(struct{ a int64 }{101}).Field(0), v2: reflect.ValueOf(struct{ a int64 }{101}).Field(0), k: reflect.Int64, e: true},
		{v1: reflect.ValueOf(struct{ a uint }{100}).Field(0), v2: reflect.ValueOf(struct{ a uint }{101}).Field(0), k: reflect.Uint},
		{v1: reflect.ValueOf(struct{ a uint }{101}).Field(0), v2: reflect.ValueOf(struct{ a uint }{101}).Field(0), k: reflect.Uint, e: true},
		{v1: reflect.ValueOf(struct{ a uint8 }{100}).Field(0), v2: reflect.ValueOf(struct{ a uint8 }{101}).Field(0), k: reflect.Uint8},
		{v1: reflect.ValueOf(struct{ a uint8 }{101}).Field(0), v2: reflect.ValueOf(struct{ a uint8 }{101}).Field(0), k: reflect.Uint8, e: true},
		{v1: reflect.ValueOf(struct{ a uint16 }{100}).Field(0), v2: reflect.ValueOf(struct{ a uint16 }{101}).Field(0), k: reflect.Uint16},
		{v1: reflect.ValueOf(struct{ a uint16 }{101}).Field(0), v2: reflect.ValueOf(struct{ a uint16 }{101}).Field(0), k: reflect.Uint16, e: true},
		{v1: reflect.ValueOf(struct{ a uint32 }{100}).Field(0), v2: reflect.ValueOf(struct{ a uint32 }{101}).Field(0), k: reflect.Uint32},
		{v1: reflect.ValueOf(struct{ a uint32 }{101}).Field(0), v2: reflect.ValueOf(struct{ a uint32 }{101}).Field(0), k: reflect.Uint32, e: true},
		{v1: reflect.ValueOf(struct{ a uint64 }{100}).Field(0), v2: reflect.ValueOf(struct{ a uint64 }{101}).Field(0), k: reflect.Uint64},
		{v1: reflect.ValueOf(struct{ a uint64 }{101}).Field(0), v2: reflect.ValueOf(struct{ a uint64 }{101}).Field(0), k: reflect.Uint64, e: true},
		{v1: reflect.ValueOf(struct{ a uintptr }{100}).Field(0), v2: reflect.ValueOf(struct{ a uintptr }{101}).Field(0), k: reflect.Uintptr},
		{v1: reflect.ValueOf(struct{ a uintptr }{101}).Field(0), v2: reflect.ValueOf(struct{ a uintptr }{101}).Field(0), k: reflect.Uintptr, e: true},
		{v1: reflect.ValueOf(struct{ a float32 }{100.25}).Field(0), v2: reflect.ValueOf(struct{ a float32 }{101.25}).Field(0), k: reflect.Float32},
		{v1: reflect.ValueOf(struct{ a float32 }{101.25}).Field(0), v2: reflect.ValueOf(struct{ a float32 }{101.25}).Field(0), k: reflect.Float32, e: true},
		{v1: reflect.ValueOf(struct{ a float64 }{100.25}).Field(0), v2: reflect.ValueOf(struct{ a float64 }{101.25}).Field(0), k: reflect.Float64},
		{v1: reflect.ValueOf(struct{ a float64 }{101.25}).Field(0), v2: reflect.ValueOf(struct{ a float64 }{101.25}).Field(0), k: reflect.Float64, e: true},
		{v1: reflect.ValueOf(struct{ a complex64 }{100.25 + 200.5i}).Field(0), v2: reflect.ValueOf(struct{ a complex64 }{101.25 + 200.5i}).Field(0), k: reflect.Complex64},
		{v1: reflect.ValueOf(struct{ a complex64 }{101.25 + 200.5i}).Field(0), v2: reflect.ValueOf(struct{ a complex64 }{101.25 + 200.5i}).Field(0), k: reflect.Complex64, e: true},
		{v1: reflect.ValueOf(struct{ a complex128 }{100.25 + 200.5i}).Field(0), v2: reflect.ValueOf(struct{ a complex128 }{101.25 + 200.5i}).Field(0), k: reflect.Complex128},
		{v1: reflect.ValueOf(struct{ a complex128 }{101.25 + 200.5i}).Field(0), v2: reflect.ValueOf(struct{ a complex128 }{101.25 + 200.5i}).Field(0), k: reflect.Complex128, e: true},
		{v1: reflect.ValueOf(struct{ a string }{""}).Field(0), v2: reflect.ValueOf(struct{ a string }{""}).Field(0), k: reflect.String, e: true},
		{v1: reflect.ValueOf(struct{ a string }{"abc"}).Field(0), v2: reflect.ValueOf(struct{ a string }{""}).Field(0), k: reflect.String},
		{v1: reflect.ValueOf(struct{ a string }{"abc"}).Field(0), v2: reflect.ValueOf(struct{ a string }{"abc"}).Field(0), k: reflect.String, e: true},
		{v1: reflect.ValueOf(struct{ a string }{"abc"}).Field(0), v2: reflect.ValueOf(struct{ a string }{"abd"}).Field(0), k: reflect.String},
		{v1: reflect.ValueOf(struct{ a chan int }{}).Field(0), v2: reflect.ValueOf(struct{ a chan int }{make(chan int)}).Field(0), k: reflect.Chan},
		{v1: reflect.ValueOf(struct{ a chan int }{make(chan int)}).Field(0), v2: reflect.ValueOf(struct{ a chan int }{make(chan int)}).Field(0), k: reflect.Chan},
		{v1: reflect.ValueOf(struct{ a chan int }{}).Field(0), v2: reflect.ValueOf(struct{ a chan int }{}).Field(0), k: reflect.Chan, e: true},
		{v1: reflect.ValueOf(struct{ a chan int }{a}).Field(0), v2: reflect.ValueOf(struct{ a chan int }{a}).Field(0), k: reflect.Chan, e: true},
		{v1: reflect.ValueOf(struct{ a <-chan int }{}).Field(0), v2: reflect.ValueOf(struct{ a <-chan int }{make(<-chan int)}).Field(0), k: reflect.Chan},
		{v1: reflect.ValueOf(struct{ a <-chan int }{make(<-chan int)}).Field(0), v2: reflect.ValueOf(struct{ a <-chan int }{make(<-chan int)}).Field(0), k: reflect.Chan},
		{v1: reflect.ValueOf(struct{ a <-chan int }{}).Field(0), v2: reflect.ValueOf(struct{ a <-chan int }{}).Field(0), k: reflect.Chan, e: true},
		{v1: reflect.ValueOf(struct{ a <-chan int }{a}).Field(0), v2: reflect.ValueOf(struct{ a <-chan int }{a}).Field(0), k: reflect.Chan, e: true},
		{v1: reflect.ValueOf(struct{ a chan<- int }{}).Field(0), v2: reflect.ValueOf(struct{ a chan<- int }{make(chan<- int)}).Field(0), k: reflect.Chan},
		{v1: reflect.ValueOf(struct{ a chan<- int }{make(chan<- int)}).Field(0), v2: reflect.ValueOf(struct{ a chan<- int }{make(chan<- int)}).Field(0), k: reflect.Chan},
		{v1: reflect.ValueOf(struct{ a chan<- int }{}).Field(0), v2: reflect.ValueOf(struct{ a chan<- int }{}).Field(0), k: reflect.Chan, e: true},
		{v1: reflect.ValueOf(struct{ a chan<- int }{a}).Field(0), v2: reflect.ValueOf(struct{ a chan<- int }{a}).Field(0), k: reflect.Chan, e: true},
		{v1: reflect.ValueOf(struct{ a func() int }{}).Field(0), v2: reflect.ValueOf(struct{ a func() int }{}).Field(0), k: reflect.Func, e: true},
		{v1: reflect.ValueOf(struct{ a func() int }{func() int { return 1 }}).Field(0), v2: reflect.ValueOf(struct{ a func() int }{}).Field(0), k: reflect.Func},
		{v1: reflect.ValueOf(struct{ a func() int }{func() int { return 1 }}).Field(0), v2: reflect.ValueOf(struct{ a func() int }{func() int { return 2 }}).Field(0), k: reflect.Func},
		{v1: reflect.ValueOf(struct{ a func() int }{b}).Field(0), v2: reflect.ValueOf(struct{ a func() int }{b}).Field(0), k: reflect.Func},
		{v1: reflect.ValueOf(struct{ a interface{} }{}).Field(0), v2: reflect.ValueOf(struct{ a interface{} }{}).Field(0), k: reflect.Interface, e: true},
		{v1: reflect.ValueOf(struct{ a interface{} }{100}).Field(0), v2: reflect.ValueOf(struct{ a interface{} }{}).Field(0), k: reflect.Interface},
		{v1: reflect.ValueOf(struct{ a interface{} }{100}).Field(0), v2: reflect.ValueOf(struct{ a interface{} }{101}).Field(0), k: reflect.Interface},
		{v1: reflect.ValueOf(struct{ a interface{} }{101}).Field(0), v2: reflect.ValueOf(struct{ a interface{} }{101}).Field(0), k: reflect.Interface, e: true},
		{v1: reflect.ValueOf(struct{ a interface{} }{100}).Field(0), v2: reflect.ValueOf(100), e: true, k: reflect.Interface, k2: reflect.Int},
		{v1: reflect.ValueOf(struct{ a interface{} }{A{}}).Field(0), v2: reflect.ValueOf(A{}), e: true, k: reflect.Interface, k2: reflect.Struct},
		{v1: reflect.ValueOf(struct{ a I }{Int(100)}).Field(0), v2: reflect.ValueOf(Int(100)), e: true, k: reflect.Interface, k2: reflect.Int},
		{v1: reflect.ValueOf(struct{ a I }{A{}}).Field(0), v2: reflect.ValueOf(A{}), e: true, k: reflect.Interface, k2: reflect.Struct},
		{v1: reflect.ValueOf(struct{ a I }{}).Field(0), v2: reflect.ValueOf(struct{ a I }{}).Field(0), k: reflect.Interface, e: true},
		{v1: reflect.ValueOf(struct{ a I }{A{}}).Field(0), v2: reflect.ValueOf(struct{ a I }{}).Field(0), k: reflect.Interface},
		{v1: reflect.ValueOf(struct{ a I }{A{a: 100}}).Field(0), v2: reflect.ValueOf(struct{ a I }{A{a: 101}}).Field(0), k: reflect.Interface},
		{v1: reflect.ValueOf(struct{ a I }{A{a: 101}}).Field(0), v2: reflect.ValueOf(struct{ a I }{A{a: 101}}).Field(0), k: reflect.Interface, e: true},
		{v1: reflect.ValueOf(struct{ a *int }{}).Field(0), v2: reflect.ValueOf(struct{ a *int }{}).Field(0), k: reflect.Ptr, e: true},
		{v1: reflect.ValueOf(struct{ a *int }{new(int)}).Field(0), v2: reflect.ValueOf(struct{ a *int }{}).Field(0), k: reflect.Ptr},
		{v1: reflect.ValueOf(struct{ a *int }{c}).Field(0), v2: reflect.ValueOf(struct{ a *int }{c}).Field(0), k: reflect.Ptr, e: true},
		{v1: reflect.ValueOf(struct{ a *int }{new(int)}).Field(0), v2: reflect.ValueOf(struct{ a *int }{new(int)}).Field(0), k: reflect.Ptr, e: true},
		{v1: reflect.ValueOf(struct{ a *int }{c}).Field(0), v2: reflect.ValueOf(struct{ a *int }{new(int)}).Field(0), k: reflect.Ptr},
		{v1: reflect.ValueOf(struct{ a unsafe.Pointer }{}).Field(0), v2: reflect.ValueOf(struct{ a unsafe.Pointer }{}).Field(0), k: reflect.UnsafePointer, e: true},
		{v1: reflect.ValueOf(struct{ a unsafe.Pointer }{unsafe.Pointer(new(int))}).Field(0), v2: reflect.ValueOf(struct{ a unsafe.Pointer }{}).Field(0), k: reflect.UnsafePointer},
		{v1: reflect.ValueOf(struct{ a unsafe.Pointer }{unsafe.Pointer(c)}).Field(0), v2: reflect.ValueOf(struct{ a unsafe.Pointer }{unsafe.Pointer(c)}).Field(0), k: reflect.UnsafePointer, e: true},
		{v1: reflect.ValueOf(struct{ a unsafe.Pointer }{unsafe.Pointer(new(int))}).Field(0), v2: reflect.ValueOf(struct{ a unsafe.Pointer }{unsafe.Pointer(new(int))}).Field(0), k: reflect.UnsafePointer},
		{v1: reflect.ValueOf(struct{ a [0]int }{}).Field(0), v2: reflect.ValueOf(struct{ a [0]int }{}).Field(0), k: reflect.Array, e: true},
		{v1: reflect.ValueOf(struct{ a [3]int }{}).Field(0), v2: reflect.ValueOf(struct{ a [3]int }{}).Field(0), k: reflect.Array, e: true},
		{v1: reflect.ValueOf(struct{ a [3]int }{[3]int{2: 1}}).Field(0), v2: reflect.ValueOf(struct{ a [3]int }{}).Field(0), k: reflect.Array},
		{v1: reflect.ValueOf(struct{ a []int }{}).Field(0), v2: reflect.ValueOf(struct{ a []int }{}).Field(0), k: reflect.Slice, e: true},
		{v1: reflect.ValueOf(struct{ a []int }{[]int{5: 1}}).Field(0), v2: reflect.ValueOf(struct{ a []int }{}).Field(0), k: reflect.Slice},
		{v1: reflect.ValueOf(struct{ a []int }{[]int{5: 1}}).Field(0), v2: reflect.ValueOf(struct{ a []int }{[]int{5: 1}}).Field(0), k: reflect.Slice, e: true},
		{v1: reflect.ValueOf(struct{ a []int }{[]int{5: 2}}).Field(0), v2: reflect.ValueOf(struct{ a []int }{[]int{5: 1}}).Field(0), k: reflect.Slice},
		{v1: reflect.ValueOf(struct{ a map[int]bool }{}).Field(0), v2: reflect.ValueOf(struct{ a map[int]bool }{}).Field(0), k: reflect.Map, e: true},
		{v1: reflect.ValueOf(struct{ a map[int]bool }{make(map[int]bool)}).Field(0), v2: reflect.ValueOf(struct{ a map[int]bool }{}).Field(0), k: reflect.Map},
		{v1: reflect.ValueOf(struct{ a map[int]bool }{make(map[int]bool)}).Field(0), v2: reflect.ValueOf(struct{ a map[int]bool }{make(map[int]bool)}).Field(0), k: reflect.Map, e: true},
		{v1: reflect.ValueOf(struct{ a map[int]bool }{make(map[int]bool)}).Field(0), v2: reflect.ValueOf(struct{ a map[int]bool }{map[int]bool{1: true, 2: false}}).Field(0), k: reflect.Map},
		{v1: reflect.ValueOf(struct{ a map[int]bool }{map[int]bool{1: true, 2: false}}).Field(0), v2: reflect.ValueOf(struct{ a map[int]bool }{map[int]bool{1: true, 2: false}}).Field(0), k: reflect.Map, e: true},
		{v1: reflect.ValueOf(struct{ a map[int]bool }{map[int]bool{1: true, 2: true}}).Field(0), v2: reflect.ValueOf(struct{ a map[int]bool }{map[int]bool{1: true, 2: false}}).Field(0), k: reflect.Map},
		{v1: reflect.ValueOf(struct{ a map[int]bool }{map[int]bool{1: true, 2: false, 3: true}}).Field(0), v2: reflect.ValueOf(struct{ a map[int]bool }{map[int]bool{1: true, 2: false}}).Field(0), k: reflect.Map},
		{v1: reflect.ValueOf(struct{ a map[int]bool }{map[int]bool{2: false, 3: true}}).Field(0), v2: reflect.ValueOf(struct{ a map[int]bool }{map[int]bool{1: true, 2: false}}).Field(0), k: reflect.Map},
		{v1: reflect.ValueOf(struct{ a struct{ a int } }{}).Field(0), v2: reflect.ValueOf(struct{ a struct{ a int } }{}).Field(0), k: reflect.Struct, e: true},
		{v1: reflect.ValueOf(struct{ a struct{ a int } }{struct{ a int }{100}}).Field(0), v2: reflect.ValueOf(struct{ a struct{ a int } }{}).Field(0), k: reflect.Struct},
		{v1: reflect.ValueOf(struct{ a A }{}).Field(0), v2: reflect.ValueOf(struct{ a A }{}).Field(0), k: reflect.Struct, e: true},
		{v1: reflect.ValueOf(struct{ a A }{A{a: 100}}).Field(0), v2: reflect.ValueOf(struct{ a A }{}).Field(0), k: reflect.Struct},
	}
	for i, c := range cs {
		ci := c.v1.IsValid() && c.v1.CanInterface() && c.v2.IsValid() && c.v2.CanInterface()
		Equal(t, c.ci, ci, "i=%v", i)
		if c.k != reflect.Invalid {
			Equal(t, c.k, c.v1.Kind(), "i=%v", i)
			if c.k2 == reflect.Invalid {
				Equal(t, c.k, c.v2.Kind(), "i=%v", i)
			} else {
				Equal(t, c.k2, c.v2.Kind(), "i=%v", i)
			}
		}
		a1 := valueEqual(c.v1, c.v2)
		a2 := valueEqual(c.v2, c.v1)
		e := c.e
		if ci {
			e = reflect.DeepEqual(c.v1.Interface(), c.v2.Interface())
		}
		Equal(t, e, a1, "i=%v, r1=\n%v\n%v", i, c.v1, c.v2)
		Equal(t, e, a2, "i=%v, r2=\n%v\n%v", i, c.v1, c.v2)
	}
}

func TestIsNil(t *testing.T) {
	True(t, isNil(nil))
	False(t, isNil(true))
	False(t, isNil(false))
	False(t, isNil(int(0)))
	False(t, isNil(int(100)))
	False(t, isNil(Int(0)))
	False(t, isNil(Int(100)))
	False(t, isNil(int8(0)))
	False(t, isNil(int8(100)))
	False(t, isNil(int16(0)))
	False(t, isNil(int16(100)))
	False(t, isNil(int32(0)))
	False(t, isNil(int32(100)))
	False(t, isNil(int64(0)))
	False(t, isNil(int64(100)))
	False(t, isNil(uint(0)))
	False(t, isNil(uint(100)))
	False(t, isNil(Uint(0)))
	False(t, isNil(Uint(100)))
	False(t, isNil(uint8(0)))
	False(t, isNil(uint8(100)))
	False(t, isNil(uint16(0)))
	False(t, isNil(uint16(100)))
	False(t, isNil(uint32(0)))
	False(t, isNil(uint32(100)))
	False(t, isNil(uint64(0)))
	False(t, isNil(uint64(100)))
	False(t, isNil(uintptr(0)))
	False(t, isNil(uintptr(100)))
	False(t, isNil(Uintptr(0)))
	False(t, isNil(Uintptr(100)))
	False(t, isNil(float32(0)))
	False(t, isNil(float32(100)))
	False(t, isNil(float64(0)))
	False(t, isNil(float64(100)))
	False(t, isNil(Float(0)))
	False(t, isNil(Float(100)))
	False(t, isNil(complex64(0)))
	False(t, isNil(complex64(100)))
	False(t, isNil(complex128(0)))
	False(t, isNil(complex128(100)))
	False(t, isNil(Complex(0)))
	False(t, isNil(Complex(100)))
	False(t, isNil(string("")))
	False(t, isNil(string("abc")))
	False(t, isNil(Str("")))
	False(t, isNil(Str("abc")))
	True(t, isNil(chan int(nil)))
	False(t, isNil(make(chan int)))
	True(t, isNil(Chan(nil)))
	False(t, isNil(make(Chan)))
	True(t, isNil((func(int) bool)(nil)))
	False(t, isNil(func(int) bool { return false }))
	True(t, isNil(Func(nil)))
	False(t, isNil(Func(func(int) bool { return false })))
	True(t, isNil((*bool)(nil)))
	False(t, isNil(new(bool)))
	True(t, isNil(Ptr(nil)))
	False(t, isNil(Ptr(new(int))))
	True(t, isNil(unsafe.Pointer(nil)))
	False(t, isNil(unsafe.Pointer(new(int))))
	True(t, isNil(UPtr(nil)))
	False(t, isNil(UPtr(new(int))))
	True(t, isNil(A{}.a))
	True(t, isNil(A{}.b))
	False(t, isNil(A{a: 100}.a))
	False(t, isNil(A{b: A{}}.b))
	False(t, isNil([...]int{}))
	False(t, isNil([...]int{1, 2, 3}))
	False(t, isNil(Array{}))
	True(t, isNil([]int(nil)))
	False(t, isNil([]int{}))
	False(t, isNil([]int{1, 2, 3}))
	True(t, isNil(Slice(nil)))
	False(t, isNil(Slice{}))
	False(t, isNil(Slice{1, 2, 3}))
	True(t, isNil(map[bool]int(nil)))
	False(t, isNil(map[bool]int{}))
	False(t, isNil(map[bool]int{true: 100}))
	True(t, isNil(Map(nil)))
	False(t, isNil(Map{}))
	False(t, isNil(Map{true: 100}))
	False(t, isNil(A{}))
	False(t, isNil(Struct{}))
}

func TestIsSameInValue(t *testing.T) {
	a := new(int)
	Z1 := "000 0000 000000 00000000 0000000000 0000000000 0000 0000 0000 0000 0000 0000 00000000000 00000000000000 0000 0000 000 00 000 000"
	N1 := "100 0000 000000 00000000 0000000000 0000000000 0000 0000 0000 0000 0000 0000 00000000000 00000000000000 0000 0000 000 00 000 000"
	cs := []interface{}{
		nil, true, false,
		int8(-0x80), int8(0), int8(0x7F), int8(-1),
		int16(-0x8000), int16(-0x80), int16(0), int16(0x7F), int16(0x7FFF), int16(-1),
		int32(-0x80000000), int32(-0x8000), int32(-0x80), int32(0), int32(0x7F), int32(0x7FFF), int32(0x7FFFFFFF), int32(-1),
		int64(-0x8000000000000000), int64(-0x80000000), int64(-0x8000), int64(-0x80), int64(0), int64(0x7F), int64(0x7FFF), int64(0x7FFFFFFF), int64(0x7FFFFFFFFFFFFFFF), int64(-1),
		int(-0x8000000000000000), int(-0x80000000), int(-0x8000), int(-0x80), int(0), int(0x7F), int(0x7FFF), int(0x7FFFFFFF), int(0x7FFFFFFFFFFFFFFF), int(-1),
		uint8(0), uint8(0x7F), uint(0x80), uint8(0xFF),
		uint16(0), uint16(0x7FFF), uint16(0x8000), uint16(0xFFFF),
		uint32(0), uint32(0x7FFFFFFF), uint32(0x80000000), uint32(0xFFFFFFFF),
		uint64(0), uint64(0x7FFFFFFFFFFFFFFF), uint64(0x8000000000000000), uint64(0xFFFFFFFFFFFFFFFF),
		uint(0), uint(0x7FFFFFFFFFFFFFFF), uint(0x8000000000000000), uint(0xFFFFFFFFFFFFFFFF),
		uintptr(0), uintptr(0x7FFFFFFFFFFFFFFF), uintptr(0x8000000000000000), uintptr(0xFFFFFFFFFFFFFFFF),
		float32(-0x8000000000000000), float32(-0x80000000), float32(-0x8000), float32(-0x80), float32(0), float32(0x7F), float32(0x7FFF), float32(0x80), float32(0x8000), float32(0xFF), float32(0xFFFF),
		float64(-0x8000000000000000), float64(-0x80000000), float64(-0x8000), float64(-0x80), float64(0), float64(0x7F), float64(0x7FFF), float64(0x7FFFFFFF), float64(0x80), float64(0x8000), float64(0x80000000), float64(0xFF), float64(0xFFFF), float64(0xFFFFFFFF),
		complex64(0), complex64(0x7F), complex64(0 + 0.25i), complex64(0x7F + 0.25i),
		complex128(0), complex128(0x7F), complex128(0 + 0.25i), complex128(0x7F + 0.25i),
		(*int)(nil), (*int)(a), (*uint)(nil),
		unsafe.Pointer(nil), unsafe.Pointer(a),
		chan int(nil), make(chan int), chan uint(nil),
		(func(int) bool)(nil), func(int) bool { return true }, (func(uint) bool)(nil),
		string(""), string("abc"), string("中文"),
		[0]int{}, [...]int{1, 2, 3}, [...]interface{}{uint(1), float32(2), complex128(3)}, [...]int{1, 2, 4}, [0]interface{}{}, [0]chan int{},
		[]int(nil), []int{}, []int{1, 2, 3}, []interface{}{uint(1), float32(2), complex128(3)}, []int{1, 2, 4}, []interface{}{}, []chan int{},
		map[int]float32(nil), map[int]float32{}, map[int]float32{10: 1.25}, map[int]float32{10: 1.50}, map[interface{}]float32{10: 1.25}, map[int]interface{}{10: 1.25}, map[interface{}]float32{}, map[int]interface{}{},
		A{}, A{a: int(100)}, A{a: uint(100)}, A{b: A{a: int(100)}}, A{b: A{a: complex(100, 0)}},
		100, "d", []byte("abc"), []byte("中文"), []rune("abc"), []rune("中文"), // custom 1
		[...]byte{'a', 'b', 'c'}, [...]rune{'a', 'b', 'c'}, [...]int{'a', 'b', 'c'}, // custom 2
		[...]byte{0xe4, 0xb8, 0xad, 0xe6, 0x96, 0x87}, [...]rune{'中', '文'}, [...]int{'中', '文'}, []int{'中', '文'}, // custom 3
		Bool(true),
	}
	rs := []string{
		"", "", "",
		"000", "000", "000", "000", // int8
		"000 0000", "000 1000", "000 0100", "000 0010", "000 0000", "000 0001", // int16
		// int32
		"000 0000 000000",
		"000 0000 100000",
		"000 1000 010000",
		"000 0100 001000",
		"000 0010 000100",
		"000 0000 000010",
		"000 0000 000000",
		"000 0001 000001",
		// int64
		"000 0000 000000 00000000",
		"000 0000 000000 10000000",
		"000 0000 100000 01000000",
		"000 1000 010000 00100000",
		"000 0100 001000 00010000",
		"000 0010 000100 00001000",
		"000 0000 000010 00000100",
		"000 0000 000000 00000010",
		"000 0000 000000 00000000",
		"000 0001 000001 00000001",
		// int
		"000 0000 000000 00000000 1000000000",
		"000 0000 000000 10000000 0100000000",
		"000 0000 100000 01000000 0010000000",
		"000 1000 010000 00100000 0001000000",
		"000 0100 001000 00010000 0000100000",
		"000 0010 000100 00001000 0000010000",
		"000 0000 000010 00000100 0000001000",
		"000 0000 000000 00000010 0000000100",
		"000 0000 000000 00000000 0000000010",
		"000 0001 000001 00000001 0000000001",
		// uint8
		"000 0100 001000 00010000 0000100000 0000100000",
		"000 0010 000100 00001000 0000010000 0000010000",
		"000 0000 000000 00000000 0000000000 0000000000",
		"000 0000 000000 00000000 0000000000 0000000000",
		// uint16
		"000 0100 001000 00010000 0000100000 0000100000 1000",
		"000 0000 000010 00000100 0000001000 0000001000 0000",
		"000 0000 000000 00000000 0000000000 0000000000 0000",
		"000 0000 000000 00000000 0000000000 0000000000 0000",
		// uint32
		"000 0100 001000 00010000 0000100000 0000100000 1000 1000",
		"000 0000 000000 00000010 0000000100 0000000100 0000 0000",
		"000 0000 000000 00000000 0000000000 0000000000 0000 0000",
		"000 0000 000000 00000000 0000000000 0000000000 0000 0000",
		// uint64
		"000 0100 001000 00010000 0000100000 0000100000 1000 1000 1000",
		"000 0000 000000 00000000 0000000010 0000000010 0000 0000 0000",
		"000 0000 000000 00000000 0000000000 0000000000 0000 0000 0000",
		"000 0000 000000 00000000 0000000000 0000000000 0000 0000 0000",
		// uint
		"000 0100 001000 00010000 0000100000 0000100000 1000 1000 1000 1000",
		"000 0000 000000 00000000 0000000010 0000000010 0000 0000 0000 0100",
		"000 0000 000000 00000000 0000000000 0000000000 0000 0000 0000 0010",
		"000 0000 000000 00000000 0000000000 0000000000 0000 0000 0000 0001",
		// uintptr
		"000 0100 001000 00010000 0000100000 0000100000 1000 1000 1000 1000 1000",
		"000 0000 000000 00000000 0000000010 0000000010 0000 0000 0000 0100 0100",
		"000 0000 000000 00000000 0000000000 0000000000 0000 0000 0000 0010 0010",
		"000 0000 000000 00000000 0000000000 0000000000 0000 0000 0000 0001 0001",
		// float32
		"000 0000 000000 00000000 1000000000 1000000000 0000 0000 0000 0000 0000 0000",
		"000 0000 000000 10000000 0100000000 0100000000 0000 0000 0000 0000 0000 0000",
		"000 0000 100000 01000000 0010000000 0010000000 0000 0000 0000 0000 0000 0000",
		"000 1000 010000 00100000 0001000000 0001000000 0000 0000 0000 0000 0000 0000",
		"000 0100 001000 00010000 0000100000 0000100000 1000 1000 1000 1000 1000 1000",
		"000 0010 000100 00001000 0000010000 0000010000 0100 0000 0000 0000 0000 0000",
		"000 0000 000010 00000100 0000001000 0000001000 0000 0100 0000 0000 0000 0000",
		"000 0000 000000 00000000 0000000000 0000000000 0010 0000 0000 0000 0000 0000",
		"000 0000 000000 00000000 0000000000 0000000000 0000 0010 0000 0000 0000 0000",
		"000 0000 000000 00000000 0000000000 0000000000 0001 0000 0000 0000 0000 0000",
		"000 0000 000000 00000000 0000000000 0000000000 0000 0001 0000 0000 0000 0000",
		// float64
		"000 0000 000000 00000000 1000000000 1000000000 0000 0000 0000 0000 0000 0000 10000000000",
		"000 0000 000000 10000000 0100000000 0100000000 0000 0000 0000 0000 0000 0000 01000000000",
		"000 0000 100000 01000000 0010000000 0010000000 0000 0000 0000 0000 0000 0000 00100000000",
		"000 1000 010000 00100000 0001000000 0001000000 0000 0000 0000 0000 0000 0000 00010000000",
		"000 0100 001000 00010000 0000100000 0000100000 1000 1000 1000 1000 1000 1000 00001000000",
		"000 0010 000100 00001000 0000010000 0000010000 0100 0000 0000 0000 0000 0000 00000100000",
		"000 0000 000010 00000100 0000001000 0000001000 0000 0100 0000 0000 0000 0000 00000010000",
		"000 0000 000000 00000010 0000000100 0000000100 0000 0000 0100 0000 0000 0000 00000000000",
		"000 0000 000000 00000000 0000000000 0000000000 0010 0000 0000 0000 0000 0000 00000001000",
		"000 0000 000000 00000000 0000000000 0000000000 0000 0010 0000 0000 0000 0000 00000000100",
		"000 0000 000000 00000000 0000000000 0000000000 0000 0000 0010 0000 0000 0000 00000000000",
		"000 0000 000000 00000000 0000000000 0000000000 0001 0000 0000 0000 0000 0000 00000000010",
		"000 0000 000000 00000000 0000000000 0000000000 0000 0001 0000 0000 0000 0000 00000000001",
		"000 0000 000000 00000000 0000000000 0000000000 0000 0000 0001 0000 0000 0000 00000000000",
		// complex64
		"000 0100 001000 00010000 0000100000 0000100000 1000 1000 1000 1000 1000 1000 00001000000 00001000000000",
		"000 0010 000100 00001000 0000010000 0000010000 0100 0000 0000 0000 0000 0000 00000100000 00000100000000",
		"000 0000 000000 00000000 0000000000 0000000000 0000 0000 0000 0000 0000 0000 00000000000 00000000000000",
		"000 0000 000000 00000000 0000000000 0000000000 0000 0000 0000 0000 0000 0000 00000000000 00000000000000",
		// complex128
		"000 0100 001000 00010000 0000100000 0000100000 1000 1000 1000 1000 1000 1000 00001000000 00001000000000 1000",
		"000 0010 000100 00001000 0000010000 0000010000 0100 0000 0000 0000 0000 0000 00000100000 00000100000000 0100",
		"000 0000 000000 00000000 0000000000 0000000000 0000 0000 0000 0000 0000 0000 00000000000 00000000000000 0010",
		"000 0000 000000 00000000 0000000000 0000000000 0000 0000 0000 0000 0000 0000 00000000000 00000000000000 0001",
		// pointer
		"100 0000 000000 00000000 0000000000 0000000000 0000 0000 0000 0000 0000 0000 00000000000 00000000000000 0000 0000",
		"000 0000 000000 00000000 0000000000 0000000000 0000 0000 0000 0000 0000 0000 00000000000 00000000000000 0000 0000",
		"100 0000 000000 00000000 0000000000 0000000000 0000 0000 0000 0000 0000 0000 00000000000 00000000000000 0000 0000",
		// unsafe pointer
		"100 0000 000000 00000000 0000000000 0000000000 0000 0000 0000 0000 0000 0000 00000000000 00000000000000 0000 0000 101",
		"000 0000 000000 00000000 0000000000 0000000000 0000 0000 0000 0000 0000 0000 00000000000 00000000000000 0000 0000 010",
		// chan
		"100 0000 000000 00000000 0000000000 0000000000 0000 0000 0000 0000 0000 0000 00000000000 00000000000000 0000 0000 000 00",
		"000 0000 000000 00000000 0000000000 0000000000 0000 0000 0000 0000 0000 0000 00000000000 00000000000000 0000 0000 000 00",
		"100 0000 000000 00000000 0000000000 0000000000 0000 0000 0000 0000 0000 0000 00000000000 00000000000000 0000 0000 000 00",
		// func
		"100 0000 000000 00000000 0000000000 0000000000 0000 0000 0000 0000 0000 0000 00000000000 00000000000000 0000 0000 000 00 000",
		"000 0000 000000 00000000 0000000000 0000000000 0000 0000 0000 0000 0000 0000 00000000000 00000000000000 0000 0000 000 00 000 00", // func is NOT self equal
		"100 0000 000000 00000000 0000000000 0000000000 0000 0000 0000 0000 0000 0000 00000000000 00000000000000 0000 0000 000 00 000",
		// string
		Z1, Z1, Z1,
		// array
		Z1 + "000",
		Z1 + "000",
		Z1 + "000 01",
		Z1 + "000",
		Z1 + "000 1",
		Z1 + "000 00001",
		// slice
		N1 + "000 000000",
		Z1 + "000 100010",
		Z1 + "000 011000",
		Z1 + "000 011000 001",
		Z1 + "000 000100",
		Z1 + "000 100011 01",
		Z1 + "000 000011 000001",
		// map
		N1 + "000 000000 0000000",
		Z1 + "000 000000 0000000",
		Z1 + "000 000000 0000000",
		Z1 + "000 000000 0000000",
		Z1 + "000 000000 0000000 001",
		Z1 + "000 000000 0000000 00101",
		Z1 + "000 000000 0000000 01",
		Z1 + "000 000000 0000000 0100001",
		// struct
		Z1 + "000 000000 0000000 00000000",
		Z1 + "000 000000 0000000 00000000",
		Z1 + "000 000000 0000000 00000000 01",
		Z1 + "000 000000 0000000 00000000",
		Z1 + "000 000000 0000000 00000000 0001",
		// custom 1
		Z1 + "000 000000 0000000 00000000 00000",
		Z1 + "000 000000 0000000 00000000 00000 1",
		Z1 + "010 000000 0000000 00000000 00000",
		Z1 + "001 000000 0000000 00000000 00000",
		Z1 + "010 000000 0000000 00000000 00000 001",
		Z1 + "001 000000 0000000 00000000 00000",
		// custom 1
		Z1 + "010 000000 0000000 00000000 00000 001010",
		Z1 + "010 000000 0000000 00000000 00000 001010 1",
		Z1 + "000 000000 0000000 00000000 00000 001010 11",
		// custom 1
		Z1 + "001 000000 0000000 00000000 00000 000100 000",
		Z1 + "001 000000 0000000 00000000 00000 000001 000",
		Z1 + "000 000000 0000000 00000000 00000 000001 000 01",
		Z1 + "000 000000 0000000 00000000 00000 000001 000 011",
		// typedef
		"010",
	}
	for i := 0; i < len(cs); i++ {
		e, ri, id := cs[i], rs[i], -1
		rr := func(b bool) bool {
			id++
			for id < len(ri) && ri[id] == ' ' {
				id++
			}
			if id < len(ri) {
				return ri[id] == '1'
			}
			return b
		}
		for j := 0; j <= i; j++ {
			r := rr(i == j)
			a := cs[j]
			Equal(t, r, isSameInValue(e, a), "i=%v, j=%v\n%T\t%[3]v\n%T\t%[4]v", i, j, e, a)
			Equal(t, r, isSameInValue(a, e), "i=%v, j=%v\n%T\t%[3]v\n%T\t%[4]v", i, j, e, a)
		}
	}
}
