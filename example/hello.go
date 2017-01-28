package main

import (
	"fmt"
	"reflect"
)

func main() {
	const f = "%T\t%[1]v\t%#[1]v\n"
	fmt.Printf(f, [...]int{1, 2, 3, 4, 5})
	fmt.Printf(f, [0]int{})
	fmt.Printf(f, map[int]string{1: "abc", 2: "xyz", 3: "aaaa"})
	fmt.Printf(f, complex(12, 5))
	fmt.Printf(f, complex(0, 5))
	fmt.Printf(f, complex(12, 0))
	fmt.Printf(f, reflect.ValueOf(complex(12, 5)))
	fmt.Printf(f, reflect.ValueOf(complex(0, 5)))
	fmt.Printf(f, reflect.ValueOf(complex(12, 0)))
	fmt.Printf(f, make(chan int))

}
