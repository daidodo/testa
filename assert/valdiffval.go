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

import "reflect"

func (vd *tValueDiffer) WriteValDiff(v1, v2 reflect.Value, tab int) {
	b1, b2 := vd.bufs()
	b1.Tab, b2.Tab = tab, tab
	vd.writeValDiff(v1, v2, false)
}

func (vd *tValueDiffer) writeValDiff(v1, v2 reflect.Value, sw bool) {
	v1, d1 := derefInterface(v1)
	v2, d2 := derefInterface(v2)
	if d1 || d2 {
		vd.writeValDiff(v1, v2, sw)
		return
	}
	if !vd.writeValDiffToNil(v1, v2, sw) &&
		!vd.writeValDiffNumbers(v1, v2, sw) &&
		!vd.writeValDiffPointers(v1, v2, sw) &&
		!vd.writeValDiffToString(v1, v2, sw) &&
		!vd.writeValDiffToComplex64(v1, v2, sw) &&
		!vd.writeValDiffToComplex128(v1, v2, sw) &&
		!vd.writeValDiffArrays(v1, v2, sw) &&
		!vd.writeValDiffMaps(v1, v2, sw) &&
		!vd.writeValDiffStructs(v1, v2, sw) {
		vd.writeValDiffC(v1, v2, sw)
	}
}

func (vd *tValueDiffer) writeValDiffToNil(v1, v2 reflect.Value, sw bool) bool {
	_, b2, i1, _ := vd.bufr(sw)
	if !v2.IsValid() {
		if t := v1.Type(); isComposite(t) {
			vd.writeTypeValue(i1, v1, false, false)
		} else if isReference(t) {
			vd.writeElem(i1, v1, true)
		} else {
			vd.writeTypeValue(i1, v1, true, false)
		}
		b2.Highlight(nil)
		return true
	} else if !v1.IsValid() {
		return vd.writeValDiffToNil(v2, v1, !sw)
	}
	return false
}

func (vd *tValueDiffer) writeValDiffNumbers(v1, v2 reflect.Value, sw bool) bool {
	_, _, i1, i2 := vd.bufr(sw)
	if isSimpleNumber(v1.Type()) && isSimpleNumber(v2.Type()) {
		if k1, k2 := v1.Kind(), v2.Kind(); k1 == reflect.Uintptr && k2 != reflect.Uintptr {
			vd.writeKeyPOD(i1, v1, true, true)
			vd.writeElem(i2, v2, true)
		} else if k1 != reflect.Uintptr && k2 == reflect.Uintptr {
			vd.writeElem(i1, v1, true)
			vd.writeKeyPOD(i2, v2, true, true)
		} else {
			vd.writeElem(i1, v1, true)
			vd.writeElem(i2, v2, true)
		}
		return true
	}
	return false
}

func (vd *tValueDiffer) writeValDiffPointers(v1, v2 reflect.Value, sw bool) bool {
	b1, b2, i1, i2 := vd.bufr(sw)
	k1, t1, t2 := v1.Kind(), v1.Type(), v2.Type()
	if k1 == reflect.Ptr && t1 == t2 {
		if !v1.IsNil() && !v2.IsNil() {
			b1.Normal("&")
			b2.Normal("&")
			vd.writeValDiff(v1.Elem(), v2.Elem(), sw)
			return true
		}
	}
	if (k1 == reflect.UnsafePointer && isSimplePointer(t2)) ||
		(v2.Kind() == reflect.UnsafePointer && isSimplePointer(t1)) {
		vd.writeElem(i1, v1, true)
		vd.writeElem(i2, v2, true)
		return true
	}
	return false
}

func (vd *tValueDiffer) writeValDiffToString(v1, v2 reflect.Value, sw bool) bool {
	b1, b2, i1, i2 := vd.bufr(sw)
	if v2.Kind() == reflect.String {
		if t := v1.Type(); isInteger(t) || isUInteger(t) {
			vd.writeElem(i1, v1, true)
			vd.writeElem(i2, v2, true)
			return true
		} else if t.Kind() == reflect.String {
			vd.writeDiffValuesString(v1, v2, sw)
			return true
		} else if isArray(t) {
			if k := t.Elem().Kind(); k == reflect.Uint8 {
				b1.Normal("[")
				b2.Normal(`"`)
				s, i, p := []byte(v2.String()), 0, rune(0)
				d := func(j int) bool {
					r := false
					for ; i < j; i++ {
						g := i < v1.Len()
						df := !g || byte(v1.Index(i).Uint()) != s[i]
						if g {
							if i > 0 {
								b1.Plain(" ")
							}
							b1.Writef(df, "%#x", byte(v1.Index(i).Uint()))
						}
						r = r || df
					}
					return r
				}
				for j, c := range v2.String() {
					if j > 0 {
						b2.Writef(d(j), "%c", p)
					}
					p = c
				}
				if i > 0 {
					b2.Writef(d(len(s)), "%c", p)
				}
				for ; i < v1.Len(); i++ {
					if i > 0 {
						b1.Plain(" ")
					}
					b1.Highlightf("%#x", byte(v1.Index(i).Uint()))
				}
				b1.Normal("]")
				b2.Normal(`"`)
				return true
			} else if k == reflect.Int32 {
				b1.Normal("[")
				b2.Normal(`"`)
				s := []rune(v2.String())
				for i := 0; i < len(s) || i < v1.Len(); i++ {
					g1, g2 := i < v1.Len(), i < len(s)
					if g1 && i > 0 {
						b1.Plain(" ")
					}
					eq := g1 && g2 && rune(v1.Index(i).Int()) == s[i]
					if g1 {
						b1.Writef(!eq, "%#x", rune(v1.Index(i).Int()))
					}
					if g2 {
						b2.Writef(!eq, "%c", s[i])
					}
				}
				b1.Normal("]")
				b2.Normal(`"`)
				return true
			}
		}
	} else if v1.Kind() == reflect.String {
		return vd.writeValDiffToString(v2, v1, !sw)
	}
	return false
}

func (vd *tValueDiffer) writeValDiffToComplex64(v1, v2 reflect.Value, sw bool) bool {
	b1, b2, i1, _ := vd.bufr(sw)
	if v2.Kind() == reflect.Complex64 {
		c2 := complex64(v2.Complex())
		r2, m2 := real(c2), imag(c2)
		if t, k := v1.Type(), v1.Kind(); isInteger(t) {
			hr, hi := float32(v1.Int()) != r2, 0 != m2
			vd.writeElem(i1, v1, hr)
			b2.Normal("(").Write(hr, r2).Plain("+").Writef(hi, "%gi", m2).Normal(")")
			return true
		} else if isUInteger(t) {
			hr, hi := float32(v1.Uint()) != r2, 0 != m2
			if k == reflect.Uintptr {
				vd.writeKeyPOD(i1, v1, hr, !hi)
			} else {
				vd.writeElem(i1, v1, hr)
			}
			b2.Normal("(").Write(hr, r2).Plain("+").Writef(hi, "%gi", m2).Normal(")")
			return true
		} else if k == reflect.Float32 {
			hr, hi := float32(v1.Float()) != r2, 0 != m2
			vd.writeElem(i1, v1, hr)
			b2.Normal("(").Write(hr, r2).Plain("+").Writef(hi, "%gi", m2).Normal(")")
			return true
		} else if k == reflect.Float64 {
			hr, hi := v1.Float() != float64(r2), 0 != m2
			vd.writeElem(i1, v1, hr)
			b2.Normal("(").Write(hr, r2).Plain("+").Writef(hi, "%gi", m2).Normal(")")
			return true
		} else if k == reflect.Complex64 {
			c1 := complex64(v1.Complex())
			r1, m1 := real(c1), imag(c1)
			hr, hi := r1 != r2, m1 != m2
			b1.Normal("(").Write(hr, r1).Plain("+").Writef(hi, "%gi", m1).Normal(")")
			b2.Normal("(").Write(hr, r2).Plain("+").Writef(hi, "%gi", m2).Normal(")")
			return true
		} else if k == reflect.Complex128 {
			c1 := v1.Complex()
			r1, m1 := real(c1), imag(c1)
			hr, hi := r1 != float64(r2), m1 != float64(m2)
			b1.Normal("(").Write(hr, r1).Plain("+").Writef(hi, "%gi", m1).Normal(")")
			b2.Normal("(").Write(hr, r2).Plain("+").Writef(hi, "%gi", m2).Normal(")")
			return true
		}
	} else if v1.Kind() == reflect.Complex64 {
		return vd.writeValDiffToComplex64(v2, v1, !sw)
	}
	return false
}

func (vd *tValueDiffer) writeValDiffToComplex128(v1, v2 reflect.Value, sw bool) bool {
	b1, b2, i1, _ := vd.bufr(sw)
	if v2.Kind() == reflect.Complex128 {
		c2 := v2.Complex()
		r2, m2 := real(c2), imag(c2)
		if t, k := v1.Type(), v1.Kind(); isInteger(t) {
			hr, hi := float64(v1.Int()) != r2, 0 != m2
			vd.writeElem(i1, v1, hr)
			b2.Normal("(").Write(hr, r2).Plain("+").Writef(hi, "%gi", m2).Normal(")")
			return true
		} else if isUInteger(t) {
			hr, hi := float64(v1.Uint()) != r2, 0 != m2
			if k == reflect.Uintptr {
				vd.writeKeyPOD(i1, v1, hr, !hi)
			} else {
				vd.writeElem(i1, v1, hr)
			}
			b2.Normal("(").Write(hr, r2).Plain("+").Writef(hi, "%gi", m2).Normal(")")
			return true
		} else if k == reflect.Float32 {
			hr, hi := v1.Float() != r2, 0 != m2
			vd.writeElem(i1, v1, hr)
			b2.Normal("(").Write(hr, r2).Plain("+").Writef(hi, "%gi", m2).Normal(")")
			return true
		} else if k == reflect.Float64 {
			hr, hi := v1.Float() != r2, 0 != m2
			vd.writeElem(i1, v1, hr)
			b2.Normal("(").Write(hr, r2).Plain("+").Writef(hi, "%gi", m2).Normal(")")
			return true
		} else if k == reflect.Complex64 {
			c1 := complex64(v1.Complex())
			r1, m1 := real(c1), imag(c1)
			hr, hi := float64(r1) != r2, float64(m1) != m2
			b1.Normal("(").Write(hr, r1).Plain("+").Writef(hi, "%gi", m1).Normal(")")
			b2.Normal("(").Write(hr, r2).Plain("+").Writef(hi, "%gi", m2).Normal(")")
			return true
		} else if k == reflect.Complex128 {
			c1 := v1.Complex()
			r1, m1 := real(c1), imag(c1)
			hr, hi := r1 != r2, m1 != m2
			b1.Normal("(").Write(hr, r1).Plain("+").Writef(hi, "%gi", m1).Normal(")")
			b2.Normal("(").Write(hr, r2).Plain("+").Writef(hi, "%gi", m2).Normal(")")
			return true
		}
	} else if v1.Kind() == reflect.Complex128 {
		return vd.writeValDiffToComplex128(v2, v1, !sw)
	}
	return false
}

func (vd *tValueDiffer) writeValDiffArrays(v1, v2 reflect.Value, sw bool) bool {
	t1, t2 := v1.Type(), v2.Type()
	k1, k2 := t1.Kind(), t2.Kind()
	if !isArray(t1) || !isArray(t2) ||
		(k1 == reflect.Array && k2 == reflect.Array && t1.Len() != t2.Len()) ||
		(k1 == reflect.Slice && v1.IsNil()) ||
		(k2 == reflect.Slice && v2.IsNil()) ||
		!convertible(t1.Elem(), t2.Elem()) {
		return false
	}
	eq := func(a, b reflect.Value) bool { return convertCompare(a, b) }
	wd := func(a, b reflect.Value, s bool) { vd.writeValDiff(a, b, s) }
	vd.writeTypeDiffValuesArray(v1, v2, sw, eq, wd)
	return true
}

func (vd *tValueDiffer) writeValDiffMaps(v1, v2 reflect.Value, sw bool) bool {
	t1, t2 := v1.Type(), v2.Type()
	if t1.Kind() == reflect.Map && !v1.IsNil() && t2.Kind() == reflect.Map && !v2.IsNil() &&
		(convertibleKeyTo(t1.Key(), t2.Key()) || convertibleKeyTo(t2.Key(), t1.Key())) &&
		convertible(t1.Elem(), t2.Elem()) {
		eq := func(a, b reflect.Value) bool { return convertCompare(a, b) }
		wd := func(a, b reflect.Value, s bool) { vd.writeValDiff(a, b, s) }
		vd.writeTypeDiffValuesMap(v1, v2, sw, eq, wd)
		return true
	}
	return false
}

func (vd *tValueDiffer) writeValDiffStructs(v1, v2 reflect.Value, sw bool) bool {
	if v1.Kind() == reflect.Struct && v1.Type() == v2.Type() {
		eq := func(a, b reflect.Value) bool { return convertCompare(a, b) }
		wd := func(a, b reflect.Value, sw bool) { vd.writeValDiff(a, b, sw) }
		vd.writeTypeDiffValuesStruct(v1, v2, sw, eq, wd)
		return true
	}
	return false
}

func (vd *tValueDiffer) writeValDiffC(v1, v2 reflect.Value, sw bool) {
	_, _, i1, i2 := vd.bufr(sw)
	if t1, t2 := v1.Type(), v2.Type(); convertible(t1, t2) {
		vd.writeTypeValue(i1, v1, false, true)
		vd.writeTypeValue(i2, v2, false, true)
	} else if isComposite(t1) && isComposite(t2) {
		vd.writeDiffKindsBeforeValue(v1, v2, convertible, sw)
		vd.writeValueAfterType(i1, v1, false)
		vd.writeValueAfterType(i2, v2, false)
	} else {
		eq := func(t1, t2 reflect.Type) bool { return t1 == t2 }
		vd.writeDiffKindsBeforeValue(v1, v2, eq, sw)
		vd.writeValueAfterType(i1, v1, false)
		vd.writeValueAfterType(i2, v2, false)
	}
}

func (vd *tValueDiffer) bufr(sw bool) (b1, b2 *tFeatureBuf, i1, i2 int) {
	if sw {
		return vd.bufi(1), vd.bufi(0), 1, 0
	}
	return vd.bufi(0), vd.bufi(1), 0, 1
}
