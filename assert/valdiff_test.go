package assert

import (
	"fmt"
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
		{v1: reflect.ValueOf(A{}).Field(1), v2: reflect.ValueOf(A{b: A{}}).Field(1), s2: "assert.A{a:<nil>, b:<nil>}"},
		//TODO
		//{v1: reflect.ValueOf(A{b: A{a: 100}}).Field(1), v2: reflect.ValueOf(A{b: A{a: 100}}).Field(1), s2: "assert.A{a:<nil>, b:<nil>}"},
	}
	for i, c := range cs {
		f := func(v1, v2 reflect.Value, s1, s2 string) {
			var d ValueDiffer
			d.writeTypeDiffValues(v1, v2)
			if s1 == "" {
				s1 = fmt.Sprintf(H("%v"), v1)
			}
			if s2 == "" {
				s2 = fmt.Sprintf(H("%v"), v2)
			}
			Caller(1).Equal(t, s1, d.String(0), "i=%v, r1=\n%v\n%v", i, d.String(0), d.String(1))
			Caller(1).Equal(t, s2, d.String(1), "i=%v, r2=\n%v\n%v", i, d.String(0), d.String(1))
			Caller(1).Equal(t, c.n1, d.Attrs[NewLine], "i=%v, NewLine0: Attrs=%v", i, d.Attrs)
			Caller(1).Equal(t, c.n2, d.Attrs[NewLine+1], "i=%v, NewLine1: Attrs=%v", i, d.Attrs)
			Caller(1).Equal(t, c.om, d.Attrs[OmitSame], "i=%v, OmitSame: Attrs=%v", i, d.Attrs)
			Caller(1).Equal(t, c.cf, d.Attrs[CompFunc], "i=%v, CompFunc: Attrs=%v", i, d.Attrs)
		}
		f(c.v1, c.v2, c.s1, c.s2)
		f(c.v2, c.v1, c.s2, c.s1)
	}
}

func TestWriteTypeDiffValuesString(t *testing.T) {
	eq := func(x1, x2 string, s1, s2 string) {
		var d ValueDiffer
		d.writeDiffValuesString(reflect.ValueOf(x1), reflect.ValueOf(x2))
		Caller(1).Equal(t, s1, d.String(0), "%s\n%s", d.String(0), d.String(1))
		Caller(1).Equal(t, s2, d.String(1), "%s\n%s", d.String(0), d.String(1))
	}
	eq("abcaadef", "accaa", "\"a\x1b[41mb\x1b[0mcaa\x1b[41mdef\x1b[0m\"", "\"a\x1b[41mc\x1b[0mcaa\"")
	eq("This is\x83这是 Chinese 中文！", "This is    这不是Chinase 汉文吗?", "\"This is\x1b[41m\\x83\x1b[0m这\x1b[41m是 \x1b[0mChin\x1b[41me\x1b[0mse \x1b[41m中\x1b[0m文\x1b[41m！\x1b[0m\"", "\"This is\x1b[41m    \x1b[0m这\x1b[41m不是\x1b[0mChin\x1b[41ma\x1b[0mse \x1b[41m汉\x1b[0m文\x1b[41m吗?\x1b[0m\"")
}
