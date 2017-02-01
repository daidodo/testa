package main

import (
	"fmt"
	"reflect"
	"unsafe"
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
}

func (A) Fun(int) string { return "3" }

type I interface {
	Fun(int) string
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
	fmt.Printf(f, &b)

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
		fmt.Println(reflect.TypeOf(a).Kind())
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
