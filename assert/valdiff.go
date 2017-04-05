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
	"fmt"
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

func (vd *tValueDiffer) WriteDiffT(t1, t2 reflect.Type, tab int) {
	b1, b2 := vd.bufs()
	b1.Tab, b2.Tab = tab, tab
	if t1 == nil {
		b1.Highlight(nil)
		vd.writeType(1, t2, false)
	} else if t2 == nil {
		vd.writeType(0, t1, false)
		b2.Highlight(nil)
	} else {
		vd.writeDiffKinds(t1, t2, func(t1, t2 reflect.Type) bool { return t1 == t2 }, false)
	}
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
	eq := func(v1, v2 reflect.Value) bool { return valueEqual(v1, v2) }
	wd := func(v1, v2 reflect.Value, sw bool) {
		if sw {
			v1, v2 = v2, v1
		}
		vd.writeDiff(v1, v2)
	}
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
		vd.writeTypeDiffValuesArray(v1, v2, false, eq, wd)
	case reflect.Slice:
		if v1.IsNil() {
			b1.Highlight(nil)
			vd.writeElem(1, v2, true)
		} else if v2.IsNil() {
			vd.writeElem(0, v1, true)
			b2.Highlight(nil)
		} else {
			vd.writeTypeDiffValuesArray(v1, v2, false, eq, wd)
		}
	case reflect.Map:
		if v1.IsNil() {
			b1.Highlight(nil)
			vd.writeElem(1, v2, true)
		} else if v2.IsNil() {
			vd.writeElem(0, v1, true)
			b2.Highlight(nil)
		} else {
			vd.writeTypeDiffValuesMap(v1, v2, false, eq, wd)
		}
	case reflect.Struct:
		vd.writeTypeDiffValuesStruct(v1, v2, false, eq, wd)
	default:
		vd.writeElem(0, v1, true)
		vd.writeElem(1, v2, true)
	}
}

func (vd *tValueDiffer) writeDiffValuesString(v1, v2 reflect.Value, sw bool) {
	b1, b2, _, _ := vd.bufr(sw)
	s1, s2 := []rune(fmt.Sprintf("%#v", v1.String())), []rune(fmt.Sprintf("%#v", v2.String()))
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
	if !v1.IsNil() && !v2.IsNil() {
		b1.Normal("&")
		b2.Normal("&")
		vd.writeTypeDiffValues(v1.Elem(), v2.Elem())
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

func (vd *tValueDiffer) writeTypeDiffValuesArray(v1, v2 reflect.Value, sw bool,
	eq func(a, b reflect.Value) bool,
	wd func(a, b reflect.Value, sw bool)) {
	b1, b2, i1, i2 := vd.bufr(sw)
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
			vd.Attrs[kNewLine+i2] = true
		}
	} else {
		b1.Normal("[")
		b2.Normal("[")
		defer b1.Normal("]")
		defer b2.Normal("]")
	}
	vd.writeDiffValuesArrayC(v1, v2, sw, tp, id, ml1, ml2, eq, wd)
}

func (vd *tValueDiffer) writeDiffValuesArrayC(v1, v2 reflect.Value, sw, tp, id, ml1, ml2 bool,
	eq func(a, b reflect.Value) bool,
	wd func(a, b reflect.Value, s bool)) {
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
				wd(e1, e2, sw)
			}
		} else if g1 {
			vd.writeElem(i1, v1.Index(i), true)
		} else {
			vd.writeElem(i2, v2.Index(i), true)
		}
	}
}

func (vd *tValueDiffer) writeTypeDiffValuesMap(v1, v2 reflect.Value, sw bool,
	eq func(v1, v2 reflect.Value) bool,
	wd func(v1, v2 reflect.Value, sw bool)) {
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
	wd func(v1, v2 reflect.Value, sw bool)) {
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
			wd(e1, e2, sw)
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
					b.Plain(",")
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

func (vd *tValueDiffer) writeTypeDiffValuesStruct(v1, v2 reflect.Value, sw bool,
	eq func(v1, v2 reflect.Value) bool,
	wd func(v1, v2 reflect.Value, sw bool)) {
	b1, b2, i1, i2 := vd.bufr(sw)
	ml1, ml2 := attrElemStruct(v1), attrElemStruct(v2)
	tp := ml1 || ml2
	if tp {
		vd.writeType(i1, v1.Type(), false)
		vd.writeType(i2, v2.Type(), false)
	}
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
	vd.writeDiffValuesStruct(v1, v2, sw, tp, ml1, ml2, eq, wd)
}

func (vd *tValueDiffer) writeDiffValuesStruct(v1, v2 reflect.Value, sw, tp, ml1, ml2 bool,
	eqf func(v1, v2 reflect.Value) bool,
	wd func(v1, v2 reflect.Value, sw bool)) {
	b1, b2, i1, i2 := vd.bufr(sw)
	id := v1.NumField() > 10
	t := v1.Type()
	for i, j := 0, 0; i < v1.NumField(); i++ {
		e1, e2 := v1.Field(i), v2.Field(i)
		eq := eqf(e1, e2)
		if eq && id {
			vd.Attrs[kOmitSame] = true
			continue
		}
		if j > 0 {
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
			vd.writeElem(i1, e1, false)
			vd.writeElem(i2, e2, false)
		} else {
			wd(v1.Field(i), v2.Field(i), sw)
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

func (vd *tValueDiffer) bufs() (b1, b2 *tFeatureBuf) {
	return vd.bufi(0), vd.bufi(1)
}
