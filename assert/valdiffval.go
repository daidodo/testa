package assert

import "reflect"

func (vd *tValueDiffer) WriteValDiff(v1, v2 reflect.Value, tab int) {
	b1, b2 := vd.bufs()
	b1.Tab, b2.Tab = tab, tab
	vd.writeValDiff(v1, v2, false)
}

func (vd *tValueDiffer) writeValDiff(v1, v2 reflect.Value, sw bool) {
	b1, b2, i1, i2 := vd.bufr(sw)
	v1, d1 := derefInterface(v1)
	v2, d2 := derefInterface(v2)
	if d1 || d2 {
		vd.writeValDiff(v1, v2, sw)
		return
	}
	if !v1.IsValid() {
		b1.Highlight(nil)
		vd.writeValue(i2, v2, false, false)
		return
	} else if !v2.IsValid() {
		vd.writeValue(i1, v1, false, false)
		b2.Highlight(nil)
		return
	}
	switch v2.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		vd.writeValDiffInt(v1, v2, sw)
		return
	default:
		vd.writeElem(i1, v1, true)
		vd.writeElem(i2, v2, true)
	}
	vd.writeValDiffC(v1, v2, sw)
}

func (vd *tValueDiffer) writeValDiffC(v1, v2 reflect.Value, sw bool) {
	_, _, i1, i2 := vd.bufr(sw)
	if t1, t2 := v1.Type(), v2.Type(); t1.ConvertibleTo(t2) || t2.ConvertibleTo(t1) {
		vd.writeValue(i1, v1, false, true)
		vd.writeValue(i2, v2, false, true)
	} else {
		vd.writeValue(i1, v1, true, false)
		vd.writeValue(i2, v2, true, false)
	}
}

func (vd *tValueDiffer) writeValue(idx int, v reflect.Value, ht, hv bool) {
}

func (vd *tValueDiffer) bufr(sw bool) (b1, b2 *tFeatureBuf, i1, i2 int) {
	if sw {
		return vd.bufi(1), vd.bufi(0), 1, 0
	}
	return vd.bufi(0), vd.bufi(1), 0, 1
}
