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
	"fmt"
	"reflect"
)

func (vd *tValueDiffer) WriteTypeValue(idx int, v reflect.Value, tab int) {
	vd.bufi(idx).Tab = tab
	vd.writeTypeValue(idx, v, false, false)
}

// writeTypeValue show full description of value v.
// It show both type and value/contents of v. If v is a non-nil pointer of composite type (array,
// slice, map or struct), it show "&" and the results of *v instead, which imitates the action of
// Package fmt.
func (vd *tValueDiffer) writeTypeValue(idx int, v reflect.Value, ht, hv bool) {
	v = vd.writeTypeBeforeValue(idx, v, ht)
	vd.writeValueAfterType(idx, v, hv)
}

func (vd *tValueDiffer) writeTypeBeforeValue(idx int, v reflect.Value, hl bool) reflect.Value {
	b := vd.bufi(idx)
	if !v.IsValid() {
		b.Write(hl, nil)
		return v
	} else if v, df := derefInterface(v); df {
		return vd.writeTypeBeforeValue(idx, v, hl)
	} else if v, df := derefPtr(v); df {
		b.Write(hl, "&")
		return vd.writeTypeBeforeValue(idx, v, hl)
	} else if isPointer(v.Type()) {
		b.Normal("(")
		defer b.Normal(")")
	}
	vd.writeType(idx, v.Type(), hl)
	return v
}

// writeType shows type string of t.
// Basically it gives the same result as Package fmt dose with "%T", if not considering the
// highlight part.
// One exception is for unnamed struct, it shows "struct" only, rather than the full definition.
func (vd *tValueDiffer) writeType(idx int, t reflect.Type, hl bool) {
	b := vd.bufi(idx)
	if t.PkgPath() == "" {
		switch t.Kind() {
		case reflect.Ptr:
			b.Write(hl, "*")
			vd.writeType(idx, t.Elem(), hl)
		case reflect.Func:
			vd.writeTypeFunc(idx, t, hl)
		case reflect.Chan:
			vd.writeTypeHeadChan(idx, t, hl, false)
			vd.writeType(idx, t.Elem(), hl)
		case reflect.Array:
			b.Write(hl, "[", t.Len(), "]")
			vd.writeType(idx, t.Elem(), hl)
		case reflect.Slice:
			b.Write(hl, "[]")
			vd.writeType(idx, t.Elem(), hl)
		case reflect.Map:
			b.Write(hl, "map[")
			vd.writeType(idx, t.Key(), hl)
			b.Write(hl, "]")
			vd.writeType(idx, t.Elem(), hl)
		case reflect.Struct: // must be unnamed
			b.Write(hl, "struct")
		default:
			b.Write(hl, t)
		}
	} else {
		b.Write(hl, t)
	}
}

func (vd *tValueDiffer) writeTypeFunc(idx int, t reflect.Type, hl bool) {
	b := vd.bufi(idx)
	b.Write(hl, "func(")
	for i := 0; i < t.NumIn(); i++ {
		if i > 0 {
			b.Write(hl, ", ")
		}
		vd.writeType(idx, t.In(i), hl)
	}
	switch t.NumOut() {
	case 0:
		b.Write(hl, ")")
	case 1:
		b.Write(hl, ") ")
	default:
		b.Write(hl, ") (")
		defer b.Write(hl, ")")
	}
	for i := 0; i < t.NumOut(); i++ {
		if i > 0 {
			b.Write(hl, ", ")
		}
		vd.writeType(idx, t.Out(i), hl)
	}
}

func (vd *tValueDiffer) writeTypeHeadChan(idx int, t reflect.Type, hl, hldir bool) {
	b := vd.bufi(idx)
	switch t.ChanDir() {
	case reflect.RecvDir:
		b.Write(hl || hldir, "<-").Write(hl, "chan")
	case reflect.SendDir:
		b.Write(hl, "chan").Write(hl || hldir, "<-")
	default:
		b.Write(hl, "chan")
	}
	b.Plain(" ")
}

func (vd *tValueDiffer) writeValueAfterType(idx int, v reflect.Value, hl bool) {
	b := vd.bufi(idx)
	if !v.IsValid() {
		return
	}
	switch v.Kind() {
	case reflect.Complex64, reflect.Complex128:
		vd.writeElem(idx, v, hl)
	case reflect.Interface:
		if v, df := derefInterface(v); !df {
			b.Normal("(").Write(hl, "nil").Normal(")")
		} else if v.IsValid() {
			vd.writeValueAfterType(idx, v, hl)
		}
	case reflect.Array:
		vd.writeValueAfterTypeArray(idx, v, hl)
	case reflect.Slice:
		vd.writeValueAfterTypeSlice(idx, v, hl)
	case reflect.Map:
		vd.writeValueAfterTypeMap(idx, v, hl)
	case reflect.Struct:
		vd.writeValueAfterTypeStruct(idx, v, hl)
	case reflect.Chan, reflect.Func, reflect.UnsafePointer, reflect.Ptr:
		if v.Pointer() == 0 {
			b.Normal("(").Write(hl, "nil").Normal(")")
			break
		}
		fallthrough
	default:
		b.Normal("(")
		vd.writeElem(idx, v, hl)
		b.Normal(")")
	}
}

func (vd *tValueDiffer) writeValueAfterTypeArray(idx int, v reflect.Value, hl bool) {
	b := vd.bufi(idx)
	_, id, ml := attrElemArray(v)
	b.Write(hl, "{")
	defer b.Write(hl, "}")
	if ml {
		b.Tab++
		defer func() { b.Tab--; b.NL() }()
	}
	vd.writeElemArrayC(idx, v, true, id, ml, hl)
}

func (vd *tValueDiffer) writeValueAfterTypeSlice(idx int, v reflect.Value, hl bool) {
	b := vd.bufi(idx)
	if v.IsNil() {
		b.Normal("(").Write(hl, "nil").Normal(")")
		return
	}
	vd.writeValueAfterTypeArray(idx, v, hl)
}

func (vd *tValueDiffer) writeValueAfterTypeMap(idx int, v reflect.Value, hl bool) {
	b := vd.bufi(idx)
	if v.IsNil() {
		b.Normal("(").Write(hl, "nil").Normal(")")
		return
	}
	_, ml := attrElemMap(v)
	b.Write(hl, "{")
	defer b.Write(hl, "}")
	if ml {
		b.Tab++
		defer func() { b.Tab--; b.NL() }()
	}
	vd.writeElemMapC(idx, v, true, ml, hl)
}

func (vd *tValueDiffer) writeValueAfterTypeStruct(idx int, v reflect.Value, hl bool) {
	b := vd.bufi(idx)
	if ml := attrElemStruct(v); ml {
		vd.writeElemStructML(idx, v, hl)
	} else {
		b.Write(hl, "{")
		defer b.Write(hl, "}")
		t := v.Type()
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				b.Write(hl, ", ")
			}
			b.Write(hl, t.Field(i).Name, ":")
			vd.writeKey(idx, v.Field(i), hl)
		}
	}
}

// writeElem formats value v to a well readable string.
// It differs from writeKey() in representation for composite types (array, slice, map or struct),
// that writeElem may produce multi line strings if their contents (keys or elements) are also of composite
// types.
func (vd *tValueDiffer) writeElem(idx int, v reflect.Value, hl bool) {
	b := vd.bufi(idx)
	if !v.IsValid() {
		b.Write(hl, nil)
	} else {
		switch v.Kind() {
		case reflect.Interface:
			if v.IsNil() {
				b.Write(hl, nil)
			} else {
				vd.writeElem(idx, v.Elem(), hl)
			}
		case reflect.Array:
			vd.writeElemArray(idx, v, hl)
		case reflect.Slice:
			vd.writeElemSlice(idx, v, hl)
		case reflect.Map:
			vd.writeElemMap(idx, v, hl)
		case reflect.Struct:
			vd.writeElemStruct(idx, v, hl)
		default: // bool, integer, float, complex, channel, function, pointer, string
			vd.writeKey(idx, v, hl)
		}
	}
}

func (vd *tValueDiffer) writeElemArray(idx int, v reflect.Value, hl bool) {
	b := vd.bufi(idx)
	tp, id, ml := attrElemArray(v)
	if tp {
		vd.writeType(idx, v.Type(), hl)
		b.Write(hl, "{")
		defer b.Write(hl, "}")
		if ml {
			b.Tab++
			defer func() { b.Tab--; b.NL() }()
			vd.Attrs[kNewLine+idx] = true
		}
	} else {
		b.Write(hl, "[")
		defer b.Write(hl, "]")
	}
	vd.writeElemArrayC(idx, v, tp, id, ml, hl)
}

func (vd *tValueDiffer) writeElemArrayC(idx int, v reflect.Value, tp, id, ml, hl bool) {
	b := vd.bufi(idx)
	p := false
	for i := 0; i < v.Len(); i++ {
		e := v.Index(i)
		t := isNonTrivialElem(e)
		t, p = (t || p || (ml && (id || i == 0))), t
		if i > 0 {
			if tp {
				b.Write(hl, ",")
			}
			if !tp || !t {
				b.Write(hl, " ")
			}
		}
		if t {
			b.NL()
		}
		if id {
			b.Write(hl, i, ":")
		}
		vd.writeElem(idx, e, hl)
	}
}

func (vd *tValueDiffer) writeElemSlice(idx int, v reflect.Value, hl bool) {
	b := vd.bufi(idx)
	if v.IsNil() {
		b.Write(hl, nil)
		return
	}
	vd.writeElemArray(idx, v, hl)
}

func (vd *tValueDiffer) writeElemMap(idx int, v reflect.Value, hl bool) {
	b := vd.bufi(idx)
	if v.IsNil() {
		b.Write(hl, nil)
		return
	}
	tp, ml := attrElemMap(v)
	if tp {
		vd.writeType(idx, v.Type(), hl)
		b.Write(hl, "{")
		defer b.Write(hl, "}")
		if ml {
			b.Tab++
			defer func() { b.Tab--; b.NL() }()
			vd.Attrs[kNewLine+idx] = true
		}
	} else {
		b.Write(hl, "map[")
		defer b.Write(hl, "]")
	}
	vd.writeElemMapC(idx, v, tp, ml, hl)
}

func (vd *tValueDiffer) writeElemMapC(idx int, v reflect.Value, tp, ml, hl bool) {
	b := vd.bufi(idx)
	for i, k := range v.MapKeys() {
		if i > 0 {
			if tp {
				b.Write(hl, ",")
			}
			if !ml {
				b.Write(hl, " ")
			}
		}
		if ml {
			b.NL()
		}
		vd.writeKey(idx, k, hl)
		b.Write(hl, ":")
		vd.writeElem(idx, v.MapIndex(k), hl)
	}
}

func (vd *tValueDiffer) writeElemStruct(idx int, v reflect.Value, hl bool) {
	if ml := attrElemStruct(v); ml {
		vd.writeType(idx, v.Type(), hl)
		vd.writeElemStructML(idx, v, hl)
		vd.Attrs[kNewLine+idx] = true
	} else {
		vd.writeKeyStruct(idx, v, hl)
	}
}

func (vd *tValueDiffer) writeElemStructML(idx int, v reflect.Value, hl bool) {
	b := vd.bufi(idx)
	b.Write(hl, "{")
	b.Tab++
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		if i > 0 {
			b.Write(hl, ",")
		}
		b.NL().Write(hl, t.Field(i).Name, ":")
		vd.writeElem(idx, v.Field(i), hl)
	}
	b.Tab--
	b.NL().Write(hl, "}")
}

// writeKey formats value v to a concise string, without type.
// Generally, it shows the real representation of v, not its Error(), GoString() or String() as
// Package fmt does.
// The only exception is for POD types (boolean and integers), it tries to show String() first,
// because that's what we expect for enumerations.
func (vd *tValueDiffer) writeKey(idx int, v reflect.Value, hl bool) {
	b := vd.bufi(idx)
	if !v.IsValid() {
		b.Write(hl, nil)
		return
	}
	switch v.Kind() {
	case reflect.String:
		b.Writef(hl, "%q", v.String())
	case reflect.Float32:
		b.Writef(hl, "%g", float32(v.Float()))
	case reflect.Float64:
		b.Writef(hl, "%g", v.Float())
	case reflect.Complex64:
		c := complex64(v.Complex())
		b.Plain("(").Writef(hl, "%g+%gi", real(c), imag(c)).Write(b.PH, ")")
	case reflect.Complex128:
		c := v.Complex()
		b.Plain("(").Writef(hl, "%g+%gi", real(c), imag(c)).Write(b.PH, ")")
	case reflect.Func, reflect.Chan, reflect.Ptr, reflect.UnsafePointer:
		if v.Pointer() == 0 {
			b.Write(hl, nil)
		} else {
			b.Writef(hl, "%#v", v.Pointer())
		}
	case reflect.Interface:
		if v.IsNil() {
			b.Write(hl, nil)
		} else {
			vd.writeKey(idx, v.Elem(), hl)
		}
	case reflect.Array:
		vd.writeKeyArray(idx, v, hl)
	case reflect.Slice:
		vd.writeKeySlice(idx, v, hl)
	case reflect.Map:
		vd.writeKeyMap(idx, v, hl)
	case reflect.Struct:
		vd.writeKeyStruct(idx, v, hl)
	default:
		vd.writeKeyPOD(idx, v, hl, false)
	}
}

func (vd *tValueDiffer) writeKeyPOD(idx int, v reflect.Value, hl, up bool) {
	b := vd.bufi(idx)
	if v.CanInterface() {
		if s, ok := v.Interface().(fmt.Stringer); ok {
			b.Write(hl, s.String())
			return
		}
	}
	switch v.Kind() {
	case reflect.Uintptr:
		b.Writef(hl, "%#x", v.Uint())
		if up {
			b.Normalf("(%d)", v.Uint())
		}
	case reflect.Bool:
		b.Writef(hl, "%t", v.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		b.Writef(hl, "%d", v.Int())
	default: // reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		b.Writef(hl, "%d", v.Uint())
	}
}

func (vd *tValueDiffer) writeKeyArray(idx int, v reflect.Value, hl bool) {
	b := vd.bufi(idx)
	b.Write(hl, "[")
	for i := 0; i < v.Len(); i++ {
		if i > 0 {
			b.Write(hl, " ")
		}
		vd.writeKey(idx, v.Index(i), hl)
	}
	b.Write(hl, "]")
}

func (vd *tValueDiffer) writeKeySlice(idx int, v reflect.Value, hl bool) {
	b := vd.bufi(idx)
	if v.IsNil() {
		b.Write(hl, nil)
		return
	}
	vd.writeKeyArray(idx, v, hl)
}

func (vd *tValueDiffer) writeKeyMap(idx int, v reflect.Value, hl bool) {
	b := vd.bufi(idx)
	if v.IsNil() {
		b.Write(hl, nil)
		return
	}
	b.Write(hl, "map[")
	for i, k := range v.MapKeys() {
		if i > 0 {
			b.Write(hl, " ")
		}
		vd.writeKey(idx, k, hl)
		b.Write(hl, ":")
		vd.writeKey(idx, v.MapIndex(k), hl)
	}
	b.Write(hl, "]")
}

func (vd *tValueDiffer) writeKeyStruct(idx int, v reflect.Value, hl bool) {
	b := vd.bufi(idx)
	b.Write(hl, "{")
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		if i > 0 {
			b.Write(hl, " ")
		}
		b.Write(hl, t.Field(i).Name, ":")
		vd.writeKey(idx, v.Field(i), hl)
	}
	b.Write(hl, "}")
}

func (vd *tValueDiffer) bufi(i int) (b *tFeatureBuf) {
	b = &vd.b[i]
	if b.w == nil {
		b.w = &vd.buf[i]
	}
	return
}
