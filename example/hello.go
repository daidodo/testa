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
	d struct {
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
		a = map[int]string{1: "abc", 2: "bcd", 3: "efg"}
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
		a.d.aa = "123"
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
	}
}
