package assert

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestStructName(t *testing.T) {
	type A struct {
		a int
		b string
	}
	Equal(t, "A", structName(reflect.ValueOf(A{})))
	a := struct {
		a int
		b string
	}{
		a: 1,
		b: "abc",
	}
	Equal(t, "struct", structName(reflect.ValueOf(a)))
}

func testWriteKey(t *testing.T, e string, a interface{}) {
	var v ValueDiffer
	v.writeKey(0, reflect.ValueOf(a))
	if e == "" {
		e = fmt.Sprint(a)
	}
	Caller(1).Equal(t, e, v.String(0))
}

func TestWriteKey(t *testing.T) {
	// bool
	testWriteKey(t, "true", true)
	testWriteKey(t, "false", false)
	// number
	testWriteKey(t, "100", 100)
	testWriteKey(t, "100", uint(100))
	testWriteKey(t, "0x64", uintptr(100))
	testWriteKey(t, "1.23", 1.23)
	testWriteKey(t, "(1.23+3.45i)", 1.23+3.45i)
	// array
	testWriteKey(t, "[]", [0]int{})
	testWriteKey(t, "[1 2 3]", [...]int{1, 2, 3})
	testWriteKey(t, "[[1 2 3] [3 4 5]]", [...][3]int{{1, 2, 3}, {3, 4, 5}})
	testWriteKey(t, "[[1 2 3] [3 4 5]]", [...][]int{{1, 2, 3}, {3, 4, 5}})
	testWriteKey(t, "[map[1:\"abc\"] map[3:\"jjl\"]]", [...]map[int]string{{1: "abc"}, {3: "jjl"}})
	testWriteKey(t, "[{a:1 b:\"abc\"} {a:3 b:\"jjl\"}]", [...]struct {
		a int
		b string
	}{{1, "abc"}, {3, "jjl"}})
	// channel
	testWriteKey(t, "<nil>", chan int(nil))
	testWriteKey(t, "", make(chan int))
	testWriteKey(t, "", make(<-chan int))
	testWriteKey(t, "", make(chan<- int))
	// function
	testWriteKey(t, "<nil>", (func(int) string)(nil))
	testWriteKey(t, "", func(int) string { return "1" })
	// interface
	testWriteKey(t, "<nil>", nil)
	// map
	testWriteKey(t, "<nil>", map[int]string(nil))
	testWriteKey(t, "map[]", map[int]string{})
	testWriteKey(t, "map[1:\"abc\"]", map[int]string{1: "abc"})
	testWriteKey(t, "map[[1 2]:\"abc\"]", map[[2]int]string{{1, 2}: "abc"})
	testWriteKey(t, "map[{a:1 b:\"kkk\"}:\"abc\"]", map[struct {
		a int
		b string
	}]string{{1, "kkk"}: "abc"})
	// pointer
	if true {
		a := true
		testWriteKey(t, "", &a)
		testWriteKey(t, "<nil>", (*bool)(nil))
		b := 100
		testWriteKey(t, "", &b)
		testWriteKey(t, "<nil>", (*int)(nil))
		c := uint(100)
		testWriteKey(t, "", &c)
		testWriteKey(t, "<nil>", (*uint)(nil))
		d := uintptr(100)
		testWriteKey(t, "", &d)
		testWriteKey(t, "<nil>", (*uintptr)(nil))
		e := 100.123
		testWriteKey(t, "", &e)
		testWriteKey(t, "<nil>", (*float32)(nil))
		f := 100.123 + 4.34i
		testWriteKey(t, "", &f)
		testWriteKey(t, "<nil>", (*complex64)(nil))
		g := make(chan int)
		testWriteKey(t, "<nil>", (*chan int)(nil))
		testWriteKey(t, "", &g)
		h := func(int) string { return "1" }
		testWriteKey(t, "<nil>", (*func(int) string)(nil))
		testWriteKey(t, "", &h)
		testWriteKey(t, "<nil>", (*interface{})(nil))
		testWriteKey(t, "<nil>", (*[3]int)(nil))
		testWriteKey(t, "&[\"Abc\" \"D e\" \"F\"]", &[3]string{"Abc", "D e", "F"})
		testWriteKey(t, "<nil>", (*map[int]string)(nil))
		testWriteKey(t, "&map[1:\"abc\"]", &map[int]string{1: "abc"})
		testWriteKey(t, "<nil>", (*[]int)(nil))
		testWriteKey(t, "&[\"Abc\" \"D e\" \"F\"]", &[]string{"Abc", "D e", "F"})
		testWriteKey(t, "<nil>", (*struct {
			a int
			b string
		})(nil))
		testWriteKey(t, "&{a:1 b:\"abc\"}", &struct {
			a int
			b string
		}{1, "abc"})
		var i unsafe.Pointer
		testWriteKey(t, "<nil>", i)
		i = unsafe.Pointer(&i)
		testWriteKey(t, "", i)
		if true {
			testWriteKey(t, "<nil>", (**[3]int)(nil))
			a := &[3]string{"Abc", "D e", "F"}
			testWriteKey(t, "", &a)
			testWriteKey(t, "<nil>", (**map[int]string)(nil))
			b := &map[int]string{1: "abc"}
			testWriteKey(t, "", &b)
			testWriteKey(t, "<nil>", (**[]int)(nil))
			c := &[]string{"Abc", "D e", "F"}
			testWriteKey(t, "", &c)
			testWriteKey(t, "<nil>", (**struct {
				a int
				b string
			})(nil))
			d := &struct {
				a int
				b string
			}{1, "abc"}
			testWriteKey(t, "", &d)
		}
	}
}
