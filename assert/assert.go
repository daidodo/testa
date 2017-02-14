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

func True(t *testing.T, actual bool, messages ...interface{}) {
	caller{1, 1}.True(t, actual, messages...)
}

func False(t *testing.T, actual bool, messages ...interface{}) {
	caller{1, 1}.False(t, actual, messages...)
}

func Equal(t *testing.T, expected, actual interface{}, messages ...interface{}) {
	caller{1, 1}.Equal(t, expected, actual, messages...)
}

func NotEqual(t *testing.T, expected, actual interface{}, messages ...interface{}) {
	caller{1, 1}.NotEqual(t, expected, actual, messages...)
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
	v.WriteDiff(reflect.ValueOf(expected), reflect.ValueOf(actual), buf.Tab+1)
	buf.NL().Normalf("Expect:\t%v", v.String(0))
	buf.NL().Normalf("Actual:\t%v", v.String(1))
	writeAttrs(buf, v)
}

func writeFailNe(buf *FeatureBuf, actual interface{}) {
	var v ValueDiffer
	v.WriteTypeValue(0, reflect.ValueOf(actual), buf.Tab+1)
	buf.NL().Normal("Expect:\t").Highlight("SAME as Actual").Finish()
	buf.NL().Normalf("Actual:\t%v", v.String(0))
	writeAttrs(buf, v)
}

func writeAttrs(buf *FeatureBuf, v ValueDiffer) {
	if v.Attrs[OmitSame] {
		buf.NL().Normal("\t(").Highlight("Only diffs are shown").Normal(")")
	}
	if v.Attrs[CompFunc] {
		buf.NL().Normal("\t(").Highlight("func can only be compared to nil").Normal(")")
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
