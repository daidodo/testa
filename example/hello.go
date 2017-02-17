package main

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"unsafe"

	"github.com/daidodo/testa/assert"
	assert2 "github.com/stretchr/testify/assert"
)

const f = "%T\t%[1]v\t%#[1]v\n"

type A struct {
	a int
	b string
	c float32
	d *struct {
		aa string
		bb [3]int
	}
	e struct {
		aa string
		bb [3]int
	}
	f interface{}
	g I
}

func (A) Fun(int) string  { return "3" }
func (A) Fun2(int) string { return "3" }

func (A) String() string {
	return "String of A"
}

func (A) GoString() string {
	return "Go String of A"
}

func (A) Error() string {
	return "Error of A"
}

type I interface {
	Fun(int) string
	Fun2(int) string
}

type I2 interface {
	Fun3(int) string
	Fun4(int) string
}

func main() {
	fmt.Printf(f, [...]int{1, 2, 3, 4, 5})
	fmt.Printf(f, [0]int{})
	fmt.Printf(f, map[int]string{1: "abc", 2: "xyz", 3: "aaaa"})
	fmt.Printf(f, complex(12, 5))
	fmt.Printf(f, complex(0, 5))
	fmt.Printf(f, complex(12, 0))
	fmt.Printf(f, reflect.ValueOf(complex(12, 5)))
	fmt.Printf(f, reflect.ValueOf(complex(0, 5)))
	fmt.Printf(f, reflect.ValueOf(complex(12, 0)))
	a := make(chan int)
	fa(a, a)
	fmt.Printf(f, interface{}(nil))
	fmt.Printf(f, nil)
	fb()
	fc()
	a = nil
	fmt.Println(reflect.DeepEqual(nil, a), " ", a == nil)
	fd()
	fmt.Printf(f, [...]string{"Ab c", "Def", "G hi"})
	ff()
	fmt.Printf(f, map[*int]string{})
	fh()
	fg()
	var b interface{} = &[]int{1, 2, 3}
	fmt.Printf(f, b)
	b = []int{1, 2, 3}
	desc(&b)
	fi()
	fj()
	fk()
	fl()
	fm()
	fn()
}

func pp(p uintptr) unsafe.Pointer {
	return unsafe.Pointer(p)
}

func fn() {
	var a A
	desc(nil == a.f)
	desc(reflect.DeepEqual(nil, a.f))
	desc(nil == a.g)
	desc(reflect.DeepEqual(nil, a.g))
	var b chan int
	desc(nil == b)
	desc(reflect.DeepEqual(nil, b))
	desc(reflect.DeepEqual(b, nil))
	var c I2
	desc(c)
	desc(errors.New("abc"))
	desc(struct{ a error }{}.a)
	desc(fmt.Errorf("abc"))
	e1, e2 := errors.New("abc"), errors.New("abc")
	desc(reflect.DeepEqual(e1, e2))
	desc(a)
	d1 := struct{ a interface{} }{100}
	d2 := 100
	desc(reflect.DeepEqual(d1.a, d2))
}

func fm() {
	//a := []unsafe.Pointer{1: unsafe.Pointer(uintptr(10)), 4: unsafe.Pointer(uintptr(20)), 5: unsafe.Pointer(uintptr(30)), 7: unsafe.Pointer(uintptr(40))}
	a := []unsafe.Pointer{1: pp(10)}
	desc(a)
	t1 := reflect.TypeOf(A{})
	t2 := reflect.TypeOf(reflect.Bool)
	t3 := reflect.TypeOf(assert.Caller(1))
	t4 := reflect.TypeOf(struct{}{})
	t5 := reflect.TypeOf(assert2.Assertions{})
	desc(t1.PkgPath())
	desc(t1.String())
	desc(t2.PkgPath())
	desc(t2.String())
	desc(t3.PkgPath())
	desc(t3.String())
	desc(t4.PkgPath())
	desc(t4.String())
	desc(t5.PkgPath())
	desc(t5.String())
	dest(int(0))
	dest(map[int]bool{})
	dest(struct{}{})
	dest(A{})
	type PInt *int
	dest(PInt(new(int)))
	desc(PInt(new(int)))
	type Int int
	desc(reflect.DeepEqual(int(100), Int(100)))
	b := struct{ a interface{} }{100}
	desc(reflect.DeepEqual(b.a, 100))
	fmt.Printf("%#v\n", unsafe.Pointer(new(int)))
}

func fl() {
	c1, c2 := new(chan int), new(chan int)
	v1 := reflect.ValueOf(c1)
	v2 := reflect.ValueOf(c2)
	desc(v1.Pointer() == v2.Pointer())
	desc(v1.Elem() == v2.Elem())
	desc(v1.Elem().Pointer() == v2.Elem().Pointer())
	desc(v1)
	desc(reflect.ValueOf(v1))
	desc(v1.Pointer())
	desc(v1.Elem())
	desc(v1.Elem().Pointer())
	desc(v2.Pointer())
	desc(v2.Elem())
	desc(v2.Elem().Pointer())
	desc(c1 == c2)
	desc(reflect.DeepEqual(c1, c2))
	fmt.Printf("%#v\n%#v\n", *c1, *c2)
	type X struct {
		a *chan int
	}
	a1 := X{new(chan int)}
	a2 := X{new(chan int)}
	fmt.Printf("%#v\n%#v\n", a1, a2)
	fmt.Printf("%#v\n%#v\n", *a1.a, *a2.a)
	desc(a1 == a2)
	desc(reflect.DeepEqual(a1, a2))

	i1, i2 := new(int), new(int)
	desc(i1 == i2)
	desc(reflect.DeepEqual(i1, i2))
	var a []int
	fmt.Println(a)
}

func fk() {
	a := struct {
		a int
		b string
	}{a: 1, b: "abc"}
	b := struct {
		a int
		b string
	}{a: 1, b: "abc"}
	desc(reflect.TypeOf(a) == reflect.TypeOf(b))
	c := struct {
		b string
		a int
	}{a: 1, b: "abc"}
	desc(reflect.TypeOf(a) == reflect.TypeOf(c))
	var d I
	var e I2
	eq := func(a, b interface{}) bool {
		return a == b
	}
	desc(eq(d, e))
	type C struct {
		a I
		b interface{}
	}
	f1 := C{a: A{}, b: 100}
	f2 := C{a: A{}, b: 100}
	desc(reflect.DeepEqual(f1, f2))
	c1 := make(chan int)
	var c2 <-chan int = c1
	var c3 chan<- int = c1
	desc(c1 == c2)
	desc(c1 == c3)
	//desc(c2 == c3)
	s1 := reflect.ValueOf(string("abc"))
	s2 := reflect.ValueOf(int(100))
	desc(s1.String())
	desc(s2.String())
	desc(reflect.Int)
	fmt.Println(reflect.ValueOf(reflect.Int))
	fmt.Println(reflect.ValueOf(reflect.Int).Interface())
}

func fj() {
	a := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}}
	b := [][]int{[]int{1, 2}, []int{5, 6}}
	desc(a, "\n")
	desc(reflect.TypeOf(a))
	desc(reflect.TypeOf(b))
	c := map[int]string{}
	d := map[uint]string{}
	desc(reflect.TypeOf(c))
	desc(reflect.TypeOf(d))
	desc(reflect.TypeOf(c) == reflect.TypeOf(d))
	e := map[int]int8{}
	f := map[int]uint8{}
	desc(reflect.TypeOf(e))
	desc(reflect.TypeOf(f))
	desc(reflect.TypeOf(e) == reflect.TypeOf(f))
	g := reflect.TypeOf(struct {
		a uint
		b string
	}{})
	h := reflect.TypeOf(struct {
		a int
		b string
	}{})
	desc(g)
	desc(h)
	desc(g == h)
	desc(reflect.ValueOf(struct{ i I }{}).Field(0).Type().Name())
	i := &[...]int{1, 2, 3}
	j := &[...]int{1, 2, 3}
	desc(i == j)

	var kk func(int, string, float32)
	desc(kk)
	var k func(int, string, float32) string
	desc(k)
	var l func(int, string, float32) (complex64, string, chan int)
	desc(l)

	var aa chan int
	desc(aa)
	desc(reflect.ValueOf(aa))
}

func fi() {
	a := [3]I{nil, A{}, &A{}}
	desc(a)
	for _, i := range a {
		desc(i)
	}
	b := [3]interface{}{nil, A{}, &A{}}
	desc(b)
	for _, i := range b {
		desc(i)
	}
	c := "abc"
	desc(&c)
	d := struct {
		a *[]int
		b *A
	}{&[]int{1, 2, 3}, &A{}}
	desc(d)
	e := [...]int{1, 2, 3}
	desc(reflect.TypeOf(a).Elem().Kind())
	desc(reflect.TypeOf(e).Elem().Kind())
	f := [...]I{A{}}
	desc(reflect.ValueOf(f).Index(0).Kind())
	g := [...]A{A{}, A{a: 100}}
	desc(g, "\n")
	type A struct {
		a uint
		b *[]int
		c map[uint]bool
	}
	h := &[]int{1, 2, 3}
	type B struct {
		a int
		b []A
		c A
	}
	desc(B{a: 100, b: []A{A{}, A{a: 200, b: h, c: map[uint]bool{20: false}}}, c: A{a: 200, b: h, c: map[uint]bool{10: true}}}, "\n")
}

func fh() {
	b := make(map[[3]int]bool)
	b[[3]int{1, 2, 3}] = true
	b[[3]int{1, 2, 3}] = false
	fmt.Println(b)

	c := make(map[*[]int]bool)
	d := []int{1, 2, 3}
	e := d
	c[&d] = true
	c[&e] = false
	fmt.Println(c)

	ff := make(map[reflect.Value]bool)
	f1, f2 := reflect.ValueOf(b), reflect.ValueOf(b)
	ff[f1] = true
	ff[f2] = false
	fmt.Printf(f, ff)

	v1 := reflect.ValueOf(b)
	v2 := reflect.ValueOf(v1)
	fmt.Printf(f, v1)
	fmt.Printf(f, v2)
}

func fg() {
	v := reflect.ValueOf(struct {
		a interface{}
	}{int32(100)})
	b := v.Field(0)
	a := b.Elem()
	fmt.Printf(f, a)
	fmt.Printf(f, a.Kind())
	fmt.Printf(f, a.Type())
	fmt.Printf(f, a.CanInterface())
}

func ff() {
	if true {
		var a A
		fmt.Printf(f, a)
		fmt.Println(reflect.ValueOf(a).Type().Name())
	}
	if true {
		a := struct {
			a int
			b string
		}{b: "Ab c"}
		fmt.Printf(f, a)
		fmt.Println(reflect.ValueOf(a).Type().Name())
	}
}

func fa(a <-chan int, b chan int) {
	fmt.Printf(f, a)
	fmt.Printf(f, b)
	c := make(chan int, 10)
	fmt.Printf(f, c)
}

func fb() {
	fmt.Printf(f, func(int) string { return "" })
	fmt.Printf(f, func(int) string { return "" })
	fmt.Printf(f, (func(int) string)(nil))
	fmt.Printf(f, (*(func(int) string))(nil))
	//a := func(int) string { return "1" }
	//b := a
	//fmt.Print(a == b)
}

func fc() {
	h := func(int) string { return "1" }
	g := func(int) string { return "2" }
	a := [...]func(int) string{h, g, nil}
	fmt.Printf(f, a)
}

func fd() {
	fmt.Println("--- fd ---")
	if true {
		a := 100
		fmt.Printf(f, a)
		pa := &a
		fmt.Printf(f, pa)
		pa = nil
		fmt.Printf(f, pa)
		ppa := &pa
		fmt.Printf(f, ppa)
		ppa = nil
		fmt.Printf(f, ppa)
	}
	if true {
		a := uint(100)
		fmt.Printf(f, a)
		pa := &a
		fmt.Printf(f, pa)
		pa = nil
		fmt.Printf(f, pa)
		ppa := &pa
		fmt.Printf(f, ppa)
		ppa = nil
		fmt.Printf(f, ppa)
	}
	if true {
		a := uintptr(100)
		fmt.Printf(f, a)
		pa := &a
		fmt.Printf(f, pa)
		pa = nil
		fmt.Printf(f, pa)
		ppa := &pa
		fmt.Printf(f, ppa)
		ppa = nil
		fmt.Printf(f, ppa)
	}
	if true {
		a := 100.123
		fmt.Printf(f, a)
		pa := &a
		fmt.Printf(f, pa)
		pa = nil
		fmt.Printf(f, pa)
		ppa := &pa
		fmt.Printf(f, ppa)
		ppa = nil
		fmt.Printf(f, ppa)
	}
	if true {
		a := 100.123 + 200.345i
		fmt.Printf(f, a)
		pa := &a
		fmt.Printf(f, pa)
		pa = nil
		fmt.Printf(f, pa)
		ppa := &pa
		fmt.Printf(f, ppa)
		ppa = nil
		fmt.Printf(f, ppa)
	}
	if true {
		var a [3]int
		fmt.Printf(f, a)
		pa := &a
		fmt.Printf(f, pa)
		pa = nil
		fmt.Printf(f, pa)
		ppa := &pa
		fmt.Printf(f, ppa)
		ppa = nil
		fmt.Printf(f, ppa)
	}
	if true {
		var a chan int
		fmt.Printf(f, a)
		a = make(chan int, 2)
		fmt.Printf(f, a)
		pa := &a
		fmt.Printf(f, pa)
		pa = nil
		fmt.Printf(f, pa)
		ppa := &pa
		fmt.Printf(f, ppa)
		ppa = nil
		fmt.Printf(f, ppa)
	}
	if true {
		var a func(int) string
		fmt.Printf(f, a)
		a = func(int) string { return "1" }
		fmt.Printf(f, a)
		pa := &a
		fmt.Printf(f, pa)
		pa = nil
		fmt.Printf(f, pa)
		ppa := &pa
		fmt.Printf(f, ppa)
		ppa = nil
		fmt.Printf(f, ppa)
	}
	if true {
		var a I
		fmt.Printf(f, a)
		a = (*A)(nil)
		fmt.Printf(f, a)
		fmt.Println("242: ", reflect.TypeOf(a).Kind())
		a = A{}
		fmt.Printf(f, a)
		fmt.Println(reflect.TypeOf(a).Kind())
		pa := &a
		fmt.Printf(f, pa)
		pa = nil
		fmt.Printf(f, pa)
		ppa := &pa
		fmt.Printf(f, ppa)
		ppa = nil
		fmt.Printf(f, ppa)
	}
	if true {
		var a map[int]string
		fmt.Printf(f, a)
		a = make(map[int]string)
		fmt.Printf(f, a)
		a[1] = "abc"
		a[2] = "bcd"
		a[3] = "ddf"
		fmt.Printf(f, a)
		pa := &a
		fmt.Printf(f, pa)
		pa = nil
		fmt.Printf(f, pa)
		ppa := &pa
		fmt.Printf(f, ppa)
		ppa = nil
		fmt.Printf(f, ppa)
	}
	if true {
		var a []int
		fmt.Printf(f, a)
		a = make([]int, 3, 4)
		fmt.Printf(f, a)
		pa := &a
		fmt.Printf(f, pa)
		pa = nil
		fmt.Printf(f, pa)
		ppa := &pa
		fmt.Printf(f, ppa)
		ppa = nil
		fmt.Printf(f, ppa)
	}
	if true {
		var a string
		fmt.Printf(f, a)
		a = "abcd"
		fmt.Printf(f, a)
		pa := &a
		fmt.Printf(f, pa)
		pa = nil
		fmt.Printf(f, pa)
		ppa := &pa
		fmt.Printf(f, ppa)
		ppa = nil
		fmt.Printf(f, ppa)
	}
	if true {
		var a A
		fmt.Printf(f, a)
		a = A{a: 100, b: "abc", c: 200.123}
		a.d = &struct {
			aa string
			bb [3]int
		}{"123", [3]int{1, 2, 3}}
		fmt.Printf(f, a)
		pa := &a
		fmt.Printf(f, pa)
		pa = nil
		fmt.Printf(f, pa)
		ppa := &pa
		fmt.Printf(f, ppa)
		ppa = nil
		fmt.Printf(f, ppa)
		var up unsafe.Pointer
		fmt.Printf(f, up)
		up = unsafe.Pointer(&a)
		fmt.Printf(f, up)
		a = A{a: 100, b: "abc", c: 200.123, d: &struct {
			aa string
			bb [3]int
		}{aa: "123", bb: [3]int{1, 2, 3}}}
	}
}

func desc(a interface{}, sep ...string) {
	if _, _, ln, ok := runtime.Caller(1); ok {
		fmt.Print(ln, ": ")
	}
	if v := reflect.ValueOf(a); v.IsValid() {
		fmt.Print(v.Kind())
	}
	s := "\t"
	if len(sep) > 0 {
		s = sep[0]
	}
	fmt.Printf("%[2]v%[1]T%[2]v%[1]v%[2]v%+[1]v%[2]v%#[1]v\n", a, s)
}

func dest(a interface{}) {
	if _, _, ln, ok := runtime.Caller(1); ok {
		fmt.Print(ln, ": ")
	}
	t := reflect.TypeOf(a)
	fmt.Printf("name: %v\tstring: %v\tpkg: %v\n", t.Name(), t.String(), t.PkgPath())
}
