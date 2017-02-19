package assert

import "reflect"

func (vd *tValueDiffer) WriteValDiff(v1, v2 reflect.Value, tab int) {
	b1, b2 := vd.bufs()
	b1.Tab, b2.Tab = tab, tab
	vd.writeValDiff(v1, v2)
}

func (vd *tValueDiffer) writeValDiff(v1, v2 reflect.Value) {
	// TODO
	if !v1.IsValid() || !v2.IsValid() || v1.Type() != v2.Type() {
		vd.writeDiffTypeValues(v1, v2)
	} else {
		vd.writeTypeDiffValues(v1, v2)
	}
}
