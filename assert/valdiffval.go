package assert

import (
	"fmt"
	"reflect"
)

func (vd *tValueDiffer) WriteValDiff(v1, v2 reflect.Value, tab int) {
	b1, b2 := vd.bufs()
	b1.Tab, b2.Tab = tab, tab
	vd.writeValDiff(v1, v2, false)
}

func (vd *tValueDiffer) writeValDiff(v1, v2 reflect.Value, sw bool) {
	v1, d1 := derefInterface(v1)
	v2, d2 := derefInterface(v2)
	if d1 || d2 {
		vd.writeValDiff(v1, v2, sw)
		return
	}
	if !v1.IsValid() {
		vd.writeValDiffToNil(v2, !sw)
		return
	} else if !v2.IsValid() {
		vd.writeValDiffToNil(v1, sw)
		return
	}
	if t, k := v2.Type(), v2.Kind(); isSimpleNumber(t) {
		vd.writeValDiffToNumber(v1, v2, sw)
	} else if k == reflect.String {
		vd.writeValDiffToString(v1, v2, sw)
	} else if k == reflect.Complex64 {
		vd.writeValDiffToComplex64(v1, v2, sw)
	} else if k == reflect.Complex128 {
		vd.writeValDiffToComplex128(v1, v2, sw)
	} else if isSimplePointer(t) {
		vd.writeValDiffToPointer(v1, v2, sw)
	} else if isArray(t) {
		vd.writeValDiffToArray(v1, v2, sw)
	} else {
		vd.writeValDiffC(v1, v2, sw)
	}
}

func (vd *tValueDiffer) writeValDiffToNil(v reflect.Value, sw bool) {
	_, b2, i1, _ := vd.bufr(sw)
	if t := v.Type(); isComposite(t) {
		vd.writeTypeValue(i1, v, false, false)
	} else if isReference(t) {
		vd.writeElem(i1, v, true)
	} else {
		vd.writeTypeValue(i1, v, true, false)
	}
	b2.Highlight(nil)
}

func (vd *tValueDiffer) writeValDiffToNumber(v1, v2 reflect.Value, sw bool) {
	_, _, i1, i2 := vd.bufr(sw)
	if t, k, k2 := v1.Type(), v1.Kind(), v2.Kind(); isSimpleNumber(t) {
		if k == reflect.Uintptr && k2 != reflect.Uintptr {
			vd.writeKeyPOD(i1, v1, true, true)
			vd.writeElem(i2, v2, true)
		} else if k != reflect.Uintptr && k2 == reflect.Uintptr {
			vd.writeElem(i1, v1, true)
			vd.writeKeyPOD(i2, v2, true, true)
		} else {
			vd.writeElem(i1, v1, true)
			vd.writeElem(i2, v2, true)
		}
	} else if k == reflect.String {
		vd.writeValDiffToString(v2, v1, !sw)
	} else if k == reflect.Complex64 {
		vd.writeValDiffToComplex64(v2, v1, !sw)
	} else if k == reflect.Complex128 {
		vd.writeValDiffToComplex128(v2, v1, !sw)
	} else {
		vd.writeValDiffC(v1, v2, sw)
	}
}

func (vd *tValueDiffer) writeValDiffToString(v1, v2 reflect.Value, sw bool) {
	_, _, i1, i2 := vd.bufr(sw)
	if t := v1.Type(); isInteger(t) || isUInteger(t) {
		vd.writeElem(i1, v1, true)
		vd.writeElem(i2, v2, true)
	} else if t.Kind() == reflect.String {
		vd.writeDiffValuesString(v1, v2, sw)
	} else {
		vd.writeValDiffC(v1, v2, sw)
	}
}

func (vd *tValueDiffer) writeValDiffToComplex64(v1, v2 reflect.Value, sw bool) {
	b1, b2, i1, _ := vd.bufr(sw)
	c2 := complex64(v2.Complex())
	r2, m2 := real(c2), imag(c2)
	if t, k := v1.Type(), v1.Kind(); isInteger(t) {
		hr, hi := float32(v1.Int()) != r2, 0 != m2
		vd.writeElem(i1, v1, hr)
		b2.Normal("(").Write(hr, r2).Plain("+").Writef(hi, "%gi", m2).Normal(")")
	} else if isUInteger(t) {
		hr, hi := float32(v1.Uint()) != r2, 0 != m2
		if k == reflect.Uintptr {
			vd.writeKeyPOD(i1, v1, hr, !hi)
		} else {
			vd.writeElem(i1, v1, hr)
		}
		b2.Normal("(").Write(hr, r2).Plain("+").Writef(hi, "%gi", m2).Normal(")")
	} else if k == reflect.Float32 {
		hr, hi := float32(v1.Float()) != r2, 0 != m2
		vd.writeElem(i1, v1, hr)
		b2.Normal("(").Write(hr, r2).Plain("+").Writef(hi, "%gi", m2).Normal(")")
	} else if k == reflect.Float64 {
		hr, hi := v1.Float() != float64(r2), 0 != m2
		vd.writeElem(i1, v1, hr)
		b2.Normal("(").Write(hr, r2).Plain("+").Writef(hi, "%gi", m2).Normal(")")
	} else if k == reflect.Complex64 {
		c1 := complex64(v1.Complex())
		r1, m1 := real(c1), imag(c1)
		hr, hi := r1 != r2, m1 != m2
		b1.Normal("(").Write(hr, r1).Plain("+").Writef(hi, "%gi", m1).Normal(")")
		b2.Normal("(").Write(hr, r2).Plain("+").Writef(hi, "%gi", m2).Normal(")")
	} else if k == reflect.Complex128 {
		c1 := v1.Complex()
		r1, m1 := real(c1), imag(c1)
		hr, hi := r1 != float64(r2), m1 != float64(m2)
		b1.Normal("(").Write(hr, r1).Plain("+").Writef(hi, "%gi", m2).Normal(")")
		b2.Normal("(").Write(hr, r2).Plain("+").Writef(hi, "%gi", m2).Normal(")")
	} else {
		vd.writeValDiffC(v1, v2, sw)
	}
}

func (vd *tValueDiffer) writeValDiffToComplex128(v1, v2 reflect.Value, sw bool) {
	b1, b2, i1, _ := vd.bufr(sw)
	c2 := v2.Complex()
	r2, m2 := real(c2), imag(c2)
	if t, k := v1.Type(), v1.Kind(); isInteger(t) {
		hr, hi := float64(v1.Int()) != r2, 0 != m2
		vd.writeElem(i1, v1, hr)
		b2.Normal("(").Write(hr, r2).Plain("+").Writef(hi, "%gi", m2).Normal(")")
	} else if isUInteger(t) {
		hr, hi := float64(v1.Uint()) != r2, 0 != m2
		if k == reflect.Uintptr {
			vd.writeKeyPOD(i1, v1, hr, !hi)
		} else {
			vd.writeElem(i1, v1, hr)
		}
		b2.Normal("(").Write(hr, r2).Plain("+").Writef(hi, "%gi", m2).Normal(")")
	} else if k == reflect.Float32 {
		hr, hi := v1.Float() != r2, 0 != m2
		vd.writeElem(i1, v1, hr)
		b2.Normal("(").Write(hr, r2).Plain("+").Writef(hi, "%gi", m2).Normal(")")
	} else if k == reflect.Float64 {
		hr, hi := v1.Float() != r2, 0 != m2
		vd.writeElem(i1, v1, hr)
		b2.Normal("(").Write(hr, r2).Plain("+").Writef(hi, "%gi", m2).Normal(")")
	} else if k == reflect.Complex64 {
		c1 := complex64(v1.Complex())
		r1, m1 := real(c1), imag(c1)
		hr, hi := float64(r1) != r2, float64(m1) != m2
		b1.Normal("(").Write(hr, r1).Plain("+").Writef(hi, "%gi", m1).Normal(")")
		b2.Normal("(").Write(hr, r2).Plain("+").Writef(hi, "%gi", m2).Normal(")")
	} else if k == reflect.Complex128 {
		c1 := v1.Complex()
		r1, m1 := real(c1), imag(c1)
		hr, hi := r1 != r2, m1 != m2
		b1.Normal("(").Write(hr, r1).Plain("+").Writef(hi, "%gi", m1).Normal(")")
		b2.Normal("(").Write(hr, r2).Plain("+").Writef(hi, "%gi", m2).Normal(")")
	} else {
		vd.writeValDiffC(v1, v2, sw)
	}
}

func (vd *tValueDiffer) writeValDiffToPointer(v1, v2 reflect.Value, sw bool) {
	_, _, i1, i2 := vd.bufr(sw)
	if t := v1.Type(); isSimplePointer(t) && (t.Kind() == reflect.UnsafePointer || v2.Kind() == reflect.UnsafePointer) {
		vd.writeElem(i1, v1, true)
		vd.writeElem(i2, v2, true)
	} else {
		vd.writeValDiffC(v1, v2, sw)
	}
}

func (vd *tValueDiffer) writeValDiffToArray(v1, v2 reflect.Value, sw bool) {
	b1, b2, i1, i2 := vd.bufr(sw)
	t1, t2 := v1.Type(), v2.Type()
	k1, k2 := t1.Kind(), t2.Kind()
	if (k1 == reflect.Array || (k1 == reflect.Slice && !v1.IsNil())) &&
		(k2 == reflect.Array || (k2 == reflect.Slice && !v2.IsNil())) &&
		convertible(t1.Elem(), t2.Elem()) {
		tp1, id1, ml1 := attrElemArray(v1)
		tp2, id2, ml2 := attrElemArray(v2)
		tp, id := tp1 || tp2, id1 || id2
		if tp {
			vd.writeType(i1, v1.Type(), false)
			vd.writeType(i2, v2.Type(), false)
			b1.Normal("{")
			b2.Normal("{")
			defer b1.Normal("}")
			defer b2.Normal("}")
			if ml1 {
				b1.Tab++
				defer func() { b1.Tab--; b1.NL() }()
				vd.Attrs[kNewLine+i1] = true
			}
			if ml2 {
				b2.Tab++
				defer func() { b2.Tab--; b2.NL() }()
				vd.Attrs[kNewLine+i1] = true
			}
		} else {
			b1.Normal("[")
			b2.Normal("[")
			defer b1.Normal("]")
			defer b2.Normal("]")
		}
		eq := func(a, b reflect.Value) bool {
			return convertCompare(a, b)
		}
		wd := func(a, b reflect.Value) {
			vd.writeValDiff(a, b, sw)
		}
		vd.writeDiffValuesArrayC(v1, v2, sw, tp, id, ml1, ml2, eq, wd)
	} else {
		vd.writeValDiffC(v1, v2, sw)
	}
}

func (vd *tValueDiffer) writeDiffValuesArrayC(v1, v2 reflect.Value, sw, tp, id, ml1, ml2 bool,
	eq func(a, b reflect.Value) bool,
	wd func(a, b reflect.Value)) {
	b1, b2, i1, i2 := vd.bufr(sw)
	var p1, p2 bool
	for i, j := 0, 0; i < v1.Len() || i < v2.Len(); i++ {
		g1, g2 := i < v1.Len(), i < v2.Len()
		eq := g1 && g2 && eq(v1.Index(i), v2.Index(i))
		if eq && id { // If equal, skip
			vd.Attrs[kOmitSame] = true
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
				vd.writeElem(i1, e1, false)
				vd.writeElem(i2, e2, false)
			} else {
				wd(e1, e2)
			}
		} else if g1 {
			vd.writeElem(i1, v1.Index(i), true)
		} else {
			vd.writeElem(i2, v2.Index(i), true)
		}
	}
}

func (vd *tValueDiffer) writeDiffValuesString(v1, v2 reflect.Value, sw bool) {
	b1, b2, _, _ := vd.bufr(sw)
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

func (vd *tValueDiffer) writeValDiffC(v1, v2 reflect.Value, sw bool) {
	_, _, i1, i2 := vd.bufr(sw)
	if t1, t2 := v1.Type(), v2.Type(); convertible(t1, t2) {
		vd.writeTypeValue(i1, v1, false, true)
		vd.writeTypeValue(i2, v2, false, true)
	} else if isComposite(t1) && isComposite(t2) {
		vd.writeDiffKindsBeforeValue(v1, v2, convertible, sw)
		vd.writeValueAfterType(i1, v1, false)
		vd.writeValueAfterType(i2, v2, false)
	} else {
		vd.writeTypeValue(i1, v1, true, false)
		vd.writeTypeValue(i2, v2, true, false)
	}
}

func (vd *tValueDiffer) bufr(sw bool) (b1, b2 *tFeatureBuf, i1, i2 int) {
	if sw {
		return vd.bufi(1), vd.bufi(0), 1, 0
	}
	return vd.bufi(0), vd.bufi(1), 0, 1
}
