package assert

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
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
		{v1: reflect.ValueOf(map[int][]int{1: []int{}, 2: nil, 3: nil}), v2: reflect.ValueOf(map[int][]int{1: []int{1, 2, 3}, 2: nil}),
			s1:  "map[int][]int{1:[], 2:<nil>\x1b[41m, 3:<nil>\x1b[0m}",
			s2:  "map[int][]int{\n\t1:[\x1b[41m1 2 3\x1b[0m],\n\t2:<nil>\n}",
			ss1: "map[int][]int{2:<nil>, 1:[]\x1b[41m, 3:<nil>\x1b[0m}",
			ss2: "map[int][]int{\n\t2:<nil>,\n\t1:[\x1b[41m1 2 3\x1b[0m]\n}",
			n2:  true},
		{v1: reflect.ValueOf(map[int]uint{1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9, 0: 0, 10: 10}),
			v2: reflect.ValueOf(map[int]uint{1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9, 0: 0, 10: 11}),
		},
	}
	for i, c := range cs {
		f := func(v1, v2 reflect.Value, s1, s2, ss1, ss2 string, n1, n2 bool) {
			var d ValueDiffer
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
