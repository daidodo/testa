package assert

import "reflect"

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
	b1, b2 := &vd.b[0], &vd.b[1]
	b1.Reset()
	b2.Reset()
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
		v1, v2 = vd.writeDiffTypesBeforeValue(v1, v2)
		vd.writeValueAfterType(0, v1)
		vd.writeValueAfterType(1, v2)
	}
}

func (vd *ValueDiffer) writeDiffTypesBeforeValue(v1, v2 reflect.Value) (r1, r2 reflect.Value) {
	if v1.Kind() == v2.Kind() {
		//TODO
		r1, r2 = v1, v2
		switch v1.Kind() {
		case reflect.Interface:
		case reflect.Ptr:
		case reflect.Chan:
		case reflect.Func:
		case reflect.Array:
		case reflect.Slice:
		case reflect.Map:
		case reflect.Struct:
		}
	} else {
		r1 = vd.writeHTypeBeforeValue(0, v1)
		r2 = vd.writeHTypeBeforeValue(1, v2)
	}
	return
}

func (vd *ValueDiffer) writeTypeDiffValues(v1, v2 reflect.Value) {
}

func (vd *ValueDiffer) writeHTypeValue(idx int, v reflect.Value) {
	v = vd.writeHTypeBeforeValue(idx, v)
	vd.writeValueAfterType(idx, v)
}

func (vd *ValueDiffer) writeHTypeBeforeValue(idx int, v reflect.Value) reflect.Value {
	b := &vd.b[idx]
	if !v.IsValid() {
		b.Highlight(nil)
	} else {
		switch v.Kind() {
		case reflect.Interface:
			if v.IsNil() {
				if n := interfaceName(v.Type()); n == "" {
					b.Highlight(nil)
				} else {
					b.Highlight(n) //TODO: test?
				}
			} else {
				v = vd.writeHTypeBeforeValue(idx, v.Elem())
			}
		case reflect.Ptr:
			if v.IsNil() {
				b.Write("(")
				if e := v.Type().Elem(); e.Kind() == reflect.Struct {
					b.Highlight("*", structName(e))
				} else {
					b.Highlight(v.Type())
				}
				b.Write(")")
			} else if e := v.Elem(); isComposite(e.Type()) {
				b.Highlight("&")
				v = vd.writeHTypeBeforeValue(idx, e)
			} else {
				b.Write("(").Highlight(v.Type()).Write(")")
			}
		case reflect.Struct:
			b.Highlight(structName(v.Type()))
		default:
			b.Highlight(v.Type())
		}
	}
	return v
}

func (vd *ValueDiffer) writeValueAfterType(idx int, v reflect.Value) {
	b := &vd.b[idx]
	if !v.IsValid() {
		return
	}
	switch v.Kind() {
	case reflect.Uintptr, reflect.String:
		b.Writef("(%#v)", v)
	case reflect.Complex64, reflect.Complex128:
		b.Write(v)
	case reflect.Chan, reflect.Func:
		if v.IsNil() {
			b.Write("(nil)")
		} else {
			b.Writef("(%v)", v)
		}
	case reflect.UnsafePointer:
		if v.Pointer() == 0 {
			b.Write("(nil)")
		} else {
			b.Writef("(%v)", v)
		}
	case reflect.Interface:
		if v.IsNil() {
			if interfaceName(v.Type()) != "" {
				b.Write("(nil)")
			}
		} else {
			panic("Should not come here")
			//vd.writeValueAfterType(idx, v.Elem()) //TODO: test?
		}
	case reflect.Ptr:
		if v.IsNil() {
			b.Write("(nil)")
		} else {
			b.Writef("(%v)", v)
		}
	case reflect.Array:
		vd.writeValueAfterTypeArray(idx, v)
	case reflect.Slice:
		vd.writeValueAfterTypeSlice(idx, v)
	case reflect.Map:
		vd.writeValueAfterTypeMap(idx, v)
	case reflect.Struct:
		vd.writeValueAfterTypeStruct(idx, v)
	default:
		b.Writef("(%v)", v)
	}
}

func (vd *ValueDiffer) writeValueAfterTypeArray(idx int, v reflect.Value) {
	if _, id, ml := attrElemArray(v); ml {
		vd.writeElemArrayML(idx, v)
	} else if id {
		vd.writeElemArrayID(idx, v)
	} else {
		vd.writeElemArrayTP(idx, v)
	}
}

func (vd *ValueDiffer) writeValueAfterTypeSlice(idx int, v reflect.Value) {
	b := &vd.b[idx]
	if v.IsNil() {
		b.Write("(nil)")
		return
	}
	vd.writeValueAfterTypeArray(idx, v)
}

func (vd *ValueDiffer) writeValueAfterTypeMap(idx int, v reflect.Value) {
	b := &vd.b[idx]
	if v.IsNil() {
		b.Write("(nil)")
		return
	}
	if _, ml := attrElemMap(v); ml {
		vd.writeElemMapML(idx, v)
	} else {
		vd.writeElemMapTP(idx, v)
	}
}

func (vd *ValueDiffer) writeValueAfterTypeStruct(idx int, v reflect.Value) {
	b := &vd.b[idx]
	if ml := attrElemStruct(v); ml {
		vd.writeElemStructML(idx, v)
	} else {
		b.Write("{")
		t := v.Type()
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				b.Write(", ")
			}
			b.Write(t.Field(i).Name, ":")
			vd.writeKey(idx, v.Field(i))
		}
		b.Write("}")
	}
}

//b := &vd.b[idx]
//if !v.IsValid() {
//    b.Highlight(nil)
//    return
//}
//switch v.Kind() {
//case reflect.Uintptr, reflect.String:
//    b.Highlight(v.Type()).Writef("(%#v)", v)
//case reflect.Complex64, reflect.Complex128:
//    b.Highlight(v.Type()).Writef("%v", v)
//case reflect.Chan, reflect.Func:
//    b.Highlight(v.Type())
//    if v.IsNil() {
//        b.Write("(nil)")
//    } else {
//        b.Writef("(%v)", v)
//    }
//case reflect.UnsafePointer:
//    b.Highlight(v.Type())
//    if v.Pointer() == 0 {
//        b.Write("(nil)")
//    } else {
//        b.Writef("(%v)", v)
//    }
//case reflect.Interface:
//    if v.IsNil() {
//        if n := interfaceName(v.Type()); n == "" {
//            b.Highlight(nil)
//        } else {
//			b.Highlight(n).Write("(nil)") //TODO: test?
//        }
//    } else {
//        vd.writeHTypeValue(idx, v.Elem())
//    }
//case reflect.Ptr:
//    if v.IsNil() {
//        b.Write("(")
//        if e := v.Type().Elem(); e.Kind() == reflect.Struct {
//            b.Highlight("*", structName(e))
//        } else {
//            b.Highlight(v.Type())
//        }
//        b.Write(")(nil)")
//    } else if e := v.Elem(); isComposite(e.Type()) {
//        b.Highlight("&")
//        vd.writeHTypeValue(idx, e)
//    } else {
//        b.Write("(").Highlight(v.Type()).Writef(")(%v)", v)
//    }
//case reflect.Array:
//    vd.writeHTypeValueArray(idx, v)
//case reflect.Slice:
//    vd.writeHTypeValueSlice(idx, v)
//case reflect.Map:
//    vd.writeHTypeValueMap(idx, v)
//case reflect.Struct:
//    vd.writeHTypeValueStruct(idx, v)
//default: // bool, integer, float
//    b.Highlight(v.Type()).Writef("(%v)", v)
//}
//}

func (vd *ValueDiffer) writeElem(idx int, v reflect.Value) {
	b := &vd.b[idx]
	if !v.IsValid() {
		b.Write(nil)
		return
	}
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

func (vd *ValueDiffer) writeElemArray(idx int, v reflect.Value) {
	b := &vd.b[idx]
	if tp, id, ml := attrElemArray(v); ml {
		b.Write(v.Type())
		vd.writeElemArrayML(idx, v)
	} else if id {
		b.Write(v.Type())
		vd.writeElemArrayID(idx, v)
	} else if tp {
		b.Write(v.Type())
		vd.writeElemArrayTP(idx, v)
	} else {
		vd.writeKeyArray(idx, v)
	}
}

func (vd *ValueDiffer) writeElemArrayML(idx int, v reflect.Value) {
	b := &vd.b[idx]
	b.Write("{")
	b.Tab++
	id, p := v.Len() > 10, false
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
}

func (vd *ValueDiffer) writeElemArrayID(idx int, v reflect.Value) {
	b := &vd.b[idx]
	b.Write("{")
	for i := 0; i < v.Len(); i++ {
		if i > 0 {
			b.Write(", ")
		}
		b.Write(i, ":")
		vd.writeElem(idx, v.Index(i))
	}
	b.Write("}")
}

func (vd *ValueDiffer) writeElemArrayTP(idx int, v reflect.Value) {
	b := &vd.b[idx]
	b.Write("{")
	for i := 0; i < v.Len(); i++ {
		if i > 0 {
			b.Write(", ")
		}
		vd.writeElem(idx, v.Index(i))
	}
	b.Write("}")
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
	if tp, ml := attrElemMap(v); ml {
		b.Write(v.Type())
		vd.writeElemMapML(idx, v)
	} else if tp {
		b.Write(v.Type())
		vd.writeElemMapTP(idx, v)
	} else {
		vd.writeKeyMap(idx, v)
	}
}

func (vd *ValueDiffer) writeElemMapML(idx int, v reflect.Value) {
	b := &vd.b[idx]
	b.Write("{")
	b.Tab++
	for i, k := range v.MapKeys() {
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
	vd.Attrs[NewLine] = true
}

func (vd *ValueDiffer) writeElemMapTP(idx int, v reflect.Value) {
	b := &vd.b[idx]
	b.Write("{")
	for i, k := range v.MapKeys() {
		if i > 0 {
			b.Write(", ")
		}
		vd.writeKey(idx, k)
		b.Write(":")
		vd.writeElem(idx, v.MapIndex(k))
	}
	b.Write("}")
}

func (vd *ValueDiffer) writeElemStruct(idx int, v reflect.Value) {
	b := &vd.b[idx]
	if ml := attrElemStruct(v); ml {
		b.Write(structName(v.Type()))
		vd.writeElemStructML(idx, v)
	} else {
		vd.writeKeyStruct(idx, v)
	}
}

func (vd *ValueDiffer) writeElemStructML(idx int, v reflect.Value) {
	b := &vd.b[idx]
	b.Write("{")
	b.Tab++
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		if i > 0 {
			b.Write(",")
		}
		b.NL().Write(t.Field(i).Name, ":")
		vd.writeElem(idx, v.Field(i))
	}
	b.Tab--
	b.NL().Write("}")
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
	b.Write("map[")
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

func isNonTrivialElem(v reflect.Value) bool {
	if !v.IsValid() || !isNonTrivial(v.Type()) {
		return false
	}
	switch v.Kind() {
	case reflect.Interface:
		return isNonTrivialElem(v.Elem())
	case reflect.Array:
		return v.Len() > 0
	case reflect.Slice, reflect.Map:
		return !v.IsNil() && v.Len() > 0
	case reflect.Struct:
		return v.NumField() > 0
	}
	panic("Should not come here!")
}

func attrElemArray(v reflect.Value) (tp, id, ml bool) {
	if v.Len() > 0 {
		id = v.Len() > 10
		for i := 0; i < v.Len() && !ml; i++ {
			ml = isNonTrivialElem(v.Index(i))
		}
		tp = id || ml || isReference(v.Type().Elem())
	}
	return
}

func attrElemMap(v reflect.Value) (tp, ml bool) {
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
	return
}

func attrElemStruct(v reflect.Value) (ml bool) {
	for i := 0; i < v.NumField() && !ml; i++ {
		ml = isNonTrivialElem(v.Field(i))
	}
	return
}

func isComposite(t reflect.Type) bool {
	k := t.Kind()
	return k == reflect.Array || k == reflect.Map || k == reflect.Slice || k == reflect.Struct
}

func isNonTrivial(t reflect.Type) bool {
	k := t.Kind()
	return k == reflect.Interface || isComposite(t)
}

func isReference(t reflect.Type) bool {
	k := t.Kind()
	return k == reflect.Chan || k == reflect.Func || k == reflect.Ptr || k == reflect.UnsafePointer || isNonTrivial(t)
}

func structName(t reflect.Type) string {
	if t.Name() == "" {
		return "struct"
	}
	return t.String()
}

func interfaceName(t reflect.Type) string {
	if t.Name() == "" {
		return ""
	}
	return t.String()
}
