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
	sf1 := &F{map[int]unsafe.Pointer{100: i}, map[uint]int{100: 101}, map[uintptr]uint{100: 101}, map[float64]uintptr{10.123: 101}, map[complex128]float64{100.1 + 200.2i: 101.123}, map[string]complex128{"A bc": 100.1 + 200.2}, map[chan int]string{g: "A bc"}, map[unsafe.Pointer]chan int{i: g}, map[interface{}]func(int) string{j: h}}
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
		ep("<nil> %[2]p %[3]p", (*map[string]complex128)(nil), &map[string]complex128{}, &map[string]complex128{"A bc": 100.1 + 200.2})
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
		//TODO
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
		ep(`[] [{a:0 b:0 c:0x0 d:0 e:(0+0i) f:"" g:<nil> h:<nil> i:<nil> j:<nil>} {a:100 b:100 c:0x64 d:1.23 e:(1.23+3.45i) f:"A bc" g:%[3]p h:%[4]p i:%[5]p j:2017}] %[3]p %[4]p %[5]p`, [...]A{}, [...]A{A{}, *sa}, g, h, i)
		ep(`[] [{a:<nil> b:<nil> c:<nil> d:<nil> e:<nil> f:<nil> g:<nil> h:<nil> i:<nil> j:<nil>} {a:%[3]p b:%[4]p c:%[5]p d:%[6]p e:%[7]p f:%[8]p g:%[9]p h:%[10]p i:%[11]p j:%[12]p}] %[3]p %[4]p %[5]p %[6]p %[7]p %[8]p %[9]p %[10]p %[11]p %[12]p`, [...]B{}, [...]B{B{}, *sb}, pa, pb, pc, pd, pe, pf, pg, ph, pi, pj)
		eq("[] [{a:[] b:[] c:[] d:[] e:[] f:[] g:[] h:[] i:[] j:[]} {a:[] b:[] c:[] d:[] e:[] f:[] g:[] h:[] i:[] j:[]}]", [...]C{}, [...]C{C{}, *sc})
		ep(`[] [{a:[0 0 0] b:[0 0 0] c:[0x0 0x0 0x0] d:[0 0 0] e:[(0+0i) (0+0i)] f:["" "" ""] g:[<nil> <nil>] h:[<nil> <nil>] i:[<nil> <nil>] j:[<nil> <nil>]} {a:[101 102 103] b:[101 102 103] c:[0x65 0x66 0x67] d:[101.123 102.234 103.345] e:[(101.1+102.2i) (103.3+104.4i)] f:["A bc" "De f" "Gh"] g:[<nil> %[3]p] h:[<nil> %[4]p] i:[<nil> %[5]p] j:[<nil> 2017]}] %[3]p %[4]p %[5]p`, [...]D{}, [...]D{D{}, *sd}, g, h, i)
		ep(`[] [{a:<nil> b:<nil> c:<nil> d:<nil> e:<nil> f:<nil> g:<nil> h:<nil> i:<nil> j:<nil>} {a:[101 102 103] b:[101 102 103] c:[0x65 0x66 0x67] d:[101.123 102.234 103.345] e:[(101.1+102.2i) (103.3+104.4i)] f:["A bc" "De f" "Gh"] g:[<nil> %[3]p] h:[<nil> %[4]p] i:[<nil> %[5]p] j:[<nil> 2017]}] %[3]p %[4]p %[5]p`, [...]E{}, [...]E{E{}, *se}, g, h, i)
		eq("[] [{a:<nil> b:<nil> c:<nil> d:<nil> e:<nil> f:<nil> g:<nil> h:<nil> i:<nil>} {a:map[] b:map[] c:map[] d:map[] e:map[] f:map[] g:map[] h:map[] i:map[]}]", [...]F{}, [...]F{F{}, *sf0})
		ep(`[] [{a:<nil> b:<nil> c:<nil> d:<nil> e:<nil> f:<nil> g:<nil> h:<nil> i:<nil>} {a:[100:%[3]p] b:[100:101] c:[0x64:101] d:[10.123:0x65] e:[(100.1+200.2i):101.123] f:["A bc":(300.3+0i)] g:[%[4]p:"A bc"] h:[%[3]p:%[4]p] i:[2017:%[5]p]}] %[3]p %[4]p %[5]p`, [...]F{}, [...]F{F{}, *sf1}, i, g, h)
		ep(`[{a:{a:0 b:0 c:0x0 d:0 e:(0+0i) f:"" g:<nil> h:<nil> i:<nil> j:<nil>} b:{a:<nil> b:<nil> c:<nil> d:<nil> e:<nil> f:<nil> g:<nil> h:<nil> i:<nil> j:<nil>} c:{a:[] b:[] c:[] d:[] e:[] f:[] g:[] h:[] i:[] j:[]} d:{a:[0 0 0] b:[0 0 0] c:[0x0 0x0 0x0] d:[0 0 0] e:[(0+0i) (0+0i)] f:["" "" ""] g:[<nil> <nil>] h:[<nil> <nil>] i:[<nil> <nil>] j:[<nil> <nil>]} e:{a:<nil> b:<nil> c:<nil> d:<nil> e:<nil> f:<nil> g:<nil> h:<nil> i:<nil> j:<nil>} f0:{a:<nil> b:<nil> c:<nil> d:<nil> e:<nil> f:<nil> g:<nil> h:<nil> i:<nil>} f1:{a:<nil> b:<nil> c:<nil> d:<nil> e:<nil> f:<nil> g:<nil> h:<nil> i:<nil>}} {a:{a:100 b:100 c:0x64 d:1.23 e:(1.23+3.45i) f:"A bc" g:%[2]p h:%[3]p i:%[4]p j:2017} b:{a:%[4]p b:%[5]p c:%[6]p d:%[7]p e:%[8]p f:%[9]p g:%[10]p h:%[11]p i:%[12]p j:%[13]p} c:{a:[] b:[] c:[] d:[] e:[] f:[] g:[] h:[] i:[] j:[]} d:{a:[101 102 103] b:[101 102 103] c:[0x65 0x66 0x67] d:[101.123 102.234 103.345] e:[(101.1+102.2i) (103.3+104.4i)] f:["A bc" "De f" "Gh"] g:[<nil> %[2]p] h:[<nil> %[3]p] i:[<nil> %[4]p] j:[<nil> 2017]} e:{a:[101 102 103] b:[101 102 103] c:[0x65 0x66 0x67] d:[101.123 102.234 103.345] e:[(101.1+102.2i) (103.3+104.4i)] f:["A bc" "De f" "Gh"] g:[<nil> %[2]p] h:[<nil> %[3]p] i:[<nil> %[4]p] j:[<nil> 2017]} f0:{a:map[] b:map[] c:map[] d:map[] e:map[] f:map[] g:map[] h:map[] i:map[]} f1:{a:[100:%[4]p] b:[100:101] c:[0x64:101] d:[10.123:0x65] e:[(100.1+200.2i):101.123] f:["A bc":(300.3+0i)] g:[%[2]p:"A bc"] h:[%[4]p:%[2]p] i:[2017:%[3]p]}}] %[2]p %[3]p %[4]p %[5]p %[6]p %[7]p %[8]p %[9]p %[10]p %[11]p %[12]p %[13]p []`, [...]G{G{}, *sg}, g, h, i, pb, pc, pd, pe, pf, pg, ph, pi, pj, [...]G{})
		//TODO
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
		//TODO
	}
	// map
	ep("<nil> map[] [100:%[4]p] %[4]p", map[int]unsafe.Pointer(nil), map[int]unsafe.Pointer{}, map[int]unsafe.Pointer{100: i}, i)
	eq("<nil> map[] [100:101]", map[uint]int(nil), map[uint]int{}, map[uint]int{100: 101})
	eq("<nil> map[] [0x64:101]", map[uintptr]uint(nil), map[uintptr]uint{}, map[uintptr]uint{100: 101})
	eq("<nil> map[] [100.123:0x65]", map[float64]uintptr(nil), map[float64]uintptr{}, map[float64]uintptr{100.123: 101})
	eq("<nil> map[] [(100.1+200.2i):101.123]", map[complex128]float64(nil), map[complex128]float64{}, map[complex128]float64{100.1 + 200.2i: 101.123})
	eq(`<nil> map[] ["A bc":(300.3+0i)]`, map[string]complex128(nil), map[string]complex128{}, map[string]complex128{"A bc": 100.1 + 200.2})
	ep(`<nil> map[] [%[4]p:"A bc"] %[4]p`, map[chan int]string(nil), map[chan int]string{}, map[chan int]string{g: "A bc"}, g)
	//map[func(int)string]...
	ep("<nil> map[] [%[4]p:%[5]p] %[4]p %[5]p", map[unsafe.Pointer]chan int(nil), map[unsafe.Pointer]chan int{}, map[unsafe.Pointer]chan int{i: g}, i, g)
	ep("<nil> map[] [2017:%[4]p] %[4]p", map[interface{}]func(int) string(nil), map[interface{}]func(int) string{}, map[interface{}]func(int) string{j: h}, h)
	// reflect.Value
	// struct
	if true {
		eq(`{a:0 b:0 c:0x0 d:0 e:(0+0i) f:"" g:<nil> h:<nil> i:<nil> j:<nil>}`, A{})
		ep(`{a:100 b:100 c:0x64 d:1.23 e:(1.23+3.45i) f:"A bc" g:%[2]p h:%[3]p i:%[4]p j:2017} %[2]p %[3]p %[4]p`, *sa, g, h, i)
		eq("{a:<nil> b:<nil> c:<nil> d:<nil> e:<nil> f:<nil> g:<nil> h:<nil> i:<nil> j:<nil>}", B{})
		ep("{a:%[2]p b:%[3]p c:%[4]p d:%[5]p e:%[6]p f:%[7]p g:%[8]p h:%[9]p i:%[10]p j:%[11]p} %[2]p %[3]p %[4]p %[5]p %[6]p %[7]p %[8]p %[9]p %[10]p %[11]p", *sb, pa, pb, pc, pd, pe, pf, pg, ph, pi, pj)
		eq("{a:[] b:[] c:[] d:[] e:[] f:[] g:[] h:[] i:[] j:[]}", C{})
		eq(`{a:[0 0 0] b:[0 0 0] c:[0x0 0x0 0x0] d:[0 0 0] e:[(0+0i) (0+0i)] f:["" "" ""] g:[<nil> <nil>] h:[<nil> <nil>] i:[<nil> <nil>] j:[<nil> <nil>]}`, D{})
		ep(`{a:[101 102 103] b:[101 102 103] c:[0x65 0x66 0x67] d:[101.123 102.234 103.345] e:[(101.1+102.2i) (103.3+104.4i)] f:["A bc" "De f" "Gh"] g:[<nil> %[2]p] h:[<nil> %[3]p] i:[<nil> %[4]p] j:[<nil> 2017]} %[2]p %[3]p %[4]p`, *sd, g, h, i)
		eq("{a:<nil> b:<nil> c:<nil> d:<nil> e:<nil> f:<nil> g:<nil> h:<nil> i:<nil> j:<nil>}", E{})
		ep(`{a:[101 102 103] b:[101 102 103] c:[0x65 0x66 0x67] d:[101.123 102.234 103.345] e:[(101.1+102.2i) (103.3+104.4i)] f:["A bc" "De f" "Gh"] g:[<nil> %[2]p] h:[<nil> %[3]p] i:[<nil> %[4]p] j:[<nil> 2017]} %[2]p %[3]p %[4]p`, *se, g, h, i)
		eq("{a:<nil> b:<nil> c:<nil> d:<nil> e:<nil> f:<nil> g:<nil> h:<nil> i:<nil>}", F{})
		eq("{a:map[] b:map[] c:map[] d:map[] e:map[] f:map[] g:map[] h:map[] i:map[]}", *sf0)
		ep(`{a:[100:%[2]p] b:[100:101] c:[0x64:101] d:[10.123:0x65] e:[(100.1+200.2i):101.123] f:["A bc":(300.3+0i)] g:[%[3]p:"A bc"] h:[%[2]p:%[3]p] i:[2017:%[4]p]} %[2]p %[3]p %[4]p`, *sf1, i, g, h)
		eq(`{a:{a:0 b:0 c:0x0 d:0 e:(0+0i) f:"" g:<nil> h:<nil> i:<nil> j:<nil>} b:{a:<nil> b:<nil> c:<nil> d:<nil> e:<nil> f:<nil> g:<nil> h:<nil> i:<nil> j:<nil>} c:{a:[] b:[] c:[] d:[] e:[] f:[] g:[] h:[] i:[] j:[]} d:{a:[0 0 0] b:[0 0 0] c:[0x0 0x0 0x0] d:[0 0 0] e:[(0+0i) (0+0i)] f:["" "" ""] g:[<nil> <nil>] h:[<nil> <nil>] i:[<nil> <nil>] j:[<nil> <nil>]} e:{a:<nil> b:<nil> c:<nil> d:<nil> e:<nil> f:<nil> g:<nil> h:<nil> i:<nil> j:<nil>} f0:{a:<nil> b:<nil> c:<nil> d:<nil> e:<nil> f:<nil> g:<nil> h:<nil> i:<nil>} f1:{a:<nil> b:<nil> c:<nil> d:<nil> e:<nil> f:<nil> g:<nil> h:<nil> i:<nil>}}`, G{})
		ep(`{a:{a:100 b:100 c:0x64 d:1.23 e:(1.23+3.45i) f:"A bc" g:%[2]p h:%[3]p i:%[4]p j:2017} b:{a:%[4]p b:%[5]p c:%[6]p d:%[7]p e:%[8]p f:%[9]p g:%[10]p h:%[11]p i:%[12]p j:%[13]p} c:{a:[] b:[] c:[] d:[] e:[] f:[] g:[] h:[] i:[] j:[]} d:{a:[101 102 103] b:[101 102 103] c:[0x65 0x66 0x67] d:[101.123 102.234 103.345] e:[(101.1+102.2i) (103.3+104.4i)] f:["A bc" "De f" "Gh"] g:[<nil> %[2]p] h:[<nil> %[3]p] i:[<nil> %[4]p] j:[<nil> 2017]} e:{a:[101 102 103] b:[101 102 103] c:[0x65 0x66 0x67] d:[101.123 102.234 103.345] e:[(101.1+102.2i) (103.3+104.4i)] f:["A bc" "De f" "Gh"] g:[<nil> %[2]p] h:[<nil> %[3]p] i:[<nil> %[4]p] j:[<nil> 2017]} f0:{a:map[] b:map[] c:map[] d:map[] e:map[] f:map[] g:map[] h:map[] i:map[]} f1:{a:[100:%[4]p] b:[100:101] c:[0x64:101] d:[10.123:0x65] e:[(100.1+200.2i):101.123] f:["A bc":(300.3+0i)] g:[%[2]p:"A bc"] h:[%[4]p:%[2]p] i:[2017:%[3]p]}} %[2]p %[3]p %[4]p %[5]p %[6]p %[7]p %[8]p %[9]p %[10]p %[11]p %[12]p %[13]p`, *sg, g, h, i, pb, pc, pd, pe, pf, pg, ph, pi, pj)
		//TODO
	}

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
		test(`[<nil> map[] [1:"abc"]]`, [...]map[int]string{nil, {}, {1: "abc"}})
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
		test(`[<nil> map[] [1:"abc"]]`, [...]interface{}{map[int]string(nil), map[int]string{}, map[int]string{1: "abc"}})
		test(`[{x:0 b:""} {a:1 y:["abc":(1.2+3.4i)]}]`, [...]interface{}{struct {
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
