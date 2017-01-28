package main

import (
	"fmt"
	"reflect"
)

const f = "%T\t%[1]v\t%#[1]v\n"

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
