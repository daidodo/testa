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

func TestWriteKey(t *testing.T) {
	eq := func(e string, xx ...interface{}) {
		if true { // value
			var v ValueDiffer
			for i, x := range xx {
				if i > 0 {
					v.b[0].Write(" ")
				}
				v.writeKey(0, reflect.ValueOf(x))
			}
			Caller(1).Equal(t, e, v.String(0))
		}
		if true { // interface
			var v ValueDiffer
			for i, x := range xx {
				if i > 0 {
					v.b[0].Write(" ")
					v.b[1].Write(" ")
				}
				s := reflect.ValueOf(struct {
					A interface{}
					a interface{}
				}{x, x})
				v.writeKey(0, s.Field(0))
				v.writeKey(1, s.Field(1))
			}
			Caller(1).Equal(t, e, v.String(0))
			Caller(1).Equal(t, e, v.String(1))
		}
	}
	ep := func(e string, xx ...interface{}) {
		e = fmt.Sprintf(e, xx...)
		if true { // value
			var v ValueDiffer
			for i, x := range xx {
				if i > 0 {
					v.b[0].Write(" ")
				}
				v.writeKey(0, reflect.ValueOf(x))
			}
			Caller(1).Equal(t, e, v.String(0))
		}
		if true { // interface
			var v ValueDiffer
			for i, x := range xx {
				if i > 0 {
					v.b[0].Write(" ")
					v.b[1].Write(" ")
				}
				s := reflect.ValueOf(struct {
					A interface{}
					a interface{}
				}{x, x})
				v.writeKey(0, s.Field(0))
				v.writeKey(1, s.Field(1))
			}
			Caller(1).Equal(t, e, v.String(0))
			Caller(1).Equal(t, e, v.String(1))
		}
	}
	a := int(100)
	pa := &a
	b := uint(100)
	pb := &b
	c := uintptr(100)
	pc := &c
	d := float64(1.23)
	pd := &d
	e := complex(float64(1.23), float64(3.45))
	pe := &e
	f := string("A bc")
	pf := &f
	g := make(chan int)
	pg := &g
	h := func(int) string { return "1" }
	ph := &h
	i := unsafe.Pointer(pa)
	pi := &i
	j := interface{}(2017)
	pj := &j
	// nil
	eq("<nil>", nil)
	// bool
	eq("true", true)
	eq("false", false)
	// number
	eq("100", a)
	eq("100", int8(100))
	eq("100", int16(100))
	eq("100", int32(100))
	eq("100", int64(100))
	eq("100", b)
	eq("100", uint8(100))
	eq("100", uint16(100))
	eq("100", uint32(100))
	eq("100", uint64(100))
	eq("0x64", c)
	eq("1.23", float32(1.23))
	eq("1.23", d)
	eq("(1.23+3.45i)", complex(float32(1.23), float32(3.45)))
	eq("(1.23+3.45i)", e)
	// string
	eq(`"A bc"`, f) // TODO: A bc?
	// channel
	eq("<nil>", chan int(nil))
	ep("%p", g)
	ep("%p", make(<-chan int))
	ep("%p", make(chan<- int))
	// function
	eq("<nil>", (func(int) string)(nil))
	ep("%p", h)
	// interface
	eq("2017", j)
	// pointer
	eq("<nil>", unsafe.Pointer(nil))
	ep("<nil> %[2]p <nil> %[4]p %[5]p", (*int)(nil), pa, (**int)(nil), &pa, unsafe.Pointer(pa))
	ep("<nil> %[2]p <nil> %[4]p %[5]p", (*uint)(nil), pb, (**uint)(nil), &pb, unsafe.Pointer(pb))
	ep("<nil> %[2]p <nil> %[4]p %[5]p", (*uintptr)(nil), pc, (**uintptr)(nil), &pc, unsafe.Pointer(pc))
	ep("<nil> %[2]p <nil> %[4]p %[5]p", (*float64)(nil), pd, (**float64)(nil), &pd, unsafe.Pointer(pd))
	ep("<nil> %[2]p <nil> %[4]p %[5]p", (*complex64)(nil), pe, (**complex64)(nil), &pe, unsafe.Pointer(pe))
	ep("<nil> %[2]p <nil> %[4]p %[5]p", (*string)(nil), pf, (**string)(nil), &pf, unsafe.Pointer(pf))
	ep("<nil> %[2]p <nil> %[4]p %[5]p", (*chan int)(nil), pg, (**chan int)(nil), &pg, unsafe.Pointer(pg))
	ep("<nil> %[2]p <nil> %[4]p %[5]p", (*func(int) string)(nil), ph, (**func(int) string)(nil), &ph, unsafe.Pointer(ph))
	ep("<nil> %[2]p <nil> %[4]p %[5]p", (*unsafe.Pointer)(nil), pi, (**unsafe.Pointer)(nil), &pi, unsafe.Pointer(pi))
	ep("<nil> %[2]p <nil> %[4]p %[5]p", (*interface{})(nil), pj, (**interface{})(nil), &pj, unsafe.Pointer(pj))
	//TODO
	// array & slice
	eq("[] [101 102 103]", [...]int{}, [...]int{101, 102, 103})
	eq("[] [101 102 103] <nil>", []int{}, []int{101, 102, 103}, []int(nil))
	eq("[] [101 102 103]", [...]uint{}, [...]uint{101, 102, 103})
	eq("[] [101 102 103] <nil>", []uint{}, []uint{101, 102, 103}, []uint(nil))
	eq("[] [0x65 0x66 0x67]", [...]uintptr{}, [...]uintptr{101, 102, 103})
	eq("[] [0x65 0x66 0x67] <nil>", []uintptr{}, []uintptr{101, 102, 103}, []uintptr(nil))
	eq("[] [101.123 102.234 103.345]", [...]float64{}, [...]float64{101.123, 102.234, 103.345})
	eq("[] [101.123 102.234 103.345] <nil>", []float64{}, []float64{101.123, 102.234, 103.345}, []float64(nil))
	eq("[] [(101.1+102.2i) (103.3+104.4i)]", [...]complex128{}, [...]complex128{101.1 + 102.2i, 103.3 + 104.4i})
	eq("[] [(101.1+102.2i) (103.3+104.4i)] <nil>", []complex128{}, []complex128{101.1 + 102.2i, 103.3 + 104.4i}, []complex128(nil))
	eq(`[] ["A bc" "De f" "Gh"]`, [...]string{}, [...]string{"A bc", "De f", "Gh"})
	eq(`[] ["A bc" "De f" "Gh"] <nil>`, []string{}, []string{"A bc", "De f", "Gh"}, []string(nil))
	ep("[] [<nil> %[3]v] %[3]v", [...]chan int{}, [...]chan int{nil, g}, g)
	ep("[] [<nil> %[3]v] %[3]v <nil>", []chan int{}, []chan int{nil, g}, g, []chan int(nil))
	ep("[] [<nil> %[3]v] %[3]v", [...]func(int) string{}, [...]func(int) string{nil, h}, h)
	ep("[] [<nil> %[3]v] %[3]v <nil>", []func(int) string{}, []func(int) string{nil, h}, h, []func(int) string(nil))
	ep("[] [<nil> %[3]v] %[3]v", [...]unsafe.Pointer{}, [...]unsafe.Pointer{nil, i}, i)
	ep("[] [<nil> %[3]v] %[3]v <nil>", []unsafe.Pointer{}, []unsafe.Pointer{nil, i}, i, []unsafe.Pointer(nil))
	ep("[] [<nil> %[3]v] %[3]v", [...]interface{}{}, [...]interface{}{nil, j}, j)
	ep("[] [<nil> %[3]v] %[3]v <nil>", []interface{}{}, []interface{}{nil, j}, j, []interface{}(nil))
	//TODO
	// map
	eq("<nil> map[] map[100:101]", map[int]int(nil), map[int]int{}, map[int]int{100: 101}) // TODO
	eq("<nil> map[] map[100:101]", map[uint]int(nil), map[uint]int{}, map[uint]int{100: 101})
	eq("<nil> map[] map[0x64:101]", map[uintptr]uint(nil), map[uintptr]uint{}, map[uintptr]uint{100: 101})
	eq("<nil> map[] map[100.123:0x65]", map[float64]uintptr(nil), map[float64]uintptr{}, map[float64]uintptr{100.123: 101})
	eq("<nil> map[] map[(100.1+200.2i):101.123]", map[complex128]float64(nil), map[complex128]float64{}, map[complex128]float64{100.1 + 200.2i: 101.123})
	eq(`<nil> map[] map["A bc":(300.3+0i)]`, map[string]complex128(nil), map[string]complex128{}, map[string]complex128{"A bc": 100.1 + 200.2})
	ep(`<nil> map[] map[%[4]p:"A bc"] %[4]p`, map[chan int]string(nil), map[chan int]string{}, map[chan int]string{g: "A bc"}, g)
	//map[func(int)string]...
	ep("<nil> map[] map[%[4]p:%[5]p] %[4]p %[5]p", map[unsafe.Pointer]chan int(nil), map[unsafe.Pointer]chan int{}, map[unsafe.Pointer]chan int{i: g}, i, g)
	ep("<nil> map[] map[2017:%[4]p] %[4]p", map[interface{}]func(int) string(nil), map[interface{}]func(int) string{}, map[interface{}]func(int) string{j: h}, h)
	// struct
	// reflect.Value

	test := func(e string, a interface{}) {
		var v ValueDiffer
		v.writeKey(0, reflect.ValueOf(a))
		if e == "" {
			e = fmt.Sprint(a)
		}
		Caller(1).Equal(t, e, v.String(0))
	}
	// struct
	test("{a:0x64 b:[1 2 3] c:<nil>}", struct {
		a uintptr
		b interface{}
		c []byte
	}{100, []int{1, 2, 3}, nil})
	// array
	if true {
		test("[]", [0]int{})
		test("[1 2 3]", [...]int{1, 2, 3})
		test(`["A bc" "De f" "Gh"]`, [...]string{"A bc", "De f", "Gh"})
		test("[[1 2 3] [3 4 5]]", [...][3]int{{1, 2, 3}, {3, 4, 5}})
		test("[<nil> [1 2 3] [3 4 5]]", [...][]int{nil, {1, 2, 3}, {3, 4, 5}})
		test(`[<nil> map[] map[1:"abc"]]`, [...]map[int]string{nil, {}, {1: "abc"}})
		test(`[{a:1 b:"abc"} {a:3 b:"jjl"}]`, [...]struct {
			a int
			b string
		}{{1, "abc"}, {3, "jjl"}})
		test("[]", [0]*int{})
		test("[<nil>]", [...]*int{nil})
		a := 100
		test(fmt.Sprintf("[%v <nil> %[1]v]", &a), [...]*int{&a, nil, &a})
		g := "abc"
		test(fmt.Sprintf("[%v <nil> %[1]v]", &g), [...]*string{&g, nil, &g})
		k := &[3]int{1, 2, 3}
		test(fmt.Sprintf("[%p <nil>]", k), [...]*[3]int{k, nil})
		l := &[]int{1, 2, 3}
		test(fmt.Sprintf("[%p <nil>]", l), [...]*[]int{l, nil})
		m := &map[int]string{1: "123"}
		test(fmt.Sprintf("[%p <nil>]", m), [...]*map[int]string{m, nil})
		n := &struct {
			a int
			b string
		}{1, "abc"}
		test(fmt.Sprintf("[%p <nil>]", n), [...]*struct {
			a int
			b string
		}{n, nil})
		b := &a
		test(fmt.Sprintf("[%v <nil> %[1]v]", &b), [...]**int{&b, nil, &b})
		h := &g
		test(fmt.Sprintf("[%v <nil> %[1]v]", &h), [...]**string{&h, nil, &h})
		c := &[3]int{1, 2, 3}
		test(fmt.Sprintf("[%v <nil>]", &c), [...]**[3]int{&c, nil})
		d := &[]int{1, 2, 3}
		test(fmt.Sprintf("[%v <nil>]", &d), [...]**[]int{&d, nil})
		e := &map[int]string{1: "abc"}
		test(fmt.Sprintf("[%v <nil>]", &e), [...]**map[int]string{&e, nil})
		f := &struct {
			a int
			b string
		}{1, "abc"}
		test(fmt.Sprintf("[%v <nil>]", &f), [...]**struct {
			a int
			b string
		}{&f, nil})
		test(`[<nil> 100 "A bc"]`, [...]interface{}{nil, 100, "A bc"})
		test(`[[] [1 2 3] ["A bc"]]`, [...]interface{}{[0]int{}, [...]int{1, 2, 3}, [...]string{"A bc"}})
		test(`[<nil> [] [1 2 3] ["A bc"]]`, [...]interface{}{[]int(nil), []int{}, []int{1, 2, 3}, []string{"A bc"}})
		test(`[<nil> map[] map[1:"abc"]]`, [...]interface{}{map[int]string(nil), map[int]string{}, map[int]string{1: "abc"}})
		test(`[{x:0 b:""} {a:1 y:map["abc":(1.2+3.4i)]}]`, [...]interface{}{struct {
			x float64
			b string
		}{}, struct {
			a int
			y map[string]complex64
		}{1, map[string]complex64{"abc": 1.2 + 3.4i}}})
		test(fmt.Sprintf(`[<nil> %v %v]`, &a, &g), [...]interface{}{(*int)(nil), &a, &g})
		test(fmt.Sprintf("[<nil> %p]", c), [...]interface{}{(*[2]int)(nil), c})
		test(fmt.Sprintf("[<nil> %p]", d), [...]interface{}{(*[]int)(nil), d})
		test(fmt.Sprintf("[<nil> %p]", e), [...]interface{}{(*map[float32][]byte)(nil), e})
		test(fmt.Sprintf("[<nil> %p]", f), [...]interface{}{(*struct {
			a int
			b string
		})(nil), f})
		test(fmt.Sprintf("[%v <nil> %[1]v]", &b), [...]interface{}{&b, (**int)(nil), &b})
		test(fmt.Sprintf("[%v <nil> %[1]v]", &h), [...]interface{}{&h, (**string)(nil), &h})
		test(fmt.Sprintf("[%v <nil>]", &c), [...]interface{}{&c, (**[4]int)(nil)})
		test(fmt.Sprintf("[%v <nil>]", &d), [...]interface{}{&d, (**[]int)(nil)})
		test(fmt.Sprintf("[%v <nil>]", &e), [...]interface{}{&e, (**map[int][]byte)(nil)})
		test(fmt.Sprintf("[%v <nil>]", &f), [...]interface{}{&f, (**struct {
			a int
			b string
		})(nil)})
		if true {
			test := func(e string, a interface{}) {
				b := reflect.ValueOf(struct {
					a interface{}
				}{a})
				var v ValueDiffer
				v.writeKey(0, b.Field(0))
				Caller(1).Equal(t, e, v.String(0))
			}
			test("100", 100)
			test("100", uint(100))
			test("0x64", uintptr(100))
			// TODO
		}
	}
	// slice // TODO
	if false {
		test("<nil>", []int(nil))
		test("[]", []int{})
		test("[1 2 3]", []int{1, 2, 3})
		test(`["A bc" "De f" "Gh"]`, []string{"A bc", "De f", "Gh"})
		test("[[1 2 3] [3 4 5]]", [][3]int{{1, 2, 3}, {3, 4, 5}})
		test("[[1 2 3] [3 4 5]]", [][]int{{1, 2, 3}, {3, 4, 5}})
		test(`[map[1:"abc"] map[3:"jjl"]]`, []map[int]string{{1: "abc"}, {3: "jjl"}})
		test(`[{a:1 b:"abc"} {a:3 b:"jjl"}]`, []struct {
			a int
			b string
		}{{1, "abc"}, {3, "jjl"}})
		test("[<nil>]", []*int{nil})
		a := 100
		test(fmt.Sprintf("[%v <nil> %[1]v]", &a), []*int{&a, nil, &a})
		test("[&[1 2 3] <nil> &[3 4 5]]", []*[3]int{&[3]int{1, 2, 3}, nil, &[3]int{3, 4, 5}})
		test("[&[1 2 3] <nil> &[3 4 5]]", []*[]int{&[]int{1, 2, 3}, nil, &[]int{3, 4, 5}})
		test(`[&map[1:"abc"] <nil> &map[3:"jjl"]]`, []*map[int]string{&map[int]string{1: "abc"}, nil, &map[int]string{3: "jjl"}})
		test(`[&{a:1 b:"abc"} <nil> &{a:3 b:"jjl"}]`, []*struct {
			a int
			b string
		}{&struct {
			a int
			b string
		}{1, "abc"}, nil, &struct {
			a int
			b string
		}{3, "jjl"}})
		b := &a
		test(fmt.Sprintf("[%v <nil> %[1]v]", &b), []**int{&b, nil, &b})
		c := &[3]int{1, 2, 3}
		test(fmt.Sprintf("[%v <nil>]", &c), []**[3]int{&c, nil})
		d := &[]int{1, 2, 3}
		test(fmt.Sprintf("[%v <nil>]", &d), []**[]int{&d, nil})
		e := &map[int]string{1: "abc"}
		test(fmt.Sprintf("[%v <nil>]", &e), []**map[int]string{&e, nil})
		f := &struct {
			a int
			b string
		}{1, "abc"}
		test(fmt.Sprintf("[%v <nil>]", &f), []**struct {
			a int
			b string
		}{&f, nil})
	}
	// map
	if false { // TODO
		test("<nil>", map[int]string(nil))
		test("map[]", map[int]string{})
		test(`map[1:"abc"]`, map[int]string{1: "abc"})
		test(`map[[1 2]:"abc"]`, map[[2]int]string{{1, 2}: "abc"})
		test(`map[{a:1 b:"kkk"}:"abc"]`, map[struct {
			a int
			b string
		}]string{{1, "kkk"}: "abc"})
		test(`map[<nil>:"abc"]`, map[*int]string{nil: "abc"})
		a := 100
		test(fmt.Sprintf(`map[%v:"abc"]`, &a), map[*int]string{&a: "abc"})
		test(`map[<nil>:"abc"]`, map[*[3]int]string{nil: "abc"})
		test(`map[&[2 3 4]:"abc"]`, map[*[3]int]string{&[3]int{2, 3, 4}: "abc"})
		test(`map[<nil>:"abc"]`, map[*[]int]string{nil: "abc"})
		test(`map[&[2 3 4]:"abc"]`, map[*[]int]string{&[]int{2, 3, 4}: "abc"})
		test(`map[<nil>:"abc"]`, map[*map[float64]int]string{nil: "abc"})
		test(`map[&map[100.456:2]:"abc"]`, map[*map[float64]int]string{&map[float64]int{100.456: 2}: "abc"})
		test(`map[<nil>:"abc"]`, map[*struct {
			a int
			b string
		}]string{nil: "abc"})
		test(`map[&{a:1 b:"kkk"}:"abc"]`, map[*struct {
			a int
			b string
		}]string{&struct {
			a int
			b string
		}{1, "kkk"}: "abc"})
		b := &[3]int{2, 3, 4}
		test(fmt.Sprintf(`map[%v:"abc"]`, &b), map[**[3]int]string{&b: "abc"})
		c := &[]int{2, 3, 4}
		test(fmt.Sprintf(`map[%v:"abc"]`, &c), map[**[]int]string{&c: "abc"})
		d := &map[float64]int{100.456: 2}
		test(fmt.Sprintf(`map[%v:"abc"]`, &d), map[**map[float64]int]string{&d: "abc"})
		e := &struct {
			a int
			b string
		}{1, "kkk"}
		test(fmt.Sprintf(`map[%v:"abc"]`, &e), map[**struct {
			a int
			b string
		}]string{&e: "abc"})
	}
	// pointer
	if false {
		a := true
		test("", &a)
		test("<nil>", (*bool)(nil))
		b := 100
		test("", &b)
		test("<nil>", (*int)(nil))
		c := uint(100)
		test("", &c)
		test("<nil>", (*uint)(nil))
		d := uintptr(100)
		test("", &d)
		test("<nil>", (*uintptr)(nil))
		e := 100.123
		test("", &e)
		test("<nil>", (*float32)(nil))
		f := 100.123 + 4.34i
		test("", &f)
		test("<nil>", (*complex64)(nil))
		g := make(chan int)
		test("<nil>", (*chan int)(nil))
		test("", &g)
		h := func(int) string { return "1" }
		test("<nil>", (*func(int) string)(nil))
		test("", &h)
		test("<nil>", (*[3]int)(nil))
		test(`&["Abc" "D e" "F"]`, &[3]string{"Abc", "D e", "F"})
		test("<nil>", (*map[int]string)(nil))
		test(`&map[1:"abc"]`, &map[int]string{1: "abc"})
		test("<nil>", (*[]int)(nil))
		test(`&["Abc" "D e" "F"]`, &[]string{"Abc", "D e", "F"})
		test("<nil>", (*struct {
			a int
			b string
		})(nil))
		test(`&{a:1 b:"abc"}`, &struct {
			a int
			b string
		}{1, "abc"})
		var i unsafe.Pointer
		test("<nil>", i)
		i = unsafe.Pointer(&i)
		test("", i)
		if true {
			test("<nil>", (**[3]int)(nil))
			a := &[3]string{"Abc", "D e", "F"}
			test("", &a)
			test("<nil>", (**map[int]string)(nil))
			b := &map[int]string{1: "abc"}
			test("", &b)
			test("<nil>", (**[]int)(nil))
			c := &[]string{"Abc", "D e", "F"}
			test("", &c)
			test("<nil>", (**struct {
				a int
				b string
			})(nil))
			d := &struct {
				a int
				b string
			}{1, "abc"}
			test("", &d)
		}
	}
}

func testWriteElem(t *testing.T, e string, a interface{}) {
	var v ValueDiffer
	v.writeElem(0, reflect.ValueOf(a))
	if e == "" {
		e = fmt.Sprint(a)
	}
	Caller(1).Equal(t, e, v.String(0))
}

func TestWriteElem(t *testing.T) {
	// bool
	testWriteElem(t, "true", true)
	testWriteElem(t, "false", false)
	// number
	testWriteElem(t, "100", 100)
	testWriteElem(t, "100", uint(100))
	testWriteElem(t, "0x64", uintptr(100))
	testWriteElem(t, "1.23", 1.23)
	testWriteElem(t, "(1.23+3.45i)", 1.23+3.45i)
	// channel
	testWriteElem(t, "<nil>", chan int(nil))
	testWriteElem(t, "", make(chan int))
	testWriteElem(t, "", make(<-chan int))
	testWriteElem(t, "", make(chan<- int))
	// function
	testWriteElem(t, "<nil>", (func(int) string)(nil))
	testWriteElem(t, "", func(int) string { return "1" })
	// interface
	testWriteElem(t, "<nil>", nil)
	// array
}
