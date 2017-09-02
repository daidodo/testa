/*
* Copyright (c) 2017 Zhao DAI <daidodo@gmail.com>
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
* http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package assert

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestWriteValDiff(t *testing.T) {
	a := new(int)
	sa := fmt.Sprintf("%p", a)
	b := make(chan int)
	sb := fmt.Sprintf("%p", b)
	c := func(int) bool { return true }
	sc := fmt.Sprintf("%p", c)
	cs := []interface{}{
		nil, int8(-1), int16(-10), int32(-100), int64(-1000), int(-10000),
		uint8(1), uint16(10), uint32(100), uint64(1000), uint(10000), uintptr(100000),
		float32(1.25), float64(10.25),
		complex64(2), complex64(-1 + 1i), complex64(1i),
		complex128(3), complex128(-2 + 2i), complex128(2i),
		(*int)(a), (*uint)(unsafe.Pointer(a)), unsafe.Pointer(P(200)),
		chan int(b), chan uint(nil), (func(int) bool)(c),
		string("中a文b"), string("a中文bc"),
		[0]int{}, [...]interface{}{uint(1), float32(2), complex128(3)},
		[]int(nil), []int{}, []interface{}{int8(1), float64(2), complex128(3i), P(100)},
		map[bool]int(nil), map[bool]uint{}, map[interface{}]float32{uint8(100): 1.25}, map[uint8]interface{}{100: 1i},
		map[interface{}]interface{}{uint8(100): 1i, "abc": false},
		A{}, A{a: 100}, A{a: "中a文b", b: A{}}, A{a: "a中文bc", b: A{b: A{}}},
	}
	rs := [][]struct {
		s1, s2 string
	}{
		{{H("<nil>"), H("int8") + "(-1)"}, {H("-1"), H("-1")}},
		{{H("<nil>"), H("int16") + "(-10)"}, {H("-1"), H("-10")}, {H("-10"), H("-10")}},
		{{H("<nil>"), H("int32") + "(-100)"}, {H("-1"), H("-100")}, {H("-10"), H("-100")}, {H("-100"), H("-100")}},
		{{H("<nil>"), H("int64") + "(-1000)"}, {H("-1"), H("-1000")}, {H("-10"), H("-1000")}, {H("-100"), H("-1000")}, {H("-1000"), H("-1000")}},
		{{H("<nil>"), H("int") + "(-10000)"}, {H("-1"), H("-10000")}, {H("-10"), H("-10000")}, {H("-100"), H("-10000")}, {H("-1000"), H("-10000")},
			{H("-10000"), H("-10000")}},
		{{H("<nil>"), H("uint8") + "(1)"}, {H("-1"), H("1")}, {H("-10"), H("1")}, {H("-100"), H("1")}, {H("-1000"), H("1")}, {H("-10000"), H("1")}, {H("1"), H("1")}},
		{{H("<nil>"), H("uint16") + "(10)"}, {H("-1"), H("10")}, {H("-10"), H("10")}, {H("-100"), H("10")}, {H("-1000"), H("10")}, {H("-10000"), H("10")},
			{H("1"), H("10")}, {H("10"), H("10")}},
		{{H("<nil>"), H("uint32") + "(100)"}, {H("-1"), H("100")}, {H("-10"), H("100")}, {H("-100"), H("100")}, {H("-1000"), H("100")}, {H("-10000"), H("100")},
			{H("1"), H("100")}, {H("10"), H("100")}, {H("100"), H("100")}},
		{{H("<nil>"), H("uint64") + "(1000)"}, {H("-1"), H("1000")}, {H("-10"), H("1000")}, {H("-100"), H("1000")}, {H("-1000"), H("1000")}, {H("-10000"), H("1000")},
			{H("1"), H("1000")}, {H("10"), H("1000")}, {H("100"), H("1000")}, {H("1000"), H("1000")}},
		{{H("<nil>"), H("uint") + "(10000)"}, {H("-1"), H("10000")}, {H("-10"), H("10000")}, {H("-100"), H("10000")}, {H("-1000"), H("10000")},
			{H("-10000"), H("10000")}, {H("1"), H("10000")}, {H("10"), H("10000")}, {H("100"), H("10000")}, {H("1000"), H("10000")}, {H("10000"), H("10000")}},
		{{H("<nil>"), H("uintptr") + "(0x186a0)"}, {H("-1"), H("0x186a0") + "(100000)"}, {H("-10"), H("0x186a0") + "(100000)"}, {H("-100"), H("0x186a0") + "(100000)"},
			{H("-1000"), H("0x186a0") + "(100000)"}, {H("-10000"), H("0x186a0") + "(100000)"}, {H("1"), H("0x186a0") + "(100000)"},
			{H("10"), H("0x186a0") + "(100000)"}, {H("100"), H("0x186a0") + "(100000)"}, {H("1000"), H("0x186a0") + "(100000)"},
			{H("10000"), H("0x186a0") + "(100000)"}, {H("0x186a0"), H("0x186a0")}},
		{{H("<nil>"), H("float32") + "(1.25)"}, {H("-1"), H("1.25")}, {H("-10"), H("1.25")}, {H("-100"), H("1.25")}, {H("-1000"), H("1.25")}, {H("-10000"), H("1.25")},
			{H("1"), H("1.25")}, {H("10"), H("1.25")}, {H("100"), H("1.25")}, {H("1000"), H("1.25")}, {H("10000"), H("1.25")}, {H("0x186a0") + "(100000)", H("1.25")},
			{H("1.25"), H("1.25")}},
		{{H("<nil>"), H("float64") + "(10.25)"}, {H("-1"), H("10.25")}, {H("-10"), H("10.25")}, {H("-100"), H("10.25")}, {H("-1000"), H("10.25")},
			{H("-10000"), H("10.25")}, {H("1"), H("10.25")}, {H("10"), H("10.25")}, {H("100"), H("10.25")}, {H("1000"), H("10.25")}, {H("10000"), H("10.25")},
			{H("0x186a0") + "(100000)", H("10.25")}, {H("1.25"), H("10.25")}, {H("10.25"), H("10.25")}},
		{{H("<nil>"), H("complex64") + "(2+0i)"}, {H("-1"), "(" + H("2") + "+0i)"}, {H("-10"), "(" + H("2") + "+0i)"}, {H("-100"), "(" + H("2") + "+0i)"},
			{H("-1000"), "(" + H("2") + "+0i)"}, {H("-10000"), "(" + H("2") + "+0i)"}, {H("1"), "(" + H("2") + "+0i)"}, {H("10"), "(" + H("2") + "+0i)"},
			{H("100"), "(" + H("2") + "+0i)"}, {H("1000"), "(" + H("2") + "+0i)"}, {H("10000"), "(" + H("2") + "+0i)"},
			{H("0x186a0") + "(100000)", "(" + H("2") + "+0i)"}, {H("1.25"), "(" + H("2") + "+0i)"}, {H("10.25"), "(" + H("2") + "+0i)"}, {"(2+0i)", "(2+0i)"}},
		{{H("<nil>"), H("complex64") + "(-1+1i)"}, {"-1", "(-1+" + H("1i") + ")"}, {H("-10"), "(" + H("-1+1i") + ")"}, {H("-100"), "(" + H("-1+1i") + ")"},
			{H("-1000"), "(" + H("-1+1i") + ")"}, {H("-10000"), "(" + H("-1+1i") + ")"}, {H("1"), "(" + H("-1+1i") + ")"}, {H("10"), "(" + H("-1+1i") + ")"},
			{H("100"), "(" + H("-1+1i") + ")"}, {H("1000"), "(" + H("-1+1i") + ")"}, {H("10000"), "(" + H("-1+1i") + ")"}, {H("0x186a0"), "(" + H("-1+1i") + ")"},
			{H("1.25"), "(" + H("-1+1i") + ")"}, {H("10.25"), "(" + H("-1+1i") + ")"}, {"(" + H("2+0i") + ")", "(" + H("-1+1i") + ")"}, {"(-1+1i)", "(-1+1i)"}},
		{{H("<nil>"), H("complex64") + "(0+1i)"}, {H("-1"), "(" + H("0+1i") + ")"}, {H("-10"), "(" + H("0+1i") + ")"}, {H("-100"), "(" + H("0+1i") + ")"},
			{H("-1000"), "(" + H("0+1i") + ")"}, {H("-10000"), "(" + H("0+1i") + ")"}, {H("1"), "(" + H("0+1i") + ")"}, {H("10"), "(" + H("0+1i") + ")"},
			{H("100"), "(" + H("0+1i") + ")"}, {H("1000"), "(" + H("0+1i") + ")"}, {H("10000"), "(" + H("0+1i") + ")"}, {H("0x186a0"), "(" + H("0+1i") + ")"},
			{H("1.25"), "(" + H("0+1i") + ")"}, {H("10.25"), "(" + H("0+1i") + ")"}, {"(" + H("2+0i") + ")", "(" + H("0+1i") + ")"},
			{"(" + H("-1") + "+1i)", "(" + H("0") + "+1i)"}, {"(0+1i)", "(0+1i)"}},
		{{H("<nil>"), H("complex128") + "(3+0i)"}, {H("-1"), "(" + H("3") + "+0i)"}, {H("-10"), "(" + H("3") + "+0i)"}, {H("-100"), "(" + H("3") + "+0i)"},
			{H("-1000"), "(" + H("3") + "+0i)"}, {H("-10000"), "(" + H("3") + "+0i)"}, {H("1"), "(" + H("3") + "+0i)"}, {H("10"), "(" + H("3") + "+0i)"},
			{H("100"), "(" + H("3") + "+0i)"}, {H("1000"), "(" + H("3") + "+0i)"}, {H("10000"), "(" + H("3") + "+0i)"},
			{H("0x186a0") + "(100000)", "(" + H("3") + "+0i)"}, {H("1.25"), "(" + H("3") + "+0i)"}, {H("10.25"), "(" + H("3") + "+0i)"},
			{"(" + H("2") + "+0i)", "(" + H("3") + "+0i)"}, {"(" + H("-1+1i") + ")", "(" + H("3+0i") + ")"}, {"(" + H("0+1i") + ")", "(" + H("3+0i") + ")"},
			{"(3+0i)", "(3+0i)"}},
		{{H("<nil>"), H("complex128") + "(-2+2i)"}, {H("-1"), "(" + H("-2+2i") + ")"}, {H("-10"), "(" + H("-2+2i") + ")"}, {H("-100"), "(" + H("-2+2i") + ")"},
			{H("-1000"), "(" + H("-2+2i") + ")"}, {H("-10000"), "(" + H("-2+2i") + ")"}, {H("1"), "(" + H("-2+2i") + ")"}, {H("10"), "(" + H("-2+2i") + ")"},
			{H("100"), "(" + H("-2+2i") + ")"}, {H("1000"), "(" + H("-2+2i") + ")"}, {H("10000"), "(" + H("-2+2i") + ")"}, {H("0x186a0"), "(" + H("-2+2i") + ")"},
			{H("1.25"), "(" + H("-2+2i") + ")"}, {H("10.25"), "(" + H("-2+2i") + ")"}, {"(" + H("2+0i") + ")", "(" + H("-2+2i") + ")"},
			{"(" + H("-1+1i") + ")", "(" + H("-2+2i") + ")"}, {"(" + H("0+1i") + ")", "(" + H("-2+2i") + ")"}, {"(" + H("3+0i") + ")", "(" + H("-2+2i") + ")"},
			{"(-2+2i)", "(-2+2i)"}},
		{{H("<nil>"), H("complex128") + "(0+2i)"}, {H("-1"), "(" + H("0+2i") + ")"}, {H("-10"), "(" + H("0+2i") + ")"}, {H("-100"), "(" + H("0+2i") + ")"},
			{H("-1000"), "(" + H("0+2i") + ")"}, {H("-10000"), "(" + H("0+2i") + ")"}, {H("1"), "(" + H("0+2i") + ")"}, {H("10"), "(" + H("0+2i") + ")"},
			{H("100"), "(" + H("0+2i") + ")"}, {H("1000"), "(" + H("0+2i") + ")"}, {H("10000"), "(" + H("0+2i") + ")"}, {H("0x186a0"), "(" + H("0+2i") + ")"},
			{H("1.25"), "(" + H("0+2i") + ")"}, {H("10.25"), "(" + H("0+2i") + ")"}, {"(" + H("2+0i") + ")", "(" + H("0+2i") + ")"},
			{"(" + H("-1+1i") + ")", "(" + H("0+2i") + ")"}, {"(0+" + H("1i") + ")", "(0+" + H("2i") + ")"}, {"(" + H("3+0i") + ")", "(" + H("0+2i") + ")"},
			{"(" + H("-2") + "+2i)", "(" + H("0") + "+2i)"}, {"(0+2i)", "(0+2i)"}},
		{{H("<nil>"), H(sa)}, {H("int8") + "(-1)", "(" + H("*int") + ")(" + sa + ")"}, {H("int16") + "(-10)", "(" + H("*int") + ")(" + sa + ")"},
			{H("int32") + "(-100)", "(" + H("*int") + ")(" + sa + ")"}, {H("int64") + "(-1000)", "(" + H("*int") + ")(" + sa + ")"},
			{H("int") + "(-10000)", "(" + H("*int") + ")(" + sa + ")"}, {H("uint8") + "(1)", "(" + H("*int") + ")(" + sa + ")"},
			{H("uint16") + "(10)", "(" + H("*int") + ")(" + sa + ")"}, {H("uint32") + "(100)", "(" + H("*int") + ")(" + sa + ")"},
			{H("uint64") + "(1000)", "(" + H("*int") + ")(" + sa + ")"}, {H("uint") + "(10000)", "(" + H("*int") + ")(" + sa + ")"},
			{H("uintptr") + "(0x186a0)", "(" + H("*int") + ")(" + sa + ")"}, {H("float32") + "(1.25)", "(" + H("*int") + ")(" + sa + ")"},
			{H("float64") + "(10.25)", "(" + H("*int") + ")(" + sa + ")"}, {H("complex64") + "(2+0i)", "(" + H("*int") + ")(" + sa + ")"},
			{H("complex64") + "(-1+1i)", "(" + H("*int") + ")(" + sa + ")"}, {H("complex64") + "(0+1i)", "(" + H("*int") + ")(" + sa + ")"},
			{H("complex128") + "(3+0i)", "(" + H("*int") + ")(" + sa + ")"}, {H("complex128") + "(-2+2i)", "(" + H("*int") + ")(" + sa + ")"},
			{H("complex128") + "(0+2i)", "(" + H("*int") + ")(" + sa + ")"}, {"&" + H("0"), "&" + H("0")}},
		{{H("<nil>"), H(sa)}, {H("int8") + "(-1)", "(" + H("*uint") + ")(" + sa + ")"}, {H("int16") + "(-10)", "(" + H("*uint") + ")(" + sa + ")"},
			{H("int32") + "(-100)", "(" + H("*uint") + ")(" + sa + ")"}, {H("int64") + "(-1000)", "(" + H("*uint") + ")(" + sa + ")"},
			{H("int") + "(-10000)", "(" + H("*uint") + ")(" + sa + ")"}, {H("uint8") + "(1)", "(" + H("*uint") + ")(" + sa + ")"},
			{H("uint16") + "(10)", "(" + H("*uint") + ")(" + sa + ")"}, {H("uint32") + "(100)", "(" + H("*uint") + ")(" + sa + ")"},
			{H("uint64") + "(1000)", "(" + H("*uint") + ")(" + sa + ")"}, {H("uint") + "(10000)", "(" + H("*uint") + ")(" + sa + ")"},
			{H("uintptr") + "(0x186a0)", "(" + H("*uint") + ")(" + sa + ")"}, {H("float32") + "(1.25)", "(" + H("*uint") + ")(" + sa + ")"},
			{H("float64") + "(10.25)", "(" + H("*uint") + ")(" + sa + ")"}, {H("complex64") + "(2+0i)", "(" + H("*uint") + ")(" + sa + ")"},
			{H("complex64") + "(-1+1i)", "(" + H("*uint") + ")(" + sa + ")"}, {H("complex64") + "(0+1i)", "(" + H("*uint") + ")(" + sa + ")"},
			{H("complex128") + "(3+0i)", "(" + H("*uint") + ")(" + sa + ")"}, {H("complex128") + "(-2+2i)", "(" + H("*uint") + ")(" + sa + ")"},
			{H("complex128") + "(0+2i)", "(" + H("*uint") + ")(" + sa + ")"}, {"(*" + H("int") + ")(" + sa + ")", "(*" + H("uint") + ")(" + sa + ")"}, {"&" + H("0"), "&" + H("0")}},
		{{H("<nil>"), H("0xc8")}, {H("int8") + "(-1)", "(" + H("unsafe.Pointer") + ")(0xc8)"},
			{H("int16") + "(-10)", "(" + H("unsafe.Pointer") + ")(0xc8)"}, {H("int32") + "(-100)", "(" + H("unsafe.Pointer") + ")(0xc8)"},
			{H("int64") + "(-1000)", "(" + H("unsafe.Pointer") + ")(0xc8)"}, {H("int") + "(-10000)", "(" + H("unsafe.Pointer") + ")(0xc8)"},
			{H("uint8") + "(1)", "(" + H("unsafe.Pointer") + ")(0xc8)"}, {H("uint16") + "(10)", "(" + H("unsafe.Pointer") + ")(0xc8)"},
			{H("uint32") + "(100)", "(" + H("unsafe.Pointer") + ")(0xc8)"}, {H("uint64") + "(1000)", "(" + H("unsafe.Pointer") + ")(0xc8)"},
			{H("uint") + "(10000)", "(" + H("unsafe.Pointer") + ")(0xc8)"}, {H("uintptr") + "(0x186a0)", "(" + H("unsafe.Pointer") + ")(0xc8)"},
			{H("float32") + "(1.25)", "(" + H("unsafe.Pointer") + ")(0xc8)"}, {H("float64") + "(10.25)", "(" + H("unsafe.Pointer") + ")(0xc8)"},
			{H("complex64") + "(2+0i)", "(" + H("unsafe.Pointer") + ")(0xc8)"}, {H("complex64") + "(-1+1i)", "(" + H("unsafe.Pointer") + ")(0xc8)"},
			{H("complex64") + "(0+1i)", "(" + H("unsafe.Pointer") + ")(0xc8)"}, {H("complex128") + "(3+0i)", "(" + H("unsafe.Pointer") + ")(0xc8)"},
			{H("complex128") + "(-2+2i)", "(" + H("unsafe.Pointer") + ")(0xc8)"}, {H("complex128") + "(0+2i)", "(" + H("unsafe.Pointer") + ")(0xc8)"},
			{H(sa), H("0xc8")}, {H(sa), H("0xc8")}, {H("0xc8"), H("0xc8")}},
		{{H("<nil>"), H(sb)}, {H("int8") + "(-1)", "(" + H("chan int") + ")(" + sb + ")"},
			{H("int16") + "(-10)", "(" + H("chan int") + ")(" + sb + ")"}, {H("int32") + "(-100)", "(" + H("chan int") + ")(" + sb + ")"},
			{H("int64") + "(-1000)", "(" + H("chan int") + ")(" + sb + ")"}, {H("int") + "(-10000)", "(" + H("chan int") + ")(" + sb + ")"},
			{H("uint8") + "(1)", "(" + H("chan int") + ")(" + sb + ")"}, {H("uint16") + "(10)", "(" + H("chan int") + ")(" + sb + ")"},
			{H("uint32") + "(100)", "(" + H("chan int") + ")(" + sb + ")"}, {H("uint64") + "(1000)", "(" + H("chan int") + ")(" + sb + ")"},
			{H("uint") + "(10000)", "(" + H("chan int") + ")(" + sb + ")"}, {H("uintptr") + "(0x186a0)", "(" + H("chan int") + ")(" + sb + ")"},
			{H("float32") + "(1.25)", "(" + H("chan int") + ")(" + sb + ")"}, {H("float64") + "(10.25)", "(" + H("chan int") + ")(" + sb + ")"},
			{H("complex64") + "(2+0i)", "(" + H("chan int") + ")(" + sb + ")"}, {H("complex64") + "(-1+1i)", "(" + H("chan int") + ")(" + sb + ")"},
			{H("complex64") + "(0+1i)", "(" + H("chan int") + ")(" + sb + ")"}, {H("complex128") + "(3+0i)", "(" + H("chan int") + ")(" + sb + ")"},
			{H("complex128") + "(-2+2i)", "(" + H("chan int") + ")(" + sb + ")"}, {H("complex128") + "(0+2i)", "(" + H("chan int") + ")(" + sb + ")"},
			{"(" + H("*int") + ")(" + sa + ")", "(" + H("chan int") + ")(" + sb + ")"}, {"(" + H("*uint") + ")(" + sa + ")", "(" + H("chan int") + ")(" + sb + ")"},
			{"(" + H("unsafe.Pointer") + ")(0xc8)", "(" + H("chan int") + ")(" + sb + ")"}, {"(chan int)(" + H(sb) + ")", "(chan int)(" + H(sb) + ")"}},
		{{H("<nil>"), H("<nil>")}, {H("int8") + "(-1)", "(" + H("chan uint") + ")(nil)"}, {H("int16") + "(-10)", "(" + H("chan uint") + ")(nil)"},
			{H("int32") + "(-100)", "(" + H("chan uint") + ")(nil)"}, {H("int64") + "(-1000)", "(" + H("chan uint") + ")(nil)"},
			{H("int") + "(-10000)", "(" + H("chan uint") + ")(nil)"}, {H("uint8") + "(1)", "(" + H("chan uint") + ")(nil)"},
			{H("uint16") + "(10)", "(" + H("chan uint") + ")(nil)"}, {H("uint32") + "(100)", "(" + H("chan uint") + ")(nil)"},
			{H("uint64") + "(1000)", "(" + H("chan uint") + ")(nil)"}, {H("uint") + "(10000)", "(" + H("chan uint") + ")(nil)"},
			{H("uintptr") + "(0x186a0)", "(" + H("chan uint") + ")(nil)"}, {H("float32") + "(1.25)", "(" + H("chan uint") + ")(nil)"},
			{H("float64") + "(10.25)", "(" + H("chan uint") + ")(nil)"}, {H("complex64") + "(2+0i)", "(" + H("chan uint") + ")(nil)"},
			{H("complex64") + "(-1+1i)", "(" + H("chan uint") + ")(nil)"}, {H("complex64") + "(0+1i)", "(" + H("chan uint") + ")(nil)"},
			{H("complex128") + "(3+0i)", "(" + H("chan uint") + ")(nil)"}, {H("complex128") + "(-2+2i)", "(" + H("chan uint") + ")(nil)"},
			{H("complex128") + "(0+2i)", "(" + H("chan uint") + ")(nil)"}, {"(" + H("*int") + ")(" + sa + ")", "(" + H("chan uint") + ")(nil)"},
			{"(" + H("*uint") + ")(" + sa + ")", "(" + H("chan uint") + ")(nil)"}, {"(" + H("unsafe.Pointer") + ")(0xc8)", "(" + H("chan uint") + ")(nil)"},
			{"(chan " + H("int") + ")(" + sb + ")", "(chan " + H("uint") + ")(nil)"}, {"(chan uint)(" + H("nil") + ")", "(chan uint)(" + H("nil") + ")"}},
		{{H("<nil>"), H(sc)}, {H("int8") + "(-1)", "(" + H("func(int) bool") + ")(" + sc + ")"},
			{H("int16") + "(-10)", "(" + H("func(int) bool") + ")(" + sc + ")"}, {H("int32") + "(-100)", "(" + H("func(int) bool") + ")(" + sc + ")"},
			{H("int64") + "(-1000)", "(" + H("func(int) bool") + ")(" + sc + ")"}, {H("int") + "(-10000)", "(" + H("func(int) bool") + ")(" + sc + ")"},
			{H("uint8") + "(1)", "(" + H("func(int) bool") + ")(" + sc + ")"}, {H("uint16") + "(10)", "(" + H("func(int) bool") + ")(" + sc + ")"},
			{H("uint32") + "(100)", "(" + H("func(int) bool") + ")(" + sc + ")"}, {H("uint64") + "(1000)", "(" + H("func(int) bool") + ")(" + sc + ")"},
			{H("uint") + "(10000)", "(" + H("func(int) bool") + ")(" + sc + ")"}, {H("uintptr") + "(0x186a0)", "(" + H("func(int) bool") + ")(" + sc + ")"},
			{H("float32") + "(1.25)", "(" + H("func(int) bool") + ")(" + sc + ")"}, {H("float64") + "(10.25)", "(" + H("func(int) bool") + ")(" + sc + ")"},
			{H("complex64") + "(2+0i)", "(" + H("func(int) bool") + ")(" + sc + ")"}, {H("complex64") + "(-1+1i)", "(" + H("func(int) bool") + ")(" + sc + ")"},
			{H("complex64") + "(0+1i)", "(" + H("func(int) bool") + ")(" + sc + ")"}, {H("complex128") + "(3+0i)", "(" + H("func(int) bool") + ")(" + sc + ")"},
			{H("complex128") + "(-2+2i)", "(" + H("func(int) bool") + ")(" + sc + ")"}, {H("complex128") + "(0+2i)", "(" + H("func(int) bool") + ")(" + sc + ")"},
			{"(" + H("*int") + ")(" + sa + ")", "(" + H("func(int) bool") + ")(" + sc + ")"},
			{"(" + H("*uint") + ")(" + sa + ")", "(" + H("func(int) bool") + ")(" + sc + ")"},
			{"(" + H("unsafe.Pointer") + ")(0xc8)", "(" + H("func(int) bool") + ")(" + sc + ")"},
			{"(" + H("chan int") + ")(" + sb + ")", "(" + H("func(int) bool") + ")(" + sc + ")"},
			{"(" + H("chan uint") + ")(nil)", "(" + H("func(int) bool") + ")(" + sc + ")"}, {"(func(int) bool)(" + H(sc) + ")", "(func(int) bool)(" + H(sc) + ")"}},
		{{H("<nil>"), H("string") + `("中a文b")`}, {H("-1"), H(`"中a文b"`)}, {H("-10"), H(`"中a文b"`)}, {H("-100"), H(`"中a文b"`)},
			{H("-1000"), H(`"中a文b"`)}, {H("-10000"), H(`"中a文b"`)}, {H("1"), H(`"中a文b"`)}, {H("10"), H(`"中a文b"`)},
			{H("100"), H(`"中a文b"`)}, {H("1000"), H(`"中a文b"`)}, {H("10000"), H(`"中a文b"`)}, {H("0x186a0"), H(`"中a文b"`)},
			{H("float32") + "(1.25)", H("string") + `("中a文b")`}, {H("float64") + "(10.25)", H("string") + `("中a文b")`},
			{H("complex64") + "(2+0i)", H("string") + `("中a文b")`}, {H("complex64") + "(-1+1i)", H("string") + `("中a文b")`},
			{H("complex64") + "(0+1i)", H("string") + `("中a文b")`}, {H("complex128") + "(3+0i)", H("string") + `("中a文b")`},
			{H("complex128") + "(-2+2i)", H("string") + `("中a文b")`}, {H("complex128") + "(0+2i)", H("string") + `("中a文b")`},
			{"(" + H("*int") + ")(" + sa + ")", H("string") + `("中a文b")`}, {"(" + H("*uint") + ")(" + sa + ")", H("string") + `("中a文b")`},
			{"(" + H("unsafe.Pointer") + ")(0xc8)", H("string") + `("中a文b")`}, {"(" + H("chan int") + ")(" + sb + ")", H("string") + `("中a文b")`},
			{"(" + H("chan uint") + ")(nil)", H("string") + `("中a文b")`}, {"(" + H("func(int) bool") + ")(" + sc + ")", H("string") + `("中a文b")`},
			{`"中a文b"`, `"中a文b"`}},
		{{s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {`"` + H("中a") + `文b"`, `"` + H("a中") + "文b" + H("c") + `"`}, {s1: "skip"}},
		{{H("<nil>"), "[0]int{}"}, {H("int8") + "(-1)", H("[0]int") + "{}"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{H("uint8") + "(1)", H("[0]int") + "{}"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {H("uintptr") + "(0x186a0)", H("[0]int") + "{}"},
			{H("float32") + "(1.25)", H("[0]int") + "{}"}, {s1: "skip"}, {H("complex64") + "(2+0i)", H("[0]int") + "{}"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{s1: "skip"}, {s1: "skip"}, {"(" + H("*int") + ")(" + sa + ")", H("[0]int") + "{}"}, {s1: "skip"},
			{"(" + H("unsafe.Pointer") + ")(0xc8)", H("[0]int") + "{}"}, {"(" + H("chan int") + ")(" + sb + ")", H("[0]int") + "{}"},
			{"(" + H("chan uint") + ")(nil)", H("[0]int") + "{}"}, {"(" + H("func(int) bool") + ")(" + sc + ")", H("[0]int") + "{}"},
			{H("string") + `("中a文b")`, H("[0]int") + "{}"}, {s1: "skip"}, {s1: "skip"}},
		{{H("<nil>"), "[3]interface {}{1, 2, (3+0i)}"}, {H("int8") + "(-1)", H("[3]interface {}") + "{1, 2, (3+0i)}"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{s1: "skip"}, {"[" + H("0") + "]int{}", "[" + H("3") + "]interface {}{1, 2, (3+0i)}"}, {s1: "skip"}},
		{{s1: "skip"}, {H("int8") + "(-1)", H("[]int") + "(nil)"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{H("uint8") + "(1)", H("[]int") + "(nil)"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {H("uintptr") + "(0x186a0)", H("[]int") + "(nil)"},
			{H("float32") + "(1.25)", H("[]int") + "(nil)"}, {s1: "skip"}, {H("complex64") + "(2+0i)", H("[]int") + "(nil)"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{s1: "skip"}, {s1: "skip"}, {"(" + H("*int") + ")(" + sa + ")", H("[]int") + "(nil)"}, {s1: "skip"},
			{"(" + H("unsafe.Pointer") + ")(0xc8)", H("[]int") + "(nil)"}, {"(" + H("chan int") + ")(" + sb + ")", H("[]int") + "(nil)"},
			{"(" + H("chan uint") + ")(nil)", H("[]int") + "(nil)"}, {"(" + H("func(int) bool") + ")(" + sc + ")", H("[]int") + "(nil)"},
			{H("string") + `("中a文b")`, H("[]int") + "(nil)"}, {s1: "skip"}, {"[0]int" + H("{}"), "[]int(" + H("nil") + ")"},
			{"[3]interface {}" + H("{1, 2, (3+0i)}"), "[]int(" + H("nil") + ")"}, {s1: "skip"}},
		{{H("<nil>"), "[]int{}"}, {H("int8") + "(-1)", H("[]int") + "{}"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{H("uint8") + "(1)", H("[]int") + "{}"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {H("uintptr") + "(0x186a0)", H("[]int") + "{}"},
			{H("float32") + "(1.25)", H("[]int") + "{}"}, {s1: "skip"}, {H("complex64") + "(2+0i)", H("[]int") + "{}"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{s1: "skip"}, {s1: "skip"}, {"(" + H("*int") + ")(" + sa + ")", H("[]int") + "{}"}, {s1: "skip"},
			{"(" + H("unsafe.Pointer") + ")(0xc8)", H("[]int") + "{}"}, {"(" + H("chan int") + ")(" + sb + ")", H("[]int") + "{}"},
			{"(" + H("chan uint") + ")(nil)", H("[]int") + "{}"}, {"(" + H("func(int) bool") + ")(" + sc + ")", H("[]int") + "{}"},
			{H("string") + `("中a文b")`, H("[]int") + "{}"}, {s1: "skip"}, {s1: "skip"}, {"[3]interface {}{" + H("1, 2, (3+0i)") + "}", "[]int{}"},
			{"[]int(" + H("nil") + ")", "[]int" + H("{}")}, {s1: "skip"}},
		{{H("<nil>"), "[]interface {}{1, 2, (0+3i), 0x64}"}, {H("int8") + "(-1)", H("[]interface {}") + "{1, 2, (0+3i), 0x64}"}, {s1: "skip"}, {s1: "skip"},
			{s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{s1: "skip"}, {s1: "skip"}, {"[0]int{}", "[]interface {}{" + H("1, 2, (0+3i), 0x64") + "}"},
			{"[3]interface {}{1, 2, (" + H("3+0i") + ")}", "[]interface {}{1, 2, (" + H("0+3i") + "), " + H("0x64") + "}"},
			{"[]int(" + H("nil") + ")", "[]interface {}" + H("{1, 2, (0+3i), 0x64}")}, {"[]int{}", "[]interface {}{" + H("1, 2, (0+3i), 0x64") + "}"}, {s1: "skip"}},
		{{s1: "skip"}, {H("int8") + "(-1)", H("map[bool]int") + "(nil)"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{H("uint8") + "(1)", H("map[bool]int") + "(nil)"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{H("uintptr") + "(0x186a0)", H("map[bool]int") + "(nil)"}, {H("float32") + "(1.25)", H("map[bool]int") + "(nil)"}, {s1: "skip"},
			{H("complex64") + "(2+0i)", H("map[bool]int") + "(nil)"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{"(" + H("*int") + ")(" + sa + ")", H("map[bool]int") + "(nil)"}, {s1: "skip"}, {"(" + H("unsafe.Pointer") + ")(0xc8)", H("map[bool]int") + "(nil)"},
			{"(" + H("chan int") + ")(" + sb + ")", H("map[bool]int") + "(nil)"}, {"(" + H("chan uint") + ")(nil)", H("map[bool]int") + "(nil)"},
			{"(" + H("func(int) bool") + ")(" + sc + ")", H("map[bool]int") + "(nil)"}, {H("string") + `("中a文b")`, H("map[bool]int") + "(nil)"}, {s1: "skip"},
			{H("[0]int") + "{}", H("map[bool]int") + "(nil)"}, {H("[3]interface {}") + "{1, 2, (3+0i)}", H("map[bool]int") + "(nil)"},
			{H("[]int") + "(nil)", H("map[bool]int") + "(nil)"}, {H("[]int") + "{}", H("map[bool]int") + "(nil)"},
			{H("[]interface {}") + "{1, 2, (0+3i), 0x64}", H("map[bool]int") + "(nil)"}, {s1: "skip"}},
		{{H("<nil>"), "map[bool]uint{}"}, {H("int8") + "(-1)", H("map[bool]uint") + "{}"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{H("uint8") + "(1)", H("map[bool]uint") + "{}"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{H("uintptr") + "(0x186a0)", H("map[bool]uint") + "{}"}, {H("float32") + "(1.25)", H("map[bool]uint") + "{}"}, {s1: "skip"},
			{H("complex64") + "(2+0i)", H("map[bool]uint") + "{}"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{"(" + H("*int") + ")(" + sa + ")", H("map[bool]uint") + "{}"}, {s1: "skip"}, {"(" + H("unsafe.Pointer") + ")(0xc8)", H("map[bool]uint") + "{}"},
			{"(" + H("chan int") + ")(" + sb + ")", H("map[bool]uint") + "{}"}, {"(" + H("chan uint") + ")(nil)", H("map[bool]uint") + "{}"},
			{"(" + H("func(int) bool") + ")(" + sc + ")", H("map[bool]uint") + "{}"}, {H("string") + `("中a文b")`, H("map[bool]uint") + "{}"},
			{s1: "skip"}, {H("[0]int") + "{}", H("map[bool]uint") + "{}"}, {H("[3]interface {}") + "{1, 2, (3+0i)}", H("map[bool]uint") + "{}"},
			{H("[]int") + "(nil)", H("map[bool]uint") + "{}"}, {H("[]int") + "{}", H("map[bool]uint") + "{}"},
			{H("[]interface {}") + "{1, 2, (0+3i), 0x64}", H("map[bool]uint") + "{}"}, {"map[bool]int(" + H("nil") + ")", "map[bool]uint" + H("{}")}, {s1: "skip"}},
		{{H("<nil>"), "map[interface {}]float32{100:1.25}"}, {H("int8") + "(-1)", H("map[interface {}]float32") + "{100:1.25}"}, {s1: "skip"}, {s1: "skip"},
			{s1: "skip"}, {s1: "skip"}, {H("uint8") + "(1)", H("map[interface {}]float32") + "{100:1.25}"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{H("uintptr") + "(0x186a0)", H("map[interface {}]float32") + "{100:1.25}"}, {H("float32") + "(1.25)", H("map[interface {}]float32") + "{100:1.25}"},
			{s1: "skip"}, {H("complex64") + "(2+0i)", H("map[interface {}]float32") + "{100:1.25}"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{s1: "skip"}, {"(" + H("*int") + ")(" + sa + ")", H("map[interface {}]float32") + "{100:1.25}"}, {s1: "skip"},
			{"(" + H("unsafe.Pointer") + ")(0xc8)", H("map[interface {}]float32") + "{100:1.25}"},
			{"(" + H("chan int") + ")(" + sb + ")", H("map[interface {}]float32") + "{100:1.25}"},
			{"(" + H("chan uint") + ")(nil)", H("map[interface {}]float32") + "{100:1.25}"},
			{"(" + H("func(int) bool") + ")(" + sc + ")", H("map[interface {}]float32") + "{100:1.25}"},
			{H("string") + `("中a文b")`, H("map[interface {}]float32") + "{100:1.25}"}, {s1: "skip"},
			{H("[0]int") + "{}", H("map[interface {}]float32") + "{100:1.25}"},
			{H("[3]interface {}") + "{1, 2, (3+0i)}", H("map[interface {}]float32") + "{100:1.25}"},
			{H("[]int") + "(nil)", H("map[interface {}]float32") + "{100:1.25}"}, {H("[]int") + "{}", H("map[interface {}]float32") + "{100:1.25}"},
			{H("[]interface {}") + "{1, 2, (0+3i), 0x64}", H("map[interface {}]float32") + "{100:1.25}"},
			{"map[bool]int(" + H("nil") + ")", "map[interface {}]float32" + H("{100:1.25}")}, {"map[bool]uint{}", "map[interface {}]float32{" + H("100:1.25") + "}"},
			{s1: "skip"}},
		{{s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{"map[" + H("bool") + "]int(nil)", "map[" + H("uint8") + "]interface {}{100:(0+1i)}"},
			{"map[" + H("bool") + "]uint{}", "map[" + H("uint8") + "]interface {}{100:(0+1i)}"},
			{"map[interface {}]float32{100:" + H("1.25") + "}", "map[uint8]interface {}{100:(" + H("0+1i") + ")}"}, {s1: "skip"}},
		{{s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{s1: "skip"}, {s1: "skip"},
			{"map[interface {}]float32{100:" + H("1.25") + "}", "map[interface {}]interface {}{100:(" + H("0+1i") + "), " + H(`"abc":false`) + "}"},
			{"map[uint8]interface {}{100:(0+1i)}", "map[interface {}]interface {}{100:(0+1i), " + H(`"abc":false`) + "}"}, {s1: "skip"}},
		{{H("<nil>"), "assert.A{a:<nil>, b:<nil>}"}, {H("int8") + "(-1)", H("assert.A") + "{a:<nil>, b:<nil>}"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{s1: "skip"}, {H("uint8") + "(1)", H("assert.A") + "{a:<nil>, b:<nil>}"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{H("uintptr") + "(0x186a0)", H("assert.A") + "{a:<nil>, b:<nil>}"}, {H("float32") + "(1.25)", H("assert.A") + "{a:<nil>, b:<nil>}"}, {s1: "skip"},
			{H("complex64") + "(2+0i)", H("assert.A") + "{a:<nil>, b:<nil>}"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{"(" + H("*int") + ")(" + sa + ")", H("assert.A") + "{a:<nil>, b:<nil>}"}, {s1: "skip"},
			{"(" + H("unsafe.Pointer") + ")(0xc8)", H("assert.A") + "{a:<nil>, b:<nil>}"},
			{"(" + H("chan int") + ")(" + sb + ")", H("assert.A") + "{a:<nil>, b:<nil>}"}, {"(" + H("chan uint") + ")(nil)", H("assert.A") + "{a:<nil>, b:<nil>}"},
			{"(" + H("func(int) bool") + ")(" + sc + ")", H("assert.A") + "{a:<nil>, b:<nil>}"}, {H("string") + `("中a文b")`, H("assert.A") + "{a:<nil>, b:<nil>}"},
			{s1: "skip"}, {H("[0]int") + "{}", H("assert.A") + "{a:<nil>, b:<nil>}"}, {H("[3]interface {}") + "{1, 2, (3+0i)}", H("assert.A") + "{a:<nil>, b:<nil>}"},
			{H("[]int") + "(nil)", H("assert.A") + "{a:<nil>, b:<nil>}"}, {H("[]int") + "{}", H("assert.A") + "{a:<nil>, b:<nil>}"},
			{H("[]interface {}") + "{1, 2, (0+3i), 0x64}", H("assert.A") + "{a:<nil>, b:<nil>}"}, {H("map[bool]int") + "(nil)", H("assert.A") + "{a:<nil>, b:<nil>}"},
			{H("map[bool]uint") + "{}", H("assert.A") + "{a:<nil>, b:<nil>}"}, {H("map[interface {}]float32") + "{100:1.25}", H("assert.A") + "{a:<nil>, b:<nil>}"},
			{H("map[uint8]interface {}") + "{100:(0+1i)}", H("assert.A") + "{a:<nil>, b:<nil>}"}, {s1: "skip"}, {s1: "skip"}},
		{{s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {"{a:" + H("<nil>") + " b:<nil>}", "{a:" + H("int") + "(100) b:<nil>}"},
			{s1: "skip"}},
		{{s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{"assert.A{a:" + H("<nil>") + ", b:assert.I(" + H("nil") + ")}",
				"assert.A{\n\ta:\x1b[41mstring\x1b[0m(\"中a文b\"),\n\tb:assert.A\x1b[41m{a:<nil>, b:<nil>}\x1b[0m\n}"},
			{"assert.A{a:" + H("100") + ", b:assert.I(" + H("nil") + ")}",
				"assert.A{\n\ta:\x1b[41m\"中a文b\"\x1b[0m,\n\tb:assert.A\x1b[41m{a:<nil>, b:<nil>}\x1b[0m\n}"},
			{s1: "skip"}},
		{{s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"}, {s1: "skip"},
			{"assert.A{\n\ta:\"\x1b[41m中a\x1b[0m文b\",\n\tb:assert.A{a:<nil>, b:assert.I(\x1b[41mnil\x1b[0m)}\n}",
				"assert.A{\n\ta:\"\x1b[41ma中\x1b[0m文b\x1b[41mc\x1b[0m\",\n\tb:assert.A{\n\t\ta:<nil>,\n\t\tb:assert.A\x1b[41m{a:<nil>, b:<nil>}\x1b[0m\n\t}\n}"},
			{s1: "skip"}},
	}
	for i := 1; i < len(cs); i++ {
		b := cs[i]
		for j := 0; j <= i; j++ {
			a := cs[j]
			var r struct{ s1, s2 string }
			if i < len(rs)+1 && j < len(rs[i-1]) {
				r = rs[i-1][j]
			}
			if r.s1 == "skip" {
				continue
			}
			var d1, d2 tValueDiffer
			d1.writeValDiff(reflect.ValueOf(a), reflect.ValueOf(b), false)
			Equal(t, r.s1, d1.String(0), "i=%v, j=%v, s1\n%v\t%T\n%v\t%T", i, j, d1.String(0), a, d1.String(1), b)
			Equal(t, r.s2, d1.String(1), "i=%v, j=%v, s2\n%v\t%T\n%v\t%T", i, j, d1.String(0), a, d1.String(1), b)
			d2.writeValDiff(reflect.ValueOf(a), reflect.ValueOf(b), true)
			Equal(t, r.s2, d2.String(0), "i=%v, j=%v, s2\n%v\n%v", i, j, d2.String(1), d2.String(0))
			Equal(t, r.s1, d2.String(1), "i=%v, j=%v, s1\n%v\n%v", i, j, d2.String(1), d2.String(0))
		}
	}
}

func TestWriteValDiffStruct(t *testing.T) {
	a := A{a: "中a文b", b: A{}}
	b := A{a: "a中文bc", b: A{a: 100, b: A{}}}
	var d tValueDiffer
	d.writeValDiff(reflect.ValueOf(a), reflect.ValueOf(b), false)
	Equal(t, "assert.A{\n\ta:\"\x1b[41m中a\x1b[0m文b\",\n\tb:assert.A{a:\x1b[41m<nil>\x1b[0m, b:assert.I(\x1b[41mnil\x1b[0m)}\n}", d.String(0), "s1\n%v\n%v", d.String(0), d.String(1))
	Equal(t, "assert.A{\n\ta:\"\x1b[41ma中\x1b[0m文b\x1b[41mc\x1b[0m\",\n\tb:assert.A{\n\t\ta:\x1b[41mint\x1b[0m(100),\n\t\tb:assert.A\x1b[41m{a:<nil>, b:<nil>}\x1b[0m\n\t}\n}", d.String(1), "s2\n%v\n%v", d.String(0), d.String(1))

}

func TestWriteValDiff2(t *testing.T) {
	te := func(a, b interface{}, s1, s2 string) {
		var d tValueDiffer
		d.writeValDiff(reflect.ValueOf(a), reflect.ValueOf(b), false)
		Caller(3).Equal(t, s1, d.String(0), "s1\n%v\n%v", d.String(0), d.String(1))
		Caller(3).Equal(t, s2, d.String(1), "s2\n%v\n%v", d.String(0), d.String(1))
	}
	tt := func(a interface{}, rs []struct{ s1, s2 string }) {
		i := 0
		eq := func(b interface{}) {
			var d tValueDiffer
			d.writeValDiff(reflect.ValueOf(a), reflect.ValueOf(b), false)
			var r struct{ s1, s2 string }
			if i < len(rs) {
				r = rs[i]
			}
			i++
			if r.s1 != "skip" {
				te(a, b, r.s1, r.s2)
			}
		}
		eq(nil)
		eq(false)
		eq(int(-1))
		eq((*int)(nil))
		eq((unsafe.Pointer)(nil))
		eq(chan int(nil))
		eq((func(bool) int)(nil))
		eq("abc")
		eq([...]int{1, 2, 3})
		eq([]int{1, 2, 3})
		eq(map[bool]int{true: 1})
		eq(A{})
	}
	tt(true, []struct{ s1, s2 string }{
		{H("bool") + "(true)", H("<nil>")},
		{"bool(" + H("true") + ")", "bool(" + H("false") + ")"},
		{H("bool") + "(true)", H("int") + "(-1)"},
		{H("bool") + "(true)", "(" + H("*int") + ")(nil)"},
		{H("bool") + "(true)", "(" + H("unsafe.Pointer") + ")(nil)"},
		{H("bool") + "(true)", "(" + H("chan int") + ")(nil)"},
		{H("bool") + "(true)", "(" + H("func(bool) int") + ")(nil)"},
		{H("bool") + "(true)", H("string") + `("abc")`},
		{H("bool") + "(true)", H("[3]int") + "{1, 2, 3}"},
		{H("bool") + "(true)", H("[]int") + "{1, 2, 3}"},
		{H("bool") + "(true)", H("map[bool]int") + "{true:1}"},
		{H("bool") + "(true)", H("assert.A") + "{a:<nil>, b:<nil>}"},
	})
	tt(Bool(true), []struct{ s1, s2 string }{
		{H("assert.Bool") + "(String of Bool)", H("<nil>")},
		{"assert.Bool(" + H("String of Bool") + ")", "bool(" + H("false") + ")"},
		{H("assert.Bool") + "(String of Bool)", H("int") + "(-1)"},
		{H("assert.Bool") + "(String of Bool)", "(" + H("*int") + ")(nil)"},
		{H("assert.Bool") + "(String of Bool)", "(" + H("unsafe.Pointer") + ")(nil)"},
		{H("assert.Bool") + "(String of Bool)", "(" + H("chan int") + ")(nil)"},
		{H("assert.Bool") + "(String of Bool)", "(" + H("func(bool) int") + ")(nil)"},
		{H("assert.Bool") + "(String of Bool)", H("string") + `("abc")`},
		{H("assert.Bool") + "(String of Bool)", H("[3]int") + "{1, 2, 3}"},
		{H("assert.Bool") + "(String of Bool)", H("[]int") + "{1, 2, 3}"},
		{H("assert.Bool") + "(String of Bool)", H("map[bool]int") + "{true:1}"},
		{"assert." + H("Bool") + "(String of Bool)", "assert." + H("A") + "{a:<nil>, b:<nil>}"},
	})
	te((func(Bool, If) (Array, Map))(nil), (func(bool, If) (Slice, Ptr))(nil),
		"(func(\x1b[41massert.Bool\x1b[0m, assert.If) (assert.\x1b[41mArray\x1b[0m, assert.\x1b[41mMap\x1b[0m))(nil)",
		"(func(\x1b[41mbool\x1b[0m, assert.If) (assert.\x1b[41mSlice\x1b[0m, assert.\x1b[41mPtr\x1b[0m))(nil)")
	a1, a2 := &[...]int{1, 2, 3}, &[]int{1, 2, 4}
	te(a1, a2, fmt.Sprintf("(*"+H("[3]int")+")(%p)", a1), fmt.Sprintf("(*"+H("[]int")+")(%p)", a2))
	te((*[]map[[12]bool]float32)(nil), (*[]map[[3]A]Bool)(nil),
		"(*[]map[[\x1b[41m12\x1b[0m]\x1b[41mbool\x1b[0m]\x1b[41mfloat32\x1b[0m)(nil)",
		"(*[]map[[\x1b[41m3\x1b[0m]\x1b[41massert.A\x1b[0m]\x1b[41massert.Bool\x1b[0m)(nil)")
	te(errors.New("abc"), errors.New("abde"), `&{s:"ab`+H("c")+`"}`, `&{s:"ab`+H("de")+`"}`)
}

func TestWriteValDiff3(t *testing.T) {
	te := func(a, b interface{}, s1, s2 string) {
		var d1, d2 tValueDiffer
		d1.writeValDiff(reflect.ValueOf(a), reflect.ValueOf(b), false)
		Caller(3).Equal(t, s1, d1.String(0), "s1\n%v\n%v", d1.String(0), d1.String(1))
		Caller(3).Equal(t, s2, d1.String(1), "s2\n%v\n%v", d1.String(0), d1.String(1))
		d2.writeValDiff(reflect.ValueOf(a), reflect.ValueOf(b), true)
		Caller(3).Equal(t, s2, d2.String(0), "s2\n%v\n%v", d2.String(1), d2.String(0))
		Caller(3).Equal(t, s1, d2.String(1), "s1\n%v\n%v", d2.String(1), d2.String(0))
	}
	tt := func(a, b string, b1, b2, c1, c2 interface{}, rs []struct{ s1, s2 string }) {
		i := 0
		eq := func(a, b interface{}) {
			var d tValueDiffer
			d.writeValDiff(reflect.ValueOf(a), reflect.ValueOf(b), false)
			var r struct{ s1, s2 string }
			if i < len(rs) {
				r = rs[i]
			}
			i++
			if r.s1 != "skip" {
				te(a, b, r.s1, r.s2)
			}
		}
		eq(a, b)
		eq(a, []rune(b))
		eq(a, b2)
		eq([]rune(a), b)
		eq(b1, b)
		eq(a, []byte(b))
		eq(a, c2)
		eq([]byte(a), b)
		eq(c1, b)
	}
	tt("中a文b", "a中文bc",
		[...]rune{0x4e2d, 0x61, 0x6587, 0x62},
		[...]rune{0x61, 0x4e2d, 0x6587, 0x62, 0x63},
		[...]byte{0xe4, 0xb8, 0xad, 0x61, 0xe6, 0x96, 0x87, 0x62},
		[...]byte{0x61, 0xe4, 0xb8, 0xad, 0xe6, 0x96, 0x87, 0x62, 0x63},
		[]struct{ s1, s2 string }{
			{`"` + H("中a") + `文b"`, `"` + H("a中") + "文b" + H("c") + `"`},
			{`"` + H("中a") + `文b"`, "[" + H("0x61 0x4e2d") + " 0x6587 0x62 " + H("0x63") + "]"},
			{`"` + H("中a") + `文b"`, "[" + H("0x61 0x4e2d") + " 0x6587 0x62 " + H("0x63") + "]"},
			{"[" + H("0x4e2d 0x61") + " 0x6587 0x62]", `"` + H("a中") + "文b" + H("c") + `"`},
			{"[" + H("0x4e2d 0x61") + " 0x6587 0x62]", `"` + H("a中") + "文b" + H("c") + `"`},
			{`"` + H("中a") + `文b"`, "[" + H("0x61 0xe4 0xb8 0xad") + " 0xe6 0x96 0x87 0x62 " + H("0x63") + "]"},
			{`"` + H("中a") + `文b"`, "[" + H("0x61 0xe4 0xb8 0xad") + " 0xe6 0x96 0x87 0x62 " + H("0x63") + "]"},
			{"[" + H("0xe4 0xb8 0xad 0x61") + " 0xe6 0x96 0x87 0x62]", `"` + H("a中") + "文b" + H("c") + `"`},
			{"[" + H("0xe4 0xb8 0xad 0x61") + " 0xe6 0x96 0x87 0x62]", `"` + H("a中") + "文b" + H("c") + `"`}})
	tt("", "a中文bc",
		[...]rune{},
		[...]rune{0x61, 0x4e2d, 0x6587, 0x62, 0x63},
		[...]byte{},
		[...]byte{0x61, 0xe4, 0xb8, 0xad, 0xe6, 0x96, 0x87, 0x62, 0x63},
		[]struct{ s1, s2 string }{
			{`""`, `"` + H("a中文bc") + `"`},
			{`""`, "[" + H("0x61 0x4e2d 0x6587 0x62 0x63") + "]"},
			{`""`, "[" + H("0x61 0x4e2d 0x6587 0x62 0x63") + "]"},
			{"[]", `"` + H("a中文bc") + `"`},
			{"[]", `"` + H("a中文bc") + `"`},
			{`""`, "[" + H("0x61 0xe4 0xb8 0xad 0xe6 0x96 0x87 0x62 0x63") + "]"},
			{`""`, "[" + H("0x61 0xe4 0xb8 0xad 0xe6 0x96 0x87 0x62 0x63") + "]"},
			{"[]", `"` + H("a中文bc") + `"`},
			{"[]", `"` + H("a中文bc") + `"`}})
	te("a中文bc", [...]byte{0x61, 0xe4, 0xb7, 0xad, 0xe6, 0x96, 0x87, 0x90, 0x63},
		`"a`+H("中")+"文"+H("b")+`c"`, "[0x61 0xe4 "+H("0xb7")+" 0xad 0xe6 0x96 0x87 "+H("0x90")+" 0x63]")
}
