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
		if true { // value, reflect.Value
			var v ValueDiffer
			for i, x := range xx {
				if i > 0 {
					v.b[0].Write(" ")
				}
				vx := reflect.ValueOf(x)
				v.writeKey(0, vx)
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
	d := float64(100.23)
	pd := &d
	e := complex(float64(100.23), float64(300.45))
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
	a0 := &[...]int{}
	a1 := &[...]int{101, 102, 103}
	as := &[]int{101, 102, 103}
	b0 := &[...]uint{}
	b1 := &[...]uint{101, 102, 103}
	bs := &[]uint{101, 102, 103}
	c0 := &[...]uintptr{}
	c1 := &[...]uintptr{101, 102, 103}
	cs := &[]uintptr{101, 102, 103}
	d0 := &[...]float64{}
	d1 := &[...]float64{101.123, 102.234, 103.345}
	ds := &[]float64{101.123, 102.234, 103.345}
	e0 := &[...]complex128{}
	e1 := &[...]complex128{101.1 + 102.2i, 103.3 + 104.4i}
	es := &[]complex128{101.1 + 102.2i, 103.3 + 104.4i}
	f0 := &[...]string{}
	f1 := &[...]string{"A bc", "De f", "Gh"}
	fs := &[]string{"A bc", "De f", "Gh"}
	g0 := &[...]chan int{}
	g1 := &[...]chan int{nil, g}
	gs := &[]chan int{nil, g}
	h0 := &[...]func(int) string{}
	h1 := &[...]func(int) string{nil, h}
	hs := &[]func(int) string{nil, h}
	i0 := &[...]unsafe.Pointer{}
	i1 := &[...]unsafe.Pointer{nil, i}
	is := &[]unsafe.Pointer{nil, i}
	j0 := &[...]interface{}{}
	j1 := &[...]interface{}{nil, j}
	js := &[]interface{}{nil, j}
	type A struct {
		a int
		b uint
		c uintptr
		d float64
		e complex128
		f string
		g chan int
		h func(int) string
		i unsafe.Pointer
		j interface{}
	}
	sa := &A{a, b, c, d, e, f, g, h, i, j}
	type B struct {
		a *int
		b *uint
		c *uintptr
		d *float64
		e *complex128
		f *string
		g *chan int
		h *func(int) string
		i *unsafe.Pointer
		j *interface{}
	}
	sb := &B{pa, pb, pc, pd, pe, pf, pg, ph, pi, pj}
	type C struct {
		a [0]int
		b [0]uint
		c [0]uintptr
		d [0]float64
		e [0]complex128
		f [0]string
		g [0]chan int
		h [0]func(int) string
		i [0]unsafe.Pointer
		j [0]interface{}
	}
	sc := &C{}
	type D struct {
		a [3]int
		b [3]uint
		c [3]uintptr
		d [3]float64
		e [2]complex128
		f [3]string
		g [2]chan int
		h [2]func(int) string
		i [2]unsafe.Pointer
		j [2]interface{}
	}
	sd := &D{*a1, *b1, *c1, *d1, *e1, *f1, *g1, *h1, *i1, *j1}
	type E struct {
		a []int
		b []uint
		c []uintptr
		d []float64
		e []complex128
		f []string
		g []chan int
		h []func(int) string
		i []unsafe.Pointer
		j []interface{}
	}
	se := &E{*as, *bs, *cs, *ds, *es, *fs, *gs, *hs, *is, *js}
	type F struct {
		a map[int]unsafe.Pointer
		b map[uint]int
		c map[uintptr]uint
		d map[float64]uintptr
		e map[complex128]float64
		f map[string]complex128
		g map[chan int]string
		h map[unsafe.Pointer]chan int
		i map[interface{}]func(int) string
	}
	sf0 := &F{map[int]unsafe.Pointer{}, map[uint]int{}, map[uintptr]uint{}, map[float64]uintptr{}, map[complex128]float64{}, map[string]complex128{}, map[chan int]string{}, map[unsafe.Pointer]chan int{}, map[interface{}]func(int) string{}}
	sf1 := &F{map[int]unsafe.Pointer{100: i}, map[uint]int{100: 101}, map[uintptr]uint{100: 101}, map[float64]uintptr{10.123: 101}, map[complex128]float64{100.1 + 200.2i: 101.123}, map[string]complex128{"A bc": 100.1 + 200.2i}, map[chan int]string{g: "A bc"}, map[unsafe.Pointer]chan int{i: g}, map[interface{}]func(int) string{j: h}}
	type G struct {
		a      A
		b      B
		c      C
		d      D
		e      E
		f0, f1 F
	}
	sg := &G{*sa, *sb, *sc, *sd, *se, *sf0, *sf1}
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
	eq("100.23", d)
	eq("(100.23+300.45i)", complex(float32(100.23), float32(300.45)))
	eq("(100.23+300.45i)", e)
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
	if true {
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
		// pointer of array
		ep("<nil> %[2]p <nil> %[4]p %[2]p", (*[0]int)(nil), a0, (**[0]int)(nil), &a0, unsafe.Pointer(a0))
		ep("<nil> %[2]p <nil> %[4]p %[2]p", (*[3]int)(nil), a1, (**[3]int)(nil), &a1, unsafe.Pointer(a1))
		ep("<nil> %[2]p <nil> %[4]p %[2]p", (*[]int)(nil), as, (**[]int)(nil), &as, unsafe.Pointer(as))
		ep("<nil> %[2]p <nil> %[4]p %[2]p", (*[0]uint)(nil), b0, (**[0]uint)(nil), &b0, unsafe.Pointer(b0))
		ep("<nil> %[2]p <nil> %[4]p %[2]p", (*[3]uint)(nil), b1, (**[3]uint)(nil), &b1, unsafe.Pointer(b1))
		ep("<nil> %[2]p <nil> %[4]p %[2]p", (*[]uint)(nil), bs, (**[]uint)(nil), &bs, unsafe.Pointer(bs))
		ep("<nil> %[2]p <nil> %[4]p %[2]p", (*[0]uintptr)(nil), c0, (**[0]uintptr)(nil), &c0, unsafe.Pointer(c0))
		ep("<nil> %[2]p <nil> %[4]p %[2]p", (*[3]uintptr)(nil), c1, (**[3]uintptr)(nil), &c1, unsafe.Pointer(c1))
		ep("<nil> %[2]p <nil> %[4]p %[2]p", (*[]uintptr)(nil), cs, (**[]uintptr)(nil), &cs, unsafe.Pointer(cs))
		ep("<nil> %[2]p <nil> %[4]p %[2]p", (*[0]float64)(nil), d0, (**[0]float64)(nil), &d0, unsafe.Pointer(d0))
		ep("<nil> %[2]p <nil> %[4]p %[2]p", (*[3]float64)(nil), d1, (**[3]float64)(nil), &d1, unsafe.Pointer(d1))
		ep("<nil> %[2]p <nil> %[4]p %[2]p", (*[]float64)(nil), ds, (**[]float64)(nil), &ds, unsafe.Pointer(ds))
		ep("<nil> %[2]p <nil> %[4]p %[2]p", (*[0]complex128)(nil), e0, (**[0]complex128)(nil), &e0, unsafe.Pointer(e0))
		ep("<nil> %[2]p <nil> %[4]p %[2]p", (*[3]complex128)(nil), e1, (**[3]complex128)(nil), &e1, unsafe.Pointer(e1))
		ep("<nil> %[2]p <nil> %[4]p %[2]p", (*[]complex128)(nil), es, (**[]complex128)(nil), &es, unsafe.Pointer(es))
		ep("<nil> %[2]p <nil> %[4]p %[2]p", (*[0]string)(nil), f0, (**[0]string)(nil), &f0, unsafe.Pointer(f0))
		ep("<nil> %[2]p <nil> %[4]p %[2]p", (*[3]string)(nil), f1, (**[3]string)(nil), &f1, unsafe.Pointer(f1))
		ep("<nil> %[2]p <nil> %[4]p %[2]p", (*[]string)(nil), fs, (**[]string)(nil), &fs, unsafe.Pointer(fs))
		ep("<nil> %[2]p <nil> %[4]p %[2]p", (*[0]chan int)(nil), g0, (**[0]chan int)(nil), &g0, unsafe.Pointer(g0))
		ep("<nil> %[2]p <nil> %[4]p %[2]p", (*[3]chan int)(nil), g1, (**[3]chan int)(nil), &g1, unsafe.Pointer(g1))
		ep("<nil> %[2]p <nil> %[4]p %[2]p", (*[]chan int)(nil), gs, (**[]chan int)(nil), &gs, unsafe.Pointer(gs))
		ep("<nil> %[2]p <nil> %[4]p %[2]p", (*[0]func(int) string)(nil), h0, (**[0]func(int) string)(nil), &h0, unsafe.Pointer(h0))
		ep("<nil> %[2]p <nil> %[4]p %[2]p", (*[3]func(int) string)(nil), h1, (**[3]func(int) string)(nil), &h1, unsafe.Pointer(h1))
		ep("<nil> %[2]p <nil> %[4]p %[2]p", (*[]func(int) string)(nil), hs, (**[]func(int) string)(nil), &hs, unsafe.Pointer(hs))
		ep("<nil> %[2]p <nil> %[4]p %[2]p", (*[0]unsafe.Pointer)(nil), i0, (**[0]unsafe.Pointer)(nil), &i0, unsafe.Pointer(i0))
		ep("<nil> %[2]p <nil> %[4]p %[2]p", (*[3]unsafe.Pointer)(nil), i1, (**[3]unsafe.Pointer)(nil), &i1, unsafe.Pointer(i1))
		ep("<nil> %[2]p <nil> %[4]p %[2]p", (*[]unsafe.Pointer)(nil), is, (**[]unsafe.Pointer)(nil), &is, unsafe.Pointer(is))
		ep("<nil> %[2]p <nil> %[4]p %[2]p", (*[0]interface{})(nil), j0, (**[0]interface{})(nil), &j0, unsafe.Pointer(j0))
		ep("<nil> %[2]p <nil> %[4]p %[2]p", (*[3]interface{})(nil), j1, (**[3]interface{})(nil), &j1, unsafe.Pointer(j1))
		ep("<nil> %[2]p <nil> %[4]p %[2]p", (*[]interface{})(nil), js, (**[]interface{})(nil), &js, unsafe.Pointer(js))
		// pointer of map
		ep("<nil> %[2]p %[3]p", (*map[int]unsafe.Pointer)(nil), &map[int]unsafe.Pointer{}, &map[int]unsafe.Pointer{100: i})
		ep("<nil> %[2]p %[3]p", (*map[uint]int)(nil), &map[uint]int{}, &map[uint]int{100: 101})
		ep("<nil> %[2]p %[3]p", (*map[uintptr]uint)(nil), &map[uintptr]uint{}, &map[uintptr]uint{100: 101})
		ep("<nil> %[2]p %[3]p", (*map[float64]uintptr)(nil), &map[float64]uintptr{}, &map[float64]uintptr{100.123: 101})
		ep("<nil> %[2]p %[3]p", (*map[complex128]float64)(nil), &map[complex128]float64{}, &map[complex128]float64{100.1 + 200.2i: 101.123})
		ep("<nil> %[2]p %[3]p", (*map[string]complex128)(nil), &map[string]complex128{}, &map[string]complex128{"A bc": 100.1 + 200.2i})
		ep("<nil> %[2]p %[3]p", (*map[chan int]string)(nil), &map[chan int]string{}, &map[chan int]string{g: "A bc"})
		ep("<nil> %[2]p %[3]p", (*map[unsafe.Pointer]chan int)(nil), &map[unsafe.Pointer]chan int{}, &map[unsafe.Pointer]chan int{i: g})
		ep("<nil> %[2]p %[3]p", (*map[interface{}]func(int) string)(nil), &map[interface{}]func(int) string{}, &map[interface{}]func(int) string{j: h})
		// pointer of struct
		ep("<nil> %[2]p <nil> %[4]p %[2]p", (*A)(nil), sa, (**A)(nil), &sa, unsafe.Pointer(sa))
		ep("<nil> %[2]p <nil> %[4]p %[2]p", (*B)(nil), sb, (**B)(nil), &sb, unsafe.Pointer(sb))
		ep("<nil> %[2]p <nil> %[4]p %[2]p", (*C)(nil), sc, (**C)(nil), &sc, unsafe.Pointer(sc))
		ep("<nil> %[2]p <nil> %[4]p %[2]p", (*D)(nil), sd, (**D)(nil), &sd, unsafe.Pointer(sd))
		ep("<nil> %[2]p <nil> %[4]p %[2]p", (*E)(nil), se, (**E)(nil), &se, unsafe.Pointer(se))
		ep("<nil> %[2]p <nil> %[4]p %[2]p", (*F)(nil), sf0, (**F)(nil), &sf0, unsafe.Pointer(sf0))
		ep("<nil> %[2]p <nil> %[4]p %[2]p", (*F)(nil), sf1, (**F)(nil), &sf1, unsafe.Pointer(sf1))
		ep("<nil> %[2]p <nil> %[4]p %[2]p", (*G)(nil), sg, (**G)(nil), &sg, unsafe.Pointer(sg))
	}
	// array & slice
	if true {
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
		// array of array
		eq("[] [[]] [] [[0 0 0] [101 102 103]] [] [<nil> [101 102 103]]", [...][0]int{}, [...][0]int{*a0}, [...][3]int{}, [...][3]int{[3]int{}, *a1}, [...][]int{}, [...][]int{nil, *as})
		eq("[] [[]] [] [[0 0 0] [101 102 103]] [] [<nil> [101 102 103]]", [...][0]uint{}, [...][0]uint{*b0}, [...][3]uint{}, [...][3]uint{[3]uint{}, *b1}, [...][]uint{}, [...][]uint{nil, *bs})
		eq("[] [[]] [] [[0x0 0x0 0x0] [0x65 0x66 0x67]] [] [<nil> [0x65 0x66 0x67]]", [...][0]uintptr{}, [...][0]uintptr{*c0}, [...][3]uintptr{}, [...][3]uintptr{[3]uintptr{}, *c1}, [...][]uintptr{}, [...][]uintptr{nil, *cs})
		eq("[] [[]] [] [[0 0 0] [101.123 102.234 103.345]] [] [<nil> [101.123 102.234 103.345]]", [...][0]float64{}, [...][0]float64{*d0}, [...][3]float64{}, [...][3]float64{[3]float64{}, *d1}, [...][]float64{}, [...][]float64{nil, *ds})
		eq("[] [[]] [] [[(0+0i) (0+0i)] [(101.1+102.2i) (103.3+104.4i)]] [] [<nil> [(101.1+102.2i) (103.3+104.4i)]]", [...][0]complex128{}, [...][0]complex128{*e0}, [...][2]complex128{}, [...][2]complex128{[2]complex128{}, *e1}, [...][]complex128{}, [...][]complex128{nil, *es})
		eq(`[] [[]] [] [["" "" ""] ["A bc" "De f" "Gh"]] [] [<nil> ["A bc" "De f" "Gh"]]`, [...][0]string{}, [...][0]string{*f0}, [...][3]string{}, [...][3]string{[3]string{}, *f1}, [...][]string{}, [...][]string{nil, *fs})
		ep("[] [[]] [] [[<nil> <nil>] [<nil> %[7]p]] [] [<nil> [<nil> %[7]p]] %[7]p", [...][0]chan int{}, [...][0]chan int{*g0}, [...][2]chan int{}, [...][2]chan int{[2]chan int{}, *g1}, [...][]chan int{}, [...][]chan int{nil, *gs}, g)
		ep("[] [[]] [] [[<nil> <nil>] [<nil> %[7]p]] [] [<nil> [<nil> %[7]p]] %[7]p", [...][0]func(int) string{}, [...][0]func(int) string{*h0}, [...][2]func(int) string{}, [...][2]func(int) string{[2]func(int) string{}, *h1}, [...][]func(int) string{}, [...][]func(int) string{nil, *hs}, h)
		ep("[] [[]] [] [[<nil> <nil>] [<nil> %[7]p]] [] [<nil> [<nil> %[7]p]] %[7]p", [...][0]unsafe.Pointer{}, [...][0]unsafe.Pointer{*i0}, [...][2]unsafe.Pointer{}, [...][2]unsafe.Pointer{[2]unsafe.Pointer{}, *i1}, [...][]unsafe.Pointer{}, [...][]unsafe.Pointer{nil, *is}, i)
		eq("[] [[]] [] [[<nil> <nil>] [<nil> 2017]] [] [<nil> [<nil> 2017]]", [...][0]interface{}{}, [...][0]interface{}{*j0}, [...][2]interface{}{}, [...][2]interface{}{[2]interface{}{}, *j1}, [...][]interface{}{}, [...][]interface{}{nil, *js})
		// array of map
		ep("[] [<nil> map[] [100:%[3]p]] %[3]p", [...]map[int]unsafe.Pointer{}, [...]map[int]unsafe.Pointer{nil, map[int]unsafe.Pointer{}, map[int]unsafe.Pointer{100: i}}, i)
		ep("[] [<nil> map[] [100:%[3]p]] %[3]p <nil>", []map[int]unsafe.Pointer{}, []map[int]unsafe.Pointer{nil, map[int]unsafe.Pointer{}, map[int]unsafe.Pointer{100: i}}, i, []map[int]unsafe.Pointer(nil))
		eq("[] [<nil> map[] [100:101]]", [...]map[uint]int{}, [...]map[uint]int{nil, map[uint]int{}, map[uint]int{100: 101}})
		eq("[] [<nil> map[] [100:101]] <nil>", []map[uint]int{}, []map[uint]int{nil, map[uint]int{}, map[uint]int{100: 101}}, []map[uint]int(nil))
		eq("[] [<nil> map[] [0x64:101]]", [...]map[uintptr]uint{}, [...]map[uintptr]uint{nil, map[uintptr]uint{}, map[uintptr]uint{100: 101}})
		eq("[] [<nil> map[] [0x64:101]] <nil>", []map[uintptr]uint{}, []map[uintptr]uint{nil, map[uintptr]uint{}, map[uintptr]uint{100: 101}}, []map[uintptr]uint(nil))
		eq("[] [<nil> map[] [100.123:0x65]]", [...]map[float64]uintptr{}, [...]map[float64]uintptr{nil, map[float64]uintptr{}, map[float64]uintptr{100.123: 101}})
		eq("[] [<nil> map[] [(100.1+200.2i):101.123]]", [...]map[complex128]float64{}, [...]map[complex128]float64{nil, map[complex128]float64{}, map[complex128]float64{100.1 + 200.2i: 101.123}})
		eq(`[] [<nil> map[] ["A bc":(100.1+200.2i)]]`, [...]map[string]complex128{}, [...]map[string]complex128{nil, map[string]complex128{}, map[string]complex128{"A bc": 100.1 + 200.2i}})
		ep(`[] [<nil> map[] [%[3]p:"A bc"]] %[3]p`, [...]map[chan int]string{}, [...]map[chan int]string{nil, map[chan int]string{}, map[chan int]string{g: "A bc"}}, g)
		ep("[] [<nil> map[] [%[3]p:%[4]p]] %[3]p %[4]p", [...]map[unsafe.Pointer]chan int{}, [...]map[unsafe.Pointer]chan int{nil, map[unsafe.Pointer]chan int{}, map[unsafe.Pointer]chan int{i: g}}, i, g)
		ep("[] [<nil> map[] [2017:%[3]p]] %[3]p", [...]map[interface{}]func(int) string{}, [...]map[interface{}]func(int) string{nil, map[interface{}]func(int) string{}, map[interface{}]func(int) string{j: h}}, h)
		// array of struct
		ep(`[] [{a:0 b:0 c:0x0 d:0 e:(0+0i) f:"" g:<nil> h:<nil> i:<nil> j:<nil>} {a:100 b:100 c:0x64 d:100.23 e:(100.23+300.45i) f:"A bc" g:%[3]p h:%[4]p i:%[5]p j:2017}] %[3]p %[4]p %[5]p`, [...]A{}, [...]A{A{}, *sa}, g, h, i)
		ep(`[] [{a:<nil> b:<nil> c:<nil> d:<nil> e:<nil> f:<nil> g:<nil> h:<nil> i:<nil> j:<nil>} {a:%[3]p b:%[4]p c:%[5]p d:%[6]p e:%[7]p f:%[8]p g:%[9]p h:%[10]p i:%[11]p j:%[12]p}] %[3]p %[4]p %[5]p %[6]p %[7]p %[8]p %[9]p %[10]p %[11]p %[12]p`, [...]B{}, [...]B{B{}, *sb}, pa, pb, pc, pd, pe, pf, pg, ph, pi, pj)
		eq("[] [{a:[] b:[] c:[] d:[] e:[] f:[] g:[] h:[] i:[] j:[]} {a:[] b:[] c:[] d:[] e:[] f:[] g:[] h:[] i:[] j:[]}]", [...]C{}, [...]C{C{}, *sc})
		ep(`[] [{a:[0 0 0] b:[0 0 0] c:[0x0 0x0 0x0] d:[0 0 0] e:[(0+0i) (0+0i)] f:["" "" ""] g:[<nil> <nil>] h:[<nil> <nil>] i:[<nil> <nil>] j:[<nil> <nil>]} {a:[101 102 103] b:[101 102 103] c:[0x65 0x66 0x67] d:[101.123 102.234 103.345] e:[(101.1+102.2i) (103.3+104.4i)] f:["A bc" "De f" "Gh"] g:[<nil> %[3]p] h:[<nil> %[4]p] i:[<nil> %[5]p] j:[<nil> 2017]}] %[3]p %[4]p %[5]p`, [...]D{}, [...]D{D{}, *sd}, g, h, i)
		ep(`[] [{a:<nil> b:<nil> c:<nil> d:<nil> e:<nil> f:<nil> g:<nil> h:<nil> i:<nil> j:<nil>} {a:[101 102 103] b:[101 102 103] c:[0x65 0x66 0x67] d:[101.123 102.234 103.345] e:[(101.1+102.2i) (103.3+104.4i)] f:["A bc" "De f" "Gh"] g:[<nil> %[3]p] h:[<nil> %[4]p] i:[<nil> %[5]p] j:[<nil> 2017]}] %[3]p %[4]p %[5]p`, [...]E{}, [...]E{E{}, *se}, g, h, i)
		eq("[] [{a:<nil> b:<nil> c:<nil> d:<nil> e:<nil> f:<nil> g:<nil> h:<nil> i:<nil>} {a:map[] b:map[] c:map[] d:map[] e:map[] f:map[] g:map[] h:map[] i:map[]}]", [...]F{}, [...]F{F{}, *sf0})
		ep(`[] [{a:<nil> b:<nil> c:<nil> d:<nil> e:<nil> f:<nil> g:<nil> h:<nil> i:<nil>} {a:[100:%[3]p] b:[100:101] c:[0x64:101] d:[10.123:0x65] e:[(100.1+200.2i):101.123] f:["A bc":(100.1+200.2i)] g:[%[4]p:"A bc"] h:[%[3]p:%[4]p] i:[2017:%[5]p]}] %[3]p %[4]p %[5]p`, [...]F{}, [...]F{F{}, *sf1}, i, g, h)
		ep(`[{a:{a:0 b:0 c:0x0 d:0 e:(0+0i) f:"" g:<nil> h:<nil> i:<nil> j:<nil>} b:{a:<nil> b:<nil> c:<nil> d:<nil> e:<nil> f:<nil> g:<nil> h:<nil> i:<nil> j:<nil>} c:{a:[] b:[] c:[] d:[] e:[] f:[] g:[] h:[] i:[] j:[]} d:{a:[0 0 0] b:[0 0 0] c:[0x0 0x0 0x0] d:[0 0 0] e:[(0+0i) (0+0i)] f:["" "" ""] g:[<nil> <nil>] h:[<nil> <nil>] i:[<nil> <nil>] j:[<nil> <nil>]} e:{a:<nil> b:<nil> c:<nil> d:<nil> e:<nil> f:<nil> g:<nil> h:<nil> i:<nil> j:<nil>} f0:{a:<nil> b:<nil> c:<nil> d:<nil> e:<nil> f:<nil> g:<nil> h:<nil> i:<nil>} f1:{a:<nil> b:<nil> c:<nil> d:<nil> e:<nil> f:<nil> g:<nil> h:<nil> i:<nil>}} {a:{a:100 b:100 c:0x64 d:100.23 e:(100.23+300.45i) f:"A bc" g:%[2]p h:%[3]p i:%[4]p j:2017} b:{a:%[4]p b:%[5]p c:%[6]p d:%[7]p e:%[8]p f:%[9]p g:%[10]p h:%[11]p i:%[12]p j:%[13]p} c:{a:[] b:[] c:[] d:[] e:[] f:[] g:[] h:[] i:[] j:[]} d:{a:[101 102 103] b:[101 102 103] c:[0x65 0x66 0x67] d:[101.123 102.234 103.345] e:[(101.1+102.2i) (103.3+104.4i)] f:["A bc" "De f" "Gh"] g:[<nil> %[2]p] h:[<nil> %[3]p] i:[<nil> %[4]p] j:[<nil> 2017]} e:{a:[101 102 103] b:[101 102 103] c:[0x65 0x66 0x67] d:[101.123 102.234 103.345] e:[(101.1+102.2i) (103.3+104.4i)] f:["A bc" "De f" "Gh"] g:[<nil> %[2]p] h:[<nil> %[3]p] i:[<nil> %[4]p] j:[<nil> 2017]} f0:{a:map[] b:map[] c:map[] d:map[] e:map[] f:map[] g:map[] h:map[] i:map[]} f1:{a:[100:%[4]p] b:[100:101] c:[0x64:101] d:[10.123:0x65] e:[(100.1+200.2i):101.123] f:["A bc":(100.1+200.2i)] g:[%[2]p:"A bc"] h:[%[4]p:%[2]p] i:[2017:%[3]p]}}] %[2]p %[3]p %[4]p %[5]p %[6]p %[7]p %[8]p %[9]p %[10]p %[11]p %[12]p %[13]p []`, [...]G{G{}, *sg}, g, h, i, pb, pc, pd, pe, pf, pg, ph, pi, pj, [...]G{})
		// array of pointer
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*int{}, [...]*int{nil, pa}, []*int{}, []*int{nil, pa}, []*int(nil), pa)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*uint{}, [...]*uint{nil, pb}, []*uint{}, []*uint{nil, pb}, []*uint(nil), pb)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*uintptr{}, [...]*uintptr{nil, pc}, []*uintptr{}, []*uintptr{nil, pc}, []*uintptr(nil), pc)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*float64{}, [...]*float64{nil, pd}, []*float64{}, []*float64{nil, pd}, []*float64(nil), pd)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*complex128{}, [...]*complex128{nil, pe}, []*complex128{}, []*complex128{nil, pe}, []*complex128(nil), pe)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*string{}, [...]*string{nil, pf}, []*string{}, []*string{nil, pf}, []*string(nil), pf)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*chan int{}, [...]*chan int{nil, pg}, []*chan int{}, []*chan int{nil, pg}, []*chan int(nil), pg)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*func(int) string{}, [...]*func(int) string{nil, ph}, []*func(int) string{}, []*func(int) string{nil, ph}, []*func(int) string(nil), ph)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*unsafe.Pointer{}, [...]*unsafe.Pointer{nil, pi}, []*unsafe.Pointer{}, []*unsafe.Pointer{nil, pi}, []*unsafe.Pointer(nil), pi)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*interface{}{}, [...]*interface{}{nil, pj}, []*interface{}{}, []*interface{}{nil, pj}, []*interface{}(nil), pj)
		// array of pointer of array
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*[0]int{}, [...]*[0]int{nil, a0}, []*[0]int{}, []*[0]int{nil, a0}, []*[0]int(nil), a0)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*[3]int{}, [...]*[3]int{nil, a1}, []*[3]int{}, []*[3]int{nil, a1}, []*[3]int(nil), a1)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*[]int{}, [...]*[]int{nil, as}, []*[]int{}, []*[]int{nil, as}, []*[]int(nil), as)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*[0]uint{}, [...]*[0]uint{nil, b0}, []*[0]uint{}, []*[0]uint{nil, b0}, []*[0]uint(nil), b0)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*[3]uint{}, [...]*[3]uint{nil, b1}, []*[3]uint{}, []*[3]uint{nil, b1}, []*[3]uint(nil), b1)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*[]uint{}, [...]*[]uint{nil, bs}, []*[]uint{}, []*[]uint{nil, bs}, []*[]uint(nil), bs)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*[0]uintptr{}, [...]*[0]uintptr{nil, c0}, []*[0]uintptr{}, []*[0]uintptr{nil, c0}, []*[0]uintptr(nil), c0)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*[3]uintptr{}, [...]*[3]uintptr{nil, c1}, []*[3]uintptr{}, []*[3]uintptr{nil, c1}, []*[3]uintptr(nil), c1)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*[]uintptr{}, [...]*[]uintptr{nil, cs}, []*[]uintptr{}, []*[]uintptr{nil, cs}, []*[]uintptr(nil), cs)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*[0]float64{}, [...]*[0]float64{nil, d0}, []*[0]float64{}, []*[0]float64{nil, d0}, []*[0]float64(nil), d0)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*[3]float64{}, [...]*[3]float64{nil, d1}, []*[3]float64{}, []*[3]float64{nil, d1}, []*[3]float64(nil), d1)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*[]float64{}, [...]*[]float64{nil, ds}, []*[]float64{}, []*[]float64{nil, ds}, []*[]float64(nil), ds)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*[0]complex128{}, [...]*[0]complex128{nil, e0}, []*[0]complex128{}, []*[0]complex128{nil, e0}, []*[0]complex128(nil), e0)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*[3]complex128{}, [...]*[2]complex128{nil, e1}, []*[2]complex128{}, []*[2]complex128{nil, e1}, []*[2]complex128(nil), e1)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*[]complex128{}, [...]*[]complex128{nil, es}, []*[]complex128{}, []*[]complex128{nil, es}, []*[]complex128(nil), es)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*[0]string{}, [...]*[0]string{nil, f0}, []*[0]string{}, []*[0]string{nil, f0}, []*[0]string(nil), f0)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*[3]string{}, [...]*[3]string{nil, f1}, []*[3]string{}, []*[3]string{nil, f1}, []*[3]string(nil), f1)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*[]string{}, [...]*[]string{nil, fs}, []*[]string{}, []*[]string{nil, fs}, []*[]string(nil), fs)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*[0]chan int{}, [...]*[0]chan int{nil, g0}, []*[0]chan int{}, []*[0]chan int{nil, g0}, []*[0]chan int(nil), g0)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*[2]chan int{}, [...]*[2]chan int{nil, g1}, []*[2]chan int{}, []*[2]chan int{nil, g1}, []*[2]chan int(nil), g1)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*[]chan int{}, [...]*[]chan int{nil, gs}, []*[]chan int{}, []*[]chan int{nil, gs}, []*[]chan int(nil), gs)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*[0]func(int) string{}, [...]*[0]func(int) string{nil, h0}, []*[0]func(int) string{}, []*[0]func(int) string{nil, h0}, []*[0]func(int) string(nil), h0)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*[2]func(int) string{}, [...]*[2]func(int) string{nil, h1}, []*[2]func(int) string{}, []*[2]func(int) string{nil, h1}, []*[2]func(int) string(nil), h1)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*[]func(int) string{}, [...]*[]func(int) string{nil, hs}, []*[]func(int) string{}, []*[]func(int) string{nil, hs}, []*[]func(int) string(nil), hs)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*[0]unsafe.Pointer{}, [...]*[0]unsafe.Pointer{nil, i0}, []*[0]unsafe.Pointer{}, []*[0]unsafe.Pointer{nil, i0}, []*[0]unsafe.Pointer(nil), i0)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*[2]unsafe.Pointer{}, [...]*[2]unsafe.Pointer{nil, i1}, []*[2]unsafe.Pointer{}, []*[2]unsafe.Pointer{nil, i1}, []*[2]unsafe.Pointer(nil), i1)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*[]unsafe.Pointer{}, [...]*[]unsafe.Pointer{nil, is}, []*[]unsafe.Pointer{}, []*[]unsafe.Pointer{nil, is}, []*[]unsafe.Pointer(nil), is)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*[0]interface{}{}, [...]*[0]interface{}{nil, j0}, []*[0]interface{}{}, []*[0]interface{}{nil, j0}, []*[0]interface{}(nil), j0)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*[2]interface{}{}, [...]*[2]interface{}{nil, j1}, []*[2]interface{}{}, []*[2]interface{}{nil, j1}, []*[2]interface{}(nil), j1)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*[]interface{}{}, [...]*[]interface{}{nil, js}, []*[]interface{}{}, []*[]interface{}{nil, js}, []*[]interface{}(nil), js)
		// array of pointer of struct
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*A{}, [...]*A{nil, sa}, []*A{}, []*A{nil, sa}, []*A(nil), sa)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*B{}, [...]*B{nil, sb}, []*B{}, []*B{nil, sb}, []*B(nil), sb)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*C{}, [...]*C{nil, sc}, []*C{}, []*C{nil, sc}, []*C(nil), sc)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*D{}, [...]*D{nil, sd}, []*D{}, []*D{nil, sd}, []*D(nil), sd)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*E{}, [...]*E{nil, se}, []*E{}, []*E{nil, se}, []*E(nil), se)
		ep("[] [<nil> %[6]p %[7]p] [] [<nil> %[6]p %[7]p] <nil> %[6]p %[7]p", [...]*F{}, [...]*F{nil, sf0, sf1}, []*F{}, []*F{nil, sf0, sf1}, []*F(nil), sf0, sf1)
		ep("[] [<nil> %[6]p] [] [<nil> %[6]p] <nil> %[6]p", [...]*G{}, [...]*G{nil, sg}, []*G{}, []*G{nil, sg}, []*G(nil), sg)
	}
	// map
	ep("<nil> map[] [100:%[4]p] %[4]p", map[int]unsafe.Pointer(nil), map[int]unsafe.Pointer{}, map[int]unsafe.Pointer{100: i}, i)
	eq("<nil> map[] [100:101]", map[uint]int(nil), map[uint]int{}, map[uint]int{100: 101})
	eq("<nil> map[] [0x64:101]", map[uintptr]uint(nil), map[uintptr]uint{}, map[uintptr]uint{100: 101})
	eq("<nil> map[] [100.123:0x65]", map[float64]uintptr(nil), map[float64]uintptr{}, map[float64]uintptr{100.123: 101})
	eq("<nil> map[] [(100.1+200.2i):101.123]", map[complex128]float64(nil), map[complex128]float64{}, map[complex128]float64{100.1 + 200.2i: 101.123})
	eq(`<nil> map[] ["A bc":(100.1+200.2i)]`, map[string]complex128(nil), map[string]complex128{}, map[string]complex128{"A bc": 100.1 + 200.2i})
	ep(`<nil> map[] [%[4]p:"A bc"] %[4]p`, map[chan int]string(nil), map[chan int]string{}, map[chan int]string{g: "A bc"}, g)
	//map[func(int)string]...
	ep("<nil> map[] [%[4]p:%[5]p] %[4]p %[5]p", map[unsafe.Pointer]chan int(nil), map[unsafe.Pointer]chan int{}, map[unsafe.Pointer]chan int{i: g}, i, g)
	ep("<nil> map[] [2017:%[4]p] %[4]p", map[interface{}]func(int) string(nil), map[interface{}]func(int) string{}, map[interface{}]func(int) string{j: h}, h)
	//TODO: key as pointer (of type, array/slice, map, struct), array, map, struct
	// reflect.Value
	// struct
	if true {
		eq(`{a:0 b:0 c:0x0 d:0 e:(0+0i) f:"" g:<nil> h:<nil> i:<nil> j:<nil>}`, A{})
		ep(`{a:100 b:100 c:0x64 d:100.23 e:(100.23+300.45i) f:"A bc" g:%[2]p h:%[3]p i:%[4]p j:2017} %[2]p %[3]p %[4]p`, *sa, g, h, i)
		eq("{a:<nil> b:<nil> c:<nil> d:<nil> e:<nil> f:<nil> g:<nil> h:<nil> i:<nil> j:<nil>}", B{})
		ep("{a:%[2]p b:%[3]p c:%[4]p d:%[5]p e:%[6]p f:%[7]p g:%[8]p h:%[9]p i:%[10]p j:%[11]p} %[2]p %[3]p %[4]p %[5]p %[6]p %[7]p %[8]p %[9]p %[10]p %[11]p", *sb, pa, pb, pc, pd, pe, pf, pg, ph, pi, pj)
		eq("{a:[] b:[] c:[] d:[] e:[] f:[] g:[] h:[] i:[] j:[]}", C{})
		eq(`{a:[0 0 0] b:[0 0 0] c:[0x0 0x0 0x0] d:[0 0 0] e:[(0+0i) (0+0i)] f:["" "" ""] g:[<nil> <nil>] h:[<nil> <nil>] i:[<nil> <nil>] j:[<nil> <nil>]}`, D{})
		ep(`{a:[101 102 103] b:[101 102 103] c:[0x65 0x66 0x67] d:[101.123 102.234 103.345] e:[(101.1+102.2i) (103.3+104.4i)] f:["A bc" "De f" "Gh"] g:[<nil> %[2]p] h:[<nil> %[3]p] i:[<nil> %[4]p] j:[<nil> 2017]} %[2]p %[3]p %[4]p`, *sd, g, h, i)
		eq("{a:<nil> b:<nil> c:<nil> d:<nil> e:<nil> f:<nil> g:<nil> h:<nil> i:<nil> j:<nil>}", E{})
		ep(`{a:[101 102 103] b:[101 102 103] c:[0x65 0x66 0x67] d:[101.123 102.234 103.345] e:[(101.1+102.2i) (103.3+104.4i)] f:["A bc" "De f" "Gh"] g:[<nil> %[2]p] h:[<nil> %[3]p] i:[<nil> %[4]p] j:[<nil> 2017]} %[2]p %[3]p %[4]p`, *se, g, h, i)
		eq("{a:<nil> b:<nil> c:<nil> d:<nil> e:<nil> f:<nil> g:<nil> h:<nil> i:<nil>}", F{})
		eq("{a:map[] b:map[] c:map[] d:map[] e:map[] f:map[] g:map[] h:map[] i:map[]}", *sf0)
		ep(`{a:[100:%[2]p] b:[100:101] c:[0x64:101] d:[10.123:0x65] e:[(100.1+200.2i):101.123] f:["A bc":(100.1+200.2i)] g:[%[3]p:"A bc"] h:[%[2]p:%[3]p] i:[2017:%[4]p]} %[2]p %[3]p %[4]p`, *sf1, i, g, h)
		eq(`{a:{a:0 b:0 c:0x0 d:0 e:(0+0i) f:"" g:<nil> h:<nil> i:<nil> j:<nil>} b:{a:<nil> b:<nil> c:<nil> d:<nil> e:<nil> f:<nil> g:<nil> h:<nil> i:<nil> j:<nil>} c:{a:[] b:[] c:[] d:[] e:[] f:[] g:[] h:[] i:[] j:[]} d:{a:[0 0 0] b:[0 0 0] c:[0x0 0x0 0x0] d:[0 0 0] e:[(0+0i) (0+0i)] f:["" "" ""] g:[<nil> <nil>] h:[<nil> <nil>] i:[<nil> <nil>] j:[<nil> <nil>]} e:{a:<nil> b:<nil> c:<nil> d:<nil> e:<nil> f:<nil> g:<nil> h:<nil> i:<nil> j:<nil>} f0:{a:<nil> b:<nil> c:<nil> d:<nil> e:<nil> f:<nil> g:<nil> h:<nil> i:<nil>} f1:{a:<nil> b:<nil> c:<nil> d:<nil> e:<nil> f:<nil> g:<nil> h:<nil> i:<nil>}}`, G{})
		ep(`{a:{a:100 b:100 c:0x64 d:100.23 e:(100.23+300.45i) f:"A bc" g:%[2]p h:%[3]p i:%[4]p j:2017} b:{a:%[4]p b:%[5]p c:%[6]p d:%[7]p e:%[8]p f:%[9]p g:%[10]p h:%[11]p i:%[12]p j:%[13]p} c:{a:[] b:[] c:[] d:[] e:[] f:[] g:[] h:[] i:[] j:[]} d:{a:[101 102 103] b:[101 102 103] c:[0x65 0x66 0x67] d:[101.123 102.234 103.345] e:[(101.1+102.2i) (103.3+104.4i)] f:["A bc" "De f" "Gh"] g:[<nil> %[2]p] h:[<nil> %[3]p] i:[<nil> %[4]p] j:[<nil> 2017]} e:{a:[101 102 103] b:[101 102 103] c:[0x65 0x66 0x67] d:[101.123 102.234 103.345] e:[(101.1+102.2i) (103.3+104.4i)] f:["A bc" "De f" "Gh"] g:[<nil> %[2]p] h:[<nil> %[3]p] i:[<nil> %[4]p] j:[<nil> 2017]} f0:{a:map[] b:map[] c:map[] d:map[] e:map[] f:map[] g:map[] h:map[] i:map[]} f1:{a:[100:%[4]p] b:[100:101] c:[0x64:101] d:[10.123:0x65] e:[(100.1+200.2i):101.123] f:["A bc":(100.1+200.2i)] g:[%[2]p:"A bc"] h:[%[4]p:%[2]p] i:[2017:%[3]p]}} %[2]p %[3]p %[4]p %[5]p %[6]p %[7]p %[8]p %[9]p %[10]p %[11]p %[12]p %[13]p`, *sg, g, h, i, pb, pc, pd, pe, pf, pg, ph, pi, pj)
		//TODO: field as pointer of struct
	}
}

func TestWriteElem(t *testing.T) {
	eq := func(e string, xx ...interface{}) {
		if true { // value
			var v ValueDiffer
			for i, x := range xx {
				if i > 0 {
					v.b[0].Write(" ")
				}
				vx := reflect.ValueOf(x)
				v.writeElem(0, vx)
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
				v.writeElem(0, s.Field(0))
				v.writeElem(1, s.Field(1))
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
				v.writeElem(0, reflect.ValueOf(x))
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
				v.writeElem(0, s.Field(0))
				v.writeElem(1, s.Field(1))
			}
			Caller(1).Equal(t, e, v.String(0))
			Caller(1).Equal(t, e, v.String(1))
		}
	}
	// nil
	eq("<nil>", nil)
	// bool
	eq("true", true)
	eq("false", false)
	// number
	eq("100", int(100))
	eq("100", int8(100))
	eq("100", int16(100))
	eq("100", int32(100))
	eq("100", int64(100))
	eq("100", uint(100))
	eq("100", uint8(100))
	eq("100", uint16(100))
	eq("100", uint32(100))
	eq("100", uint64(100))
	eq("0x64", uintptr(100))
	eq("1.23", float32(1.23))
	eq("100.23", float64(100.23))
	eq("(100.23+300.45i)", complex(float32(100.23), float32(300.45)))
	eq("(100.23+300.45i)", complex(float64(100.23), float64(300.45)))
	// string
	eq(`"A bc"`, string("A bc"))
	// channel
	eq("<nil>", chan int(nil))
	ep("%p", make(chan int))
	ep("%p", make(<-chan int))
	ep("%p", make(chan<- int))
	// function
	eq("<nil>", (func(int) string)(nil))
	ep("%p", func(int) string { return "1" })
	// unsafe pointer
	eq("<nil>", unsafe.Pointer(nil))
	ep("%p", unsafe.Pointer(&[]int{}))
	// pointer
	if true {
		a := 100
		var b interface{} = &[0]int{}
		ep("<nil> %[2]p %[3]p", (*int)(nil), &a, &b)
		c := "A bc"
		ep("%[1]p", &c)
		ep("%[1]p", &[...]int{1, 2, 3})
		ep("%[1]p", &[]int{1, 2, 3})
		ep("%[1]p", &map[int]string{1: "abc"})
		ep("%[1]p", &struct {
			a int
			b string
		}{})
	}
	// array
	if true {
		eq("[]", [0]int{})
		// short
		eq("[1 2 3]", [...]int{1, 2, 3})
		a := 100
		ep("[2]*int{<nil>, %[2]p} %[2]p", [...]*int{nil, &a}, &a)
		eq("[2]map[int]string{<nil>, map[]}", [...]map[int]string{nil, map[int]string{}})
		eq(`[6]interface {}{
	<nil>, map[],
	[1 2 3],
	[97 98],
	[], <nil>
}`, [...]interface{}{nil, map[int]string{}, [...]int{1, 2, 3}, [...]byte{'a', 'b'}, [0]int{}, nil})
		// long
		eq("[11]int{0:0, 1:0, 2:0, 3:0, 4:0, 5:0, 6:0, 7:0, 8:0, 9:0, 10:0}", [11]int{})
		eq(`[11]interface {}{
	0:<nil>,
	1:map[],
	2:[1 2 3],
	3:[97 98],
	4:[],
	5:[1.2 0],
	6:<nil>,
	7:<nil>,
	8:<nil>,
	9:<nil>,
	10:<nil>
}`, [11]interface{}{nil, map[int]string{}, [...]int{1, 2, 3}, [...]byte{'a', 'b'}, [0]int{}, [2]float32{1.2}})
	}
	// slice
	if true {
		eq("<nil>", []int(nil))
		eq("[]", []int{})
		// short
		eq("[1 2 3]", []int{1, 2, 3})
		a := 100
		ep("[]*int{<nil>, %[2]p} %[2]p", []*int{nil, &a}, &a)
		eq("[]map[int]string{<nil>, map[]}", []map[int]string{nil, map[int]string{}})
		eq(`[]interface {}{
	<nil>, map[],
	[1 2 3],
	[97 98],
	[], <nil>
}`, []interface{}{nil, map[int]string{}, [...]int{1, 2, 3}, [...]byte{'a', 'b'}, [0]int{}, nil})
		// long
		var b [11]int
		eq("[]int{0:0, 1:0, 2:0, 3:0, 4:0, 5:0, 6:0, 7:0, 8:0, 9:0, 10:0}", b[:])
		c := [11]interface{}{nil, map[int]string{}, [...]int{1, 2, 3}, [...]byte{'a', 'b'}, [0]int{}, [2]float32{1.2}}
		eq(`[]interface {}{
	0:<nil>,
	1:map[],
	2:[1 2 3],
	3:[97 98],
	4:[],
	5:[1.2 0],
	6:<nil>,
	7:<nil>,
	8:<nil>,
	9:<nil>,
	10:<nil>
}`, c[:])
	}
	// map
	if true {
		eq("<nil>", map[int]string(nil))
		eq("map[]", map[int]string{})
		eq(`map[10:"A bc"]`, map[int]string{10: "A bc"})
		eq("map[interface {}]int{[]:30}", map[interface{}]int{[0]string{}: 30})
		eq(`map[interface {}]int{
	["A bc" "B cd"]:30
}`, map[interface{}]int{[...]string{"A bc", "B cd"}: 30})
	}
	// struct
}
