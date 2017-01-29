package assert

import "reflect"

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
	b := vd.b[idx]
	if !v.IsValid() {
		b.Highlight(nil)
		return
	}
	switch v.Kind() {
	case reflect.Uintptr, reflect.String:
		b.Highlight(v.Type()).Writef("(%#v)", v)
	case reflect.Complex64, reflect.Complex128:
		b.Highlight(v.Type()).Writef("%v", v)
	case reflect.Array, reflect.Slice:
		vd.writeArrayHType(idx, v)
	case reflect.Chan, reflect.Func, reflect.UnsafePointer:
		b.Write("(").Highlight(v.Type())
		if v.IsNil() {
			b.Write(")(nil)")
		} else {
			b.Writef(")(%v)", v)
		}
	case reflect.Interface:
		panic("Should not com here")
	case reflect.Map:
		vd.writeMapHType(idx, v)
	case reflect.Ptr:
		vd.writePtrHType(idx, v)
	case reflect.Struct:
		vd.writeStructHType(idx, v)
	default: // bool, signed/unsigned integers, floats
		b.Highlight(v.Type).Writef("(%v)", v)
	}
}

func writeArrayHType(idx int, v reflect.Value) {

}

func writeMapHType(idx int, v reflect.Value) {

}

func writePtrHType(idx int, v reflect.Value) {

}

func writeStructHType(idx int, v reflect.Value) {

}

func (vd *ValueDiffer) writeTypeDiffValues(v1, v2 reflect.Value) {
}
