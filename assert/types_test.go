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

func TestConvertible(t *testing.T) {
	cs := []interface{}{
		nil, true,
		int8(1), int16(1), int32(1), int64(1), int(1),
		uint8(1), uint16(1), uint32(1), uint64(1), uint(1), uintptr(1),
		float32(1), float64(1),
		complex64(1), complex128(1),
		new(int), new(uint), unsafe.Pointer(nil),
		make(chan int), make(chan uint), func(int) bool { return true },
		string("abc"),
		[...]int{1, 2, 3}, [...]float32{1, 2, 3}, [...]byte{1, 2, 3}, [...]rune{1, 2, 3}, [...]interface{}{1, 2, 3}, [...]I{nil},
		[]int{1, 2, 3}, []float32{1, 2, 3}, []byte{1, 2, 3}, []rune{1, 2, 3}, []interface{}{1, 2, 3}, []I{nil},
		map[int]uint{}, map[uint]bool{}, map[bool]float32{}, map[bool]bool{}, map[uint]float32{}, map[interface{}]bool{}, map[bool]interface{}{}, map[interface{}]interface{}{},
		A{}, B{},
	}
	rs := []string{
		"", "",
		"00 11111", "00 11111", "00 11111", "00 11111", "00 11111",
		"00 11111 111111", "00 11111 111111", "00 11111 111111", "00 11111 111111", "00 11111 111111", "00 11111 111111",
		"00 11111 111111 11", "00 11111 111111 11",
		"00 11111 111111 11 11", "00 11111 111111 11 11",
		"00 00000 000000 00 00 101", "00 00000 000000 00 00 011", "00 00000 000000 00 00 111",
		"00 00000 000000 00 00 000", "00 00000 000000 00 00 000", "00 00000 000000 00 00 000",
		"00 11111 111111 00 00 000 000",
		// array
		"00 00000 000000 00 00 000 000 0",
		"00 00000 000000 00 00 000 000 0 1",
		"00 00000 000000 00 00 000 000 1 11",
		"00 00000 000000 00 00 000 000 1 111",
		"00 00000 000000 00 00 000 000 0 1111",
		"00 00000 000000 00 00 000 000 0 00000",
		// slice
		"00 00000 000000 00 00 000 000 0 111110",
		"00 00000 000000 00 00 000 000 0 111110 1",
		"00 00000 000000 00 00 000 000 1 111110 11",
		"00 00000 000000 00 00 000 000 1 111110 111",
		"00 00000 000000 00 00 000 000 0 111111 1111",
		"00 00000 000000 00 00 000 000 0 000011 00001",
		// map
		"00 00000 000000 00 00 000 000 0 000000 000000",
		"00 00000 000000 00 00 000 000 0 000000 000000",
		"00 00000 000000 00 00 000 000 0 000000 000000",
		"00 00000 000000 00 00 000 000 0 000000 000000",
		"00 00000 000000 00 00 000 000 0 000000 000000 1",
		"00 00000 000000 00 00 000 000 0 000000 000000 0101",
		"00 00000 000000 00 00 000 000 0 000000 000000 001101",
		"00 00000 000000 00 00 000 000 0 000000 000000 1111111",
		// struct
		"00 00000 000000 00 00 000 000 0 000000 000000 00000000",
		"00 00000 000000 00 00 000 000 0 000000 000000 00000000",
	}
	for i := 0; i < len(cs); i++ {
		e, ri, id := cs[i], rs[i], -1
		t1 := reflect.TypeOf(e)
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
			t2 := reflect.TypeOf(a)
			c := convertible(t1, t2)
			Equal(t, r, c, "i=%v, j=%v\n%T\t%[3]v\n%T\t%[4]v", i, j, e, a)
			if t1 != nil && t2 != nil {
				if t1.ConvertibleTo(t2) {
					True(t, c, "i=%v, j=%v\n%T\t%[3]v\n%T\t%[4]v", i, j, e, a)
				}
				if t2.ConvertibleTo(t1) {
					True(t, c, "i=%v, j=%v\n%T\t%[3]v\n%T\t%[4]v", i, j, e, a)
				}
			}
		}
	}
}

func TestMapKeyDiff(t *testing.T) {
	v1 := reflect.ValueOf(map[uint8]interface{}{100: (0 + 1i)})
	v2 := reflect.ValueOf(map[interface{}]float32{uint8(100): 1.25})
	True(t, convertibleKeyTo(v1.Type().Key(), v2.Type().Key()))
	False(t, convertibleKeyTo(v2.Type().Key(), v1.Type().Key()))
	s1, s2 := v1.MapKeys(), v2.MapKeys()
	True(t, reflect.DeepEqual(s1[0].Interface(), s2[0].Interface()))
	True(t, valueEqual(s1[0], s2[0]))
	ks, ks1, ks2 := mapKeyDiff(v1, v2)
	Equal(t, 1, len(ks), "ks=%v\nks1=%v\nks2=%v", ks, ks1, ks2)
}
