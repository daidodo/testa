/*
* Copyright (c) 2017 Zhao DAI <daidodo@gmail.com>
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
* http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package assert

import "reflect"

func valueEqual(v1, v2 reflect.Value) bool {
	if !v1.IsValid() || !v2.IsValid() {
		return v1.IsValid() == v2.IsValid()
	}
	if v1.CanInterface() && v2.CanInterface() {
		return reflect.DeepEqual(v1.Interface(), v2.Interface())
	}
	v1, d1 := derefInterface(v1)
	v2, d2 := derefInterface(v2)
	if d1 || d2 {
		return valueEqual(v1, v2)
	}
	if v1.Type() != v2.Type() {
		return false
	}
	switch v1.Kind() {
	case reflect.Bool:
		return v1.Bool() == v2.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v1.Int() == v2.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v1.Uint() == v2.Uint()
	case reflect.Float32, reflect.Float64:
		return v1.Float() == v2.Float()
	case reflect.Complex64, reflect.Complex128:
		return v1.Complex() == v2.Complex()
	case reflect.String:
		return v1.String() == v2.String()
	case reflect.Chan, reflect.UnsafePointer:
		return v1.Pointer() == v2.Pointer()
	case reflect.Func:
		return v1.IsNil() && v2.IsNil()
	case reflect.Ptr:
		if v1.IsNil() || v2.IsNil() {
			return v1.IsNil() && v2.IsNil()
		}
		if v1.Pointer() == v2.Pointer() {
			return true
		}
		return valueEqual(v1.Elem(), v2.Elem())
	case reflect.Interface:
		if v1.IsNil() || v2.IsNil() {
			return v1.IsNil() == v2.IsNil()
		}
		return valueEqual(v1.Elem(), v2.Elem())
	case reflect.Slice:
		if v1.IsNil() != v2.IsNil() {
			return false
		}
		if v1.Len() != v2.Len() {
			return false
		}
		if v1.Pointer() == v2.Pointer() {
			return true
		}
		fallthrough
	case reflect.Array:
		for i := 0; i < v1.Len(); i++ {
			if !valueEqual(v1.Index(i), v2.Index(i)) {
				return false
			}
		}
		return true
	case reflect.Map:
		if v1.IsNil() != v2.IsNil() {
			return false
		}
		if v1.Len() != v2.Len() {
			return false
		}
		if v1.Pointer() == v2.Pointer() {
			return true
		}
		for _, k := range v1.MapKeys() {
			if e1, e2 := v1.MapIndex(k), v2.MapIndex(k); !valueEqual(e1, e2) {
				return false
			}
		}
		return true
	case reflect.Struct:
		for i, n := 0, v1.NumField(); i < n; i++ {
			if !valueEqual(v1.Field(i), v2.Field(i)) {
				return false
			}
		}
		return true
	default: // reflect.Invalid
	}
	return false
}

func isNil(a interface{}) bool {
	if a == nil {
		return true
	}
	return isNilForValue(reflect.ValueOf(a))
}

func isNilForValue(v reflect.Value) bool {
	if !v.IsValid() {
		return true
	}
	if isPointer(v.Type()) {
		return v.Pointer() == 0
	} else if k := v.Kind(); k != reflect.Array && k != reflect.Struct && isNonTrivial(v.Type()) {
		return v.IsNil()
	}
	return false
}

func isSameInValue(e, a interface{}) bool {
	if reflect.DeepEqual(e, a) {
		return true
	}
	if e == nil || a == nil {
		return isNil(e) && isNil(a)
	}
	return convertCompare(reflect.ValueOf(e), reflect.ValueOf(a))
}

func convertCompare(v1, v2 reflect.Value) bool {
	v1, _ = derefInterface(v1)
	v2, _ = derefInterface(v2)
	if !v1.IsValid() || !v2.IsValid() {
		return isNilForValue(v1) && isNilForValue(v2)
	}
	return convertCompareB(v1, v2) || convertCompareB(v2, v1)
}

func convertCompareB(f, t reflect.Value) bool {
	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return convertCompareInt(f, t)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return convertCompareUint(f, t)
	case reflect.Float32, reflect.Float64:
		return convertCompareFloat(f, t)
	case reflect.Complex64, reflect.Complex128:
		return convertCompareComplex(f, t)
	case reflect.String:
		return convertCompareString(f, t)
	case reflect.Ptr, reflect.UnsafePointer:
		return convertComparePtr(f, t)
	case reflect.Array, reflect.Slice:
		return convertCompareArray(f, t)
	case reflect.Map:
		return convertCompareMap(f, t)
	case reflect.Struct:
		return convertCompareStruct(f, t)
	}
	return convertCompareC(f, t)
}

func convertCompareInt(f, t reflect.Value) bool {
	switch f.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return f.Int() == t.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return convertCompareUint(t, f)
	case reflect.Float32, reflect.Float64:
		return convertCompareFloat(t, f)
	case reflect.Complex64, reflect.Complex128:
		return convertCompareComplex(t, f)
	}
	return convertCompareC(f, t)
}

func convertCompareUint(f, t reflect.Value) bool {
	v := t.Uint()
	switch f.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return f.Int() >= 0 && uint64(f.Int()) == v
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return f.Uint() == v
	case reflect.Float32, reflect.Float64:
		return convertCompareFloat(t, f)
	case reflect.Complex64, reflect.Complex128:
		return convertCompareComplex(t, f)
	}
	return convertCompareC(f, t)
}

func convertCompareFloat(f, t reflect.Value) bool {
	v := t.Float()
	switch f.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(f.Int()) == v
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return float64(f.Uint()) == v
	case reflect.Float32, reflect.Float64:
		return f.Float() == v
	case reflect.Complex64, reflect.Complex128:
		return convertCompareComplex(t, f)
	}
	return convertCompareC(f, t)
}

func convertCompareComplex(f, t reflect.Value) bool {
	v := t.Complex()
	switch f.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return imag(v) == 0 && float64(f.Int()) == real(v)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return imag(v) == 0 && float64(f.Uint()) == real(v)
	case reflect.Float32, reflect.Float64:
		return imag(v) == 0 && f.Float() == real(v)
	case reflect.Complex64, reflect.Complex128:
		return f.Complex() == v
	}
	return convertCompareC(f, t)
}

func convertCompareString(f, t reflect.Value) bool {
	switch f.Kind() {
	case reflect.String:
		return f.String() == t.String()
	case reflect.Array, reflect.Slice:
		if e := f.Type().Elem().Kind(); e == reflect.Uint8 {
			return convertCompareArray(f, reflect.ValueOf([]byte(t.String())))
		} else if e == reflect.Int32 {
			return convertCompareArray(f, reflect.ValueOf([]rune(t.String())))
		}
	}
	return convertCompareC(f, t)
}

func convertComparePtr(f, t reflect.Value) bool {
	v := t.Pointer()
	switch f.Kind() {
	case reflect.UnsafePointer:
		return f.Pointer() == v
	case reflect.Ptr:
		return t.Kind() == reflect.UnsafePointer && f.Pointer() == v // diff type pointers are NOT equal
	}
	return convertCompareC(f, t)
}

func convertCompareArray(f, t reflect.Value) bool {
	if t.Kind() != reflect.Slice || !t.IsNil() {
		switch f.Kind() {
		case reflect.Slice:
			if f.IsNil() {
				break
			}
			fallthrough
		case reflect.Array:
			if f.Len() != t.Len() {
				return false
			}
			if f.Len() == 0 {
				return convertible(f.Type().Elem(), t.Type().Elem())
			}
			for i := 0; i < f.Len(); i++ {
				if !convertCompare(f.Index(i), t.Index(i)) {
					return false
				}
			}
			return true
		}
	}
	return convertCompareC(f, t)
}

// NOTE: Map keys must be exactly equal (both type and value), e.g. int(100) and uint(100) are
// different keys.
func convertCompareMap(f, t reflect.Value) bool {
	if !t.IsNil() && f.Kind() == reflect.Map && !f.IsNil() {
		if f.Len() != t.Len() || !convertibleKeyTo(f.Type().Key(), t.Type().Key()) {
			return false
		}
		if f.Len() == 0 {
			return convertible(f.Type().Elem(), t.Type().Elem())
		}
		ks := t.MapKeys()
		find := func(k1 reflect.Value) bool {
			for _, k2 := range ks {
				if valueEqual(k1, k2) {
					return true
				}
			}
			return false
		}
		for _, k := range f.MapKeys() {
			if !find(k) {
				return false
			}
			if !convertCompare(f.MapIndex(k), t.MapIndex(k)) {
				return false
			}
		}
		return true
	}
	return convertCompareC(f, t)
}

func convertCompareStruct(f, t reflect.Value) bool {
	if f.Type() == t.Type() {
		for i := 0; i < f.NumField(); i++ {
			if !convertCompare(f.Field(i), t.Field(i)) {
				return false
			}
		}
		return true
	}
	return convertCompareC(f, t)
}

func convertCompareC(f, t reflect.Value) bool {
	if !f.Type().ConvertibleTo(t.Type()) {
		return false
	}
	a := f.Convert(t.Type())
	return valueEqual(a, t)
}
