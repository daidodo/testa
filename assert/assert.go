package assert

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"unsafe"
)

func Equal(t *testing.T, expected, actual interface{}, messages ...interface{}) {
	caller{1, 1}.Equal(t, expected, actual, messages...)
}

func Caller(level int) caller {
	return caller{0, level}
}

func (c caller) True(t *testing.T, actual bool, messages ...interface{}) {
	if actual {
		return
	}
	fail(c, t, true, actual, messages...)
}

func (c caller) False(t *testing.T, actual bool, messages ...interface{}) {
	if !actual {
		return
	}
	fail(c, t, false, actual, messages...)
}

func (c caller) Equal(t *testing.T, expected, actual interface{}, messages ...interface{}) {
	if reflect.DeepEqual(expected, actual) {
		return
	}
	fail(c, t, expected, actual, messages...)
}

//func NotEqual(t *testing.T, expected, actual interface{}, messages ...interface{}) {
//    if !reflect.DeepEqual(expected, actual) {
//        return
//    }
// TODO
//    fail(c, t, expected, actual, messages...)
//}

type caller struct {
	from, to int
}

func fail(c caller, t *testing.T, expected, actual interface{}, messages ...interface{}) {
	var buf FeatureBuf
	if kHOOK {
		buf.Write("\t")
		buf.Tab = 2
	} else {
		buf.Write("\n")
	}
	writeCodeInfo(c, &buf)
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

func narrow(i *int, min, max int) {
	if *i < min {
		*i = min
	}
	if *i > max {
		*i = max
	}
}

func writeCodeInfo(c caller, buf *FeatureBuf) {
	narrow(&c.from, 0, 100)
	narrow(&c.to, c.from, 100)
	for find := false; c.to >= c.from; c.to-- {
		if find {
			buf.NL()
		}
		pc, file, line, ok := runtime.Caller(3 + c.to)
		if ok {
			buf.Writef("%v:%v:", lastPartOf(file), line)
			if fp := runtime.FuncForPC(pc); fp != nil {
				buf.Writef(" in %v:", lastPartOf(fp.Name()))
			}
			find = true
		} else if find || c.to == c.from {
			buf.Write("???:1:")
		}
	}
}

func writeVariables(buf *FeatureBuf, expected, actual interface{}) {
	if expected == nil {
		buf.NL().Write("expected:\t").Highlight(nil)
		buf.NL().Write("  actual:\t")
		// TODO
		return
	} else if actual == nil {
		buf.NL().Write("expected:\t")
		// TODO
		buf.NL().Write("  actual:\t").Highlight(nil)
		return
	}
	e, a := reflect.ValueOf(expected), reflect.ValueOf(actual)
	var b1, b2 FeatureBuf
	b1.Tab = buf.Tab + 1
	b2.Tab = buf.Tab + 1
	var nl, omit, fn bool
	if e.Type() != a.Type() {
		nl, omit, fn = writeDiffTypeValues(&b1, &b2, e, a)
	} else {
		nl, omit, fn = writeDiffValues(&b1, &b2, e, a)
	}
	if nl {
		buf.NL().Write("expected:\t")
		if omit {
			buf.Write("(").Highlight("Only diffs are shown").Write(")")
		}
		buf.Tab++
		buf.NL().Write(b1.String())
		buf.Tab--
		buf.NL().Write("  actual:\t")
		buf.Tab++
		buf.NL().Write(b2.String())
		//if fn {
		//    buf.NL().Write("(").Highlight("func can only be compared to nil").Write(")")
		//}
		buf.Tab--
	} else {
		buf.NL().Writef("expected:\t%v", b1.String())
		buf.NL().Writef("  actual:\t%v", b2.String())
		if omit {
			buf.NL().Write("\t\t(").Highlight("Only diffs are shown").Write(")")
		}
		if fn {
			buf.NL().Write("\t\t(").Highlight("func can only be compared to nil").Write(")")
		}
	}
	/*
		e, a := reflect.ValueOf(expected), reflect.ValueOf(actual)
		var v ValueDiffer
		v.WriteDiff(e, a, buf.Tab+1)
		if v.Attrs[NewLine] {
			buf.NL().Write("expected:\t")
			if v.Attrs[Omit] {
				buf.Write("(").Highlight("Only diffs are shown").Write(")")
			}
			buf.Tab++
			buf.NL().Write(v.String(0))
			buf.Tab--
			buf.NL().Write("  actual:\t")
			buf.Tab++
			buf.NL().Write(v.String(0))
			if v.Attrs[CompFunc] {
				buf.NL().Write("(").Highlight("func can only be compared to nil").Write(")")
			}
			buf.Tab--
		} else {
			buf.NL().Writef("expected:\t%v", v.String1())
			buf.NL().Writef("  actual:\t%v", v.String2())
			if v.Attrs[Omit] {
				buf.NL().Write("\t\t(").Highlight("Only diffs are shown").Write(")")
			}
			if v.Attrs[CompFunc] {
				buf.NL().Write("\t\t(").Highlight("func can only be compared to nil").Write(")")
			}
		}
	*/
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

func writeDiffTypeValues(b1, b2 *FeatureBuf, v1, v2 reflect.Value) (nl, omit, fn bool) {
	return
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

func writeDiffValues(b1, b2 *FeatureBuf, v1, v2 reflect.Value) (nl, omit, fn bool) {
	switch v1.Kind() {
	case reflect.Complex64, reflect.Complex128:
		writeDiffComplexValue(b1, b2, v1, v2)
	case reflect.Array:
		nl, omit, fn = writeDiffArrayValue(b1, b2, v1, v2)
	case reflect.Chan:
		b1.Writef("(%v)(", v1.Type()).Highlight(v1).Write(")")
		b2.Writef("(%v)(", v2.Type()).Highlight(v2).Write(")")
	case reflect.Func:
		fn = writeDiffFuncValue(b1, b2, v1, v2)
	case reflect.Map:
		writeDiffMapValue(b1, b2, v1, v2)
	case reflect.String:
		writeDiffStringValue(b1, b2, v1, v2)
	default:
		b1.Highlightf("%#v", v1)
		b2.Highlightf("%#v", v2)
	}
	return
}

func writeDiffStringValue(b1, b2 *FeatureBuf, v1, v2 reflect.Value) {
	s1, s2 := fmt.Sprintf("%#v", v1), fmt.Sprintf("%#v", v2)
	for i := 0; i < len(s1) || i < len(s2); i++ {
		if i >= len(s1) {
			b2.Highlight(s2[i:])
			break
		} else if i >= len(s2) {
			b1.Highlight(s1[i:])
			break
		} else if s1[i] == s2[i] {
			b1.Write(s1[i : i+1])
			b2.Write(s2[i : i+1])
		} else {
			b1.Highlight(s1[i : i+1])
			b2.Highlight(s2[i : i+1])
		}
	}
}

func writeDiffMapValue(b1, b2 *FeatureBuf, v1, v2 reflect.Value) {
	b1.Write(v1.Type(), "{")
	b2.Write(v2.Type(), "{")
	var m1, m2 []reflect.Value
	i1, i2 := 0, 0
	for _, k := range v1.MapKeys() {
		e1, e2 := v1.MapIndex(k), v2.MapIndex(k)
		if !e2.IsValid() {
			m1 = append(m1, k)
		} else if !reflect.DeepEqual(e1.Interface(), e2.Interface()) {
			if i1 > 0 {
				b1.Write(", ")
			}
			i1++
			writeValue(b1, k)
			b1.Write(":")
			if i2 > 0 {
				b2.Write(", ")
			}
			i2++
			writeValue(b2, k)
			b2.Write(":")
			writeDiffValues(b1, b2, e1, e2)
		}
	}
	// TODO
	_ = m2
	b1.Write("}")
	b2.Write("}")
}

func writeDiffFuncValue(b1, b2 *FeatureBuf, v1, v2 reflect.Value) (fn bool) {
	t := fmt.Sprint(v1.Type())
	b1.Writef("(%v)(", t)
	b2.Writef("(%v)(", t)
	fn = writeDiffFuncValueShort(b1, b2, v1, v2)
	b1.Writef(")")
	b2.Writef(")")
	return
}

func writeDiffFuncValueShort(b1, b2 *FeatureBuf, v1, v2 reflect.Value) (fn bool) {
	p1, p2 := "nil", "nil"
	if !v1.IsNil() {
		p1 = fmt.Sprint(v1)
	}
	if !v2.IsNil() {
		p2 = fmt.Sprint(v2)
	}
	fn = p1 == p2
	b1.Highlight(p1)
	b2.Highlight(p2)
	return
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

func writeDiffArrayValue(b1, b2 *FeatureBuf, v1, v2 reflect.Value) (nl, omit, fn bool) {
	if v1.Len() < 1 || (v1.Len() <= 10 && isSimpleType(v1.Index(0).Kind())) {
		writeDiffArrayShort(b1, b2, v1, v2)
	} else if isSimpleType(v1.Index(0).Kind()) {
		omit = writeDiffArrayLong(b1, b2, v1, v2)
	} else if v1.Index(0).Kind() == reflect.Func {
		omit, fn = writeDiffArrayFunc(b1, b2, v1, v2)
	} else {
		omit = writeDiffArrayComposite(b1, b2, v1, v2)
		nl = true
	}
	return
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

func writeDiffArrayLong(b1, b2 *FeatureBuf, v1, v2 reflect.Value) (omit bool) {
	b1.Write(v1.Type(), "{")
	b2.Write(v2.Type(), "{")
	for i, j := 0, 0; i < v1.Len(); i++ {
		e1, e2 := v1.Index(i), v2.Index(i)
		if reflect.DeepEqual(e1.Interface(), e2.Interface()) {
			omit = true
			continue // Only diffs are shown
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
	return
}

func writeDiffArrayFunc(b1, b2 *FeatureBuf, v1, v2 reflect.Value) (omit, fn bool) {
	if v1.Len() <= 10 {
		fn = writeDiffArrayFuncShort(b1, b2, v1, v2)
	} else {
		omit, fn = writeDiffArrayFuncLong(b1, b2, v1, v2)
	}
	return
}

func writeDiffArrayFuncShort(b1, b2 *FeatureBuf, v1, v2 reflect.Value) (fn bool) {
	b1.Write(v1.Type(), "[")
	b2.Write(v2.Type(), "[")
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
			if writeDiffFuncValueShort(b1, b2, e1, e2) {
				fn = true
			}
		}
	}
	b1.Write("]")
	b2.Write("]")
	return
}

func writeDiffArrayFuncLong(b1, b2 *FeatureBuf, v1, v2 reflect.Value) (omit, fn bool) {
	b1.Write(v1.Type(), "{")
	b2.Write(v1.Type(), "{")
	for i, j := 0, 0; i < v1.Len(); i++ {
		e1, e2 := v1.Index(i), v2.Index(i)
		if reflect.DeepEqual(e1.Interface(), e2.Interface()) {
			omit = true
			continue
		}
		if j > 0 {
			b1.Write(", ")
			b2.Write(", ")
		}
		j++
		b1.Write(i, ":")
		b2.Write(i, ":")
		if writeDiffFuncValueShort(b1, b2, e1, e2) {
			fn = true
		}
	}
	b1.Write("}")
	b2.Write("}")
	return
}

func writeDiffArrayComposite(b1, b2 *FeatureBuf, v1, v2 reflect.Value) (omit bool) {
	b1.Write(v1.Type(), "{")
	b2.Write(v2.Type(), "{")
	b1.Tab++
	b2.Tab++
	idx := v1.Len() > 10
	for i := 0; i < v1.Len(); i++ {
		e1, e2 := v1.Index(i), v2.Index(i)
		eq := reflect.DeepEqual(e1.Interface(), e2.Interface())
		if eq && idx {
			omit = true
			continue // Only diffs are shown
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
	return
}

func writeValue(b *FeatureBuf, v reflect.Value) {
	switch v.Kind() {
	case reflect.Array:
		writeArrayValue(b, v)
	case reflect.Func:
		b.Write(v)
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
	return k != reflect.Array && k != reflect.Func
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
