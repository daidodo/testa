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
		buf.NL().Normal("Expected:")
		if v.Attrs[OmitSame] {
			buf.Normal("\t(").Highlight("Only diffs are shown").Normal(")")
		}
		buf.Tab++
		buf.NL().Normal(v.String(0))
		buf.Tab--
		buf.NL().Normal("  Actual:")
		buf.Tab++
		buf.NL().Normal(v.String(0))
		if v.Attrs[CompFunc] {
			buf.NL().Normal("(").Highlight("func can only be compared to nil").Normal(")")
		}
		buf.Tab--
	} else {
		buf.NL().Normalf("Expected:\t%v", v.String(0))
		buf.NL().Normalf("  Actual:\t%v", v.String(1))
		if v.Attrs[OmitSame] {
			buf.NL().Normal("\t\t(").Highlight("Only diffs are shown").Normal(")")
		}
		if v.Attrs[CompFunc] {
			buf.NL().Normal("\t\t(").Highlight("func can only be compared to nil").Normal(")")
		}
	}
}

func writeFailNe(buf *FeatureBuf, actual interface{}) {
	var v ValueDiffer
	v.WriteTypeValue(0, reflect.ValueOf(actual), buf.Tab)
	if v.Attrs[NewLine] {
		buf.NL().Normal("Expected:\t").Highlight("SAME as Actual")
		buf.NL().Normal("  Actual:")
		buf.Tab++
		buf.NL().Normal(v.String(0))
		buf.Tab--
	} else {
		buf.NL().Normal("Expected:\t").Highlight("SAME as Actual")
		buf.NL().Normalf("  Actual:\t%v", v.String(0))
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
			buf.Normalf("%v:%v:", lastPartOf(file), line)
			if fp := runtime.FuncForPC(pc); fp != nil {
				buf.Normalf(" in %v:", lastPartOf(fp.Name()))
			}
			find = true
		} else if find || c.to == c.from {
			buf.Normal("???:1:")
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
	buf.Normal(m)
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
		b1.Normalf("(%v)(", v1.Type()).Highlight(v1).Normal(")")
		b2.Normalf("(%v)(", v2.Type()).Highlight(v2).Normal(")")
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
			b1.Normal(s1[i : i+1])
			b2.Normal(s2[i : i+1])
		} else {
			b1.Highlight(s1[i : i+1])
			b2.Highlight(s2[i : i+1])
		}
	}
}

func writeDiffMapValue(b1, b2 *FeatureBuf, v1, v2 reflect.Value) {
	b1.Normal(v1.Type(), "{")
	b2.Normal(v2.Type(), "{")
	var m1, m2 []reflect.Value
	i1, i2 := 0, 0
	for _, k := range v1.MapKeys() {
		e1, e2 := v1.MapIndex(k), v2.MapIndex(k)
		if !e2.IsValid() {
			m1 = append(m1, k)
		} else if !reflect.DeepEqual(e1.Interface(), e2.Interface()) {
			if i1 > 0 {
				b1.Normal(", ")
			}
			i1++
			writeValue(b1, k)
			b1.Normal(":")
			if i2 > 0 {
				b2.Normal(", ")
			}
			i2++
			writeValue(b2, k)
			b2.Normal(":")
			writeDiffValues(b1, b2, e1, e2)
		}
	}
	// TODO
	_ = m2
	b1.Normal("}")
	b2.Normal("}")
}

func writeDiffFuncValue(b1, b2 *FeatureBuf, v1, v2 reflect.Value) (fn bool) {
	t := fmt.Sprint(v1.Type())
	b1.Normalf("(%v)(", t)
	b2.Normalf("(%v)(", t)
	fn = writeDiffFuncValueShort(b1, b2, v1, v2)
	b1.Normalf(")")
	b2.Normalf(")")
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
	b1.Normal("(")
	b2.Normal("(")
	if real(c1) == real(c2) {
		b1.Normal(real(c1))
		b2.Normal(real(c2))
	} else {
		b1.Highlight(real(c1))
		b2.Highlight(real(c2))
	}
	b1.Normal("+")
	b2.Normal("+")
	if imag(c1) == imag(c2) {
		b1.Normal(imag(c1))
		b2.Normal(imag(c2))
	} else {
		b1.Highlight(imag(c1))
		b2.Highlight(imag(c2))
	}
	b1.Normal(")")
	b2.Normal(")")
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
	b1.Normal("[")
	b2.Normal("[")
	for i := 0; i < v1.Len(); i++ {
		e1, e2 := v1.Index(i), v2.Index(i)
		if reflect.DeepEqual(e1.Interface(), e2.Interface()) {
			if i > 0 {
				b1.Normal(" ")
				b2.Normal(" ")
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
	b1.Normal("]")
	b2.Normal("]")
}

func writeDiffArrayLong(b1, b2 *FeatureBuf, v1, v2 reflect.Value) (omit bool) {
	b1.Normal(v1.Type(), "{")
	b2.Normal(v2.Type(), "{")
	for i, j := 0, 0; i < v1.Len(); i++ {
		e1, e2 := v1.Index(i), v2.Index(i)
		if reflect.DeepEqual(e1.Interface(), e2.Interface()) {
			omit = true
			continue // Only diffs are shown
		}
		if j > 0 {
			b1.Normal(", ")
			b2.Normal(", ")
		}
		j++
		b1.Normal(i, ":")
		b2.Normal(i, ":")
		writeDiffValues(b1, b2, e1, e2)
	}
	b1.Normal("}")
	b2.Normal("}")
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
	b1.Normal(v1.Type(), "[")
	b2.Normal(v2.Type(), "[")
	for i := 0; i < v1.Len(); i++ {
		e1, e2 := v1.Index(i), v2.Index(i)
		if reflect.DeepEqual(e1.Interface(), e2.Interface()) {
			if i > 0 {
				b1.Normal(" ")
				b2.Normal(" ")
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
	b1.Normal("]")
	b2.Normal("]")
	return
}

func writeDiffArrayFuncLong(b1, b2 *FeatureBuf, v1, v2 reflect.Value) (omit, fn bool) {
	b1.Normal(v1.Type(), "{")
	b2.Normal(v1.Type(), "{")
	for i, j := 0, 0; i < v1.Len(); i++ {
		e1, e2 := v1.Index(i), v2.Index(i)
		if reflect.DeepEqual(e1.Interface(), e2.Interface()) {
			omit = true
			continue
		}
		if j > 0 {
			b1.Normal(", ")
			b2.Normal(", ")
		}
		j++
		b1.Normal(i, ":")
		b2.Normal(i, ":")
		if writeDiffFuncValueShort(b1, b2, e1, e2) {
			fn = true
		}
	}
	b1.Normal("}")
	b2.Normal("}")
	return
}

func writeDiffArrayComposite(b1, b2 *FeatureBuf, v1, v2 reflect.Value) (omit bool) {
	b1.Normal(v1.Type(), "{")
	b2.Normal(v2.Type(), "{")
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
			b1.Normal(i, ":")
			b2.Normal(i, ":")
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
	b1.NL().Normal("}")
	b2.NL().Normal("}")
	return
}

func writeValue(b *FeatureBuf, v reflect.Value) {
	switch v.Kind() {
	case reflect.Array:
		writeArrayValue(b, v)
	case reflect.Func:
		b.Normal(v)
	default:
		b.Normalf("%#v", v)
	}
}

func writeArrayValue(b *FeatureBuf, v reflect.Value) {
	if v.Len() < 1 || (v.Len() <= 10 && isSimpleType(v.Index(0).Kind())) {
		b.Normal(v)
	} else if isSimpleType(v.Index(0).Kind()) {
		b.Normal(v.Type(), "{")
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				b.Normal(", ")
			}
			b.Normal(i, ":")
			writeValue(b, v.Index(i))
		}
		b.Normal("}")
	} else {
		idx := v.Len() > 10
		b.Normal(v.Type(), "{")
		b.Tab++
		for i := 0; i < v.Len(); i++ {
			b.NL()
			if idx {
				b.Normal(i, ":")
			}
			writeValue(b, v.Index(i))
		}
		b.Tab--
		b.NL().Normal("}")
	}
}

func isSimpleType(k reflect.Kind) bool {
	return k != reflect.Array && k != reflect.Func
}
