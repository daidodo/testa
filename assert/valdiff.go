package assert

import (
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
	b     [2]FeatureBuf
	Attrs [kAttrSize]bool
}

func (vd *ValueDiffer) String(i int) string {
	return vd.b[i].String()
}

func (vd *ValueDiffer) WriteDiff(v1, v2 reflect.Value, tab int) {
	b1, b2 := vd.b[0], vd.b[1]
	b1.Reset()
	b2.Reset()
	b1.Tab, b2.Tab = tab, tab
	if !v1.IsValid() || !v2.IsValid() || v1.Kind() != v2.Kind() {
		vd.writeDiffKindValues(v1, v2)
	} else if v1.Type() != v2.Type() {
		vd.writeDiffTypeValues(v1, v2)
	} else {
		vd.writeTypeDiffValues(v1, v2)
	}
}

func (vd *ValueDiffer) writeDiffKindValues(v1, v2 reflect.Value) {
	vd.writeHTypeValue(0, v1)
	vd.writeHTypeValue(1, v2)
}

func (vd *ValueDiffer) writeHTypeValue(idx int, v reflect.Value) {
	b := &vd.b[idx]
	if !v.IsValid() {
		b.Highlight(nil)
		return
	}
	switch v.Kind() {
	case reflect.Uintptr, reflect.String:
		b.Highlight(v.Type()).Writef("(%#v)", v)
	case reflect.Complex64, reflect.Complex128:
		b.Highlight(v.Type()).Writef("%v", v)
	case reflect.Array:
		vd.writeHTypeValueArray(idx, v)
	case reflect.Chan, reflect.Func, reflect.UnsafePointer:
		b.Write("(").Highlight(v.Type())
		if v.IsNil() {
			b.Write(")(nil)")
		} else {
			b.Writef(")(%v)", v)
		}
	case reflect.Interface:
		panic(fmt.Sprintf("Should not come here, v=%v(%#v)", v.Type(), v))
	case reflect.Map:
		vd.writeHTypeValueMap(idx, v)
	case reflect.Ptr:
		vd.writeHTypeValuePtr(idx, v)
	case reflect.Slice:
		vd.writeHTypeValueSlice(idx, v)
	case reflect.Struct:
		vd.writeHTypeValueStruct(idx, v)
	default: // bool, integer, float
		b.Highlight(v.Type).Writef("(%v)", v)
	}
}

func (vd *ValueDiffer) writeHTypeValueArray(idx int, v reflect.Value) {
	b := vd.b[idx]
	b.Highlight(v.Type()).Write("{")
	if v.Len() > 0 {
		id := v.Len() > 10
		if isComposite(v.Index(0).Type()) {
			b.Tab++
			for i := 0; i < v.Len(); i++ {
				b.NL()
				if id {
					b.Write(i, ":")
				}
				vd.writeValue(idx, v.Index(i)) // writeElem?
			}
			b.Tab--
			b.NL()
			vd.Attrs[NewLine] = true
		} else {
			for i := 0; i < v.Len(); i++ {
				if i > 0 {
					b.Write(", ")
				}
				if id {
					b.Write(i, ":")
				}
				vd.writeValue(idx, v.Index(i)) // TODO: writeElem?
			}
		}
	}
	b.Write("}")
}

func (vd *ValueDiffer) writeHTypeValueSlice(idx int, v reflect.Value) {
}

func (vd *ValueDiffer) writeHTypeValueMap(idx int, v reflect.Value) {
}

func (vd *ValueDiffer) writeHTypeValuePtr(idx int, v reflect.Value) {
}

func (vd *ValueDiffer) writeHTypeValueStruct(idx int, v reflect.Value) {
}

// TODO: writeField?
func (vd *ValueDiffer) writeValue(idx int, v reflect.Value) {
	b := &vd.b[idx]
	if !v.IsValid() {
		b.Write(nil)
	} else {
		switch v.Kind() {
		case reflect.Array:
			vd.writeValueArray(idx, v)
		case reflect.Interface:
			panic(fmt.Sprintf("Should not come here, v=%v(%#v)", v.Type(), v))
		case reflect.Map:
			vd.writeValueMap(idx, v)
		case reflect.Ptr:
			vd.writeValuePtr(idx, v)
		case reflect.Slice:
			vd.writeValueSlice(idx, v)
		case reflect.Struct:
			vd.writeValueStruct(idx, v)
		default: // bool, integer, float, complex, channel, function, unsafe pointer, string
			vd.writeKey(idx, v)
		}
	}
}

func (vd *ValueDiffer) writeValueArray(idx int, v reflect.Value) {
	b := &vd.b[idx]
	if v.Len() < 1 {
		b.Write("[]")
		return
	}
	id := v.Len() > 10
	if isComposite(v.Index(0).Type()) {
		b.Write(v.Type(), "{")
		b.Tab++
		for i := 0; i < v.Len(); i++ {
			b.NL()
			if id {
				b.Write(i, ":")
			}
			vd.writeElem(idx, v.Index(i))
		}
		b.Tab--
		b.NL().Write("}")
		vd.Attrs[NewLine] = true
	} else if id {
		b.Write(v.Type(), "{")
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				b.Write(", ")
			}
			b.Write(i, ":")
			vd.writeElem(idx, v.Index(i))
		}
		b.Write("}")
	} else {
		b.Write("[")
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				b.Write(", ")
			}
			vd.writeElem(idx, v.Index(i))
		}
		b.Write("]")
	}
}

func (vd *ValueDiffer) writeValuePtr(idx int, v reflect.Value) {
}

//func (vd *ValueDiffer) writeElemPtr(idx int, v reflect.Value) {
//    b := &vd.b[idx]
//    if v.IsNil() {
//        b.Write(v)
//    } else if e := v.Elem(); isComposite(e.Type()) {
//        b.Write("&")
//        vd.writeElem(idx, e)
//    } else {
//        b.Write(v)
//    }
//}

func (vd *ValueDiffer) writeValueSlice(idx int, v reflect.Value) {
	b := &vd.b[idx]
	if v.IsNil() {
		b.Write(v)
		return
	}
	vd.writeValueArray(idx, v)
}

func (vd *ValueDiffer) writeValueMap(idx int, v reflect.Value) {
	b := &vd.b[idx]
	if v.IsNil() {
		b.Write(v) //TODO: nil?
		return
	} else if v.Len() < 1 {
		b.Write("map[]")
		return
	}
	b.Write(v.Type(), "{")
	keys := v.MapKeys()
	if isComposite(keys[0].Type()) || isComposite(v.MapIndex(keys[0]).Type()) {
		b.Tab++
		for _, k := range keys {
			b.NL()
			vd.writeKey(idx, k)
			b.Write(":")
			vd.writeElem(idx, v.MapIndex(k))
		}
		b.Tab--
		b.NL()
	} else {
		for i, k := range keys {
			if i > 0 {
				b.Write(", ")
			}
			vd.writeKey(idx, k)
			b.Write(":")
			vd.writeElem(idx, v.MapIndex(k))
		}
	}
	b.Write("}")
}

func (vd *ValueDiffer) writeValueStruct(idx int, v reflect.Value) {
	b := &vd.b[idx]
	b.Write(structName(v), "{")
	comp := false
	for i := 0; i < v.NumField(); i++ {
		if isComposite(v.Field(i).Type()) {
			comp = true
			break
		}
	}
	t := v.Type()
	if comp {
		b.Tab++
		for i := 0; i < v.NumField(); i++ {
			b.NL().Write(t.Field(i).Name, ":")
			vd.writeValue(idx, v.Field(i))
		}
		b.Tab--
		b.NL()
	} else {
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				b.Write(", ")
			}
			b.Write(t.Field(i).Name, ":")
			vd.writeValue(idx, v.Field(i))
		}
	}
	b.Write("}")
}

func (vd *ValueDiffer) writeElem(idx int, v reflect.Value) {
	b := &vd.b[idx]
	if !v.IsValid() {
		b.Write(nil)
	} else {
		switch v.Kind() {
		case reflect.Interface:
			if v.IsNil() {
				b.Write(nil)
			} else {
				vd.writeElem(idx, v.Elem())
			}
		case reflect.Array:
			vd.writeElemArray(idx, v)
		case reflect.Slice:
			vd.writeElemSlice(idx, v)
		case reflect.Map:
			vd.writeElemMap(idx, v)
		case reflect.Struct:
			vd.writeElemStruct(idx, v)
		default: // bool, integer, float, complex, channel, function, pointer, string
			vd.writeKey(idx, v)
		}
	}
}

func (vd *ValueDiffer) writeElemArray(idx int, v reflect.Value) {
	b := &vd.b[idx]
	var id, tp, ml bool
	if v.Len() > 0 {
		id = v.Len() > 10
		for i := 0; i < v.Len() && !ml; i++ {
			ml = isNonTrivialElem(v.Index(i))
		}
		tp = id || ml || isReference(v.Type().Elem())
	}
	if ml {
		b.Write(v.Type(), "{")
		b.Tab++
		p := false
		for i := 0; i < v.Len(); i++ {
			e := v.Index(i)
			t := isNonTrivialElem(e)
			t, p = (id || i == 0 || p || t), t
			if i > 0 {
				b.Write(",")
			}
			if t {
				b.NL()
			} else {
				b.Write(" ")
			}

			if id {
				b.Write(i, ":")
			}
			vd.writeElem(idx, e)
		}
		b.Tab--
		b.NL().Write("}")
		vd.Attrs[NewLine] = true
	} else if id {
		b.Write(v.Type(), "{")
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				b.Write(", ")
			}
			b.Write(i, ":")
			vd.writeElem(idx, v.Index(i))
		}
		b.Write("}")
	} else if tp {
		b.Write(v.Type(), "{")
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				b.Write(", ")
			}
			vd.writeElem(idx, v.Index(i))
		}
		b.Write("}")
	} else {
		b.Write("[")
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				b.Write(" ")
			}
			vd.writeElem(idx, v.Index(i))
		}
		b.Write("]")
	}
}

func (vd *ValueDiffer) writeElemSlice(idx int, v reflect.Value) {
	b := &vd.b[idx]
	if v.IsNil() {
		b.Write(nil)
		return
	}
	vd.writeElemArray(idx, v)
}

func (vd *ValueDiffer) writeElemMap(idx int, v reflect.Value) {
	b := &vd.b[idx]
	if v.IsNil() {
		b.Write(nil)
		return
	}
	var tp, ml bool
	ks := v.MapKeys()
	if v.Len() > 0 {
		for _, k := range ks {
			if ml = isNonTrivialElem(k); ml {
				break
			}
			if ml = isNonTrivialElem(v.MapIndex(k)); ml {
				break
			}
		}
		t := v.Type()
		tp = ml || isReference(t.Key()) || isReference(t.Elem())
	}
	if ml {
		b.Write(v.Type(), "{")
		b.Tab++
		for i, k := range ks {
			if i > 0 {
				b.Write(",")
			}
			b.NL()
			vd.writeKey(idx, k)
			b.Write(":")
			vd.writeElem(idx, v.MapIndex(k))
		}
		b.Tab--
		b.NL().Write("}")
	} else if tp {
		b.Write(v.Type(), "{")
		for i, k := range ks {
			if i > 0 {
				b.Write(", ")
			}
			vd.writeKey(idx, k)
			b.Write(":")
			vd.writeElem(idx, v.MapIndex(k))
		}
		b.Write("}")
	} else {
		b.Write("map[")
		for i, k := range ks {
			if i > 0 {
				b.Write(" ")
			}
			vd.writeKey(idx, k)
			b.Write(":")
			vd.writeElem(idx, v.MapIndex(k))
		}
		b.Write("]")
	}
}

func (vd *ValueDiffer) writeElemStruct(idx int, v reflect.Value) {
}

func isNonTrivialElem(v reflect.Value) bool {
	return !isTrivialElem(v)
}

func isTrivialElem(v reflect.Value) bool {
	if !v.IsValid() || !isComposite(v.Type()) {
		return true
	}
	switch v.Kind() {
	case reflect.Interface:
		return isTrivialElem(v.Elem())
	case reflect.Array:
		return v.Len() < 1
	case reflect.Slice, reflect.Map:
		return v.IsNil() || v.Len() < 1
	}
	return false // struct
}

func (vd *ValueDiffer) writeKey(idx int, v reflect.Value) {
	b := &vd.b[idx]
	if !v.IsValid() {
		b.Write(nil)
	} else {
		switch v.Kind() {
		case reflect.Uintptr, reflect.String:
			b.Writef("%#v", v)
		case reflect.Array:
			vd.writeKeyArray(idx, v)
		case reflect.Interface:
			if v.IsNil() {
				b.Write(nil)
			} else {
				vd.writeKey(idx, v.Elem())
			}
		case reflect.Map:
			vd.writeKeyMap(idx, v)
		case reflect.Ptr:
			if v.IsNil() {
				b.Write(nil)
			} else {
				b.Writef("%#v", v.Pointer())
			}
		case reflect.Slice:
			vd.writeKeySlice(idx, v)
		case reflect.Struct:
			vd.writeKeyStruct(idx, v)
		default: // bool, integer, float, complex, channel, function, pointer
			b.Write(v)
		}
	}
}

func (vd *ValueDiffer) writeKeyArray(idx int, v reflect.Value) {
	b := &vd.b[idx]
	b.Write("[")
	for i := 0; i < v.Len(); i++ {
		if i > 0 {
			b.Write(" ")
		}
		vd.writeKey(idx, v.Index(i))
	}
	b.Write("]")
}

func (vd *ValueDiffer) writeKeySlice(idx int, v reflect.Value) {
	b := &vd.b[idx]
	if v.IsNil() {
		b.Write(nil)
		return
	}
	vd.writeKeyArray(idx, v)
}

func (vd *ValueDiffer) writeKeyMap(idx int, v reflect.Value) {
	b := &vd.b[idx]
	if v.IsNil() {
		b.Write(nil)
		return
	}
	ks := v.MapKeys()
	if len(ks) < 1 {
		b.Write("map")
	}
	b.Write("[")
	for i, k := range v.MapKeys() {
		if i > 0 {
			b.Write(" ")
		}
		vd.writeKey(idx, k)
		b.Write(":")
		vd.writeKey(idx, v.MapIndex(k))
	}
	b.Write("]")
}

func (vd *ValueDiffer) writeKeyStruct(idx int, v reflect.Value) {
	b := &vd.b[idx]
	b.Write("{")
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		if i > 0 {
			b.Write(" ")
		}
		b.Write(t.Field(i).Name, ":")
		vd.writeKey(idx, v.Field(i))
	}
	b.Write("}")
}

func (vd *ValueDiffer) writeDiffTypeValues(v1, v2 reflect.Value) {
}

func (vd *ValueDiffer) writeTypeDiffValues(v1, v2 reflect.Value) {
}

func isComposite(t reflect.Type) bool {
	k := t.Kind()
	return k == reflect.Array || k == reflect.Map || k == reflect.Slice || k == reflect.Struct || k == reflect.Interface
}

func isReference(t reflect.Type) bool {
	k := t.Kind()
	return k == reflect.Chan || k == reflect.Func || k == reflect.Ptr || k == reflect.UnsafePointer || isComposite(t)
}

func structName(v reflect.Value) string {
	n := v.Type().Name()
	if n != "" {
		return n
	}
	return "struct"
}
