package assert

import (
	"bytes"
	"fmt"
	"reflect"
)

const (
	NewLine = iota
	_
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
	vd.writeDiff(v1, v2)
}

func (vd *ValueDiffer) writeDiff(v1, v2 reflect.Value) {
	if !v1.IsValid() || !v2.IsValid() || v1.Type() != v2.Type() {
		vd.writeDiffTypeValues(v1, v2)
	} else {
		vd.writeTypeDiffValues(v1, v2)
	}
}

func (vd *ValueDiffer) writeDiffTypeValues(v1, v2 reflect.Value) {
	if !v1.IsValid() || !v2.IsValid() || v2.Kind() != v2.Kind() {
		v1 = vd.writeTypeBeforeValue(0, v1, true)
		v2 = vd.writeTypeBeforeValue(1, v2, true)
	} else {
		v1, v2 = vd.writeDiffTypesBeforeValue(v1, v2)
	}
	vd.writeValueAfterType(0, v1)
	vd.writeValueAfterType(1, v2)
}

func (vd *ValueDiffer) writeDiffTypesBeforeValue(v1, v2 reflect.Value) (r1, r2 reflect.Value) {
	b1, b2 := vd.bufs()
	r1, r2 = v1, v2
	switch v1.Kind() {
	case reflect.Interface:
		if vd.writeTypeBeforeInterfaceNil(0, v1, true) {
			r2 = vd.writeTypeBeforeValue(1, v2, true)
		} else if vd.writeTypeBeforeInterfaceNil(1, v2, true) {
			r1 = vd.writeTypeBeforeValue(0, v1, true)
		} else {
			r1, r2 = vd.writeDiffTypesBeforeValue(v1.Elem(), v2.Elem())
		}
	case reflect.Ptr:
		b1.Normal("(*")
		b2.Normal("(*")
		vd.writeDiffTypes(v1.Type().Elem(), v2.Type().Elem())
		b1.Normal(")")
		b2.Normal(")")
	case reflect.Func, reflect.Chan:
		b1.Normal("(")
		b2.Normal("(")
		vd.writeDiffTypes(v1.Type(), v2.Type())
		b1.Normal(")")
		b2.Normal(")")
	default:
		vd.writeDiffTypes(v1.Type(), v2.Type())
	}
	return
}

func (vd *ValueDiffer) writeDiffKinds(t1, t2 reflect.Type) {
	if t1 == nil || t2 == nil {
		panic("Should not come here!")
	}
	if t1 == t2 {
		vd.writeType(0, t1, false)
		vd.writeType(1, t2, false)
	} else if t1.Kind() == t2.Kind() {
		vd.writeDiffTypes(t1, t2)
	} else {
		vd.writeType(0, t1, true)
		vd.writeType(1, t2, true)
	}
}

func (vd *ValueDiffer) writeDiffTypes(t1, t2 reflect.Type) {
	b1, b2 := vd.bufs()
	switch t1.Kind() {
	case reflect.Ptr:
		b1.Normal("*")
		b2.Normal("*")
		vd.writeDiffKinds(t1.Elem(), t2.Elem())
	case reflect.Func:
		vd.writeDiffTypesFunc(t1, t2)
	case reflect.Chan:
		hd := t1.ChanDir() != t2.ChanDir()
		vd.writeTypeHeadChan(0, t1, false, hd)
		vd.writeTypeHeadChan(1, t2, false, hd)
		vd.writeDiffKinds(t1.Elem(), t2.Elem())
	case reflect.Array:
		h := t1.Len() == t2.Len()
		b1.Normal("[").Write(h, t1.Len()).Normal("]")
		b2.Normal("[").Write(h, t2.Len()).Normal("]")
		vd.writeDiffKinds(t1.Elem(), t2.Elem())
	case reflect.Slice:
		b1.Normal("[]")
		b2.Normal("[]")
		vd.writeDiffKinds(t1.Elem(), t2.Elem())
	case reflect.Map:
		b1.Normal("map[")
		b2.Normal("map[")
		vd.writeDiffKinds(t1.Key(), t2.Key())
		b1.Normal("]")
		b2.Normal("]")
		vd.writeDiffKinds(t1.Elem(), t2.Elem())
	case reflect.Struct:
		b1.Highlight(structName(t1))
		b2.Highlight(structName(t2))
	default:
		b1.Highlight(t1)
		b2.Highlight(t2)
	}
}

func (vd *ValueDiffer) writeDiffTypesFunc(t1, t2 reflect.Type) {
	b1, b2 := vd.bufs()
	b1.Normal("func(")
	b2.Normal("func(")
	for i := 0; i < t1.NumIn() || i < t2.NumIn(); i++ {
		if i >= t1.NumIn() {
			if i > 0 {
				b2.Plain(", ")
			}
			vd.writeType(1, t2.In(i), true)
		} else if i >= t2.NumIn() {
			if i > 0 {
				b1.Plain(", ")
			}
			vd.writeType(0, t1.In(i), true)
		} else {
			if i > 0 {
				b1.Normal(", ")
				b2.Normal(", ")
			}
			vd.writeDiffKinds(t1.In(i), t2.In(i))
		}
	}
	switch t1.NumOut() {
	case 0:
		b1.Normal(")")
	case 1:
		b1.Normal(") ")
	default:
		b1.Normal(") (")
		defer b1.Normal(")")
	}
	switch t2.NumOut() {
	case 0:
		b2.Normal(")")
	case 1:
		b2.Normal(") ")
	default:
		b2.Normal(") (")
		defer b2.Normal(")")
	}
	for i := 0; i < t1.NumOut() || i < t2.NumOut(); i++ {
		if i >= t1.NumOut() {
			if i > 0 {
				b2.Plain(", ")
			}
			vd.writeType(1, t2.Out(i), true)
		} else if i >= t2.NumOut() {
			if i > 0 {
				b1.Plain(", ")
			}
			vd.writeType(0, t1.Out(i), true)
		} else {
			if i > 0 {
				b1.Normal(", ")
				b2.Normal(", ")
			}
			vd.writeDiffKinds(t1.Out(i), t2.Out(i))
		}
	}
}

func (vd *ValueDiffer) writeTypeDiffValues(v1, v2 reflect.Value) {
	b1, b2 := vd.bufs()
	switch v1.Kind() {
	case reflect.Complex64, reflect.Complex128:
		c1, c2 := v1.Complex(), v2.Complex()
		hr, hi := real(c1) != real(c2), imag(c1) != imag(c2)
		b1.Normal("(").Write(hr, real(c1)).Normal("+").Write(hi, imag(c1)).Normal(")")
		b2.Normal("(").Write(hr, real(c2)).Normal("+").Write(hi, imag(c2)).Normal(")")
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
	case reflect.String:
		vd.writeDiffValuesString(v1, v2)
	case reflect.Array:
		vd.writeTypeDiffValuesArray(v1, v2, false)
	case reflect.Slice:
		if v1.IsNil() {
			b1.Highlight(nil)
			vd.writeElem(1, v2, true)
		} else if v2.IsNil() {
			vd.writeElem(0, v1, true)
			b2.Highlight(nil)
		} else {
			vd.writeTypeDiffValuesArray(v1, v2, true)
		}
	case reflect.Map:
		if v1.IsNil() {
			b1.Highlight(nil)
			vd.writeElem(1, v2, true)
		} else if v2.IsNil() {
			vd.writeElem(0, v1, true)
			b2.Highlight(nil)
		} else {
			vd.writeTypeDiffValuesMap(v1, v2)
		}
	case reflect.Struct:
		vd.writeTypeDiffValuesStruct(v1, v2)
	default:
		vd.writeElem(0, v1, true)
		vd.writeElem(1, v2, true)
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
			defer func() { b1.Tab--; b1.NL() }()
			vd.Attrs[NewLine+0] = true
		}
		if ml2 {
			b2.Tab++
			defer func() { b2.Tab--; b2.NL() }()
			vd.Attrs[NewLine+1] = true
		}
	} else {
		b1.Normal("[")
		b2.Normal("[")
		defer b1.Normal("]")
		defer b2.Normal("]")
	}
	if slice {
		vd.writeDiffValuesSlice(v1, v2, tp, id, ml1, ml2)
	} else {
		vd.writeDiffValuesArray(v1, v2, tp, id, ml1, ml2)
	}
}

func (vd *ValueDiffer) writeDiffValuesArray(v1, v2 reflect.Value, tp, id, ml1, ml2 bool) {
	b1, b2 := vd.bufs()
	var p1, p2 bool
	for i, j := 0, 0; i < v1.Len(); i++ {
		e1, e2 := v1.Index(i), v2.Index(i)
		eq := valueEqual(e1, e2)
		if eq && id {
			vd.Attrs[OmitSame] = true
			continue
		}
		t1, t2 := isNonTrivialElem(e1), isNonTrivialElem(e2)
		t1, p1 = (t1 || p1 || (ml1 && (id || i == 0))), t1
		t2, p2 = (t2 || p2 || (ml2 && (id || i == 0))), t2
		if j > 0 {
			if tp {
				b1.Plain(",")
				b2.Plain(",")
			}
			if !tp || !t1 {
				b1.Plain(" ")
			}
			if !tp || !t2 {
				b2.Plain(" ")
			}
		}
		j++
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
		if eq {
			vd.writeElem(0, e1, false)
			vd.writeElem(1, e2, false)
		} else {
			vd.writeDiff(e1, e2)
		}
	}
}

func (vd *ValueDiffer) writeDiffValuesSlice(v1, v2 reflect.Value, tp, id, ml1, ml2 bool) {
	b1, b2 := vd.bufs()
	var p1, p2 bool
	for i, j := 0, 0; i < v1.Len() || i < v2.Len(); i++ {
		g1, g2 := i < v1.Len(), i < v2.Len()
		eq := g1 && g2 && valueEqual(v1.Index(i), v2.Index(i))
		if eq && id { // If equal, skip
			vd.Attrs[OmitSame] = true
			// If all elems are skipped, show last elem's index (if it's NOT empty):
			// IDX:...
			if i+1 == v1.Len() && j == 0 {
				if ml1 {
					b1.NL()
				}
				b1.Normal(v1.Len()-1, ":...")
			}
			if i+1 == v2.Len() && j == 0 {
				if ml2 {
					b2.NL()
				}
				b2.Normal(v2.Len()-1, ":...")
			}
			continue
		}
		t1, t2 := g1 && isNonTrivialElem(v1.Index(i)), g2 && isNonTrivialElem(v2.Index(i))
		t1, p1 = g1 && (t1 || p1 || (ml1 && (id || i == 0))), t1
		t2, p2 = g2 && (t2 || p2 || (ml2 && (id || i == 0))), t2
		if j > 0 {
			if tp {
				if g1 {
					b1.Plain(",")
				}
				if g2 {
					b2.Plain(",")
				}
			}
			if g1 && (!tp || !t1) {
				b1.Plain(" ")
			}
			if g2 && (!tp || !t2) {
				b2.Plain(" ")
			}
		}
		j++
		if t1 {
			b1.NL()
		}
		if t2 {
			b2.NL()
		}
		if id {
			if g1 {
				b1.Write(!g2, i, ":")
			}
			if g2 {
				b2.Write(!g1, i, ":")
			}
		}
		if g1 && g2 {
			if e1, e2 := v1.Index(i), v2.Index(i); eq {
				vd.writeElem(0, e1, false)
				vd.writeElem(1, e2, false)
			} else {
				vd.writeDiff(e1, e2)
			}
		} else if g1 {
			vd.writeElem(0, v1.Index(i), true)
		} else {
			vd.writeElem(1, v2.Index(i), true)
		}
	}
}

func (vd *ValueDiffer) writeTypeDiffValuesMap(v1, v2 reflect.Value) {
	b1, b2 := vd.bufs()
	tp1, ml1 := attrElemMap(v1)
	tp2, ml2 := attrElemMap(v2)
	tp := tp1 || tp2
	if tp {
		vd.writeType(0, v1.Type(), false)
		vd.writeType(1, v2.Type(), false)
		b1.Normal("{")
		b2.Normal("{")
		defer b1.Normal("}")
		defer b2.Normal("}")
		if ml1 {
			b1.Tab++
			defer func() { b1.Tab--; b1.NL() }()
			vd.Attrs[NewLine+0] = true
		}
		if ml2 {
			b2.Tab++
			defer func() { b2.Tab--; b2.NL() }()
			vd.Attrs[NewLine+1] = true
		}
	} else {
		b1.Normal("map[")
		b2.Normal("map[")
		defer b1.Normal("]")
		defer b2.Normal("]")
	}
	vd.writeDiffValuesMap(v1, v2, tp, ml1, ml2)
}

func (vd *ValueDiffer) writeDiffValuesMap(v1, v2 reflect.Value, tp, ml1, ml2 bool) {
	b1, b2 := vd.bufs()
	var ks, ks1, ks2 []reflect.Value
	for _, k := range v1.MapKeys() {
		if v2.MapIndex(k).IsValid() {
			ks = append(ks, k)
		} else {
			ks1 = append(ks1, k)
		}
	}
	for _, k := range v2.MapKeys() {
		if !v1.MapIndex(k).IsValid() {
			ks2 = append(ks2, k)
		}
	}
	i := 0
	for _, k := range ks {
		if i > 0 {
			if tp {
				b1.Plain(",")
				b2.Plain(",")
			}
			if !ml1 {
				b1.Plain(" ")
			}
			if !ml2 {
				b2.Plain(" ")
			}
		}
		if ml1 {
			b1.NL()
		}
		if ml2 {
			b2.NL()
		}
		vd.writeKey(0, k, false)
		vd.writeKey(1, k, false)
		b1.Normal(":")
		b2.Normal(":")
		if e1, e2 := v1.MapIndex(k), v2.MapIndex(k); valueEqual(e1, e2) {
			vd.writeElem(0, e1, false)
			vd.writeElem(1, e2, false)
		} else {
			vd.writeDiff(e1, e2)
		}
		i++
	}
	f := func(idx int, v reflect.Value, ks []reflect.Value, ml bool, i int) {
		b := vd.bufi(idx)
		for _, k := range ks {
			if i > 0 {
				if tp {
					b.Highlight(",")
				}
				if !ml {
					b.Plain(" ")
				}
			}
			if ml {
				b.NL()
			}
			vd.writeKey(idx, k, true)
			b.Highlight(":")
			vd.writeElem(idx, v.MapIndex(k), true)
			i++
		}
	}
	f(0, v1, ks1, ml1, i)
	f(1, v2, ks2, ml2, i)
}

func (vd *ValueDiffer) writeTypeDiffValuesStruct(v1, v2 reflect.Value) {
	b1, b2 := vd.bufs()
	ml1, ml2 := attrElemStruct(v1), attrElemStruct(v2)
	if ml1 || ml2 {
		vd.writeType(0, v1.Type(), false)
		vd.writeType(1, v2.Type(), false)
	}
	b1.Normal("{")
	b2.Normal("{")
	defer b1.Normal("}")
	defer b2.Normal("}")
	if ml1 {
		b1.Tab++
		defer func() { b1.Tab--; b1.NL() }()
		vd.Attrs[NewLine+0] = true
	}
	if ml2 {
		b2.Tab++
		defer func() { b2.Tab--; b2.NL() }()
		vd.Attrs[NewLine+1] = true
	}
	vd.writeDiffValuesStruct(v1, v2, ml1, ml2)
}

func (vd *ValueDiffer) writeDiffValuesStruct(v1, v2 reflect.Value, ml1, ml2 bool) {
	b1, b2 := vd.bufs()
	t := v1.Type()
	for i := 0; i < v1.NumField(); i++ {
		if i > 0 {
			if ml1 {
				b1.Normal(",")
			} else {
				b1.Normal(" ")
			}
			if ml2 {
				b2.Normal(",")
			} else {
				b2.Normal(" ")
			}
		}
		if ml1 {
			b1.NL()
		}
		if ml2 {
			b2.NL()
		}
		n := t.Field(i).Name
		b1.Normal(n, ":")
		b2.Normal(n, ":")
		if e1, e2 := v1.Field(i), v2.Field(i); valueEqual(e1, e2) {
			vd.writeElem(0, e1, false)
			vd.writeElem(1, e2, false)
		} else {
			vd.writeDiff(v1.Field(i), v2.Field(i))
		}
	}
}

func (vd *ValueDiffer) writeDiffValuesString(v1, v2 reflect.Value) {
	b1, b2 := vd.bufs()
	s1, s2 := []rune(fmt.Sprintf("%#v", v1)), []rune(fmt.Sprintf("%#v", v2))
	s1, s2 = s1[1:len(s1)-1], s2[1:len(s2)-1] // skip front and end "
	b1.Normal(`"`)
	b2.Normal(`"`)
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
	b1.Normal(`"`)
	b2.Normal(`"`)
}

func (vd *ValueDiffer) bufs() (b1, b2 *FeatureBuf) {
	return vd.bufi(0), vd.bufi(1)
}

func valueEqual(v1, v2 reflect.Value) bool {
	if !v1.IsValid() || !v2.IsValid() {
		return v1.IsValid() == v2.IsValid()
	}
	if v1.Type() != v2.Type() {
		return false
	}
	if v1.CanInterface() && v2.CanInterface() {
		return reflect.DeepEqual(v1.Interface(), v2.Interface())
	}
	switch v1.Kind() {
	case reflect.Bool:
		return v1.Bool() == v2.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v1.Int() == v2.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v1.Uint() == v2.Uint()
	case reflect.Float32, reflect.Float64:
		return v1.Float() == v2.Float()
	case reflect.Complex64, reflect.Complex128:
		return v1.Complex() == v2.Complex()
	case reflect.String:
		return v1.String() == v2.String()
	case reflect.Chan, reflect.UnsafePointer:
		return v1.Pointer() == v2.Pointer()
	case reflect.Func:
		return v1.IsNil() && v2.IsNil()
	case reflect.Ptr:
		if v1.IsNil() || v2.IsNil() {
			return v1.IsNil() && v2.IsNil()
		}
		if v1.Pointer() == v2.Pointer() {
			return true
		}
		return valueEqual(v1.Elem(), v2.Elem())
	case reflect.Interface:
		if v1.IsNil() || v2.IsNil() {
			return v1.IsNil() == v2.IsNil()
		}
		return valueEqual(v1.Elem(), v2.Elem())
	case reflect.Slice:
		if v1.IsNil() != v2.IsNil() {
			return false
		}
		if v1.Len() != v2.Len() {
			return false
		}
		if v1.Pointer() == v2.Pointer() {
			return true
		}
		fallthrough
	case reflect.Array:
		for i := 0; i < v1.Len(); i++ {
			if !valueEqual(v1.Index(i), v2.Index(i)) {
				return false
			}
		}
		return true
	case reflect.Map:
		if v1.IsNil() != v2.IsNil() {
			return false
		}
		if v1.Len() != v2.Len() {
			return false
		}
		if v1.Pointer() == v2.Pointer() {
			return true
		}
		for _, k := range v1.MapKeys() {
			if e1, e2 := v1.MapIndex(k), v2.MapIndex(k); !valueEqual(e1, e2) {
				return false
			}
		}
		return true
	case reflect.Struct:
		for i, n := 0, v1.NumField(); i < n; i++ {
			if !valueEqual(v1.Field(i), v2.Field(i)) {
				return false
			}
		}
		return true
	default: // reflect.Invalid
	}
	return false
}
