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

func isInteger(t reflect.Type) bool {
	if t == nil {
		return false
	}
	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	}
	return false
}

func isUInteger(t reflect.Type) bool {
	if t == nil {
		return false
	}
	switch t.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return true
	}
	return false
}

func isSimpleNumber(t reflect.Type) bool {
	if t == nil {
		return false
	}
	return isInteger(t) || isUInteger(t) || t.Kind() == reflect.Float32 || t.Kind() == reflect.Float64
}

func isMath(t reflect.Type) bool {
	if t == nil {
		return false
	}
	return isSimpleNumber(t) || t.Kind() == reflect.Complex64 || t.Kind() == reflect.Complex128
}

func isSimplePointer(t reflect.Type) bool {
	return t.Kind() == reflect.Ptr || t.Kind() == reflect.UnsafePointer
}

func isPointer(t reflect.Type) bool {
	if t == nil {
		return false
	}
	return isSimplePointer(t) || t.Kind() == reflect.Chan || t.Kind() == reflect.Func
}

func isArray(t reflect.Type) bool {
	if t == nil {
		return false
	}
	return t.Kind() == reflect.Array || t.Kind() == reflect.Slice
}

func isCharacter(t reflect.Type) bool {
	if t == nil {
		return false
	}
	return t.Kind() == reflect.Uint8 || t.Kind() == reflect.Int32
}

func isString(t reflect.Type) bool {
	if t == nil {
		return false
	}
	return t.Kind() == reflect.String || (isArray(t) && isCharacter(t.Elem()))
}
func isComposite(t reflect.Type) bool {
	if t == nil {
		return false
	}
	return isArray(t) || t.Kind() == reflect.Map || t.Kind() == reflect.Struct
}

func isNonTrivial(t reflect.Type) bool {
	if t == nil {
		return false
	}
	k := t.Kind()
	return k == reflect.Interface || isComposite(t)
}

func isReference(t reflect.Type) bool {
	return isPointer(t) || isNonTrivial(t)
}

func convertible(t1, t2 reflect.Type) bool {
	if t1 == nil || t2 == nil {
		return t1 == t2
	}
	k1, k2 := t1.Kind(), t2.Kind()
	if isMath(t1) && isMath(t2) {
		return true
	} else if isArray(t1) && isArray(t2) {
		if k1 == reflect.Array && k2 == reflect.Array && t1.Len() != t2.Len() {
			return false
		}
		return convertible(t1.Elem(), t2.Elem())
	} else if isSimplePointer(t1) && isSimplePointer(t2) {
		return t1 == t2 || t1.Kind() == reflect.UnsafePointer || t2.Kind() == reflect.UnsafePointer
	} else if isString(t1) && isString(t2) {
		return true
	} else if t1.Kind() == reflect.Map && t2.Kind() == reflect.Map {
		return convertible(t1.Key(), t2.Key()) && convertible(t1.Elem(), t2.Elem())
	}
	return t1.ConvertibleTo(t2) || t2.ConvertibleTo(t1)
}

func convertibleKeyTo(f, t reflect.Type) bool {
	if f == t {
		return true
	}
	if f == nil || t == nil || t.Kind() != reflect.Interface {
		return false
	}
	return f.ConvertibleTo(t)
}

func mapKeyDiff(v1, v2 reflect.Value) (ks, ks1, ks2 []reflect.Value) {
	if t1, t2 := v1.Type().Key(), v2.Type().Key(); t1 == t2 {
		for _, k := range v1.MapKeys() {
			if v2.MapIndex(k).IsValid() {
				ks = append(ks, k)
			} else {
				ks1 = append(ks1, k)
			}
		}
		for _, k := range v2.MapKeys() {
			if !v1.MapIndex(k).IsValid() {
				ks2 = append(ks2, k)
			}
		}
	} else if convertibleKeyTo(t1, t2) {
		s1, s2 := v1.MapKeys(), v2.MapKeys()
		find := func(k1 reflect.Value, s []reflect.Value) bool {
			for _, k2 := range s {
				if valueEqual(k1, k2) {
					return true
				}
			}
			return false
		}
		for _, k := range s1 {
			if find(k, s2) {
				ks = append(ks, k)
			} else {
				ks1 = append(ks1, k)
			}
		}
		for _, k := range s2 {
			if !find(k, s1) {
				ks2 = append(ks2, k)
			}
		}
	} else if convertibleKeyTo(t2, t1) {
		ks, ks2, ks1 = mapKeyDiff(v2, v1)
	} else {
		ks1, ks2 = v1.MapKeys(), v2.MapKeys()
	}
	return
}

func derefInterface(v reflect.Value) (r reflect.Value, d bool) {
	if v.IsValid() && v.Kind() == reflect.Interface {
		if !v.IsNil() {
			return v.Elem(), true
		} else if v.Type().Name() == "" {
			return r, true
		}
	}
	return v, d
}

func derefPtr(v reflect.Value) (r reflect.Value, d bool) {
	if v.IsValid() && v.Kind() == reflect.Ptr && !v.IsNil() {
		if e := v.Elem(); isComposite(e.Type()) {
			return e, true
		}
	}
	return v, false
}
