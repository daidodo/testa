package assert

import (
	"reflect"
	"runtime"
	"strings"
	"testing"
	"unsafe"
)

func True(t *testing.T, actual bool, messages ...interface{}) {
	if actual {
		return
	}
	fail(t, true, actual, messages...)
}

func False(t *testing.T, actual bool, messages ...interface{}) {
	if !actual {
		return
	}
	fail(t, false, actual, messages...)
}

func Equal(t *testing.T, expected, actual interface{}, messages ...interface{}) {
	if reflect.DeepEqual(expected, actual) {
		return
	}
	fail(t, expected, actual, messages...)
}

func fail(t *testing.T, expected, actual interface{}, messages ...interface{}) {
	var buf FeatureBuf
	if kHOOK {
		buf.Write("\t")
		buf.Tab = 2
	} else {
		buf.Write("\n")
	}
	writeCodeInfo(&buf)
	writeVariables(&buf, expected, actual)
	writeMessages(&buf, messages...)
	if kHOOK {
		buf.Write("\n")
		output := buf.Bytes()
		tt := (*common)(unsafe.Pointer(t))
		tt.su.Lock()
		tt.output = output
		tt.su.Unlock()
		t.FailNow()
	} else {
		t.Fatal(buf.String())
	}
}

func writeCodeInfo(buf *FeatureBuf) {
	pc, file, line, ok := runtime.Caller(3)
	if ok {
		buf.Writef("%v:%v:", lastPartOf(file), line)
		if fp := runtime.FuncForPC(pc); fp != nil {
			buf.Writef(" in %v:", lastPartOf(fp.Name()))
		}
	} else {
		buf.Write("???:1:")
	}
}

func writeVariables(buf *FeatureBuf, expected, actual interface{}) {
	e, a := reflect.ValueOf(expected), reflect.ValueOf(actual)
	var b1, b2 FeatureBuf
	b1.Tab = buf.Tab + 1
	b2.Tab = buf.Tab + 1
	nl := false
	if e.Type() != a.Type() {
		nl = writeDiffTypeValues(&b1, &b2, e, a)
	} else {
		nl = writeDiffValues(&b1, &b2, e, a)
	}
	if nl {
		buf.NL().Write("expected:\t")
		buf.Tab++
		buf.NL().Write(b1.String())
		buf.Tab--
		buf.NL().Write("  actual:\t")
		buf.Tab++
		buf.NL().Write(b2.String())
		buf.Tab--
	} else {
		buf.NL().Write("expected:\t")
		buf.Write(b1.String())
		buf.NL().Write("  actual:\t")
		buf.Write(b2.String())
	}
}

func writeMessages(buf *FeatureBuf, messages ...interface{}) {
	if len(messages) < 1 {
		return
	}
	buf.NL()
	if s, ok := messages[0].(string); ok {
		buf.Writef(s, messages[1:]...)
	} else {
		buf.Write(messages...)
	}
}

func writeDiffTypeValues(b1, b2 *FeatureBuf, v1, v2 reflect.Value) bool {
	return false
	//switch v1.Kind() {
	//case reflect.Array:

	//default:
	//    b1.Highlight(v1.Type()).Writef("(%v)", valueString(v1))
	//    b2.Highlight(v2.Type()).Writef("(%v)", valueString(v2))
	//}
}

//func typeValueString(val reflect.Value) string {
//    var b FeatureBuf
//    switch val.Kind() {
//    case reflect.Complex64, reflect.Complex128:
//        b.Highlightf("%T", val.Interface()).Writef("%v", valueString(val))
//    case reflect.Array:
//        b.Highlightf("%T", val.Interface()).Write("{")
//        idx := val.Len() > 16
//        for i := 0; i < val.Len(); i++ {
//            if i > 0 {
//                b.Write(", ")
//            }
//            if idx {
//                b.Writef("%v:", i)
//            }
//            b.Writef("%v", valueString(val.Index(i)))
//        }
//        b.Write("}")
//TODO
//    default:
//        b.Highlightf("%T", val.Interface()).Writef("(%v)", valueString(val))
//    }
//    return b.String()
//}

func writeDiffValues(b1, b2 *FeatureBuf, v1, v2 reflect.Value) bool {
	switch v1.Kind() {
	case reflect.Complex64, reflect.Complex128:
		writeDiffComplexValue(b1, b2, v1, v2)
	case reflect.Array:
		return writeDiffArrayValue(b1, b2, v1, v2)
	default:
		b1.Highlightf("%#v", v1)
		b2.Highlightf("%#v", v2)
	}
	return false
}

func writeDiffComplexValue(b1, b2 *FeatureBuf, v1, v2 reflect.Value) {
	c1, c2 := v1.Complex(), v2.Complex()
	b1.Write("(")
	b2.Write("(")
	if real(c1) == real(c2) {
		b1.Write(real(c1))
		b2.Write(real(c2))
	} else {
		b1.Highlight(real(c1))
		b2.Highlight(real(c2))
	}
	b1.Write("+")
	b2.Write("+")
	if imag(c1) == imag(c2) {
		b1.Write(imag(c1))
		b2.Write(imag(c2))
	} else {
		b1.Highlight(imag(c1))
		b2.Highlight(imag(c2))
	}
	b1.Write(")")
	b2.Write(")")
}

func writeDiffArrayValue(b1, b2 *FeatureBuf, v1, v2 reflect.Value) bool {
	if v1.Len() < 1 || (v1.Len() <= 10 && isSimpleType(v1.Index(0).Kind())) {
		writeDiffArrayShort(b1, b2, v1, v2)
	} else if isSimpleType(v1.Index(0).Kind()) {
		writeDiffArrayLong(b1, b2, v1, v2)
	} else {
		writeDiffArrayComposite(b1, b2, v1, v2)
		return true
	}
	return false
}

func writeDiffArrayShort(b1, b2 *FeatureBuf, v1, v2 reflect.Value) {
	b1.Write("[")
	b2.Write("[")
	for i := 0; i < v1.Len(); i++ {
		e1, e2 := v1.Index(i), v2.Index(i)
		if reflect.DeepEqual(e1.Interface(), e2.Interface()) {
			if i > 0 {
				b1.Write(" ")
				b2.Write(" ")
			}
			writeValue(b1, e1)
			writeValue(b2, e2)
		} else {
			if i > 0 {
				b1.Plain(" ")
				b2.Plain(" ")
			}
			writeDiffValues(b1, b2, e1, e2)
		}
	}
	b1.Write("]")
	b2.Write("]")
}

func writeDiffArrayLong(b1, b2 *FeatureBuf, v1, v2 reflect.Value) {
	b1.Write(v1.Type(), "{")
	b2.Write(v2.Type(), "{")
	for i, j := 0, 0; i < v1.Len(); i++ {
		e1, e2 := v1.Index(i), v2.Index(i)
		if reflect.DeepEqual(e1.Interface(), e2.Interface()) {
			continue
		}
		if j > 0 {
			b1.Write(", ")
			b2.Write(", ")
		}
		j++
		b1.Write(i, ":")
		b2.Write(i, ":")
		writeDiffValues(b1, b2, e1, e2)
	}
	b1.Write("}")
	b2.Write("}")
}

func writeDiffArrayComposite(b1, b2 *FeatureBuf, v1, v2 reflect.Value) {
	b1.Write(v1.Type(), "{")
	b2.Write(v2.Type(), "{")
	b1.Tab++
	b2.Tab++
	idx := v1.Len() > 10
	for i := 0; i < v1.Len(); i++ {
		e1, e2 := v1.Index(i), v2.Index(i)
		eq := reflect.DeepEqual(e1.Interface(), e2.Interface())
		if eq && idx {
			continue
		}
		b1.NL()
		b2.NL()
		if idx {
			b1.Write(i, ":")
			b2.Write(i, ":")
		}
		if eq {
			writeValue(b1, e1)
			writeValue(b2, e2)
		} else {
			writeDiffValues(b1, b2, e1, e2)
		}
	}
	b1.Tab--
	b2.Tab--
	b1.NL().Write("}")
	b2.NL().Write("}")
}

func writeValue(b *FeatureBuf, v reflect.Value) {
	switch v.Kind() {
	case reflect.Array:
		writeArrayValue(b, v)
	default:
		b.Writef("%#v", v)
	}
}

func writeArrayValue(b *FeatureBuf, v reflect.Value) {
	if v.Len() < 1 || (v.Len() <= 10 && isSimpleType(v.Index(0).Kind())) {
		b.Write(v)
	} else if isSimpleType(v.Index(0).Kind()) {
		b.Write(v.Type(), "{")
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				b.Write(", ")
			}
			b.Write(i, ":")
			writeValue(b, v.Index(i))
		}
		b.Write("}")
	} else {
		idx := v.Len() > 10
		b.Write(v.Type(), "{")
		b.Tab++
		for i := 0; i < v.Len(); i++ {
			b.NL()
			if idx {
				b.Write(i, ":")
			}
			writeValue(b, v.Index(i))
		}
		b.Tab--
		b.NL().Write("}")
	}
}

func isSimpleType(k reflect.Kind) bool {
	return !(k == reflect.Array)
}

func lastPartOf(str string) string {
	if index := strings.LastIndex(str, "/"); index >= 0 {
		return str[index+1:]
	} else if index = strings.LastIndex(str, "\\"); index >= 0 {
		return str[index+1:]
	}
	return str
}

/*
	// expected, actual
	t := reflect.TypeOf(expected)
	if t != reflect.TypeOf(actual) {
		return []string{
			fmt.Sprintf("expected:\t%v", toTypeStr(expected)),
			fmt.Sprintf("  actual:\t%v", toTypeStr(actual)),
		}
	}
	switch t.Kind() {
	case reflect.Array:
		e, a := diffStr(expected, actual)
		return []string{
			fmt.Sprintf("expected:\t%v", e),
			fmt.Sprintf("  actual:\t%v", a),
		}
	}
	return []string{
		fmt.Sprintf("expected:\t%v", toStr(expected)),
		fmt.Sprintf("  actual:\t%v", toStr(actual)),
	}

	var content []string
	var m string
	if len(function) > 0 {
		m = fmt.Sprintf("%v:%v: in %v:", file, line, function)
	} else {
		m = fmt.Sprintf("%v:%v:", file, line)
	}
	if kHOOK {
		content = append(content, "\t"+m)
	} else {
		content = append(content, "\n"+m)
	}
	content = append(content, expectAndActual(expected, actual)...)
	if m := formatMsg(messages...); len(m) > 0 {
		content = append(content, m)
	}
	if kHOOK {
		content[len(content)-1] += "\n"
		output := strings.Join(content, "\n\t\t")
		tt := (*common)(unsafe.Pointer(t))
		tt.su.Lock()
		tt.output = append([]byte(nil), output...)
		tt.su.Unlock()
		t.FailNow()
	} else {
		t.Fatal(strings.Join(content, "\n"))
	}
}

func codeInfo() (function, file string, line int) {
	pc, file, line, ok := runtime.Caller(3)
	if ok {
		if fp := runtime.FuncForPC(pc); fp != nil {
			function = lastPartOf(fp.Name())
		}
		file = lastPartOf(file)
	} else {
		file = "???"
		line = 1
	}
	return
}

func expectAndActual(expected, actual interface{}) []string {
	t := reflect.TypeOf(expected)
	if t != reflect.TypeOf(actual) {
		return []string{
			fmt.Sprintf("expected:\t%v", toTypeStr(expected)),
			fmt.Sprintf("  actual:\t%v", toTypeStr(actual)),
		}
	}
	switch t.Kind() {
	case reflect.Array:
		e, a := diffStr(expected, actual)
		return []string{
			fmt.Sprintf("expected:\t%v", e),
			fmt.Sprintf("  actual:\t%v", a),
		}
	}
	return []string{
		fmt.Sprintf("expected:\t%v", toStr(expected)),
		fmt.Sprintf("  actual:\t%v", toStr(actual)),
	}
}

func toTypeStr(val interface{}) string {
	switch reflect.TypeOf(val).Kind() {
	case reflect.Uintptr:
		return fmt.Sprintf("%T(%#v)", val, val)
	}
	return fmt.Sprintf("%T(%v)", val, val)
}

func toStr(val interface{}) string {
	switch reflect.TypeOf(val).Kind() {
	case reflect.Uintptr:
		return fmt.Sprintf("%#v", val)
	}
	return fmt.Sprintf("%v", val)
}

func diffStr(var1, var2 interface{}) (string, string) {
	v1, v2 := reflect.ValueOf(var1), reflect.ValueOf(var2)
	if v1.Type() != v2.Type() {
		panic(fmt.Sprintf("Should be the same type, but var1 is %T, var2 is %T", var1, var2))
	}
	var b1, b2 []byte
	switch v1.Kind() {
	case reflect.Array:
		for i := 0; i < v1.Len(); i++ { // v1.Len() == v2.Len()
			e1, e2 := v1.Index(i), v2.Index(i)
			if reflect.DeepEqual(e1.Interface(), e2.Interface()) {
				b1 = appendValueStr(b1, toStr(e1), false)
				b2 = appendValueStr(b2, toStr(e2), false)
			} else {
				b1 = appendValueStr(b1, toStr(e1), true)
				b2 = appendValueStr(b2, toStr(e2), true)
			}
		}
		return string(b1), string(b2)
	}
	return toStr(var1), toStr(var2)
}

func appendValueStr(buf []byte, str string, highlight bool) []byte {
	if len(str) < 1 {
		return buf
	}
	sp := len(buf) > 0
	if highlight {
		buf = append(buf, 033, '[', '3', '1', 'm')
	}
	if sp {
		buf = append(buf, ' ')
	}
	buf = append(buf, str...)
	if highlight {
		buf = append(buf, 033, '[', '0', 'm')
	}
	return buf
}

func formatMsg(m ...interface{}) string {
	if len(m) < 1 {
		return ""
	}
	if s, ok := m[0].(string); ok {
		return fmt.Sprintf(s, m[1:]...)
	}
	return fmt.Sprint(m...)
}
*/
