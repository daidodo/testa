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

import (
	"bytes"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"unsafe"
)

/* TODO:
Nil/NotNil
Error/NoError
Contain/NotContain
Empty/NotEmpty
EqualValue/NotEqualValue
EqualError
*/

// True asserts that a is true.
func True(t *testing.T, a bool, messages ...interface{}) {
	CallerT{1, 1}.True(t, a, messages...)
}

// False asserts that a is false.
func False(t *testing.T, a bool, messages ...interface{}) {
	CallerT{1, 1}.False(t, a, messages...)
}

// Equal asserts that e and a are exactly the same, both type and value.
func Equal(t *testing.T, e, a interface{}, messages ...interface{}) {
	CallerT{1, 1}.Equal(t, e, a, messages...)
}

// NotEqual asserts that e and a are not the same, either type or value.
func NotEqual(t *testing.T, e, a interface{}, messages ...interface{}) {
	CallerT{1, 1}.NotEqual(t, e, a, messages...)
}

// CallerT is useful for customizing caller information shown for assertions.
type CallerT struct {
	from, to int
}

// Caller changes caller information shown for assertions.
func Caller(lv int) CallerT {
	return CallerT{0, lv}
}

// True asserts that a is true.
func (c CallerT) True(t *testing.T, actual bool, messages ...interface{}) {
	if actual {
		return
	}
	fail(c, t, true, actual, true, messages...)
}

// False asserts that a is false.
func (c CallerT) False(t *testing.T, actual bool, messages ...interface{}) {
	if !actual {
		return
	}
	fail(c, t, false, actual, true, messages...)
}

// Equal asserts that e and a are exactly the same, both type and value.
func (c CallerT) Equal(t *testing.T, expected, actual interface{}, messages ...interface{}) {
	if reflect.DeepEqual(expected, actual) {
		return
	}
	fail(c, t, expected, actual, true, messages...)
}

// NotEqual asserts that e and a are not the same, either type or value.
func (c CallerT) NotEqual(t *testing.T, expected, actual interface{}, messages ...interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		return
	}
	fail(c, t, expected, actual, false, messages...)
}

func fail(c CallerT, t *testing.T, expected, actual interface{}, eq bool, msg ...interface{}) {
	var buf bytes.Buffer
	b := tFeatureBuf{w: &buf, Tab: 0}
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

func writeFailEq(buf *tFeatureBuf, expected, actual interface{}) {
	var v tValueDiffer
	v.WriteDiff(reflect.ValueOf(expected), reflect.ValueOf(actual), buf.Tab+1)
	buf.NL().Normalf("Expect:\t%v", v.String(0))
	buf.NL().Normalf("Actual:\t%v", v.String(1))
	writeAttrs(buf, v)
}

func writeFailNe(buf *tFeatureBuf, actual interface{}) {
	var v tValueDiffer
	v.WriteTypeValue(0, reflect.ValueOf(actual), buf.Tab+1)
	buf.NL().Normal("Expect:\t").Highlight("SAME as Actual").Finish()
	buf.NL().Normalf("Actual:\t%v", v.String(0))
	writeAttrs(buf, v)
}

func writeAttrs(buf *tFeatureBuf, v tValueDiffer) {
	if v.Attrs[kOmitSame] {
		buf.NL().Normal("\t(").Highlight("Only diffs are shown").Normal(")")
	}
	if v.Attrs[kCompFunc] {
		buf.NL().Normal("\t(").Highlight("func can only be compared to nil").Normal(")")
	}
}

func writeCodeInfo(c CallerT, buf *tFeatureBuf) {
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

func writeMessages(buf *tFeatureBuf, messages ...interface{}) {
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
