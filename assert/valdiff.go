package assert

import (
	"fmt"
	"reflect"
)

const (
	NewLine = iota
	Omit
	CompFunc
)

type ValueDiffer struct {
	b     [2]FeatureBuf
	Attrs map[int]bool
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
		vd.writeArrayHTypeValue(idx, v)
	case reflect.Chan, reflect.Func, reflect.UnsafePointer:
		b.Write("(").Highlight(v.Type())
		if v.IsNil() {
			b.Write(")(nil)")
		} else {
			b.Writef(")(%v)", v)
		}
	case reflect.Interface:
		panic(fmt.Sprintf("Should not come here, v=%T(%#v)", v))
	case reflect.Map:
		vd.writeMapHTypeValue(idx, v)
	case reflect.Ptr:
		vd.writePtrHTypeValue(idx, v)
	case reflect.Slice:
		vd.writeSliceHTypeValue(idx, v)
	case reflect.Struct:
		vd.writeStructHTypeValue(idx, v)
	default: // bool, integer, float
		b.Highlight(v.Type).Writef("(%v)", v)
	}
}

func (vd *ValueDiffer) writeArrayHTypeValue(idx int, v reflect.Value) {
	b := vd.b[idx]
	b.Highlight(v.Type()).Write("{")
	if v.Len() > 0 {
		id := v.Len() > 10
		if isComposite(v.Index(0).Kind()) {
			b.Tab++
			for i := 0; i < v.Len(); i++ {
				b.NL()
				if id {
					b.Write(i, ":")
				}
				vd.writeValue(idx, v.Index(i))
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
				vd.writeValue(idx, v.Index(i))
			}
		}
	}
	b.Write("}")
}

func (vd *ValueDiffer) writeSliceHTypeValue(idx int, v reflect.Value) {
}

func (vd *ValueDiffer) writeMapHTypeValue(idx int, v reflect.Value) {
}

func (vd *ValueDiffer) writePtrHTypeValue(idx int, v reflect.Value) {
}

func (vd *ValueDiffer) writeStructHTypeValue(idx int, v reflect.Value) {
}

func (vd *ValueDiffer) writeValue(idx int, v reflect.Value) {
	b := &vd.b[idx]
	if !v.IsValid() {
		b.Write(v)
	} else {
		switch v.Kind() {
		case reflect.Uintptr, reflect.String:
			b.Writef("%#v", v)
		case reflect.Array:
			vd.writeArrayValue(idx, v)
		case reflect.Interface:
			panic(fmt.Sprintf("Should not come here, v=%T(%#v)", v))
		case reflect.Map:
			vd.writeMapValue(idx, v)
		case reflect.Slice:
			vd.writeSliceValue(idx, v)
		case reflect.Struct:
			vd.writeStructValue(idx, v)
		default: // bool, integer, float, complex, channel, function, pointer
			b.Write(v)
		}
	}
}

func (vd *ValueDiffer) writeArrayValue(idx int, v reflect.Value) {
	b := &vd.b[idx]
	if v.Len() < 1 {
		b.Write("[]")
		return
	}
	id := v.Len() > 10
	if isComposite(v.Index(0).Kind()) {
		b.Write(v.Type(), "{")
		b.Tab++
		for i := 0; i < v.Len(); i++ {
			b.NL()
			if id {
				b.Write(i, ":")
			}
			vd.writeValue(idx, v.Index(i))
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
			vd.writeValue(idx, v.Index(i))
		}
		b.Write("}")
	} else {
		b.Write("[")
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				b.Write(", ")
			}
			vd.writeValue(idx, v.Index(i))
		}
		b.Write("]")
	}
}

func (vd *ValueDiffer) writeSliceValue(idx int, v reflect.Value) {
	b := &vd.b[idx]
	if v.IsNil() {
		b.Write(v)
		return
	}
	vd.writeArrayValue(idx, v)
}

func (vd *ValueDiffer) writeMapValue(idx int, v reflect.Value) {
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
	if isComposite(keys[0].Kind()) || isComposite(v.MapIndex(keys[0]).Kind()) {
		b.Tab++
		for _, k := range keys {
			b.NL()
			vd.writeKey(idx, k)
			b.Write(":")
			vd.writeValue(idx, v.MapIndex(k))
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
			vd.writeValue(idx, v.MapIndex(k))
		}
	}
	b.Write("}")
}

func (vd *ValueDiffer) writeStructValue(idx int, v reflect.Value) {
	b := &vd.b[idx]
	b.Write(structName(v), "{")
	comp := false
	for i := 0; i < v.NumField(); i++ {
		if isComposite(v.Field(i).Kind()) {
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

func (vd *ValueDiffer) writeKey(idx int, v reflect.Value) {
	b := &vd.b[idx]
	if !v.IsValid() {
		b.Write(nil)
	} else {
		switch v.Kind() {
		case reflect.Uintptr, reflect.String:
			b.Writef("%#v", v)
		case reflect.Array:
			vd.writeArrayKey(idx, v)
		case reflect.Interface:
			panic(fmt.Sprintf("Should not come here, v=%T(%#v)", v))
		case reflect.Map:
			vd.writeMapKey(idx, v)
		case reflect.Ptr:
			vd.writePtrKey(idx, v)
		case reflect.Slice:
			vd.writeSliceKey(idx, v)
		case reflect.Struct:
			vd.writeStructKey(idx, v)
		default: // bool, integer, float, complex, channel, function, pointer
			b.Write(v)
		}
	}
}

func (vd *ValueDiffer) writeArrayKey(idx int, v reflect.Value) {
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

func (vd *ValueDiffer) writePtrKey(idx int, v reflect.Value) {
	b := &vd.b[idx]
	if v.IsNil() {
		b.Write(v)
		return
	}
	// TODO: isComposite?
	e := v.Elem()
	switch e.Kind() {
	case reflect.Array, reflect.Slice:
		b.Write("&")
		vd.writeArrayKey(idx, e)
	case reflect.Map:
		b.Write("&")
		vd.writeMapKey(idx, e)
	case reflect.Struct:
		b.Write("&")
		vd.writeStructKey(idx, e)
	default:
		b.Write(v)
	}
}

func (vd *ValueDiffer) writeSliceKey(idx int, v reflect.Value) {
	b := &vd.b[idx]
	if v.IsNil() {
		b.Write(v)
		return
	}
	vd.writeArrayKey(idx, v)
}

func (vd *ValueDiffer) writeMapKey(idx int, v reflect.Value) {
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

func (vd *ValueDiffer) writeStructKey(idx int, v reflect.Value) {
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

func isComposite(k reflect.Kind) bool {
	return k == reflect.Array || k == reflect.Map || k == reflect.Slice || k == reflect.Struct
}

func structName(v reflect.Value) string {
	n := v.Type().Name()
	if n != "" {
		return n
	}
	return "struct"
}
