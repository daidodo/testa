package assert

import (
	"testing"
	"unsafe"
)

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
	cs := []interface{}{
		nil, true, false,
		int8(-100), int8(0), int8(100),
		int16(-10000), int16(-100), int16(0), int16(100), int16(10000),
		int32(-1000000000), int32(-10000), int32(-100), int32(0), int32(100), int32(10000), int32(1000000000),
		//int64(-1000000000000), int64(-1000000000), int64(-10000), int64(-100), int64(0), int64(100), int64(10000), int64(1000000000), int64(1000000000000),
	}
	rs := []string{
		"100 000 00000 0000000 000000000",
		"010 000 00000 0000000 000000000",
		"001 000 00000 0000000 000000000",

		"000 100 01000 0010000 000100000",
		"000 010 00100 0001000 000010000",
		"000 001 00010 0000100 000001000",

		"000 000 10000 0100000 001000000",
		"000 100 01000 0010000 000100000",
		"000 010 00100 0001000 000010000",
		"000 001 00010 0000100 000001000",
		"000 000 00001 0000010 000000100",

		"000 000 00000 1000000 010000000",
		"000 000 10000 0100000 001000000",
		"000 100 01000 0010000 000100000",
		"000 010 00100 0001000 000010000",
		"000 001 00010 0000100 000001000",
		"000 000 00001 0000010 000000100",
		"000 000 00000 0000001 000000010",
	}
	for i := 0; i < len(cs); i++ {
		e, ri, id := cs[i], "", -1
		if i < len(rs) {
			ri = rs[i]
		}
		rr := func() bool {
			id++
			for id < len(ri) && ri[id] == ' ' {
				id++
			}
			return id < len(ri) && ri[id] == '1'
		}
		for j := 0; j < len(cs); j++ {
			r := rr()
			a := cs[j]
			Equal(t, r, isSameInValue(e, a), "i=%v, j=%v\n%T\t%[3]v\n%T\t%[4]v", i, j, e, a)
			Equal(t, r, isSameInValue(a, e), "i=%v, j=%v\n%T\t%[3]v\n%T\t%[4]v", i, j, e, a)
		}
	}

	//t1 := func(a interface{}, r string) {
	//    i := 0
	//    rr := func() (b bool) {
	//        if i < len(r) {
	//            if r[i] == ' ' {
	//                i++
	//            }
	//            b = r[i] == '1'
	//        }
	//        i++
	//        return
	//    }
	//    tt(a, nil, rr())
	//    tt(a, true, rr())
	//    tt(a, false, rr())

	//    tt(a, int8(-100), rr())
	//    tt(a, int8(0), rr())
	//    tt(a, int8(100), rr())

	//    tt(a, int16(-10000), rr())
	//    tt(a, int16(-100), rr())
	//    tt(a, int16(0), rr())
	//    tt(a, int16(100), rr())
	//    tt(a, int16(10000), rr())

	//    tt(a, int32(-10000), rr())
	//    tt(a, int32(-10000), rr())
	//    tt(a, int32(-100), rr())
	//    tt(a, int32(0), rr())
	//    tt(a, int32(100), rr())
	//    tt(a, int32(10000), rr())

	//    tt(a, int(100), rr())
	//    tt(a, uint(100), rr())

	//    tt(a, uintptr(100), rr())
	//    tt(a, string("100"), rr())
	//    tt(a, chan int(nil), rr())
	//    tt(a, make(chan int), rr())

	//    tt(a, (func(bool) int)(nil), rr())
	//    tt(a, func(bool) int { return 0 }, rr())
	//    tt(a, (*int)(nil), rr())
	//    tt(a, new(int), rr())

	//    tt(a, unsafe.Pointer(nil), rr())
	//    tt(a, unsafe.Pointer(new(int)), rr())
	//    tt(a, [0]int{}, rr())
	//    tt(a, [...]int{1, 2, 3}, rr())

	//    tt(a, []int(nil), rr())
	//    tt(a, []int{}, rr())
	//    tt(a, []int{1, 2, 3}, rr())
	//    tt(a, map[bool]int(nil), rr())

	//    tt(a, map[bool]int{}, rr())
	//    tt(a, map[bool]int{true: 10}, rr())
	//    tt(a, A{}, rr())
	//}
	//t1(nil, "1000 0010 1010 1000 1001 000")
}
