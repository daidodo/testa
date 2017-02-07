package assert

import (
	"bytes"
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

type caller struct {
	from, to int
}

func Caller(level int) caller {
	return caller{0, level}
}

func (c caller) True(t *testing.T, actual bool, messages ...interface{}) {
	if actual {
		return
	}
	fail(c, t, true, actual, true, messages...)
}

func (c caller) False(t *testing.T, actual bool, messages ...interface{}) {
	if !actual {
		return
	}
	fail(c, t, false, actual, true, messages...)
}

func (c caller) Equal(t *testing.T, expected, actual interface{}, messages ...interface{}) {
	if reflect.DeepEqual(expected, actual) {
		return
	}
	fail(c, t, expected, actual, true, messages...)
}

func (c caller) NotEqual(t *testing.T, expected, actual interface{}, messages ...interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		return
	}
	fail(c, t, expected, actual, false, messages...)
}

func fail(c caller, t *testing.T, expected, actual interface{}, eq bool, msg ...interface{}) {
	var buf bytes.Buffer
	b := FeatureBuf{w: &buf, Tab: 0}
	writeCodeInfo(c, &b)
	b.Tab++
	if eq {
		writeFailEq(&b, expected, actual)
	} else {
		writeFailNe(&b, actual)
	}
	writeMessages(&b, msg...)
	b.Tab--
	b.Finish()
	flushLog(t, &buf)
	t.FailNow()
}

func writeFailEq(buf *FeatureBuf, expected, actual interface{}) {
	var v ValueDiffer
	v.WriteDiff(reflect.ValueOf(expected), reflect.ValueOf(actual), buf.Tab)
	if v.Attrs[NewLine] {
		buf.NL().Write("Expected:")
		if v.Attrs[OmitSame] {
			buf.Write("\t(").Highlight("Only diffs are shown").Write(")")
		}
		buf.Tab++
		buf.NL().Write(v.String(0))
		buf.Tab--
		buf.NL().Write("  Actual:")
		buf.Tab++
		buf.NL().Write(v.String(0))
		if v.Attrs[CompFunc] {
			buf.NL().Write("(").Highlight("func can only be compared to nil").Write(")")
		}
		buf.Tab--
	} else {
		buf.NL().Writef("Expected:\t%v", v.String(0))
		buf.NL().Writef("  Actual:\t%v", v.String(1))
		if v.Attrs[OmitSame] {
			buf.NL().Write("\t\t(").Highlight("Only diffs are shown").Write(")")
		}
		if v.Attrs[CompFunc] {
			buf.NL().Write("\t\t(").Highlight("func can only be compared to nil").Write(")")
		}
	}
}

func writeFailNe(buf *FeatureBuf, actual interface{}) {
	var v ValueDiffer
	v.WriteTypeValue(0, reflect.ValueOf(actual))
	if v.Attrs[NewLine] {
		buf.NL().Write("Expected:\t").Highlight("SAME as Actual")
		buf.NL().Write("  Actual:")
		buf.Tab++
		buf.NL().Write(v.String(0))
		buf.Tab--
	} else {
		buf.NL().Write("Expected:\t").Highlight("SAME as Actual")
		buf.NL().Writef("  Actual:\t%v", v.String(0))
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

func writeMessages(buf *FeatureBuf, messages ...interface{}) {
	if len(messages) < 1 {
		return
	}
	var m, h string
	if s, ok := messages[0].(string); ok {
		m = fmt.Sprintf(s, messages[1:]...)
	} else {
		m = fmt.Sprint(messages...)
	}
	for i := 0; i < buf.Tab; i++ {
		h = h + "\t"
	}
	m = format(h, m)
	buf.Write(m)
}

func flushLog(t *testing.T, buf *bytes.Buffer) {
	if kHOOK {
		buf.WriteByte('\n')
		output := buf.Bytes()
		tt := (*common)(unsafe.Pointer(t))
		tt.su.Lock()
		tt.output = output
		tt.su.Unlock()
	} else {
		t.Log("\n" + buf.String())
	}
	t.FailNow()
}

func narrow(i *int, min, max int) {
	if *i < min {
		*i = min
	}
	if *i > max {
		*i = max
	}
}

func lastPartOf(str string) string {
	if index := strings.LastIndex(str, "/"); index >= 0 {
		return str[index+1:]
	} else if index = strings.LastIndex(str, "\\"); index >= 0 {
		return str[index+1:]
	}
	return str
}

func format(h, s string) string {
	if h == "" {
		return s
	}
	var buf bytes.Buffer
	for _, l := range strings.Split(s, "\n") {
		buf.WriteString("\n")
		if l != "" {
			buf.WriteString(h)
		}
		buf.WriteString(l)
	}
	return buf.String()
}

//-----------old

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
