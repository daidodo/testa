package assert

import "reflect"

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

func convertCompareMap(f, t reflect.Value) bool {
	if !t.IsNil() && f.Kind() == reflect.Map && !f.IsNil() {
		if f.Len() != t.Len() {
			return false
		}
		if f.Len() == 0 {
			return convertible(f.Type().Key(), t.Type().Key()) &&
				convertible(f.Type().Elem(), t.Type().Elem())
		}
		ks := t.MapKeys()
		find := func(v reflect.Value) (reflect.Value, bool) {
			for _, k := range ks {
				if convertCompare(v, k) {
					return k, true
				}
			}
			return reflect.Value{}, false
		}
		for _, k := range f.MapKeys() {
			kk, ok := find(k)
			if !ok {
				return false
			}
			if !convertCompare(f.MapIndex(k), t.MapIndex(kk)) {
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
