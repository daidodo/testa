/*
 * Copyright (c) 2017 Zhao DAI <daidodo@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or any
 * later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see accompanying file LICENSE.txt
 * or <http://www.gnu.org/licenses/>.
 */

package assert

import (
	"bytes"
	"reflect"
	"strings"
)

type tValueDiffer struct {
	buf   [2]bytes.Buffer
	b     [2]tFeatureBuf
	Attrs [kAttrSize]bool
}

func (vd *tValueDiffer) String(i int) string {
	vd.b[i].Finish()
	return vd.buf[i].String()
}

func (vd *tValueDiffer) WriteDiff(v1, v2 reflect.Value, tab int) {
	b1, b2 := vd.bufs()
	b1.Tab, b2.Tab = tab, tab
	vd.writeDiff(v1, v2)
}

func (vd *tValueDiffer) writeDiff(v1, v2 reflect.Value) {
	if !v1.IsValid() || !v2.IsValid() || v1.Type() != v2.Type() {
		vd.writeDiffTypeValues(v1, v2)
	} else {
		vd.writeTypeDiffValues(v1, v2)
	}
}

func (vd *tValueDiffer) writeDiffTypeValues(v1, v2 reflect.Value) {
	b1, b2 := vd.bufs()
	v1, r1 := derefInterface(v1)
	v2, r2 := derefInterface(v2)
	if r1 || r2 {
		vd.writeDiff(v1, v2)
	} else {
		e1, r1 := derefPtr(v1)
		e2, r2 := derefPtr(v2)
		if r1 && r2 {
			b1.Normal("&")
			b2.Normal("&")
			vd.writeDiff(e1, e2)
		} else if !v1.IsValid() || !v2.IsValid() {
			vd.writeTypeValue(0, v1, true, false)
			vd.writeTypeValue(1, v2, true, false)
		} else {
			vd.writeDiffKindsBeforeValue(v1, v2, func(t1, t2 reflect.Type) bool { return t1 == t2 }, false)
			vd.writeValueAfterType(0, v1, false)
			vd.writeValueAfterType(1, v2, false)
		}
	}
}

func (vd *tValueDiffer) writeDiffKindsBeforeValue(v1, v2 reflect.Value, eq func(t1, t2 reflect.Type) bool, sw bool) {
	b1, b2, _, _ := vd.bufr(sw)
	t1, t2 := v1.Type(), v2.Type()
	if isPointer(t1) {
		b1.Normal("(")
		defer b1.Normal(")")
	}
	if isPointer(t2) {
		b2.Normal("(")
		defer b2.Normal(")")
	}
	vd.writeDiffKinds(t1, t2, eq, sw)
}

func (vd *tValueDiffer) writeDiffKinds(t1, t2 reflect.Type, eq func(v1, v2 reflect.Type) bool, sw bool) {
	_, _, i1, i2 := vd.bufr(sw)
	if t1 == nil || t2 == nil {
		panic("Should not come here!")
	}
	if eq(t1, t2) {
		vd.writeType(i1, t1, false)
		vd.writeType(i2, t2, false)
	} else if t1.PkgPath() == "" && t2.PkgPath() == "" && t1.Kind() == t2.Kind() {
		vd.writeDiffTypes(t1, t2, eq, sw)
	} else if t1.PkgPath() == "" || t2.PkgPath() == "" {
		vd.writeType(i1, t1, true)
		vd.writeType(i2, t2, true)
	} else {
		vd.writeDiffPkgTypes(t1, t2, sw)
	}
}

func (vd *tValueDiffer) writeDiffPkgTypes(t1, t2 reflect.Type, sw bool) {
	b1, b2, _, _ := vd.bufr(sw)
	if t1.PkgPath() == t2.PkgPath() {
		p := lastPartOf(t1.PkgPath())
		b1.Normal(p, ".").Highlight(t1.Name())
		b2.Normal(p, ".").Highlight(t2.Name())
	} else {
		p1 := strings.Split(t1.PkgPath(), "/")
		p2 := strings.Split(t2.PkgPath(), "/")
		i := 1
		for ; i <= len(p1) && i <= len(p2) && p1[len(p1)-i] == p2[len(p2)-i]; i++ {
		}
		if i < len(p1) {
			p1 = p1[len(p1)-i:]
		}
		if i < len(p2) {
			p2 = p2[len(p2)-i:]
		}
		pt := func(b *tFeatureBuf, p []string, nh bool) {
			if !nh {
				b.Highlight(p[0])
				p = p[1:]
				if len(p) > 0 {
					b.Normal("/")
				}
			}
			for i, c := range p {
				if i > 0 {
					b.Normal("/")
				}
				b.Normal(c)
			}
			b.Plain(".")
		}
		pt(b1, p1, len(p1) < len(p2))
		pt(b2, p2, len(p2) < len(p1))
		h := t1.Name() != t2.Name()
		b1.Write(h, t1.Name())
		b2.Write(h, t2.Name())
	}
}

func (vd *tValueDiffer) writeDiffTypes(t1, t2 reflect.Type, eq func(t1, t2 reflect.Type) bool, sw bool) {
	b1, b2, i1, i2 := vd.bufr(sw)
	switch t1.Kind() {
	case reflect.Ptr:
		b1.Normal("*")
		b2.Normal("*")
		vd.writeDiffKinds(t1.Elem(), t2.Elem(), eq, sw)
	case reflect.Func:
		vd.writeDiffTypesFunc(t1, t2, eq, sw)
	case reflect.Chan:
		hd := t1.ChanDir() != t2.ChanDir()
		vd.writeTypeHeadChan(i1, t1, false, hd)
		vd.writeTypeHeadChan(i2, t2, false, hd)
		vd.writeDiffKinds(t1.Elem(), t2.Elem(), eq, sw)
	case reflect.Array:
		h := t1.Len() != t2.Len()
		b1.Normal("[").Write(h, t1.Len()).Normal("]")
		b2.Normal("[").Write(h, t2.Len()).Normal("]")
		vd.writeDiffKinds(t1.Elem(), t2.Elem(), eq, sw)
	case reflect.Slice:
		b1.Normal("[]")
		b2.Normal("[]")
		vd.writeDiffKinds(t1.Elem(), t2.Elem(), eq, sw)
	case reflect.Map:
		b1.Normal("map[")
		b2.Normal("map[")
		vd.writeDiffKinds(t1.Key(), t2.Key(), eq, sw)
		b1.Normal("]")
		b2.Normal("]")
		vd.writeDiffKinds(t1.Elem(), t2.Elem(), eq, sw)
	case reflect.Struct: // must be unnamed struct
		b1.Highlight("struct")
		b2.Highlight("struct")
	default:
		b1.Highlight(t1)
		b2.Highlight(t2)
	}
}

func (vd *tValueDiffer) writeDiffTypesFunc(t1, t2 reflect.Type, eq func(t1, t2 reflect.Type) bool, sw bool) {
	b1, b2, i1, i2 := vd.bufr(sw)
	b1.Normal("func(")
	b2.Normal("func(")
	for i := 0; i < t1.NumIn() || i < t2.NumIn(); i++ {
		g1, g2 := i < t1.NumIn(), i < t2.NumIn()
		if i > 0 {
			if g1 {
				b1.Plain(", ")
			}
			if g2 {
				b2.Plain(", ")
			}
		}
		if g1 && g2 {
			vd.writeDiffKinds(t1.In(i), t2.In(i), eq, sw)
		} else if g1 {
			vd.writeType(i1, t1.In(i), true)
		} else {
			vd.writeType(i2, t2.In(i), true)
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
		g1, g2 := i < t1.NumOut(), i < t2.NumOut()
		if i > 0 {
			if g1 {
				b1.Plain(", ")
			}
			if g2 {
				b2.Plain(", ")
			}
		}
		if g1 && g2 {
			vd.writeDiffKinds(t1.Out(i), t2.Out(i), eq, sw)
		} else if g1 {
			vd.writeType(i1, t1.Out(i), true)
		} else {
			vd.writeType(i2, t2.Out(i), true)
		}
	}
}

func (vd *tValueDiffer) writeTypeDiffValues(v1, v2 reflect.Value) {
	b1, b2 := vd.bufs()
	switch v1.Kind() {
	case reflect.Complex64:
		c1, c2 := complex64(v1.Complex()), complex64(v2.Complex())
		hr, hi := real(c1) != real(c2), imag(c1) != imag(c2)
		b1.Plain("(").Write(hr, real(c1))
		b2.Plain("(").Write(hr, real(c2))
		h1, h2 := b1.PH, b2.PH
		b1.Plain("+").Writef(hi, "%gi", imag(c1)).Write(h1, ")")
		b2.Plain("+").Writef(hi, "%gi", imag(c2)).Write(h2, ")")
	case reflect.Complex128:
		c1, c2 := v1.Complex(), v2.Complex()
		hr, hi := real(c1) != real(c2), imag(c1) != imag(c2)
		b1.Plain("(").Write(hr, real(c1))
		b2.Plain("(").Write(hr, real(c2))
		h1, h2 := b1.PH, b2.PH
		b1.Plain("+").Writef(hi, "%gi", imag(c1)).Write(h1, ")")
		b2.Plain("+").Writef(hi, "%gi", imag(c2)).Write(h2, ")")
	case reflect.String:
		vd.writeDiffValuesString(v1, v2, false)
	case reflect.Func:
		vd.writeDiffValuesFunc(v1, v2)
	case reflect.Interface:
		vd.writeDiffValuesInterface(v1, v2)
	case reflect.Ptr:
		vd.writeDiffValuesPtr(v1, v2)
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
			eq := func(v1, v2 reflect.Value) bool { return valueEqual(v1, v2) }
			wd := func(v1, v2 reflect.Value) { vd.writeDiff(v1, v2) }
			vd.writeTypeDiffValuesMap(v1, v2, false, eq, wd)
		}
	case reflect.Struct:
		vd.writeTypeDiffValuesStruct(v1, v2)
	default:
		vd.writeElem(0, v1, true)
		vd.writeElem(1, v2, true)
	}
}

func (vd *tValueDiffer) writeDiffValuesInterface(v1, v2 reflect.Value) {
	b1, b2 := vd.bufs()
	if v1.IsNil() {
		b1.Highlight(nil)
		vd.writeTypeValue(1, v2.Elem(), false, false)
	} else if v2.IsNil() {
		vd.writeTypeValue(0, v1.Elem(), false, false)
		b2.Highlight(nil)
	} else {
		vd.writeDiff(v1.Elem(), v2.Elem())
	}
}

func (vd *tValueDiffer) writeDiffValuesPtr(v1, v2 reflect.Value) {
	b1, b2 := vd.bufs()
	e1, d1 := derefPtr(v1)
	e2, d2 := derefPtr(v2)
	if d1 && d2 {
		b1.Normal("&")
		b2.Normal("&")
		vd.writeTypeDiffValues(e1, e2)
	} else {
		vd.writeElem(0, v1, true)
		vd.writeElem(1, v2, true)
	}
}

func (vd *tValueDiffer) writeDiffValuesFunc(v1, v2 reflect.Value) {
	b1, b2 := vd.bufs()
	if v1.IsNil() && v2.IsNil() {
		b1.Normal("nil")
		b2.Normal("nil")
	} else {
		b1.Highlight(v1)
		b2.Highlight(v2)
		if v1.Pointer() == v2.Pointer() {
			vd.Attrs[kCompFunc] = true
		}
	}
}

func (vd *tValueDiffer) writeTypeDiffValuesArray(v1, v2 reflect.Value, slice bool) {
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
			vd.Attrs[kNewLine+0] = true
		}
		if ml2 {
			b2.Tab++
			defer func() { b2.Tab--; b2.NL() }()
			vd.Attrs[kNewLine+1] = true
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

func (vd *tValueDiffer) writeDiffValuesArray(v1, v2 reflect.Value, tp, id, ml1, ml2 bool) {
	b1, b2 := vd.bufs()
	var p1, p2 bool
	for i, j := 0, 0; i < v1.Len(); i++ {
		e1, e2 := v1.Index(i), v2.Index(i)
		eq := valueEqual(e1, e2)
		if eq && id {
			vd.Attrs[kOmitSame] = true
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

func (vd *tValueDiffer) writeDiffValuesSlice(v1, v2 reflect.Value, tp, id, ml1, ml2 bool) {
	eq := func(a, b reflect.Value) bool {
		return valueEqual(a, b)
	}
	wd := func(a, b reflect.Value) {
		vd.writeDiff(a, b)
	}
	vd.writeDiffValuesArrayC(v1, v2, false, tp, id, ml1, ml2, eq, wd)
}

func (vd *tValueDiffer) writeTypeDiffValuesMap(v1, v2 reflect.Value, sw bool,
	eq func(v1, v2 reflect.Value) bool,
	wd func(v1, v2 reflect.Value)) {
	b1, b2, i1, i2 := vd.bufr(sw)
	tp1, ml1 := attrElemMap(v1)
	tp2, ml2 := attrElemMap(v2)
	tp := tp1 || tp2
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
			vd.Attrs[kNewLine+i2] = true
		}
	} else {
		b1.Normal("map[")
		b2.Normal("map[")
		defer b1.Normal("]")
		defer b2.Normal("]")
	}
	vd.writeDiffValuesMap(v1, v2, sw, tp, ml1, ml2, eq, wd)
}

func (vd *tValueDiffer) writeDiffValuesMap(v1, v2 reflect.Value, sw, tp, ml1, ml2 bool,
	eqf func(v1, v2 reflect.Value) bool,
	wd func(v1, v2 reflect.Value)) {
	b1, b2, i1, i2 := vd.bufr(sw)
	ks, ks1, ks2 := mapKeyDiff(v1, v2)
	id := v1.Len() > 10 || v2.Len() > 10
	i := 0
	for _, k := range ks {
		e1, e2 := v1.MapIndex(k), v2.MapIndex(k)
		eq := eqf(e1, e2)
		if eq && id {
			vd.Attrs[kOmitSame] = true
			continue
		}
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
		vd.writeKey(i1, k, false)
		vd.writeKey(i2, k, false)
		b1.Normal(":")
		b2.Normal(":")
		if eq {
			vd.writeElem(i1, e1, false)
			vd.writeElem(i2, e2, false)
		} else {
			wd(e1, e2)
		}
		i++
	}
	// If all are skipped, show "...:..." if NOT empty
	if len(ks) > 0 && i == 0 {
		if len(ks1) == 0 {
			if ml1 {
				b1.NL()
			}
			b1.Normal("...:...")
		} else if len(ks2) == 0 {
			if ml2 {
				b2.NL()
			}
			b2.Normal("...:...")
		}
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
	f(i1, v1, ks1, ml1, i)
	f(i2, v2, ks2, ml2, i)
}

func (vd *tValueDiffer) writeTypeDiffValuesStruct(v1, v2 reflect.Value) {
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
		vd.Attrs[kNewLine+0] = true
	}
	if ml2 {
		b2.Tab++
		defer func() { b2.Tab--; b2.NL() }()
		vd.Attrs[kNewLine+1] = true
	}
	vd.writeDiffValuesStruct(v1, v2, ml1, ml2)
}

func (vd *tValueDiffer) writeDiffValuesStruct(v1, v2 reflect.Value, ml1, ml2 bool) {
	b1, b2 := vd.bufs()
	id := v1.NumField() > 10
	t := v1.Type()
	for i, j := 0, 0; i < v1.NumField(); i++ {
		e1, e2 := v1.Field(i), v2.Field(i)
		eq := valueEqual(e1, e2)
		if eq && id {
			vd.Attrs[kOmitSame] = true
			continue
		}
		if j > 0 {
			if ml1 {
				b1.Plain(",")
			} else {
				b1.Plain(" ")
			}
			if ml2 {
				b2.Plain(",")
			} else {
				b2.Plain(" ")
			}
		}
		j++
		if ml1 {
			b1.NL()
		}
		if ml2 {
			b2.NL()
		}
		n := t.Field(i).Name
		b1.Normal(n, ":")
		b2.Normal(n, ":")
		if eq {
			vd.writeElem(0, e1, false)
			vd.writeElem(1, e2, false)
		} else {
			vd.writeDiff(v1.Field(i), v2.Field(i))
		}
	}
}

const (
	kNewLine = iota
	_
	kOmitSame
	kCompFunc

	kAttrSize
)

func valueEqual(v1, v2 reflect.Value) bool {
	if !v1.IsValid() || !v2.IsValid() {
		return v1.IsValid() == v2.IsValid()
	}
	if v1.CanInterface() && v2.CanInterface() {
		return reflect.DeepEqual(v1.Interface(), v2.Interface())
	}
	v1, d1 := derefInterface(v1)
	v2, d2 := derefInterface(v2)
	if d1 || d2 {
		return valueEqual(v1, v2)
	}
	if v1.Type() != v2.Type() {
		return false
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

func mapKeyDiff(v1, v2 reflect.Value) (ks, ks1, ks2 []reflect.Value) {
	if t1, t2 := v1.Type().Key(), v2.Type().Key(); t1 == t2 {
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
	} else if convertibleKeyTo(t1, t2) {
		s1, s2 := v1.MapKeys(), v2.MapKeys()
		find := func(k1 reflect.Value, s []reflect.Value) bool {
			for _, k2 := range s {
				if valueEqual(k1, k2) {
					return true
				}
			}
			return false
		}
		for _, k := range s1 {
			if find(k, s2) {
				ks = append(ks, k)
			} else {
				ks1 = append(ks1, k)
			}
		}
		for _, k := range s2 {
			if !find(k, s1) {
				ks2 = append(ks2, k)
			}
		}
	} else if convertibleKeyTo(t2, t1) {
		ks, ks2, ks1 = mapKeyDiff(v2, v1)
	} else {
		ks1, ks2 = v1.MapKeys(), v2.MapKeys()
	}
	return
}

func derefInterface(v reflect.Value) (r reflect.Value, d bool) {
	if v.IsValid() && v.Kind() == reflect.Interface {
		if !v.IsNil() {
			return v.Elem(), true
		} else if v.Type().Name() == "" {
			return r, true
		}
	}
	return v, d
}

func derefPtr(v reflect.Value) (r reflect.Value, d bool) {
	if v.IsValid() && v.Kind() == reflect.Ptr && !v.IsNil() {
		if e := v.Elem(); isComposite(e.Type()) {
			return e, true
		}
	}
	return v, false
}

func (vd *tValueDiffer) bufs() (b1, b2 *tFeatureBuf) {
	return vd.bufi(0), vd.bufi(1)
}
