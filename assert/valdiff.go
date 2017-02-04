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

func (vd *ValueDiffer) WriteDiff(v1, v2 reflect.Value, tab int) {
	b1, b2 := vd.bufs()
	b1.Tab, b2.Tab = tab, tab
	if !v1.IsValid() {
		b1.Highlight(nil)
		vd.writeHTypeValue(1, v2)
	} else if !v2.IsValid() {
		vd.writeHTypeValue(0, v1)
		b2.Highlight(nil)
	} else if v1.Type() == v2.Type() {
		vd.writeTypeDiffValues(v1, v2)
	} else {
		vd.writeHTypeValue(0, v1)
		vd.writeHTypeValue(1, v2)
		//v1, v2 = vd.writeDiffTypesBeforeValue(v1, v2)
		//vd.writeValueAfterType(0, v1)
		//vd.writeValueAfterType(1, v2)
	}
}

func (vd *ValueDiffer) writeTypeDiffValues(v1, v2 reflect.Value) {
	b1, b2 := vd.bufs()
	switch v1.Kind() {
	case reflect.Chan:
		b1.Writef("%v(", v1.Type()).Highlight(v1).Write(")")
		b2.Writef("%v(", v2.Type()).Highlight(v2).Write(")")
	case reflect.Complex64, reflect.Complex128:
		vd.writeTypeDiffValuesComplex(v1, v2)
	case reflect.Func:
		vd.writeTypeDiffValuesFunc(v1, v2)
	case reflect.String:
		vd.writeTypeDiffValuesString(v1, v2)
	case reflect.Ptr:
		//TODO
	case reflect.Interface:
		//TODO
	case reflect.Array:
		//vd.writeTypeDiffValuesArray(v1, v2)
	case reflect.Slice:
		//TODO
	case reflect.Map:
		//vd.writeTypeDiffValuesMap(v1, v2)
	case reflect.Struct:
		//TODO
	default:
		b1.Highlightf("%#v", v1)
		b2.Highlightf("%#v", v2)
	}
}

func (vd *ValueDiffer) writeTypeDiffValuesComplex(v1, v2 reflect.Value) {
	b1, b2 := vd.bufs()
	c1, c2 := v1.Complex(), v2.Complex()
	b1.Write("(")
	b2.Write("(")
	if real(c1) == real(c2) {
		b1.Write(real(c1))
		b2.Write(real(c2))
	} else {
		b1.Highlight(real(c1))
		b2.Highlight(real(c2))
	}
	b1.Write("+")
	b2.Write("+")
	if imag(c1) == imag(c2) {
		b1.Write(imag(c1))
		b2.Write(imag(c2))
	} else {
		b1.Highlight(imag(c1))
		b2.Highlight(imag(c2))
	}
	b1.Write(")")
	b2.Write(")")
}

func (vd *ValueDiffer) writeTypeDiffValuesFunc(v1, v2 reflect.Value) {
	b1, b2 := vd.bufs()
	b1.Writef("(%v)(", v1.Type())
	b2.Writef("(%v)(", v2.Type())
	vd.writeDiffValuesFunc(v1, v2)
	b1.Writef(")")
	b2.Writef(")")
}

func (vd *ValueDiffer) writeDiffValuesFunc(v1, v2 reflect.Value) {
	b1, b2 := vd.bufs()
	p1, p2 := "nil", "nil"
	if !v1.IsNil() {
		p1 = fmt.Sprint(v1)
	}
	if !v2.IsNil() {
		p2 = fmt.Sprint(v2)
	}
	b1.Highlight(p1)
	b2.Highlight(p2)
	if p1 == p2 {
		vd.Attrs[CompFunc] = true
	}
}

func (vd *ValueDiffer) writeTypeDiffValuesString(v1, v2 reflect.Value) {
	b1, b2 := vd.bufs()
	s1, s2 := []rune(fmt.Sprintf("%#v", v1)), []rune(fmt.Sprintf("%#v", v2))
	for i := 0; i < len(s1) || i < len(s2); i++ {
		if i >= len(s1) {
			b2.Highlightf("%c", s2[i])
		} else if i >= len(s2) {
			b1.Highlightf("%c", s1[i])
		} else if s1[i] == s2[i] {
			b1.Writef("%c", s1[i])
			b2.Writef("%c", s2[i])
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
