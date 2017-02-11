package assert

import "reflect"

func (vd *ValueDiffer) WriteTypeValue(idx int, v reflect.Value, tab int) {
	vd.bufi(idx).Tab = tab
	vd.writeTypeValue(idx, v)
}

func (vd *ValueDiffer) writeTypeValue(idx int, v reflect.Value) {
	v = vd.writeTypeBeforeValue(idx, v, false)
	vd.writeValueAfterType(idx, v)
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
			b.Write(hl, nil)
		} else {
			b.Write(hl, n)
		}
	}
	return
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
	vd.writeElemArrayC(idx, v, true, id, ml, false)
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
	vd.writeElemMapC(idx, v, true, ml, false)
}

func (vd *ValueDiffer) writeValueAfterTypeStruct(idx int, v reflect.Value) {
	b := vd.bufi(idx)
	if ml := attrElemStruct(v); ml {
		vd.writeElemStructML(idx, v, false)
	} else {
		b.Normal("{")
		t := v.Type()
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				b.Normal(", ")
			}
			b.Normal(t.Field(i).Name, ":")
			vd.writeKey(idx, v.Field(i), false)
		}
		b.Normal("}")
	}
}

func (vd *ValueDiffer) writeElem(idx int, v reflect.Value, hl bool) {
	b := vd.bufi(idx)
	if !v.IsValid() {
		b.Write(hl, nil)
		return
	}
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

func (vd *ValueDiffer) writeElemArray(idx int, v reflect.Value, hl bool) {
	b := vd.bufi(idx)
	tp, id, ml := attrElemArray(v)
	if tp {
		vd.writeType(idx, v.Type(), hl)
		b.Write(hl, "{")
		defer b.Write(hl, "}")
		if ml {
			b.Tab++
			defer func() { b.Tab--; b.NL() }()
			vd.Attrs[NewLine+idx] = true
		}
	} else {
		b.Write(hl, "[")
		defer b.Write(hl, "]")
	}
	vd.writeElemArrayC(idx, v, tp, id, ml, hl)
}

func (vd *ValueDiffer) writeElemArrayC(idx int, v reflect.Value, tp, id, ml, hl bool) {
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
				b.Plain(" ")
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

func (vd *ValueDiffer) writeElemSlice(idx int, v reflect.Value, hl bool) {
	b := vd.bufi(idx)
	if v.IsNil() {
		b.Write(hl, nil)
		return
	}
	vd.writeElemArray(idx, v, hl)
}

func (vd *ValueDiffer) writeElemMap(idx int, v reflect.Value, hl bool) {
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
			defer func() { b.Tab-- }()
			vd.Attrs[NewLine+idx] = true
		}
	} else {
		b.Write(hl, "map[")
		defer b.Write(hl, "]")
	}
	vd.writeElemMapC(idx, v, tp, ml, hl)
}

func (vd *ValueDiffer) writeElemMapC(idx int, v reflect.Value, tp, ml, hl bool) {
	b := vd.bufi(idx)
	for i, k := range v.MapKeys() {
		if i > 0 {
			if tp {
				b.Write(hl, ",")
			}
			if !ml {
				b.Plain(" ")
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

func (vd *ValueDiffer) writeElemStruct(idx int, v reflect.Value, hl bool) {
	b := vd.bufi(idx)
	if ml := attrElemStruct(v); ml {
		b.Write(hl, structName(v.Type()))
		vd.writeElemStructML(idx, v, hl)
		vd.Attrs[NewLine+idx] = true
	} else {
		vd.writeKeyStruct(idx, v, hl)
	}
}

func (vd *ValueDiffer) writeElemStructML(idx int, v reflect.Value, hl bool) {
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

func (vd *ValueDiffer) writeKey(idx int, v reflect.Value, hl bool) {
	b := vd.bufi(idx)
	if !v.IsValid() {
		b.Write(hl, nil)
	} else {
		switch v.Kind() {
		case reflect.Uintptr, reflect.String:
			b.Writef(hl, "%#v", v)
		case reflect.Array:
			vd.writeKeyArray(idx, v, hl)
		case reflect.Interface:
			if v.IsNil() {
				b.Write(hl, nil)
			} else {
				vd.writeKey(idx, v.Elem(), hl)
			}
		case reflect.Map:
			vd.writeKeyMap(idx, v, hl)
		case reflect.Ptr:
			if v.IsNil() {
				b.Write(hl, nil)
			} else {
				b.Writef(hl, "%#v", v.Pointer())
			}
		case reflect.Slice:
			vd.writeKeySlice(idx, v, hl)
		case reflect.Struct:
			vd.writeKeyStruct(idx, v, hl)
		default: // bool, integer, float, complex, channel, function, unsafe pointer
			b.Write(hl, v)
		}
	}
}

func (vd *ValueDiffer) writeKeyArray(idx int, v reflect.Value, hl bool) {
	b := vd.bufi(idx)
	b.Write(hl, "[")
	for i := 0; i < v.Len(); i++ {
		if i > 0 {
			b.Plain(" ")
		}
		vd.writeKey(idx, v.Index(i), hl)
	}
	b.Plain("]")
}

func (vd *ValueDiffer) writeKeySlice(idx int, v reflect.Value, hl bool) {
	b := vd.bufi(idx)
	if v.IsNil() {
		b.Write(hl, nil)
		return
	}
	vd.writeKeyArray(idx, v, hl)
}

func (vd *ValueDiffer) writeKeyMap(idx int, v reflect.Value, hl bool) {
	b := vd.bufi(idx)
	if v.IsNil() {
		b.Write(hl, nil)
		return
	}
	b.Write(hl, "map[")
	for i, k := range v.MapKeys() {
		if i > 0 {
			b.Plain(" ")
		}
		vd.writeKey(idx, k, hl)
		b.Plain(":")
		vd.writeKey(idx, v.MapIndex(k), hl)
	}
	b.Plain("]")
}

func (vd *ValueDiffer) writeKeyStruct(idx int, v reflect.Value, hl bool) {
	b := vd.bufi(idx)
	b.Write(hl, "{")
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		if i > 0 {
			b.Plain(" ")
		}
		b.Plain(t.Field(i).Name, ":")
		vd.writeKey(idx, v.Field(i), hl)
	}
	b.Plain("}")
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

func isComposite(t reflect.Type) bool {
	if t == nil {
		return false
	}
	k := t.Kind()
	return k == reflect.Array || k == reflect.Map || k == reflect.Slice || k == reflect.Struct
}

func isNonTrivial(t reflect.Type) bool {
	if t == nil {
		return false
	}
	k := t.Kind()
	return k == reflect.Interface || isComposite(t)
}

func isReference(t reflect.Type) bool {
	if t == nil {
		return false
	}
	k := t.Kind()
	return k == reflect.Chan || k == reflect.Func || k == reflect.Ptr || k == reflect.UnsafePointer || isNonTrivial(t)
}

func structName(t reflect.Type) string {
	if t == nil {
		return ""
	} else if t.Name() == "" {
		return "struct"
	}
	return t.String()
}

func interfaceName(t reflect.Type) string {
	if t == nil || t.Name() == "" {
		return ""
	}
	return t.String()
}

func (vd *ValueDiffer) bufi(i int) (b *FeatureBuf) {
	b = &vd.b[i]
	if b.w == nil {
		b.w = &vd.buf[i]
	}
	return
}
