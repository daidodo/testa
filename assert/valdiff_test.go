package assert

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"

	assert2 "github.com/stretchr/testify/assert"
)

func P(p uintptr) unsafe.Pointer {
	return unsafe.Pointer(p)
}

func TestValueEqual(t *testing.T) {
	a := make(chan int)
	b := func() int { return 0 }
	c := new(int)
	*c = 100
	cs := []struct {
		v1, v2 reflect.Value
		e, ci  bool
		k      reflect.Kind
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
		//TODO
		//{v1: reflect.ValueOf(struct{ a interface{} }{100}).Field(0), v2: reflect.ValueOf(100), k: reflect.Int},
	}
	for i, c := range cs {
		ci := c.v1.IsValid() && c.v1.CanInterface() && c.v2.IsValid() && c.v2.CanInterface()
		Equal(t, c.ci, ci, "i=%v", i)
		if c.k != reflect.Invalid {
			Equal(t, c.k, c.v1.Kind(), "i=%v", i)
			Equal(t, c.k, c.v2.Kind(), "i=%v", i)
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

func TestWriteTypeDiffValues(t *testing.T) {
	a := func() int { return 1 }
	cs := []struct {
		v1, v2         reflect.Value
		s1, s2         string
		ss1, ss2       string
		n1, n2, om, cf bool
	}{
		{v1: reflect.ValueOf(true), v2: reflect.ValueOf(false)},
		{v1: reflect.ValueOf(int(100)), v2: reflect.ValueOf(int(101))},
		{v1: reflect.ValueOf(int8(100)), v2: reflect.ValueOf(int8(101))},
		{v1: reflect.ValueOf(int16(100)), v2: reflect.ValueOf(int16(101))},
		{v1: reflect.ValueOf(int32(100)), v2: reflect.ValueOf(int32(101))},
		{v1: reflect.ValueOf(int64(100)), v2: reflect.ValueOf(int64(101))},
		{v1: reflect.ValueOf(uint(100)), v2: reflect.ValueOf(uint(101))},
		{v1: reflect.ValueOf(uint8(100)), v2: reflect.ValueOf(uint8(101))},
		{v1: reflect.ValueOf(uint16(100)), v2: reflect.ValueOf(uint16(101))},
		{v1: reflect.ValueOf(uint32(100)), v2: reflect.ValueOf(uint32(101))},
		{v1: reflect.ValueOf(uint64(100)), v2: reflect.ValueOf(uint64(101))},
		{v1: reflect.ValueOf(uintptr(100)), v2: reflect.ValueOf(uintptr(101)), s1: H("0x64"), s2: H("0x65")},
		{v1: reflect.ValueOf(float32(100.25)), v2: reflect.ValueOf(float32(101.25))},
		{v1: reflect.ValueOf(float64(100.25)), v2: reflect.ValueOf(float64(101.25))},
		{v1: reflect.ValueOf(complex64(100.25 + 200.5i)), v2: reflect.ValueOf(complex64(101.25 + 200.5i)), s1: "(" + H("100.25") + "+200.5)", s2: "(" + H("101.25") + "+200.5)"},
		{v1: reflect.ValueOf(complex64(100.25 + 200.5i)), v2: reflect.ValueOf(complex64(100.25 + 201.5i)), s1: "(100.25+" + H("200.5") + ")", s2: "(100.25+" + H("201.5") + ")"},
		{v1: reflect.ValueOf(complex64(100.25 + 200.5i)), v2: reflect.ValueOf(complex64(101.25 + 201.5i)), s1: "(\x1b[41m100.25\x1b[0m+\x1b[41m200.5\x1b[0m)", s2: "(\x1b[41m101.25\x1b[0m+\x1b[41m201.5\x1b[0m)"},
		{v1: reflect.ValueOf(complex128(100.25 + 200.5i)), v2: reflect.ValueOf(complex128(101.25 + 200.5i)), s1: "(" + H("100.25") + "+200.5)", s2: "(" + H("101.25") + "+200.5)"},
		{v1: reflect.ValueOf(complex128(100.25 + 200.5i)), v2: reflect.ValueOf(complex128(100.25 + 201.5i)), s1: "(100.25+" + H("200.5") + ")", s2: "(100.25+" + H("201.5") + ")"},
		{v1: reflect.ValueOf(complex128(100.25 + 200.5i)), v2: reflect.ValueOf(complex128(101.25 + 201.5i)), s1: "(\x1b[41m100.25\x1b[0m+\x1b[41m200.5\x1b[0m)", s2: "(\x1b[41m101.25\x1b[0m+\x1b[41m201.5\x1b[0m)"},
		{v1: reflect.ValueOf(string("")), v2: reflect.ValueOf(string("abc")), s1: `""`, s2: `"` + H("abc") + `"`},
		{v1: reflect.ValueOf(string("accde")), v2: reflect.ValueOf(string("abc")), s1: `"a` + H("c") + "c" + H("de") + `"`, s2: `"a` + H("b") + `c"`},
		{v1: reflect.ValueOf(string("a中cd文？繁體")), v2: reflect.ValueOf(string("abc！字？")), s1: `"a` + H("中") + "c" + H("d文") + "？" + H("繁體") + `"`, s2: `"a` + H("b") + "c" + H("！字") + `？"`},
		{v1: reflect.ValueOf(chan int(nil)), v2: reflect.ValueOf(make(chan int))},
		{v1: reflect.ValueOf(make(chan int)), v2: reflect.ValueOf(make(chan int))},
		{v1: reflect.ValueOf((func() int)(nil)), v2: reflect.ValueOf(func() int { return 1 })},
		{v1: reflect.ValueOf(func() int { return 2 }), v2: reflect.ValueOf(func() int { return 1 })},
		{v1: reflect.ValueOf(a), v2: reflect.ValueOf(a), cf: true},
		{v1: reflect.ValueOf((*int)(nil)), v2: reflect.ValueOf(new(int))},
		{v1: reflect.ValueOf(new(int)), v2: reflect.ValueOf(new(int))},
		{v1: reflect.ValueOf(unsafe.Pointer(nil)), v2: reflect.ValueOf(unsafe.Pointer(new(int)))},
		{v1: reflect.ValueOf(unsafe.Pointer(new(int))), v2: reflect.ValueOf(unsafe.Pointer(new(int)))},
		{v1: reflect.ValueOf(A{}).Field(0), v2: reflect.ValueOf(A{a: 100}).Field(0), s2: "int(100)"},
		{v1: reflect.ValueOf(A{a: 101}).Field(0), v2: reflect.ValueOf(A{a: 100}).Field(0), s1: H("101"), s2: H("100")},
		{v1: reflect.ValueOf(A{}).Field(1), v2: reflect.ValueOf(A{b: A{}}).Field(1), s2: "assert.A{a:<nil>, b:<nil>}"},
		{v1: reflect.ValueOf(A{b: A{a: 100}}).Field(1), v2: reflect.ValueOf(A{b: A{a: 101}}).Field(1), s1: "{a:" + H("100") + " b:<nil>}", s2: "{a:" + H("101") + " b:<nil>}"},
		{v1: reflect.ValueOf([8]int{}), v2: reflect.ValueOf([8]int{1: 10, 4: 20, 5: 30, 7: 40}), s1: "[0 \x1b[41m0\x1b[0m 0 0 \x1b[41m0 0\x1b[0m 0 \x1b[41m0\x1b[0m]", s2: "[0 \x1b[41m10\x1b[0m 0 0 \x1b[41m20 30\x1b[0m 0 \x1b[41m40\x1b[0m]"},
		{v1: reflect.ValueOf([8]unsafe.Pointer{}),
			v2: reflect.ValueOf([8]unsafe.Pointer{1: P(10), 4: P(20), 5: P(30), 7: P(40)}),
			s1: "[8]unsafe.Pointer{<nil>, \x1b[41m<nil>\x1b[0m, <nil>, <nil>, \x1b[41m<nil>, <nil>\x1b[0m, <nil>, \x1b[41m<nil>\x1b[0m}",
			s2: "[8]unsafe.Pointer{<nil>, \x1b[41m0xa\x1b[0m, <nil>, <nil>, \x1b[41m0x14, 0x1e\x1b[0m, <nil>, \x1b[41m0x28\x1b[0m}"},
		{v1: reflect.ValueOf([11]unsafe.Pointer{}),
			v2: reflect.ValueOf([11]unsafe.Pointer{1: P(10), 4: P(20), 5: P(30), 7: P(40)}),
			s1: "[11]unsafe.Pointer{1:\x1b[41m<nil>\x1b[0m, 4:\x1b[41m<nil>\x1b[0m, 5:\x1b[41m<nil>\x1b[0m, 7:\x1b[41m<nil>\x1b[0m}",
			s2: "[11]unsafe.Pointer{1:\x1b[41m0xa\x1b[0m, 4:\x1b[41m0x14\x1b[0m, 5:\x1b[41m0x1e\x1b[0m, 7:\x1b[41m0x28\x1b[0m}",
			om: true},
		{v1: reflect.ValueOf([8][]int{}),
			v2: reflect.ValueOf([8][]int{1: []int{}, 4: []int{1}, 5: []int{2, 3}, 7: []int{4, 5, 6}}),
			s1: "[8][]int{<nil>, \x1b[41m<nil>\x1b[0m, <nil>, <nil>, \x1b[41m<nil>, <nil>\x1b[0m, <nil>, \x1b[41m<nil>\x1b[0m}",
			s2: "[8][]int{\n\t<nil>, \x1b[41m[]\x1b[0m, <nil>, <nil>,\n\t\x1b[41m[1]\x1b[0m,\n\t\x1b[41m[2 3]\x1b[0m,\n\t<nil>,\n\t\x1b[41m[4 5 6]\x1b[0m\n}",
			n2: true},
		{v1: reflect.ValueOf([11][]int{}),
			v2: reflect.ValueOf([11][]int{1: []int{}, 4: []int{1}, 5: []int{2, 3}, 7: []int{4, 5, 6}}),
			s1: "[11][]int{1:\x1b[41m<nil>\x1b[0m, 4:\x1b[41m<nil>\x1b[0m, 5:\x1b[41m<nil>\x1b[0m, 7:\x1b[41m<nil>\x1b[0m}",
			s2: "[11][]int{\n\t1:\x1b[41m[]\x1b[0m,\n\t4:\x1b[41m[1]\x1b[0m,\n\t5:\x1b[41m[2 3]\x1b[0m,\n\t7:\x1b[41m[4 5 6]\x1b[0m\n}",
			n2: true, om: true},
		{v1: reflect.ValueOf([]int(nil)), v2: reflect.ValueOf([]int{}), s1: H("<nil>")},
		{v1: reflect.ValueOf([]int{1, 2, 3}), v2: reflect.ValueOf([]int{}), s1: "[" + H("1 2 3") + "]", s2: "[]"},
		{v1: reflect.ValueOf(make([]int, 10)), v2: reflect.ValueOf([]int{1: 10, 4: 20, 5: 30, 7: 40}),
			s1: "[0 \x1b[41m0\x1b[0m 0 0 \x1b[41m0 0\x1b[0m 0 \x1b[41m0 0 0\x1b[0m]",
			s2: "[0 \x1b[41m10\x1b[0m 0 0 \x1b[41m20 30\x1b[0m 0 \x1b[41m40\x1b[0m]"},
		{v1: reflect.ValueOf(make([]unsafe.Pointer, 6)),
			v2: reflect.ValueOf([]unsafe.Pointer{1: P(10), 4: P(20), 5: P(30), 7: P(40)}),
			s1: "[]unsafe.Pointer{<nil>, \x1b[41m<nil>\x1b[0m, <nil>, <nil>, \x1b[41m<nil>, <nil>\x1b[0m}",
			s2: "[]unsafe.Pointer{<nil>, \x1b[41m0xa\x1b[0m, <nil>, <nil>, \x1b[41m0x14, 0x1e, <nil>, 0x28\x1b[0m}"},
		{v1: reflect.ValueOf(make([]int, 11)), v2: reflect.ValueOf([]int{1: 10, 4: 20, 5: 30, 7: 40}),
			s1: "[]int{1:\x1b[41m0\x1b[0m, 4:\x1b[41m0\x1b[0m, 5:\x1b[41m0\x1b[0m, 7:\x1b[41m0, 8:0, 9:0, 10:0\x1b[0m}",
			s2: "[]int{1:\x1b[41m10\x1b[0m, 4:\x1b[41m20\x1b[0m, 5:\x1b[41m30\x1b[0m, 7:\x1b[41m40\x1b[0m}",
			om: true},
		{v1: reflect.ValueOf(make([][]int, 6)),
			v2: reflect.ValueOf([][]int{1: []int{}, 4: []int{1}, 5: []int{2, 3}, 7: []int{4, 5, 6}}),
			s1: "[][]int{<nil>, \x1b[41m<nil>\x1b[0m, <nil>, <nil>, \x1b[41m<nil>, <nil>\x1b[0m}",
			s2: "[][]int{\n\t<nil>, \x1b[41m[]\x1b[0m, <nil>, <nil>,\n\t\x1b[41m[1]\x1b[0m,\n\t\x1b[41m[2 3]\x1b[0m,\n\t\x1b[41m<nil>\x1b[0m,\n\t\x1b[41m[4 5 6]\x1b[0m\n}",
			n2: true},
		{v1: reflect.ValueOf(make([][]int, 11)),
			v2: reflect.ValueOf([][]int{1: []int{}, 4: []int{1}, 5: []int{2, 3}, 7: []int{4, 5, 6}}),
			s1: "[][]int{1:\x1b[41m<nil>\x1b[0m, 4:\x1b[41m<nil>\x1b[0m, 5:\x1b[41m<nil>\x1b[0m, 7:\x1b[41m<nil>, 8:<nil>, 9:<nil>, 10:<nil>\x1b[0m}",
			s2: "[][]int{\n\t1:\x1b[41m[]\x1b[0m,\n\t4:\x1b[41m[1]\x1b[0m,\n\t5:\x1b[41m[2 3]\x1b[0m,\n\t7:\x1b[41m[4 5 6]\x1b[0m\n}",
			n2: true, om: true},
		{v1: reflect.ValueOf([][]int{3: []int(nil)}),
			v2: reflect.ValueOf([][]int{4: []int{1}, 5: []int{2, 3}, 7: []int{4, 5, 6}, 10: []int(nil)}),
			s1: "[][]int{3:...}",
			s2: "[][]int{\n\t\x1b[41m4:[1]\x1b[0m,\n\t\x1b[41m5:[2 3]\x1b[0m,\n\t\x1b[41m6:<nil>\x1b[0m,\n\t\x1b[41m7:[4 5 6]\x1b[0m,\n\t\x1b[41m8:<nil>\x1b[0m,\n\t\x1b[41m9:<nil>\x1b[0m,\n\t\x1b[41m10:<nil>\x1b[0m\n}",
			n2: true, om: true},
		{v1: reflect.ValueOf([][]int{4: []int{1}}),
			v2: reflect.ValueOf([][]int{4: []int{1}, 5: []int{2, 3}, 7: []int{4, 5, 6}, 10: []int(nil)}),
			s1: "[][]int{\n\t4:...\n}",
			s2: "[][]int{\n\t\x1b[41m5:[2 3]\x1b[0m,\n\t\x1b[41m6:<nil>\x1b[0m,\n\t\x1b[41m7:[4 5 6]\x1b[0m,\n\t\x1b[41m8:<nil>\x1b[0m,\n\t\x1b[41m9:<nil>\x1b[0m,\n\t\x1b[41m10:<nil>\x1b[0m\n}",
			n1: true, n2: true, om: true},
		{v1: reflect.ValueOf(map[int]bool(nil)), v2: reflect.ValueOf(map[int]bool{}), s1: H("<nil>")},
		{v1: reflect.ValueOf(map[int]bool{1: true}), v2: reflect.ValueOf(map[int]bool{}), s1: "map[" + H("1:true") + "]", s2: "map[]"},
		{v1: reflect.ValueOf(map[int]bool{1: true, 2: false}), v2: reflect.ValueOf(map[int]bool{1: false, 2: false}),
			s1:  "map[1:\x1b[41mtrue\x1b[0m 2:false]",
			s2:  "map[1:\x1b[41mfalse\x1b[0m 2:false]",
			ss1: "map[2:false 1:\x1b[41mtrue\x1b[0m]",
			ss2: "map[2:false 1:\x1b[41mfalse\x1b[0m]"},
		{v1: reflect.ValueOf(map[int]bool{1: true, 2: false}), v2: reflect.ValueOf(map[int]bool{3: true, 2: false}),
			s1: "map[2:false \x1b[41m1:true\x1b[0m]",
			s2: "map[2:false \x1b[41m3:true\x1b[0m]"},
		{v1: reflect.ValueOf(map[int]bool{1: true, 2: true}), v2: reflect.ValueOf(map[int]bool{3: true, 2: false}),
			s1: "map[2:\x1b[41mtrue 1:true\x1b[0m]",
			s2: "map[2:\x1b[41mfalse 3:true\x1b[0m]"},
		{v1: reflect.ValueOf(map[int]uint{1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9, 0: 0, 10: 10}),
			v2: reflect.ValueOf(map[int]uint{1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9, 0: 0, 10: 11}),
			s1: "map[10:\x1b[41m10\x1b[0m]",
			s2: "map[10:\x1b[41m11\x1b[0m]",
			om: true},
		{v1: reflect.ValueOf(map[int]uint{1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9, 0: 0}),
			v2: reflect.ValueOf(map[int]uint{1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9, 0: 0, 10: 11}),
			s1: "map[...:...]",
			s2: "map[\x1b[41m10:11\x1b[0m]",
			om: true},
		{v1: reflect.ValueOf(map[int][]int{1: []int{}, 2: nil, 3: nil}), v2: reflect.ValueOf(map[int][]int{1: []int{1, 2, 3}, 2: nil}),
			s1:  "map[int][]int{1:[], 2:<nil>\x1b[41m, 3:<nil>\x1b[0m}",
			s2:  "map[int][]int{\n\t1:[\x1b[41m1 2 3\x1b[0m],\n\t2:<nil>\n}",
			ss1: "map[int][]int{2:<nil>, 1:[]\x1b[41m, 3:<nil>\x1b[0m}",
			ss2: "map[int][]int{\n\t2:<nil>,\n\t1:[\x1b[41m1 2 3\x1b[0m]\n}",
			n2:  true},
		{v1: reflect.ValueOf(map[int][]int{1: []int{}, 2: nil, 3: nil, 4: []int{}, 5: nil, 6: nil, 7: nil, 8: nil, 9: nil, 10: nil, 11: nil}),
			v2: reflect.ValueOf(map[int][]int{1: []int{}, 2: nil, 3: nil, 4: []int{}, 5: nil, 6: nil, 7: nil, 8: nil, 9: nil, 10: []int{1, 2, 3}, 12: nil}),
			s1: "map[int][]int{10:\x1b[41m<nil>, 11:<nil>\x1b[0m}",
			s2: "map[int][]int{\n\t10:\x1b[41m[1 2 3],\x1b[0m\n\t\x1b[41m12:<nil>\x1b[0m\n}",
			n2: true, om: true},
		{v1: reflect.ValueOf(map[int][]int{1: []int{}, 2: nil, 3: nil, 4: []int{}, 5: nil, 6: nil, 7: nil, 8: nil, 9: nil}),
			v2:  reflect.ValueOf(map[int][]int{1: []int{}, 2: nil, 3: nil, 4: []int{}, 5: nil, 6: nil, 7: nil, 8: nil, 9: nil, 10: []int{1, 2, 3}, 12: nil}),
			s1:  "map[int][]int{...:...}",
			s2:  "map[int][]int{\n\t\x1b[41m10:[1 2 3],\x1b[0m\n\t\x1b[41m12:<nil>\x1b[0m\n}",
			ss2: "map[int][]int{\n\t\x1b[41m12:<nil>,\x1b[0m\n\t\x1b[41m10:[1 2 3]\x1b[0m\n}",
			n2:  true, om: true},
		{v1: reflect.ValueOf(map[int][]int{1: []int{1}, 2: nil, 3: nil, 4: []int{}, 5: nil, 6: nil, 7: nil, 8: nil, 9: nil}),
			v2:  reflect.ValueOf(map[int][]int{1: []int{1}, 2: nil, 3: nil, 4: []int{}, 5: nil, 6: nil, 7: nil, 8: nil, 9: nil, 10: []int{1, 2, 3}, 12: nil}),
			s1:  "map[int][]int{\n\t...:...\n}",
			s2:  "map[int][]int{\n\t\x1b[41m10:[1 2 3],\x1b[0m\n\t\x1b[41m12:<nil>\x1b[0m\n}",
			ss2: "map[int][]int{\n\t\x1b[41m12:<nil>,\x1b[0m\n\t\x1b[41m10:[1 2 3]\x1b[0m\n}",
			n1:  true, n2: true, om: true},
		{v1: reflect.ValueOf(struct {
			a int
			b unsafe.Pointer
		}{100, P(100)}),
			v2: reflect.ValueOf(struct {
				a int
				b unsafe.Pointer
			}{101, P(101)}),
			s1: "{a:\x1b[41m100\x1b[0m b:\x1b[41m0x64\x1b[0m}",
			s2: "{a:\x1b[41m101\x1b[0m b:\x1b[41m0x65\x1b[0m}",
		},
		{v1: reflect.ValueOf(struct{ a, b, c, d, e, f, g, h, i, j, k []int }{}),
			v2: reflect.ValueOf(struct{ a, b, c, d, e, f, g, h, i, j, k []int }{b: []int{}, e: []int{}, f: []int{}}),
			s1: "{b:\x1b[41m<nil>\x1b[0m e:\x1b[41m<nil>\x1b[0m f:\x1b[41m<nil>\x1b[0m}",
			s2: "{b:\x1b[41m[]\x1b[0m e:\x1b[41m[]\x1b[0m f:\x1b[41m[]\x1b[0m}",
			om: true},
		{v1: reflect.ValueOf(struct{ a, b, c []int }{}),
			v2: reflect.ValueOf(struct{ a, b, c []int }{b: []int{1}, c: []int{2, 3}}),
			s1: "struct{a:<nil> b:\x1b[41m<nil>\x1b[0m c:\x1b[41m<nil>\x1b[0m}",
			s2: "struct{\n\ta:<nil>,\n\tb:\x1b[41m[1]\x1b[0m,\n\tc:\x1b[41m[2 3]\x1b[0m\n}",
			n2: true},
		{v1: reflect.ValueOf(struct{ a, b, c, d, e, f, g, h, i, j, k []int }{}),
			v2: reflect.ValueOf(struct{ a, b, c, d, e, f, g, h, i, j, k []int }{b: []int{}, e: []int{1}, f: []int{2, 3}}),
			s1: "struct{b:\x1b[41m<nil>\x1b[0m e:\x1b[41m<nil>\x1b[0m f:\x1b[41m<nil>\x1b[0m}",
			s2: "struct{\n\tb:\x1b[41m[]\x1b[0m,\n\te:\x1b[41m[1]\x1b[0m,\n\tf:\x1b[41m[2 3]\x1b[0m\n}",
			n2: true, om: true},
		{v1: reflect.ValueOf(reflect.Array), v2: reflect.ValueOf(reflect.Bool), s1: H("array"), s2: H("bool")},
		{v1: reflect.ValueOf(struct{ a reflect.Kind }{reflect.Array}).Field(0), v2: reflect.ValueOf(struct{ a reflect.Kind }{reflect.Bool}).Field(0)},
		{v1: reflect.ValueOf(PInt(100)), v2: reflect.ValueOf(PInt(101)), s1: H("String of PInt"), s2: H("String of PInt")},
		{v1: reflect.ValueOf(struct{ a PInt }{100}).Field(0), v2: reflect.ValueOf(struct{ a PInt }{101}).Field(0), s1: H("0x64"), s2: H("0x65")},
		{v1: reflect.ValueOf(PStr(100)), v2: reflect.ValueOf(PStr(101)), s1: H("Go String of PStr"), s2: H("Go String of PStr")},
		{v1: reflect.ValueOf(struct{ a PStr }{100}).Field(0), v2: reflect.ValueOf(struct{ a PStr }{101}).Field(0), s1: H("0x64"), s2: H("0x65")},
	}
	for i, c := range cs {
		f := func(v1, v2 reflect.Value, s1, s2, ss1, ss2 string, n1, n2 bool) {
			var d tValueDiffer
			d.writeTypeDiffValues(v1, v2)
			if s1 == "" {
				s1 = H(fmt.Sprintf("%v", v1))
			}
			if s2 == "" {
				s2 = H(fmt.Sprintf("%v", v2))
			}
			if ss1 != "" && ss1 == d.String(0) {
				Caller(1).Equal(t, ss1, d.String(0), "i=%v, s1\n%v\n%v", i, d.String(0), d.String(1))
			} else {
				Caller(1).Equal(t, s1, d.String(0), "i=%v, s1\n%v\n%v", i, d.String(0), d.String(1))
			}
			if ss2 != "" && ss2 == d.String(1) {
				Caller(1).Equal(t, ss2, d.String(1), "i=%v, s1\n%v\n%v", i, d.String(0), d.String(1))
			} else {
				Caller(1).Equal(t, s2, d.String(1), "i=%v, s2\n%v\n%v", i, d.String(0), d.String(1))
			}
			Caller(1).Equal(t, n1, d.Attrs[NewLine], "i=%v, NewLine1: Attrs=%v", i, d.Attrs)
			Caller(1).Equal(t, n2, d.Attrs[NewLine+1], "i=%v, NewLine2: Attrs=%v", i, d.Attrs)
			Caller(1).Equal(t, c.om, d.Attrs[OmitSame], "i=%v, OmitSame: Attrs=%v", i, d.Attrs)
			Caller(1).Equal(t, c.cf, d.Attrs[CompFunc], "i=%v, CompFunc: Attrs=%v", i, d.Attrs)
		}
		f(c.v1, c.v2, c.s1, c.s2, c.ss1, c.ss2, c.n1, c.n2)
		f(c.v2, c.v1, c.s2, c.s1, c.ss2, c.ss1, c.n2, c.n1)
	}
}

func TestWriteTypeDiffValuesArrayShort2(t *testing.T) {
	var a [10][10]int
	for i := range a {
		for j := range a[0] {
			a[i][j] = 100 + i + j*j
		}
	}
	b := a
	b[1][1]++
	b[len(b)/2][len(b[0])/2]++
	b[len(b)-1][len(b[0])-2]++
	b[len(b)-1][len(b[0])-1]++
	var d tValueDiffer
	d.writeTypeDiffValues(reflect.ValueOf(a), reflect.ValueOf(b))
	s1 := "[10][10]int{\n\t[100 101 104 109 116 125 136 149 164 181],\n\t[101 \x1b[41m102\x1b[0m 105 110 117 126 137 150 165 182],\n\t[102 103 106 111 118 127 138 151 166 183],\n\t[103 104 107 112 119 128 139 152 167 184],\n\t[104 105 108 113 120 129 140 153 168 185],\n\t[105 106 109 114 121 \x1b[41m130\x1b[0m 141 154 169 186],\n\t[106 107 110 115 122 131 142 155 170 187],\n\t[107 108 111 116 123 132 143 156 171 188],\n\t[108 109 112 117 124 133 144 157 172 189],\n\t[109 110 113 118 125 134 145 158 \x1b[41m173 190\x1b[0m]\n}"
	s2 := "[10][10]int{\n\t[100 101 104 109 116 125 136 149 164 181],\n\t[101 \x1b[41m103\x1b[0m 105 110 117 126 137 150 165 182],\n\t[102 103 106 111 118 127 138 151 166 183],\n\t[103 104 107 112 119 128 139 152 167 184],\n\t[104 105 108 113 120 129 140 153 168 185],\n\t[105 106 109 114 121 \x1b[41m131\x1b[0m 141 154 169 186],\n\t[106 107 110 115 122 131 142 155 170 187],\n\t[107 108 111 116 123 132 143 156 171 188],\n\t[108 109 112 117 124 133 144 157 172 189],\n\t[109 110 113 118 125 134 145 158 \x1b[41m174 191\x1b[0m]\n}"
	Equal(t, s1, d.String(0), "%v\n%v", d.String(0), d.String(1))
	Equal(t, s2, d.String(1), "%v\n%v", d.String(0), d.String(1))
}

func TestWriteTypeDiffValuesArrayLong2(t *testing.T) {
	var a [14][14]int
	for i := range a {
		for j := range a[0] {
			a[i][j] = 100 + i + j*j
		}
	}
	b := a
	b[1][1]++
	b[len(b)/2][len(b[0])/2]++
	b[len(b)-1][len(b[0])-2]++
	b[len(b)-1][len(b[0])-1]++
	var d tValueDiffer
	d.writeTypeDiffValues(reflect.ValueOf(a), reflect.ValueOf(b))
	s1 := "[14][14]int{\n\t1:[14]int{1:\x1b[41m102\x1b[0m},\n\t7:[14]int{7:\x1b[41m156\x1b[0m},\n\t13:[14]int{12:\x1b[41m257\x1b[0m, 13:\x1b[41m282\x1b[0m}\n}"
	s2 := "[14][14]int{\n\t1:[14]int{1:\x1b[41m103\x1b[0m},\n\t7:[14]int{7:\x1b[41m157\x1b[0m},\n\t13:[14]int{12:\x1b[41m258\x1b[0m, 13:\x1b[41m283\x1b[0m}\n}"
	Equal(t, s1, d.String(0), "%v\n%v", d.String(0), d.String(1))
	Equal(t, s2, d.String(1), "%v\n%v", d.String(0), d.String(1))
}

func TestWriteTypeDiffValuesArrayShort3(t *testing.T) {
	var a [5][5][5]int
	for i := range a {
		for j := range a[0] {
			for k := range a[0][0] {
				a[i][j][k] = 100 + i + j*j + k*k*k
			}
		}
	}
	b := a
	b[1][1][1]++
	b[len(b)/2][len(b[0])/2][len(b[0][0])/3]++
	b[len(b)-1][len(b[0])-1][len(b[0][0])-2]++
	b[len(b)-1][len(b[0])-1][len(b[0][0])-1]++
	var d tValueDiffer
	d.writeTypeDiffValues(reflect.ValueOf(a), reflect.ValueOf(b))
	s1 := "[5][5][5]int{\n\t[5][5]int{\n\t\t[100 101 108 127 164],\n\t\t[101 102 109 128 165],\n\t\t[104 105 112 131 168],\n\t\t[109 110 117 136 173],\n\t\t[116 117 124 143 180]\n\t},\n\t[5][5]int{\n\t\t[101 102 109 128 165],\n\t\t[102 \x1b[41m103\x1b[0m 110 129 166],\n\t\t[105 106 113 132 169],\n\t\t[110 111 118 137 174],\n\t\t[117 118 125 144 181]\n\t},\n\t[5][5]int{\n\t\t[102 103 110 129 166],\n\t\t[103 104 111 130 167],\n\t\t[106 \x1b[41m107\x1b[0m 114 133 170],\n\t\t[111 112 119 138 175],\n\t\t[118 119 126 145 182]\n\t},\n\t[5][5]int{\n\t\t[103 104 111 130 167],\n\t\t[104 105 112 131 168],\n\t\t[107 108 115 134 171],\n\t\t[112 113 120 139 176],\n\t\t[119 120 127 146 183]\n\t},\n\t[5][5]int{\n\t\t[104 105 112 131 168],\n\t\t[105 106 113 132 169],\n\t\t[108 109 116 135 172],\n\t\t[113 114 121 140 177],\n\t\t[120 121 128 \x1b[41m147 184\x1b[0m]\n\t}\n}"
	s2 := "[5][5][5]int{\n\t[5][5]int{\n\t\t[100 101 108 127 164],\n\t\t[101 102 109 128 165],\n\t\t[104 105 112 131 168],\n\t\t[109 110 117 136 173],\n\t\t[116 117 124 143 180]\n\t},\n\t[5][5]int{\n\t\t[101 102 109 128 165],\n\t\t[102 \x1b[41m104\x1b[0m 110 129 166],\n\t\t[105 106 113 132 169],\n\t\t[110 111 118 137 174],\n\t\t[117 118 125 144 181]\n\t},\n\t[5][5]int{\n\t\t[102 103 110 129 166],\n\t\t[103 104 111 130 167],\n\t\t[106 \x1b[41m108\x1b[0m 114 133 170],\n\t\t[111 112 119 138 175],\n\t\t[118 119 126 145 182]\n\t},\n\t[5][5]int{\n\t\t[103 104 111 130 167],\n\t\t[104 105 112 131 168],\n\t\t[107 108 115 134 171],\n\t\t[112 113 120 139 176],\n\t\t[119 120 127 146 183]\n\t},\n\t[5][5]int{\n\t\t[104 105 112 131 168],\n\t\t[105 106 113 132 169],\n\t\t[108 109 116 135 172],\n\t\t[113 114 121 140 177],\n\t\t[120 121 128 \x1b[41m148 185\x1b[0m]\n\t}\n}"
	Equal(t, s1, d.String(0), "%v\n%v", d.String(0), d.String(1))
	Equal(t, s2, d.String(1), "%v\n%v", d.String(0), d.String(1))
}

func TestWriteTypeDiffValuesArrayLong3(t *testing.T) {
	var a [15][15][15]int
	for i := range a {
		for j := range a[0] {
			for k := range a[0][0] {
				a[i][j][k] = 100 + i + j*j + k*k*k
			}
		}
	}
	b := a
	b[1][1][1]++
	b[len(b)/2][len(b[0])/2][len(b[0][0])/3]++
	b[len(b)-1][len(b[0])-1][len(b[0][0])-2]++
	b[len(b)-1][len(b[0])-1][len(b[0][0])-1]++
	var d tValueDiffer
	d.writeTypeDiffValues(reflect.ValueOf(a), reflect.ValueOf(b))
	s1 := "[15][15][15]int{\n\t1:[15][15]int{\n\t\t1:[15]int{1:\x1b[41m103\x1b[0m}\n\t},\n\t7:[15][15]int{\n\t\t7:[15]int{5:\x1b[41m281\x1b[0m}\n\t},\n\t14:[15][15]int{\n\t\t14:[15]int{13:\x1b[41m2507\x1b[0m, 14:\x1b[41m3054\x1b[0m}\n\t}\n}"
	s2 := "[15][15][15]int{\n\t1:[15][15]int{\n\t\t1:[15]int{1:\x1b[41m104\x1b[0m}\n\t},\n\t7:[15][15]int{\n\t\t7:[15]int{5:\x1b[41m282\x1b[0m}\n\t},\n\t14:[15][15]int{\n\t\t14:[15]int{13:\x1b[41m2508\x1b[0m, 14:\x1b[41m3055\x1b[0m}\n\t}\n}"
	Equal(t, s1, d.String(0), "%v\n%v", d.String(0), d.String(1))
	Equal(t, s2, d.String(1), "%v\n%v", d.String(0), d.String(1))
}

func TestWriteTypeDiffValuesArrayShort3Dense(t *testing.T) {
	var a, b [5][5][5]int
	for i := range a {
		for j := range a[0] {
			for k := range a[0][0] {
				a[i][j][k] = 1 + i + j*j + k*k*k
				b[i][j][k] = 1 + i*i + j*j*j + k
			}
		}
	}
	var d tValueDiffer
	d.writeTypeDiffValues(reflect.ValueOf(a), reflect.ValueOf(b))
	s1 := "[5][5][5]int{\n\t[5][5]int{\n\t\t[1 2 \x1b[41m9 28 65\x1b[0m],\n\t\t[2 3 \x1b[41m10 29 66\x1b[0m],\n\t\t[\x1b[41m5 6 13 32 69\x1b[0m],\n\t\t[\x1b[41m10 11 18 37 74\x1b[0m],\n\t\t[\x1b[41m17 18 25 44 81\x1b[0m]\n\t},\n\t[5][5]int{\n\t\t[2 3 \x1b[41m10 29 66\x1b[0m],\n\t\t[3 4 \x1b[41m11 30 67\x1b[0m],\n\t\t[\x1b[41m6 7 14 33 70\x1b[0m],\n\t\t[\x1b[41m11 12 19 38 75\x1b[0m],\n\t\t[\x1b[41m18 19 26 45 82\x1b[0m]\n\t},\n\t[5][5]int{\n\t\t[\x1b[41m3 4 11 30 67\x1b[0m],\n\t\t[\x1b[41m4 5 12 31 68\x1b[0m],\n\t\t[\x1b[41m7 8\x1b[0m 15 \x1b[41m34 71\x1b[0m],\n\t\t[\x1b[41m12 13 20 39 76\x1b[0m],\n\t\t[\x1b[41m19 20 27 46 83\x1b[0m]\n\t},\n\t[5][5]int{\n\t\t[\x1b[41m4 5\x1b[0m 12 \x1b[41m31 68\x1b[0m],\n\t\t[\x1b[41m5 6\x1b[0m 13 \x1b[41m32 69\x1b[0m],\n\t\t[\x1b[41m8 9 16 35 72\x1b[0m],\n\t\t[\x1b[41m13 14 21\x1b[0m 40 \x1b[41m77\x1b[0m],\n\t\t[\x1b[41m20 21 28 47 84\x1b[0m]\n\t},\n\t[5][5]int{\n\t\t[\x1b[41m5 6 13 32 69\x1b[0m],\n\t\t[\x1b[41m6 7 14 33 70\x1b[0m],\n\t\t[\x1b[41m9 10 17 36 73\x1b[0m],\n\t\t[\x1b[41m14 15 22 41 78\x1b[0m],\n\t\t[\x1b[41m21 22 29 48\x1b[0m 85]\n\t}\n}"
	s2 := "[5][5][5]int{\n\t[5][5]int{\n\t\t[1 2 \x1b[41m3 4 5\x1b[0m],\n\t\t[2 3 \x1b[41m4 5 6\x1b[0m],\n\t\t[\x1b[41m9 10 11 12 13\x1b[0m],\n\t\t[\x1b[41m28 29 30 31 32\x1b[0m],\n\t\t[\x1b[41m65 66 67 68 69\x1b[0m]\n\t},\n\t[5][5]int{\n\t\t[2 3 \x1b[41m4 5 6\x1b[0m],\n\t\t[3 4 \x1b[41m5 6 7\x1b[0m],\n\t\t[\x1b[41m10 11 12 13 14\x1b[0m],\n\t\t[\x1b[41m29 30 31 32 33\x1b[0m],\n\t\t[\x1b[41m66 67 68 69 70\x1b[0m]\n\t},\n\t[5][5]int{\n\t\t[\x1b[41m5 6 7 8 9\x1b[0m],\n\t\t[\x1b[41m6 7 8 9 10\x1b[0m],\n\t\t[\x1b[41m13 14\x1b[0m 15 \x1b[41m16 17\x1b[0m],\n\t\t[\x1b[41m32 33 34 35 36\x1b[0m],\n\t\t[\x1b[41m69 70 71 72 73\x1b[0m]\n\t},\n\t[5][5]int{\n\t\t[\x1b[41m10 11\x1b[0m 12 \x1b[41m13 14\x1b[0m],\n\t\t[\x1b[41m11 12\x1b[0m 13 \x1b[41m14 15\x1b[0m],\n\t\t[\x1b[41m18 19 20 21 22\x1b[0m],\n\t\t[\x1b[41m37 38 39\x1b[0m 40 \x1b[41m41\x1b[0m],\n\t\t[\x1b[41m74 75 76 77 78\x1b[0m]\n\t},\n\t[5][5]int{\n\t\t[\x1b[41m17 18 19 20 21\x1b[0m],\n\t\t[\x1b[41m18 19 20 21 22\x1b[0m],\n\t\t[\x1b[41m25 26 27 28 29\x1b[0m],\n\t\t[\x1b[41m44 45 46 47 48\x1b[0m],\n\t\t[\x1b[41m81 82 83 84\x1b[0m 85]\n\t}\n}"
	Equal(t, s1, d.String(0), "%v\n%v", d.String(0), d.String(1))
	Equal(t, s2, d.String(1), "%v\n%v", d.String(0), d.String(1))
}

func TestWriteDiffKinds(t *testing.T) {
	cs := []struct {
		t1, t2 reflect.Type
		s1, s2 string
	}{
		{t1: reflect.TypeOf(100), t2: reflect.TypeOf(101), s1: "int", s2: "int"},
		{t1: reflect.TypeOf(100), t2: reflect.TypeOf(101.2), s1: H("int"), s2: H("float64")},
		{t1: reflect.TypeOf(Int(0)), t2: reflect.TypeOf(int(100)), s1: H("assert.Int"), s2: H("int")},
		{t1: reflect.TypeOf(make(chan int)), t2: reflect.TypeOf(make(chan int)), s1: "chan int", s2: "chan int"},
		{t1: reflect.TypeOf(make(<-chan int)), t2: reflect.TypeOf(make(chan int)), s1: H("<-") + "chan int", s2: "chan int"},
		{t1: reflect.TypeOf(make(chan<- int)), t2: reflect.TypeOf(make(chan int)), s1: "chan" + H("<-") + " int", s2: "chan int"},
		{t1: reflect.TypeOf(make(chan<- int)), t2: reflect.TypeOf(make(<-chan int)), s1: "chan" + H("<-") + " int", s2: H("<-") + "chan int"},
		{t1: reflect.TypeOf(make(<-chan chan int)), t2: reflect.TypeOf(make(chan int)), s1: H("<-") + "chan " + H("chan int"), s2: "chan " + H("int")},
		{t1: reflect.TypeOf(make(chan<- chan int)), t2: reflect.TypeOf(make(chan int)), s1: "chan" + H("<- chan int"), s2: "chan " + H("int")},
		{t1: reflect.TypeOf(make(chan<- chan int)), t2: reflect.TypeOf(make(<-chan int)), s1: "chan" + H("<- chan int"), s2: H("<-") + "chan " + H("int")},
		{t1: reflect.TypeOf(make(chan Array)), t2: reflect.TypeOf(make(chan int)), s1: "chan " + H("assert.Array"), s2: "chan " + H("int")},
		{t1: reflect.TypeOf(make(Chan)), t2: reflect.TypeOf(make(chan int)), s1: H("assert.Chan"), s2: H("chan int")},
		{t1: reflect.TypeOf(func() {}), t2: reflect.TypeOf(func() {}), s1: "func()", s2: "func()"},
		{t1: reflect.TypeOf(func(int) float32 { return .1 }), t2: reflect.TypeOf(func() {}), s1: "func(" + H("int") + ") " + H("float32"), s2: "func()"},
		{t1: reflect.TypeOf(func(int, string) (uint, float32) { return 1, .1 }), t2: reflect.TypeOf(func() {}), s1: "func(" + H("int, string") + ") (" + H("uint, float32") + ")", s2: "func()"},
		{t1: reflect.TypeOf(func(int, string) (uint, float32) { return 1, .1 }),
			t2: reflect.TypeOf(func(int) uint { return 2 }),
			s1: "func(int, " + H("string") + ") (uint, " + H("float32") + ")",
			s2: "func(int) uint"},
		{t1: reflect.TypeOf(func(a, b, c, d, e, f, g, h, i int) (j, k, l, m, n, o, p, q int) { return }),
			t2: reflect.TypeOf(func(a int, b uint, c, d int, e float32, f string, g int) (j uintptr, k int, l complex64, m string, n int, o chan int) {
				return
			}),
			s1: "func(int, \x1b[41mint\x1b[0m, int, int, \x1b[41mint, int\x1b[0m, int, \x1b[41mint, int\x1b[0m) (\x1b[41mint\x1b[0m, int, \x1b[41mint, int\x1b[0m, int, \x1b[41mint, int, int\x1b[0m)",
			s2: "func(int, \x1b[41muint\x1b[0m, int, int, \x1b[41mfloat32, string\x1b[0m, int) (\x1b[41muintptr\x1b[0m, int, \x1b[41mcomplex64, string\x1b[0m, int, \x1b[41mchan int\x1b[0m)"},
		{t1: reflect.TypeOf(func(a int, b uint, c, d int, e float32, f string, g int) (j uintptr, k int, l complex64, m string, n int, o chan int) {
			return
		}),
			t2: reflect.TypeOf(func(a int, b uint, c, d int, e float32, f string, g int) (j uintptr, k int, l complex64, m string, n int, o chan int) {
				return
			}),
			s1: "func(int, uint, int, int, float32, string, int) (uintptr, int, complex64, string, int, chan int)",
			s2: "func(int, uint, int, int, float32, string, int) (uintptr, int, complex64, string, int, chan int)"},
		{t1: reflect.TypeOf(func(interface{}, I, interface{}, I) {}),
			t2: reflect.TypeOf(func(I, interface{}, interface{}, I) {}),
			s1: "func(" + H("interface {}, assert.I") + ", interface {}, assert.I)",
			s2: "func(" + H("assert.I, interface {}") + ", interface {}, assert.I)"},
		{t1: reflect.TypeOf(func(Int) Float { return .1 }), t2: reflect.TypeOf(func(int) float64 { return .2 }),
			s1: "func(" + H("assert.Int") + ") " + H("assert.Float"),
			s2: "func(" + H("int") + ") " + H("float64")},
		{t1: reflect.TypeOf(func(struct{ x bool }, struct{ y int }) (interface{}, I) { return 100, A{} }),
			t2: reflect.TypeOf(func(struct{ a bool }, struct{ y uint }) (I, interface{}) { return A{}, 101 }),
			s1: "func(" + H("struct, struct") + ") (" + H("interface {}, assert.I") + ")",
			s2: "func(" + H("struct, struct") + ") (" + H("assert.I, interface {}") + ")"},
		{t1: reflect.TypeOf(Func(nil)), t2: reflect.TypeOf(func() {}), s1: H("assert.Func"), s2: H("func()")},
		{t1: reflect.TypeOf(new(int)), t2: reflect.TypeOf(new(float32)), s1: "*" + H("int"), s2: "*" + H("float32")},
		{t1: reflect.TypeOf(new(Array)), t2: reflect.TypeOf(new(float32)), s1: "*" + H("assert.Array"), s2: "*" + H("float32")},
		{t1: reflect.TypeOf(Ptr(nil)), t2: reflect.TypeOf(new(float32)), s1: H("assert.Ptr"), s2: H("*float32")},
		{t1: reflect.TypeOf([0]int{}), t2: reflect.TypeOf([10]int{}), s1: "[" + H("0") + "]int", s2: "[" + H("10") + "]int"},
		{t1: reflect.TypeOf([10]chan int{}), t2: reflect.TypeOf([10]int{}), s1: "[10]" + H("chan int"), s2: "[10]" + H("int")},
		{t1: reflect.TypeOf([10]chan int{}), t2: reflect.TypeOf([100]int{}), s1: "[" + H("10") + "]" + H("chan int"), s2: "[" + H("100") + "]" + H("int")},
		{t1: reflect.TypeOf([10]chan int{}), t2: reflect.TypeOf([10]chan uint{}), s1: "[10]chan " + H("int"), s2: "[10]chan " + H("uint")},
		{t1: reflect.TypeOf([0]Slice{}), t2: reflect.TypeOf([10]int{}), s1: "[" + H("0") + "]" + H("assert.Slice"), s2: "[" + H("10") + "]" + H("int")},
		{t1: reflect.TypeOf(Array{}), t2: reflect.TypeOf([10]int{}), s1: H("assert.Array"), s2: H("[10]int")},
		{t1: reflect.TypeOf([]int{}), t2: reflect.TypeOf(make([]int, 10)), s1: "[]int", s2: "[]int"},
		{t1: reflect.TypeOf([]int{}), t2: reflect.TypeOf(make([]float32, 10)), s1: "[]" + H("int"), s2: "[]" + H("float32")},
		{t1: reflect.TypeOf([]Map{}), t2: reflect.TypeOf(make([]float32, 10)), s1: "[]" + H("assert.Map"), s2: "[]" + H("float32")},
		{t1: reflect.TypeOf([]chan int{}), t2: reflect.TypeOf(make([]chan float32, 10)), s1: "[]chan " + H("int"), s2: "[]chan " + H("float32")},
		{t1: reflect.TypeOf(Slice(nil)), t2: reflect.TypeOf(make([]int, 10)), s1: H("assert.Slice"), s2: H("[]int")},
		{t1: reflect.TypeOf(map[bool]int{}), t2: reflect.TypeOf(make(map[bool]int)), s1: "map[bool]int", s2: "map[bool]int"},
		{t1: reflect.TypeOf(map[bool]uint{}), t2: reflect.TypeOf(make(map[bool]int)), s1: "map[bool]" + H("uint"), s2: "map[bool]" + H("int")},
		{t1: reflect.TypeOf(map[bool]chan uint{}), t2: reflect.TypeOf(make(map[bool]chan int)), s1: "map[bool]chan " + H("uint"), s2: "map[bool]chan " + H("int")},
		{t1: reflect.TypeOf(map[string]int{}), t2: reflect.TypeOf(make(map[bool]int)), s1: "map[" + H("string") + "]int", s2: "map[" + H("bool") + "]int"},
		{t1: reflect.TypeOf(map[*string]int{}), t2: reflect.TypeOf(make(map[*bool]int)), s1: "map[*" + H("string") + "]int", s2: "map[*" + H("bool") + "]int"},
		{t1: reflect.TypeOf(map[string]uint{}), t2: reflect.TypeOf(make(map[bool]int)), s1: "map[" + H("string") + "]" + H("uint"), s2: "map[" + H("bool") + "]" + H("int")},
		{t1: reflect.TypeOf(map[*string]*uint{}), t2: reflect.TypeOf(make(map[*bool]*int)), s1: "map[*" + H("string") + "]*" + H("uint"), s2: "map[*" + H("bool") + "]*" + H("int")},
		{t1: reflect.TypeOf(map[String]If{}), t2: reflect.TypeOf(make(map[bool]int)), s1: "map[" + H("assert.String") + "]" + H("assert.If"), s2: "map[" + H("bool") + "]" + H("int")},
		{t1: reflect.TypeOf(Map{}), t2: reflect.TypeOf(make(map[bool]int)), s1: H("assert.Map"), s2: H("map[bool]int")},
		{t1: reflect.TypeOf(struct{}{}), t2: reflect.TypeOf(struct{}{}), s1: "struct", s2: "struct"},
		{t1: reflect.TypeOf(struct{ a int }{}), t2: reflect.TypeOf(struct{ a int }{}), s1: "struct", s2: "struct"},
		{t1: reflect.TypeOf(struct{ b int }{}), t2: reflect.TypeOf(struct{ a int }{}), s1: H("struct"), s2: H("struct")},
		{t1: reflect.TypeOf(struct{ a uint }{}), t2: reflect.TypeOf(struct{ a int }{}), s1: H("struct"), s2: H("struct")},
		{t1: reflect.TypeOf(A{}), t2: reflect.TypeOf(struct{ a int }{}), s1: H("assert.A"), s2: H("struct")},
		{t1: reflect.TypeOf(A{}), t2: reflect.TypeOf(Struct{}), s1: "assert." + H("A"), s2: "assert." + H("Struct")},
		{t1: reflect.TypeOf(A{}), t2: reflect.TypeOf(B{}), s1: "assert." + H("A"), s2: "assert." + H("B")},
	}
	for i, c := range cs {
		var d tValueDiffer
		d.writeDiffKinds(c.t1, c.t2)
		Equal(t, c.s1, d.String(0), "i=%v, s1\n%v\n%v", i, d.String(0), d.String(1))
		Equal(t, c.s2, d.String(1), "i=%v, s2\n%v\n%v", i, d.String(0), d.String(1))
	}
}

func TestWriteDiffKindsBeforeValue(t *testing.T) {
	cs := []struct {
		v1, v2 reflect.Value
		s1, s2 string
	}{
		{v1: reflect.ValueOf(func() {}), v2: reflect.ValueOf(func(int) uint { return 1 }), s1: "(func())", s2: "(func(" + H("int") + ") " + H("uint") + ")"},
		{v1: reflect.ValueOf(make(chan int)), v2: reflect.ValueOf(make(chan uint)), s1: "(chan " + H("int") + ")", s2: "(chan " + H("uint") + ")"},
		{v1: reflect.ValueOf(new(int)), v2: reflect.ValueOf(new(uint)), s1: "(*" + H("int") + ")", s2: "(*" + H("uint") + ")"},
		{v1: reflect.ValueOf(unsafe.Pointer(new(int))), v2: reflect.ValueOf(100), s1: "(" + H("unsafe.Pointer") + ")", s2: H("int")},
		{v1: reflect.ValueOf(Func(nil)), v2: reflect.ValueOf(100), s1: "(" + H("assert.Func") + ")", s2: H("int")},
		{v1: reflect.ValueOf(Chan(nil)), v2: reflect.ValueOf(100), s1: "(" + H("assert.Chan") + ")", s2: H("int")},
		{v1: reflect.ValueOf(Ptr(nil)), v2: reflect.ValueOf(100), s1: "(" + H("assert.Ptr") + ")", s2: H("int")},
		{v1: reflect.ValueOf(UPtr(nil)), v2: reflect.ValueOf(100), s1: "(" + H("assert.UPtr") + ")", s2: H("int")},
	}
	for i, c := range cs {
		var d1, d2 tValueDiffer
		d1.writeDiffKindsBeforeValue(c.v1, c.v2)
		d2.writeDiffKindsBeforeValue(c.v2, c.v1)
		Equal(t, c.s1, d1.String(0), "i=%v, s1\n%v\n%v", i, d1.String(0), d1.String(1))
		Equal(t, c.s2, d1.String(1), "i=%v, s2\n%v\n%v", i, d1.String(0), d1.String(1))
		Equal(t, c.s1, d2.String(1), "i=%v, rs2\n%v\n%v", i, d2.String(1), d2.String(0))
		Equal(t, c.s2, d2.String(0), "i=%v, rs1\n%v\n%v", i, d2.String(1), d2.String(0))
	}
}

type Assertions struct{}
type Kind uint

func TestWriteDiffPkgTypes(t *testing.T) {
	cs := []struct {
		t1, t2 reflect.Type
		s1, s2 string
	}{
		{t1: reflect.TypeOf(reflect.Value{}), t2: reflect.TypeOf(make(Chan)), s1: H("reflect.Value"), s2: H("assert.Chan")},
		{t1: reflect.TypeOf(Int(100)), t2: reflect.TypeOf(make(Chan)), s1: "assert." + H("Int"), s2: "assert." + H("Chan")},
		{t1: reflect.TypeOf(reflect.Int), t2: reflect.TypeOf(Kind(100)), s1: H("reflect") + ".Kind", s2: H("assert") + ".Kind"},
		{t1: reflect.TypeOf(Assertions{}), t2: reflect.TypeOf(assert2.Assertions{}), s1: H("testa") + "/assert.Assertions", s2: H("testify") + "/assert.Assertions"},
		{t1: reflect.TypeOf(Int(100)), t2: reflect.TypeOf(assert2.Assertions{}), s1: H("testa") + "/assert." + H("Int"), s2: H("testify") + "/assert." + H("Assertions")},
	}
	for i, c := range cs {
		var d tValueDiffer
		d.writeDiffKinds(c.t1, c.t2)
		Equal(t, c.s1, d.String(0), "i=%v, s1\n%v\n%v", i, d.String(0), d.String(1))
		Equal(t, c.s2, d.String(1), "i=%v, s2\n%v\n%v", i, d.String(0), d.String(1))
	}
}
