package assert

import "reflect"

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
		h := t1.ChanDir() == t2.ChanDir()
		vd.writeTypeHeadChan(0, t1, false, h)
		vd.writeTypeHeadChan(1, t2, false, h)
		vd.writeDiffKinds(t2.Elem(), t2.Elem())
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

func (vd *ValueDiffer) writeType(idx int, t reflect.Type, hl bool) {
	b := vd.bufi(idx)
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
	case reflect.Struct:
		b.Write(hl, structName(t))
	default:
		b.Write(hl, t)
	}
}

func (vd *ValueDiffer) writeTypeFunc(idx int, t reflect.Type, hl bool) {
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
			b.Normal(", ")
		}
		vd.writeType(idx, t.Out(i), hl)
	}
}

func (vd *ValueDiffer) writeTypeHeadChan(idx int, t reflect.Type, hl, hldir bool) {
	b := vd.bufi(idx)
	switch t.ChanDir() {
	case reflect.RecvDir:
		b.Write(hl || hldir, "<-").Write(hl, "chan ")
	case reflect.SendDir:
		b.Write(hl, "chan").Write(hl || hldir, "<- ")
	default:
		b.Write(hl, "chan ")
	}
}

func (vd *ValueDiffer) writeTypeBeforeValue(idx int, v reflect.Value, hl bool) reflect.Value {
	b := vd.bufi(idx)
	if !v.IsValid() {
		b.Write(hl, nil)
	} else {
		switch v.Kind() {
		case reflect.Interface:
			if !vd.writeTypeBeforeInterfaceNil(idx, v, hl) {
				v = vd.writeTypeBeforeValue(idx, v.Elem(), hl)
			}
		case reflect.Ptr, reflect.Func, reflect.Chan:
			b.Normal("(")
			vd.writeType(idx, v.Type(), hl)
			b.Normal(")")
		default:
			vd.writeType(idx, v.Type(), hl)
		}
	}
	return v
}

func (vd *ValueDiffer) writeTypeBeforeInterfaceNil(idx int, v reflect.Value, hl bool) (isNil bool) {
	b := vd.bufi(idx)
	if isNil = v.IsNil(); isNil {
		if n := interfaceName(v.Type()); n == "" {
			if hl {
				b.Highlight(nil)
			} else {
				b.Normal(nil)
			}
		} else {
			if hl {
				b.Highlight(n)
			} else {
				b.Normal(n)
			}
		}
	}
	return
}

func (vd *ValueDiffer) writeValueAfterType(idx int, v reflect.Value) {
	b := vd.bufi(idx)
	if !v.IsValid() {
		return
	}
	switch v.Kind() {
	case reflect.Uintptr, reflect.String:
		b.Normalf("(%#v)", v)
	case reflect.Complex64, reflect.Complex128:
		b.Normal(v)
	case reflect.Chan, reflect.Func:
		if v.IsNil() {
			b.Normal("(nil)")
		} else {
			b.Normalf("(%v)", v)
		}
	case reflect.UnsafePointer:
		if v.Pointer() == 0 {
			b.Normal("(nil)")
		} else {
			b.Normalf("(%v)", v)
		}
	case reflect.Interface:
		if v.IsNil() {
			if interfaceName(v.Type()) != "" {
				b.Normal("(nil)")
			}
		} else {
			panic("Should not come here!")
		}
	case reflect.Ptr:
		if v.IsNil() {
			b.Normal("(nil)")
		} else {
			b.Normalf("(%v)", v)
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
		b.Normalf("(%v)", v)
	}
}

func (vd *ValueDiffer) writeValueAfterTypeArray(idx int, v reflect.Value) {
	b := vd.bufi(idx)
	_, id, ml := attrElemArray(v)
	b.Normal("{")
	defer b.Normal("}")
	if ml {
		b.Tab++
		defer func() { b.Tab-- }()
	}
	vd.writeElemArrayC(idx, v, true, id, ml)
}

func (vd *ValueDiffer) writeValueAfterTypeSlice(idx int, v reflect.Value) {
	b := vd.bufi(idx)
	if v.IsNil() {
		b.Normal("(nil)")
		return
	}
	vd.writeValueAfterTypeArray(idx, v)
}

func (vd *ValueDiffer) writeValueAfterTypeMap(idx int, v reflect.Value) {
	b := vd.bufi(idx)
	if v.IsNil() {
		b.Normal("(nil)")
		return
	}
	_, ml := attrElemMap(v)
	b.Normal("{")
	defer b.Normal("}")
	if ml {
		b.Tab++
		defer func() { b.Tab-- }()
	}
	vd.writeElemMapC(idx, v, true, ml)
}

func (vd *ValueDiffer) writeValueAfterTypeStruct(idx int, v reflect.Value) {
	b := vd.bufi(idx)
	if ml := attrElemStruct(v); ml {
		vd.writeElemStructML(idx, v)
	} else {
		b.Normal("{")
		t := v.Type()
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				b.Normal(", ")
			}
			b.Normal(t.Field(i).Name, ":")
			vd.writeKey(idx, v.Field(i))
		}
		b.Normal("}")
	}
}

func (vd *ValueDiffer) writeElem(idx int, v reflect.Value) {
	b := vd.bufi(idx)
	if !v.IsValid() {
		b.Normal(nil)
		return
	}
	switch v.Kind() {
	case reflect.Interface:
		if v.IsNil() {
			b.Normal(nil)
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
	b := vd.bufi(idx)
	tp, id, ml := attrElemArray(v)
	if tp {
		vd.writeType(idx, v.Type(), false)
		b.Normal("{")
		defer b.Normal("}")
	} else {
		b.Normal("[")
		defer b.Normal("]")
	}
	if ml {
		b.Tab++
		defer func() { b.Tab-- }()
	}
	vd.writeElemArrayC(idx, v, tp, id, ml)
}

func (vd *ValueDiffer) writeElemArrayC(idx int, v reflect.Value, tp, id, ml bool) {
	b := vd.bufi(idx)
	p := false
	for i := 0; i < v.Len(); i++ {
		e := v.Index(i)
		t := isNonTrivialElem(e)
		t, p = (t || p || (ml && (id || i == 0))), t
		if i > 0 {
			if tp {
				b.Normal(",")
			}
			if !tp || !t {
				b.Normal(" ")
			}
		}
		if t {
			b.NL()
		}
		if id {
			b.Normal(i, ":")
		}
		vd.writeElem(idx, e)
	}
}

func (vd *ValueDiffer) writeElemSlice(idx int, v reflect.Value) {
	b := vd.bufi(idx)
	if v.IsNil() {
		b.Normal(nil)
		return
	}
	vd.writeElemArray(idx, v)
}

func (vd *ValueDiffer) writeElemMap(idx int, v reflect.Value) {
	b := vd.bufi(idx)
	if v.IsNil() {
		b.Normal(nil)
		return
	}
	tp, ml := attrElemMap(v)
	if tp {
		vd.writeType(idx, v.Type(), false)
		b.Normal("{")
		defer b.Normal("}")
	} else {
		b.Normal("map[")
		defer b.Normal("]")
	}
	if ml {
		b.Tab++
		defer func() { b.Tab-- }()
	}
	vd.writeElemMapC(idx, v, tp, ml)
}

func (vd *ValueDiffer) writeElemMapC(idx int, v reflect.Value, tp, ml bool) {
	b := vd.bufi(idx)
	for i, k := range v.MapKeys() {
		if i > 0 {
			if tp {
				b.Normal(",")
			}
			if !ml {
				b.Normal(" ")
			}
		}
		if ml {
			b.NL()
		}
		vd.writeKey(idx, k)
		b.Normal(":")
		vd.writeElem(idx, v.MapIndex(k))
	}
}

func (vd *ValueDiffer) writeElemStruct(idx int, v reflect.Value) {
	b := vd.bufi(idx)
	if ml := attrElemStruct(v); ml {
		b.Normal(structName(v.Type()))
		vd.writeElemStructML(idx, v)
	} else {
		vd.writeKeyStruct(idx, v)
	}
}

func (vd *ValueDiffer) writeElemStructML(idx int, v reflect.Value) {
	b := vd.bufi(idx)
	b.Normal("{")
	b.Tab++
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		if i > 0 {
			b.Normal(",")
		}
		b.NL().Normal(t.Field(i).Name, ":")
		vd.writeElem(idx, v.Field(i))
	}
	b.Tab--
	b.NL().Normal("}")
}

func (vd *ValueDiffer) writeKey(idx int, v reflect.Value) {
	b := vd.bufi(idx)
	if !v.IsValid() {
		b.Normal(nil)
	} else {
		switch v.Kind() {
		case reflect.Uintptr, reflect.String:
			b.Normalf("%#v", v)
		case reflect.Array:
			vd.writeKeyArray(idx, v)
		case reflect.Interface:
			if v.IsNil() {
				b.Normal(nil)
			} else {
				vd.writeKey(idx, v.Elem())
			}
		case reflect.Map:
			vd.writeKeyMap(idx, v)
		case reflect.Ptr:
			if v.IsNil() {
				b.Normal(nil)
			} else {
				b.Normalf("%#v", v.Pointer())
			}
		case reflect.Slice:
			vd.writeKeySlice(idx, v)
		case reflect.Struct:
			vd.writeKeyStruct(idx, v)
		default: // bool, integer, float, complex, channel, function, pointer
			b.Normal(v)
		}
	}
}

func (vd *ValueDiffer) writeKeyArray(idx int, v reflect.Value) {
	b := vd.bufi(idx)
	b.Normal("[")
	for i := 0; i < v.Len(); i++ {
		if i > 0 {
			b.Normal(" ")
		}
		vd.writeKey(idx, v.Index(i))
	}
	b.Normal("]")
}

func (vd *ValueDiffer) writeKeySlice(idx int, v reflect.Value) {
	b := vd.bufi(idx)
	if v.IsNil() {
		b.Normal(nil)
		return
	}
	vd.writeKeyArray(idx, v)
}

func (vd *ValueDiffer) writeKeyMap(idx int, v reflect.Value) {
	b := vd.bufi(idx)
	if v.IsNil() {
		b.Normal(nil)
		return
	}
	b.Normal("map[")
	for i, k := range v.MapKeys() {
		if i > 0 {
			b.Normal(" ")
		}
		vd.writeKey(idx, k)
		b.Normal(":")
		vd.writeKey(idx, v.MapIndex(k))
	}
	b.Normal("]")
}

func (vd *ValueDiffer) writeKeyStruct(idx int, v reflect.Value) {
	b := vd.bufi(idx)
	b.Normal("{")
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		if i > 0 {
			b.Normal(" ")
		}
		b.Normal(t.Field(i).Name, ":")
		vd.writeKey(idx, v.Field(i))
	}
	b.Normal("}")
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
