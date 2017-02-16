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
	False(t, isNil(String("")))
	False(t, isNil(String("abc")))
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
