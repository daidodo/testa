package assert

import (
	"bytes"
	"fmt"
	"reflect"
)

const (
	NewLine = iota
	OmitSame
	CompFunc

	kAttrSize
)

type ValueDiffer struct {
	buf   [2]bytes.Buffer
	b     [2]FeatureBuf
	Attrs [kAttrSize]bool
}

func (vd *ValueDiffer) String(i int) string {
	vd.b[i].Finish()
	return vd.buf[i].String()
}

func (vd *ValueDiffer) WriteTypeValue(idx int, v reflect.Value, tab int) {
	vd.bufi(idx).Tab = tab
	vd.writeTypeValue(idx, v)
}

func (vd *ValueDiffer) WriteDiff(v1, v2 reflect.Value, tab int) {
	b1, b2 := vd.bufs()
	b1.Tab, b2.Tab = tab, tab
	vd.writeDiff(v1, v2)
}

func (vd *ValueDiffer) writeTypeValue(idx int, v reflect.Value) {
	v = vd.writeTypeBeforeValue(idx, v, false)
	vd.writeValueAfterType(idx, v)
}

func (vd *ValueDiffer) writeDiff(v1, v2 reflect.Value) {
	if !v1.IsValid() || !v2.IsValid() || v1.Type() != v2.Type() {
		vd.writeDiffTypeValues(v1, v2)
	} else {
		vd.writeTypeDiffValues(v1, v2)
	}
}

func (vd *ValueDiffer) writeTypeDiffValues(v1, v2 reflect.Value) {
	b1, b2 := vd.bufs()
	switch v1.Kind() {
	case reflect.Func:
		vd.writeDiffValuesFunc(v1, v2)
	case reflect.Interface:
		if v1.IsNil() {
			b1.Highlight(nil)
			vd.writeTypeValue(1, v2.Elem())
		} else if v2.IsNil() {
			vd.writeTypeValue(0, v1.Elem())
			b2.Highlight(nil)
		} else {
			vd.writeDiff(v1.Elem(), v2.Elem())
		}
	case reflect.Complex64, reflect.Complex128:
		c1, c2 := v1.Complex(), v2.Complex()
		hr, hi := real(c1) != real(c2), imag(c1) != imag(c2)
		b1.Normal("(").Write(hr, real(c1)).Normal("+").Write(hi, imag(c1)).Normal(")")
		b2.Normal("(").Write(hr, real(c2)).Normal("+").Write(hi, imag(c2)).Normal(")")
	case reflect.String:
		vd.writeDiffValuesString(v1, v2)
	case reflect.Array:
		vd.writeTypeDiffValuesArray(v1, v2, false)
	case reflect.Slice:
		vd.writeTypeDiffValuesArray(v1, v2, true)
	//case reflect.Map:
	//vd.writeTypeDiffValuesMap(v1, v2)
	//case reflect.Struct:
	//TODO
	default:
		b1.Highlightf("%#v", v1)
		b2.Highlightf("%#v", v2)
	}
}

func (vd *ValueDiffer) writeDiffValuesFunc(v1, v2 reflect.Value) {
	b1, b2 := vd.bufs()
	if v1.IsNil() && v2.IsNil() {
		b1.Normal("nil")
		b2.Normal("nil")
	} else {
		b1.Highlight(v1)
		b2.Highlight(v2)
		if v1.Pointer() == v2.Pointer() {
			vd.Attrs[CompFunc] = true
		}
	}
}

func (vd *ValueDiffer) writeDiffValuesString(v1, v2 reflect.Value) {
	b1, b2 := vd.bufs()
	s1, s2 := []rune(fmt.Sprintf("%#v", v1)), []rune(fmt.Sprintf("%#v", v2))
	s1, s2 = s1[1:len(s1)-1], s2[1:len(s2)-1] // skip front and end "
	b1.Normal(`"`)
	b2.Normal(`"`)
	vd.writeDiffPlainRunes(s1, s2)
	b1.Normal(`"`)
	b2.Normal(`"`)
}

func (vd *ValueDiffer) writeTypeDiffValuesArray(v1, v2 reflect.Value, slice bool) {
	b1, b2 := vd.bufs()
	tp1, id1, ml1 := attrElemArray(v1)
	tp2, id2, ml2 := attrElemArray(v2)
	tp, id := tp1 || tp2, id1 || id2
	if tp {
		vd.writeType(0, v1.Type(), false)
		vd.writeType(1, v2.Type(), false)
		b1.Normal("{")
		b2.Normal("{")
		defer b1.Normal("}")
		defer b2.Normal("}")
		if ml1 {
			b1.Tab++
			defer func() { b1.Tab-- }()
		}
		if ml2 {
			b2.Tab++
			defer func() { b2.Tab-- }()
		}
	} else {
		b1.Normal("[")
		b2.Normal("[")
		defer b1.Normal("]")
		defer b2.Normal("]")
	}
	if slice {
		vd.writeDiffValuesSliceC(v1, v2, tp, id, ml1, ml2)
	} else {
		vd.writeDiffValuesArrayC(v1, v2, tp, id, ml1, ml2)
	}
}

func (vd *ValueDiffer) writeDiffValuesSliceC(v1, v2 reflect.Value, tp, id, ml1, ml2 bool) {
	b1, b2 := vd.bufs()
	var p1, p2 bool
	for i := 0; i < v1.Len() || i < v2.Len(); i++ {
		g1, g2 := i < v1.Len(), i < v2.Len()
		t1, t2 := g1 && isNonTrivialElem(v1.Index(i)), g2 && isNonTrivialElem(v2.Index(i))
		t1, p1 = (t1 || p1 || (ml1 && (id || i == 0))), t1
		t2, p2 = (t2 || p2 || (ml2 && (id || i == 0))), t2
		if i > 0 {
			if tp {
				if g1 {
					b1.Normal(",")
				}
				if g2 {
					b2.Normal(",")
				}
			}
			if g1 && (!tp || !t1) {
				b1.Normal(" ")
			}
			if g2 && (!tp || !t2) {
				b2.Normal(" ")
			}
		}
		if t1 {
			b1.NL()
		}
		if t2 {
			b2.NL()
		}
		if id {
			if g1 {
				b1.Normal(i, ":")
			}
			if g2 {
				b2.Normal(i, ":")
			}
		}
		if g1 && g2 {
			if e1, e2 := v1.Index(i), v2.Index(i); reflect.DeepEqual(e1.Interface(), e2.Interface()) {
				vd.writeElem(0, e1)
				vd.writeElem(1, e2)
			} else {
				vd.writeTypeDiffValues(e1, e2)
			}
		} else if g1 {
			vd.writeElem(0, v1.Index(i)) //TODO: hl
		} else {
			vd.writeElem(1, v2.Index(i)) //TODO: hl
		}
	}
}

func (vd *ValueDiffer) writeDiffValuesArrayC(v1, v2 reflect.Value, tp, id, ml1, ml2 bool) {
	b1, b2 := vd.bufs()
	var p1, p2 bool
	for i := 0; i < v1.Len(); i++ {
		e1, e2 := v1.Index(i), v2.Index(i)
		t1, t2 := isNonTrivialElem(e1), isNonTrivialElem(e2)
		t1, p1 = (t1 || p1 || (ml1 && (id || i == 0))), t1
		t2, p2 = (t2 || p2 || (ml2 && (id || i == 0))), t2
		if i > 0 {
			if tp {
				b1.Normal(",")
				b2.Normal(",")
			}
			if !tp || !t1 {
				b1.Normal(" ")
			}
			if !tp || !t2 {
				b2.Normal(" ")
			}
		}
		if t1 {
			b1.NL()
		}
		if t2 {
			b2.NL()
		}
		if id {
			b1.Normal(i, ":")
			b2.Normal(i, ":")
		}
		if reflect.DeepEqual(e1.Interface(), e2.Interface()) {
			vd.writeElem(0, e1)
			vd.writeElem(1, e2)
		} else {
			vd.writeTypeDiffValues(e1, e2)
		}
	}
}

func (vd *ValueDiffer) writeDiffPlain(v1, v2 interface{}) {
	vd.writeDiffPlainRunes([]rune(fmt.Sprint(v1)), []rune(fmt.Sprint(v2)))
}

func (vd *ValueDiffer) writeDiffPlainf(format string, v1, v2 interface{}) {
	vd.writeDiffPlainRunes([]rune(fmt.Sprintf(format, v1)), []rune(fmt.Sprintf(format, v2)))
}

func (vd *ValueDiffer) writeDiffPlainString(v1, v2 string) {
	vd.writeDiffPlainRunes([]rune(v1), []rune(v2))
}

func (vd *ValueDiffer) writeDiffPlainRunes(s1, s2 []rune) {
	b1, b2 := vd.bufs()
	for i := 0; i < len(s1) || i < len(s2); i++ {
		if i >= len(s1) {
			b2.Highlightf("%c", s2[i])
		} else if i >= len(s2) {
			b1.Highlightf("%c", s1[i])
		} else if s1[i] == s2[i] {
			b1.Normalf("%c", s1[i])
			b2.Normalf("%c", s2[i])
		} else {
			b1.Highlightf("%c", s1[i])
			b2.Highlightf("%c", s2[i])
		}
	}
}

func (vd *ValueDiffer) bufi(i int) (b *FeatureBuf) {
	b = &vd.b[i]
	if b.w == nil {
		b.w = &vd.buf[i]
	}
	return
}

func (vd *ValueDiffer) bufs() (b1, b2 *FeatureBuf) {
	return vd.bufi(0), vd.bufi(1)
}
